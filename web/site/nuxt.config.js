// https://nuxt.com/docs/api/configuration/nuxt-config
console.warn('process.argv: ', process.argv)
console.warn('process.env.NODE_ENV: ', process.env.NODE_ENV)
export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      baseURL: process.env.NUXT_PUBLIC_BASE_URL,
    },
    app: {
      appId: process.env.NUXT_APP_ID,
      appSecret: process.env.NUXT_APP_SECRET,
    },
  },
  ssr: true,
  devtools: { enabled: true },
  experimental: {
    payloadExtraction: true,
  },

  modules: [
    '@nuxtjs/eslint-module',
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
    classSuffix: '',
  },

  arco: {
    importPrefix: 'A',
    hookPrefix: 'Arco',
    locales: ['getLocale'],
    localePrefix: 'Arco',
    theme: '@arco-themes/vue-digitforce',
  },

  css: [
    '~/assets/css/index.scss',
  ],

  pinia: {
    storesDirs: ['./stores/**'],
  },

  imports: {
    autoImport: true,
    dirs: [
      'stores',
    ],
  },

  nitro: {
    devProxy: {
      '/api': {
        target: process.env.NUXT_PUBLIC_BASE_URL,
        changeOrigin: true,
        prependPath: true,
      },
    },
  },

  compatibilityDate: '2025-04-22',
})
