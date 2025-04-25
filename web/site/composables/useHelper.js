export function isClient() {
  return typeof window !== 'undefined'
}

export function isServer() {
  return !isClient()
}

export function isDev() {
  return process.env.NODE_ENV === 'development'
}

export const useHelper = {
  isClient,
  isServer,
  isDev,
}
