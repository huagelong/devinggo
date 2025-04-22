export function test(query) {
  return useHttp.get('test.test', '/api/test', query)
}
