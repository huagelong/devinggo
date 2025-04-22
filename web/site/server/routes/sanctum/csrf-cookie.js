import Mock from 'mockjs'

export default defineEventHandler((event) => {
//   const { setCookie } = useCookie()
  const uuid = Mock.mock('@guid')
  setCookie(event, 'XSRF-TOKEN', uuid, {
    maxAge: 60 * 60 * 24 * 7, // Expires in 1 week
    sameSite: 'strict', // Only send cookie on same-site requests
  })
})
