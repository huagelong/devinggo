export async function test(query) {
  return await useHttp().get('test.test', '/api/test', query)
}
