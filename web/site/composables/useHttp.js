// useHttp.js
import { useHelper } from './useHelper'

function _throttle(fn, delay = 500) {
  let timer = null
  return function (...args) {
    if (timer)
      return
    timer = setTimeout(() => {
      fn.apply(this, args)
      timer = null
    }, delay)
  }
}

async function getMD5() {
  if (useHelper.isClient()) {
    const CryptoJS = await import('crypto-js')
    return CryptoJS.MD5
  }

  if (process.env.NODE_ENV === 'production') {
    const crypto = await import('node:crypto')
    return text => crypto.createHash('md5').update(text).digest('hex')
  }

  const CryptoJS = await import('crypto-js')
  return CryptoJS.MD5
}

async function getConfig(nuxtApp) {
  return nuxtApp.runWithContext(async () => {
    const runtimeConfig = useRuntimeConfig()
    return {
      baseURL: runtimeConfig.public.baseURL,
      appId: runtimeConfig.app.appId,
      appSecret: runtimeConfig.app.appSecret,
      defaultLang: 'zh_CN',
    }
  })
}

async function generateSignature(appSecret) {
  const xtimestamp = Date.now().toString()
  const xnonce = Math.floor(Math.random() * 999999999 + 99999).toString()
  const md5Func = await getMD5()
  const xsign = md5Func(appSecret + xtimestamp + xnonce).toString()
  return { timestamp: xtimestamp, nonce: xnonce, sign: xsign }
}

async function getToken(nuxtApp) {
  return nuxtApp.runWithContext(async () => {
    try {
      const tokenCookie = useCookie('token')
      const expireCookie = useCookie('token_expire')

      if (tokenCookie.value && expireCookie.value) {
        const expireTime = Number.parseInt(expireCookie.value)
        if (Date.now() < expireTime - 60000)
          return tokenCookie.value
        return await refreshToken(nuxtApp, tokenCookie.value)
      }

      const config = await getConfig(nuxtApp)
      const signatureParams = await generateSignature(config.appSecret)
      if (!signatureParams) {
        if (useHelper.isClient()) {
          const Arco = await import('@arco-design/web-vue')
          Arco.Message.error('签名生成失败')
        }
        return null
      }

      const langCookie = useCookie('language')
      const language = langCookie.value || config.defaultLang

      let response
      if (useHelper.isClient()) {
        response = await $fetch('/api/getToken', {
          baseURL: config.baseURL,
          method: 'GET',
          params: {
            app_id: config.appId,
            signature: signatureParams.sign,
            timestamp: signatureParams.timestamp,
            nonce: signatureParams.nonce,
            language,
          },
        })
      }
      else {
        const { data } = await useFetch('/api/getToken', {
          baseURL: config.baseURL,
          method: 'GET',
          params: {
            app_id: config.appId,
            signature: signatureParams.sign,
            timestamp: signatureParams.timestamp,
            nonce: signatureParams.nonce,
            language,
          },
          key: 'server:getToken',
        })
        response = data.value
      }

      if (response?.code === 0 && response.data?.token) {
        tokenCookie.value = response.data.token
        expireCookie.value = (response.data.expire * 1000).toString()
        return response.data.token
      }
      return null
    }
    catch (error) {
      console.error('获取Token失败:', error)
      return null
    }
  })
}

async function refreshToken(nuxtApp, currentToken) {
  return nuxtApp.runWithContext(async () => {
    try {
      // 确保useCookie在正确的上下文中调用
      const tokenCookie = useCookie('token')
      const expireCookie = useCookie('token_expire')
      const langCookie = useCookie('language')
      const config = await getConfig(nuxtApp)
      const language = langCookie.value || config.defaultLang

      let response
      if (useHelper.isClient()) {
        response = await $fetch('/api/refreshToken', {
          baseURL: config.baseURL,
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${currentToken}`,
            'Accept-Language': language,
            'X-App-Id': config.appId,
          },
        })
      }
      else {
        const { data } = await useFetch('/api/refreshToken', {
          baseURL: config.baseURL,
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${currentToken}`,
            'Accept-Language': language,
            'X-App-Id': config.appId,
          },
          key: 'server:refreshToken',
        })
        response = data.value
      }

      if (response?.code === 0 && response.data?.token) {
        tokenCookie.value = response.data.token
        expireCookie.value = (response.data.expire * 1000).toString()
        return response.data.token
      }
      return currentToken
    }
    catch (error) {
      console.error('刷新Token失败:', error)
      return currentToken
    }
  })
}

