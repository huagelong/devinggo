export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.config.errorHandler = (error) => {
    console.error(error)
    // console.log(instance)
    // console.log(info)
  }
})
