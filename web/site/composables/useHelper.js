export function isClient() {
  return process.client
}

export function isServer() {
  return process.server
}

export function isDev() {
  return process.env.NODE_ENV === 'development'
}

export const useHelper = {
  isClient,
  isServer,
  isDev,
}