async function applyOptions(nuxtApp, options = {}) {
  return nuxtApp.runWithContext(async () => {
    const config = await getConfig(nuxtApp)
    options.baseURL = options.baseURL ?? config.baseURL
    options.initialCache = options.initialCache ?? false
    options.headers = options.headers || {}
    options.method = options.method || 'GET'
    options.timeout = 3000

    // 确保useCookie在正确的上下文中调用
    const langCookie = useCookie('language')
    const language = langCookie.value || config.defaultLang

    let headers = {
      'accept': 'application/json',
      'Accept-Language': language,
      'X-App-Id': config.appId,
    }

    const tokenCookie = await nuxtApp.runWithContext(async () => useCookie('token'))
    const token = tokenCookie.value
    if (!token) {
      const newToken = await getToken(nuxtApp)
      if (newToken)
        headers.Authorization = `Bearer ${newToken}`
    }
    else {
      headers.Authorization = `Bearer ${token}`
    }

    if (useHelper.isServer()) {
      const serverHeaders = useRequestHeaders(['referer', 'cookie'])
      headers = { ...headers, ...serverHeaders }
    }

    options.headers = { ...headers, ...options.headers }
    return options
  })
}

// 创建带上下文的节流函数
function createThrottleWithContext(nuxtApp) {
  return function (fn, delay = 500) {
    let timer = null
    return function (...args) {
      if (timer)
        return
      timer = setTimeout(() => {
        nuxtApp.runWithContext(() => fn.apply(this, args))
        timer = null
      }, delay)
    }
  }
}

function handleError(nuxtApp, response) {
  // 创建带上下文的工具函数
  // 增强版上下文绑定
  const withContext = function (fn) {
    return function (...args) {
      const contextApp = nuxtApp
      return contextApp.runWithContext(() => fn.call(this, contextApp, ...args))
    }
  }

  // 带上下文的错误显示
  const showError = withContext(async (nuxtApp, text) => {
    if (useHelper.isClient()) {
      const Arco = await import('@arco-design/web-vue')
      Arco.Message.error(text || '未知错误')
    }
    else {
      console.error('服务端错误:', text)
    }
  })

  // 带上下文的清理token
  const clearTokens = withContext((nuxtApp) => {
    nuxtApp.runWithContext(() => {
      const tokenCookie = useCookie('token')
      const expireCookie = useCookie('token_expire')
      tokenCookie.value = null
      expireCookie.value = null
    })
  })

  // 创建带上下文的节流实例
  const throttle = createThrottleWithContext(nuxtApp)

  switch (response?.code) {
    case 1000:
      throttle(() => {
        clearTokens()
        showError('登录状态过期')
      })()
      break
    case 65:
      showError('资源不存在')
      break
    case 50:
      showError('服务器错误')
      break
    case 1002:
      showError('token已过期')
      break
    case 1001:
    case 61:
      showError('无访问权限')
      break
    default:
      showError(response?.message || '未知错误')
  }
}

async function fetch(nuxtApp, key, url, options) {
  return nuxtApp.runWithContext(async () => {
    try {
      options = await applyOptions(nuxtApp, { ...options, key })
      let response

      if (useHelper.isClient()) {
        response = await $fetch(url, options)
        if (response?.code !== 0)
          handleError(nuxtApp, response)
      }
      else {
        const { data } = await useFetch(url, {
          ...options,
          transform: res => ({ ...res, fetchedAt: new Date() }),
        })
        response = data.value
        if (response?.code !== 0)
          handleError(nuxtApp, response)
      }

      return response
    }
    catch (error) {
      console.error('请求异常:', error)
      handleError(nuxtApp, { code: 500, message: error.message })
      return null
    }
  })
}

export function useHttp() {
  const nuxtApp = useNuxtApp()

  const withContext = fn => (...args) =>
    nuxtApp.runWithContext(() => fn(nuxtApp, ...args))

  return {
    get: withContext((nuxtApp, key, url, params, options) =>
      fetch(nuxtApp, key, url, { method: 'GET', params, ...options })),
    post: withContext((nuxtApp, key, url, body, options) =>
      fetch(nuxtApp, key, url, { method: 'POST', body, ...options })),
    put: withContext((nuxtApp, key, url, body, options) =>
      fetch(nuxtApp, key, url, { method: 'PUT', body, ...options })),
    delete: withContext((nuxtApp, key, url, options) =>
      fetch(nuxtApp, key, url, { method: 'DELETE', ...options })),
    getToken: withContext(() => getToken(nuxtApp)),
    refreshToken: withContext(currentToken => refreshToken(nuxtApp, currentToken)),
  }
}
