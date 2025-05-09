// 移除显式导入，使用Nuxt3自动导入机制
import { useHelper } from './useHelper'

// 导入节流函数
function throttle(fn, delay = 500) {
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

// 创建一个异步函数来获取MD5
async function getMD5() {
  // 检查是否在客户端环境
  if (useHelper.isClient()) {
    // 客户端环境使用crypto-js
    const CryptoJS = await import('crypto-js')
    return CryptoJS.MD5
  }

  // 服务端环境
  try {
    if (process.env.NODE_ENV === 'production') {
      // 生产环境使用Node.js原生crypto模块
      const crypto = await import('node:crypto')
      return text => crypto.createHash('md5').update(text).digest('hex')
    }
    // 开发环境降级方案
    const CryptoJS = await import('crypto-js')
    return CryptoJS.MD5
  }
  catch (error) {
    console.error('加密模块加载失败:', error)
    throw createError({
      statusCode: 500,
      message: '加密服务初始化失败',
      data: {
        code: 'ENCRYPT_INIT_FAILED',
        env: process.env.NODE_ENV,
      },
    })
  }
}

function getConfig() {
  const runtimeConfig = useRuntimeConfig()
  return {
    baseURL: runtimeConfig.public.baseURL, // 使用本地代理路径
    appId: runtimeConfig.app.appId,
    appSecret: runtimeConfig.app.appSecret,
    defaultLang: 'zh_CN',
  }
}
// 生成签名
async function generateSignature(appSecret) {
  try {
    const xtimestamp = String(new Date().getTime())
    const xnonce = String(Math.floor(Math.random() * 999999999 + 99999))
    // 获取MD5函数
    const md5Func = await getMD5()
    // 使用MD5加密算法生成签名
    const xsign = md5Func(appSecret + xtimestamp + xnonce).toString()
    return {
      timestamp: xtimestamp,
      nonce: xnonce,
      sign: xsign,
    }
  }
  catch (error) {
    console.error('Error generating signature:', error)
    // 如果出错，返回一个简单的字符串作为备用
    return undefined
  }
}

// 获取token
async function getToken() {
  const nuxtApp = useNuxtApp()
  return nuxtApp.runWithContext(async () => {
    try {
      const tokenCookie = useCookie('token')
      const expireCookie = useCookie('token_expire')

      // 如果token存在且未过期，直接返回
      if (tokenCookie.value && expireCookie.value) {
        const expireTime = Number.parseInt(expireCookie.value)
        // 提前一分钟刷新token
        if (Date.now() < (expireTime - 60000)) {
          return tokenCookie.value
        }
        else {
          // 如果token即将过期，尝试刷新
          return await refreshToken(tokenCookie.value)
        }
      }
      // 获取新token
      const signatureParams = await generateSignature(getConfig().appSecret)
      if (!signatureParams) {
        AMessage.error('Error generating signature')
        return
      }

      // 从cookie中获取语言设置，如果不存在则使用默认值
      const langCookie = useCookie('language')
      const language = langCookie.value || getConfig().defaultLang
      const response = await $fetch('/api/getToken', {
        baseURL: getConfig().baseURL,
        method: 'GET',
        params: {
          app_id: getConfig().appId,
          signature: signatureParams.sign,
          timestamp: signatureParams.timestamp,
          nonce: signatureParams.nonce,
          language,
        },
      })

      if (response && response.success && response.data && response.data.token) {
        const token = response.data.token
        const expire = response.data.expire * 1000 // 转换为毫秒

        // 保存token到cookie
        tokenCookie.value = token
        expireCookie.value = expire.toString()

        return token
      }

      return null
    }
    catch (error) {
      console.error('获取token失败:', error)
      return null
    }
  })
}

// 刷新token
async function refreshToken(currentToken) {
  const nuxtApp = useNuxtApp()
  return nuxtApp.runWithContext(async () => {
    try {
      const tokenCookie = useCookie('token')
      const expireCookie = useCookie('token_expire')

      // 从cookie中获取语言设置，如果不存在则使用默认值
      const langCookie = useCookie('language')
      const language = langCookie.value || getConfig().defaultLang

      const response = await $fetch('/api/refreshToken', {
        baseURL: getConfig().baseURL,
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${currentToken}`,
          'Accept-Language': language,
          'X-App-Id': getConfig().appId,
        },
      })

      if (response && response.success && response.data && response.data.token) {
        const token = response.data.token
        const expire = response.data.expire * 1000 // 转换为毫秒

        // 更新token到cookie
        tokenCookie.value = token
        expireCookie.value = expire.toString()

        return token
      }

      return currentToken
    }
    catch (error) {
      console.error('刷新token失败:', error)
      return currentToken
    }
  })
}

// 请求体封装
async function applyOptions(options = {}) {
  options.baseURL = options.baseURL ?? getConfig().baseURL
  options.initialCache = options.initialCache ?? false
  options.headers = options.headers || {}
  options.method = options.method || 'GET'
  options.timeout = 3000
  // options.credentials = 'include'

  // 从cookie中获取语言设置，如果不存在则使用默认值
  const langCookie = useCookie('language')
  const language = langCookie.value || getConfig().defaultLang

  let headers = {
    'accept': 'application/json',
    'Accept-Language': language,
    'X-App-Id': getConfig().appId,
  }

  // 获取token
  const token = useCookie('token')
  if (!token.value) {
    // 如果没有token，尝试获取
    const newToken = await getToken()
    if (newToken) {
      headers.Authorization = `Bearer ${newToken}`
    }
  }
  else {
    headers.Authorization = `Bearer ${token.value}`
  }

  const csrfToken = useCookie('XSRF-TOKEN')
  if (csrfToken) {
    headers['X-XSRF-TOKEN'] = csrfToken
  }

  // 在服务端运行时添加请求头
  if (useHelper.isServer()) {
    try {
      const serverHeaders = useRequestHeaders(['referer', 'cookie'])
      headers = {
        ...headers,
        ...serverHeaders,
      }
    }
    catch (error) {
      console.error('Error getting server headers:', error)
    }
  }

  options.headers = {
    ...headers,
    ...options.headers,
  }

  return options
}

function handleError(response) {
  const err = async (text) => {
    if (useHelper.isClient()) {
      try {
        // 动态导入Arco组件
        const Arco = await import('@arco-design/web-vue')
        Arco.Message.error(text || '未知错误')
      }
      catch (error) {
        console.error('Error showing client error message:', error)
        console.error(text || '未知错误')
      }
    }
    else {
      // 服务端环境
      try {
        throw createError({ statusCode: 500, message: text })
      }
      catch (error) {
        console.error('Error showing server error message:', error)
        console.error(text || '未知错误')
      }
    }
    return false
  }

  // 检查响应数据结构
  const responseData = response.data?.value || response.data

  if (!responseData) {
    return err('响应数据格式错误')
  }

  // 根据不同的错误码进行处理
  switch (responseData.code) {
    case 0:
      // 请求成功
      return false
    case 1000:
      // 未登录或登录状态已过期，需要重新登录
      throttle(() => {
        err('未登录或登录状态已过期，需要重新登录')
        // 清除token
        const tokenCookie = useCookie('token')
        const expireCookie = useCookie('token_expire')
        tokenCookie.value = null
        expireCookie.value = null

        // 这里可以添加重定向到登录页的逻辑
      })()
      break
    case 65:
      err('服务器资源不存在')
      break
    case 50:
      err('服务器内部错误')
      break
    case 1001:
    case 61:
      err('没有权限访问')
      break
    default:
      err(responseData.message || '未知错误')
  }

  return true
}

async function fetch(key, url, options) {
  options.key = key
  options = await applyOptions(options)
  if (useHelper.isClient()) {
    const data = ref(null)
    const error = ref(null)
    return await $fetch(url, options).then((res) => {
      // 检查是否需要处理错误
      if (res && res.code !== 0 && res.code !== undefined) {
        const errorHandled = handleError({ data: res })
        if (errorHandled) {
          error.value = res.message || '请求失败'
          return { data, error }
        }
      }

      data.value = res
      return {
        data,
        error,
      }
    }).catch((err) => {
      const msg = err?.data?.message || err?.data?.data
      if (useHelper.isClient()) {
        try {
          AMessage.error(msg || '服务端错误')
        }
        catch (error) {
          console.error('Error showing error message:', error)
          console.error(msg || '服务端错误')
        }
      }
      error.value = msg
      return {
        data,
        error,
      }
    })
  }

  // 过期时间毫秒
  if (options.expire && options.expire > 0) {
    options.getCachedData = (key) => {
      const data = nuxtApp.payload.data[key] || nuxtApp.static.data[key]
      if (!data)
        return null

      const expirationDate = new Date(data.fetchedAt)
      expirationDate.setTime(expirationDate.getTime() + options.expire)
      const isExpired = expirationDate.getTime() < Date.now()
      if (isExpired)
        return null
      return data
    }
  }

  // 添加响应拦截器
  options.onResponse = async ({ request, response, options }) => {
    // 检查token是否过期
    if (response._data && response._data.code === 1002) {
      // token过期，尝试重新获取token
      const newToken = await getToken()
      if (newToken) {
        // 更新请求头中的token
        options.headers.Authorization = `Bearer ${newToken}`
        // 重新发送请求
        return $fetch(request, options)
      }
    }
    return response._data
  }

  const res = await useFetch(url, {
    transform(res) {
      return {
        ...res,
        fetchedAt: new Date(),
      }
    },
    // 相当于响应拦截器
    ...options,
  })

  // 客户端错误处理
  if (useHelper.isClient() && res.error.value) {
    const msg = res.error.value?.data?.message || res.error.value?.data?.data
    if (!options.lazy) {
      try {
        AMessage.error(msg || '服务端错误')
      }
      catch (error) {
        console.error('Error showing error message:', error)
        console.error(msg || '服务端错误')
      }
    }
  }

  // 处理业务错误
  if (!res.error.value && res.data.value) {
    const hasError = handleError(res)
    if (hasError && !options.skipErrorHandling) {
      // 如果有业务错误且不跳过错误处理，将错误信息设置到error中
      res.error.value = { message: res.data.value.message || '请求失败' }
    }
  }

  return res
}

// 自动导出
export const useHttp = {
  // GET请求
  get(key, url, params, option) {
    return fetch(key, url, { method: 'get', params, ...option })
  },

  // POST请求
  post(key, url, body, option) {
    return fetch(key, url, { method: 'post', body, ...option })
  },

  // PUT请求
  put(key, url, body, option) {
    return fetch(key, url, { method: 'put', body, ...option })
  },

  // DELETE请求
  delete(key, url, body, option) {
    return fetch(key, url, { method: 'delete', body, ...option })
  },

  // 获取token的方法暴露出去，以便其他地方可以直接调用
  getToken,

  // 刷新token的方法暴露出去
  refreshToken,
}
