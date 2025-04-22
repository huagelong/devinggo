// https://nuxt.com/docs/api/configuration/nuxt-config
console.log('process.argv: ', process.argv)
console.log('process.env.NODE_ENV: ', process.env.NODE_ENV)
export default defineNuxtConfig({
  ssr: true,
  devtools: { enabled: true },
  experimental: {
    payloadExtraction: true,
  },

  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxtjs/color-mode',
    '@pinia/nuxt',
    'arco-design-nuxt-module',
    ['nuxt-lazy-load', {
        images: true,
        videos: true,
        audios: true,
        iframes: true,
        native: true,
        directiveOnly: false,
        // Default image must be in the public folder
        // defaultImage: '/images/default-image.jpg',
        // To remove class set value to false
        loadingClass: 'isLoading',
        loadedClass: 'isLoaded',
        appendClass: 'lazyLoad',

        observerConfig: {
          // See IntersectionObserver documentation
        },
      }],
  ],

  colorMode: {
    classSuffix: ''
  },

  arco: {
    importPrefix: 'A',
    hookPrefix: 'Arco',
    locales: ['getLocale'],
    localePrefix: 'Arco',
    theme:'@arco-themes/vue-digitforce',
  },

  css: [
    '~/assets/css/index.scss',
  ],

  pinia: {
    autoImports: ['defineStore', 'storeToRefs'],
  },

  imports: {
    autoImport: true,
    dirs: [
      'stores'
    ],
  },

  nitro: {
    routeRules: {
        //使用mock数据先注释
    //   '/api/**': {
    //     proxy: `${import.meta.env.SERVER_URL}/api/**`,
    //   },
    },
  },

  buildModules: ['@nuxtjs/eslint-module'],

  runtimeConfig: {
    // The private keys which are only available server-side
    apiSecret: '12121wwwassdfsfde2',
    env: process.env.NODE_ENV
  },

  compatibilityDate: '2025-04-22'
})