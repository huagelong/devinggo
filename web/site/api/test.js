export function test(query) {
  const nuxtApp = useNuxtApp()
  return nuxtApp.runWithContext(async () => {
    return useHttp().get('test.test', '/api/test', query)
  })
}
