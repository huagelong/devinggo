export default defineEventHandler(async (event) => {
  console.warn(`New request: ${getRequestURL(event)}`)
})
