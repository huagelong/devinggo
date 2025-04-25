// 使用Nuxt的composables
import { $fetch, ref, showError, useCookie, useFetch, useRequestHeaders } from '#app'
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

// AMessage组件，根据环境动态导入
let AMessage
if (useHelper.isClient()) {
  // 客户端环境
  import('@arco-design/web-vue').then((Arco) => {
    AMessage = Arco.Message
  }).catch((err) => {
    console.error('Failed to import Arco components:', err)
  })
}

// 创建一个异步函数来获取MD5
async function getMD5() {
  // 检查是否在客户端环境
  if (useHelper.isClient()) {
    // 客户端环境
    const CryptoJS = await import('crypto-js')
    return CryptoJS.MD5
  }
  else {
    // 服务端环境
    try {
      const CryptoJS = await import('crypto-js')
      return CryptoJS.MD5
    }
    catch (error) {
      console.error('Failed to import crypto-js in server environment:', error)
      // 返回一个简单的替代函数，仅用于服务端
      return text => ({
        toString: () => `server-side-md5-${text}`,
      })
    }
  }
}

const CONFIG = {
  baseURL: process.env.BASE_URL,
  appId: process.env.APP_ID, // 默认应用ID
  appSecret: process.env.APP_SECRET, // 应用密钥，实际项目中应该从环境变量或配置中获取
  defaultLang: 'zh_CN', // 默认语言，当cookie中没有语言设置时使用
}

// 生成签名
async function generateSignature(appId, appSecret) {
  try {
    // 获取MD5函数
    const md5Func = await getMD5()
    // 使用MD5加密算法生成签名
    return md5Func(appId + appSecret).toString()
  }
  catch (error) {
    console.error('Error generating signature:', error)
    // 如果出错，返回一个简单的字符串作为备用
    return `${appId}${appSecret}`
  }
}

// 获取token
async function getToken() {
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
    const signature = await generateSignature(CONFIG.appId, CONFIG.appSecret)

    // 从cookie中获取语言设置，如果不存在则使用默认值
    const langCookie = useCookie('language')
    const language = langCookie.value || CONFIG.defaultLang

    const response = await $fetch('/api/getToken', {
      method: 'GET',
      params: {
        app_id: CONFIG.appId,
        signature,
        language, // 添加语言参数
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
}

// 刷新token
async function refreshToken(currentToken) {
  try {
    const tokenCookie = useCookie('token')
    const expireCookie = useCookie('token_expire')

    // 从cookie中获取语言设置，如果不存在则使用默认值
    const langCookie = useCookie('language')
    const language = langCookie.value || CONFIG.defaultLang

    const response = await $fetch('/api/refreshToken', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${currentToken}`,
        'Accept-Language': language,
        'X-App-Id': CONFIG.appId,
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
}

// 请求体封装
async function applyOptions(options = {}) {
  options.baseURL = options.baseURL ?? CONFIG.baseURL
  options.initialCache = options.initialCache ?? false
  options.headers = options.headers || {}
  options.method = options.method || 'GET'
  options.timeout = 3000
  options.credentials = 'include'

  // 从cookie中获取语言设置，如果不存在则使用默认值
  const langCookie = useCookie('language')
  const language = langCookie.value || CONFIG.defaultLang

  let headers = {
    'accept': 'application/json',
    'Accept-Language': language,
    'X-App-Id': CONFIG.appId,
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
  const err = function (text) {
    if (useHelper.isClient()) {
      // 确保AMessage在客户端环境中可用
      try {
        AMessage.error(text || '未知错误')
      }
      catch (error) {
        console.error('Error showing client error message:', error)
        console.error(text || '未知错误')
      }
    }
    else {
      // 服务端环境
      try {
        showError(text)
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

  if (options.$) {
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
    if (response._data && response._data.code === 1000) {
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
