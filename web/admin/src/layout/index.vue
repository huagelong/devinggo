<template>
  <a-watermark
      v-if="watermarkEnabled"
      :content="[userStore.user.nickname, currentDate]"
      :font-size="14"
      :line-height="14"
      :gap="[80, 80]"
      :z-index="9999"
      :rotate="-22"
      class="h-full main-container"
  >
  <a-layout-content class="h-full" >

      <columns-layout v-if="appStore.layout === 'columns'" />
      <classic-layout v-if="appStore.layout === 'classic'" />
      <banner-layout v-if="appStore.layout === 'banner'" />
      <mixed-layout v-if="appStore.layout === 'mixed'" />

      <setting ref="settingRef" />

      <transition name="ma-slide-down" mode="out-in">
        <system-search ref="systemSearchRef" v-show="appStore.searchOpen" />
      </transition>

      <ma-button-menu />

      <div class="max-size-exit" @click="tagExitMaxSize"><icon-close /></div>

    </a-layout-content>
  </a-watermark>
<a-layout-content v-else class="h-full main-container">

      <columns-layout v-if="appStore.layout === 'columns'" />
      <classic-layout v-if="appStore.layout === 'classic'" />
      <banner-layout v-if="appStore.layout === 'banner'" />
      <mixed-layout v-if="appStore.layout === 'mixed'" />

      <setting ref="settingRef" />

      <transition name="ma-slide-down" mode="out-in">
        <system-search ref="systemSearchRef" v-show="appStore.searchOpen" />
      </transition>

      <ma-button-menu />

      <div class="max-size-exit" @click="tagExitMaxSize"><icon-close /></div>

  </a-layout-content>
</template>
<script setup>
  import { onMounted, ref, watch, computed } from 'vue'
  import { useAppStore, useUserStore } from '@/store'
  import dayjs from 'dayjs'

  import ColumnsLayout from './components/columns/index.vue'
  import ClassicLayout from './components/classic/index.vue'
  import BannerLayout from './components/banner/index.vue'
  import MixedLayout from './components/mixed/index.vue'
  import Setting from './setting.vue'
  import SystemSearch from './search.vue'
  import MaButtonMenu from './components/ma-buttonMenu.vue'

  const appStore = useAppStore()
  const userStore = useUserStore()

  const currentDate = computed(() => dayjs().format('YYYY-MM-DD'))

  // 从环境变量中读取水印开关配置
  const watermarkEnabled = computed(() => {
    const enabled = import.meta.env.VITE_APP_WATERMARK_ENABLED
    return enabled === 'true' || enabled === true
  })

  const settingRef = ref()
  const systemSearchRef = ref()
  watch(() => appStore.settingOpen, vl => {
    if (vl === true) {
      settingRef.value.open()
      appStore.settingOpen = false
    }
  })

  const tagExitMaxSize = () => {
    document.getElementById('app').classList.remove('max-size')
  }

  onMounted(() => {
    document.addEventListener('keydown', e => {
      const keyCode = e.keyCode ?? e.which ?? e.charCode
      const altKey = e.altKey ?? e.metaKey
      if(altKey && keyCode === 83) {
        appStore.searchOpen =  true
        return
      }

      if (keyCode === 27) {
        appStore.searchOpen = false
        return
      }
    })
  })

</script>
