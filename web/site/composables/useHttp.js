const CONFIG = {
  baseURL: process.env.BASE_URL_TEST,
}

// 请求体封装
function applyOptions(options = {}) {
  options.baseURL = options.baseURL ?? CONFIG.baseURL
  options.initialCache = options.initialCache ?? false
  options.headers = options.headers || {}
  options.method = options.method || 'GET'
  options.timeout = 3000
  options.credentials = 'include'
  let headers = {
    accept: 'application/json',
  }
  const token = useCookie('token')
  if (token.value) {
    headers.Authorization = `Bearer ${token.value}`
  }
  const csrfToken = useCookie('XSRF-TOKEN')
  if (csrfToken) {
    headers['X-XSRF-TOKEN'] = csrfToken
  }
  if (process.server) {
    headers = {
      ...headers,
      ...useRequestHeaders(['referer', 'cookie']),
    }
  }
  options.headers = {
    ...headers,
    ...options.headers,
  }
  return options
}

function handleError(res) {
  const err = function (text) {
    if (useHelper.isClient()) {
      AMessage.error(text || '未知错误')
    }
    else {
      showError(text)
    }

    return false
  }
  if (!res.data.value.status) {
    err('请求超时，服务器无响应！')
    return false
  }
  const handleMap = {
    400() { return err(res.data.value.msg) },
    404() { return err('接口不存在') },
    401() { return err('需要登录') },
  }
  if (handleMap[res.data.value.status]) {
    return handleMap[res.data.value.status]()
  }
  return true
}

async function fetch(key, url, options) {
  options.key = key
  options = applyOptions(options)

  if (options.$) {
    const data = ref(null)
    const error = ref(null)
    return await $fetch(url, options).then((res) => {
      data.value = res
      return {
        data,
        error,
      }
    }).catch((err) => {
      const msg = err?.data?.data
      if (useHelper.isClient()) {
        AMessage.error(msg || '服务端错误')
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
    const msg = res.error.value?.data?.data
    if (!options.lazy) {
      AMessage.error(msg || '服务端错误')
    }
  }

  handleError(res)
  return res
}

// 自动导出
export const useHttp = {
  get(key, url, params, option) {
    return fetch(key, url, { method: 'get', params, ...option })
  },
  post(url, body, option) {
    return fetch(key, url, { method: 'post', body, ...option })
  },
  put(url, body, option) {
    return fetch(key, url, { method: 'put', body, ...option })
  },
  delete(url, body, option) {
    return fetch(key, url, { method: 'delete', body, ...option })
  },
}
