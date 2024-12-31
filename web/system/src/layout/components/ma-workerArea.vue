<template>
  <a-layout-content class="work-area customer-scrollbar relative">
    <div class="h-full" :class="{ 'p-3': $route.path.indexOf('maIframe') === -1 }">
      <router-view v-slot="{ Component }">
        <transition :name="appStore.animation" mode="out-in">
          <keep-alive :include="keepStore.keepAlives">
            <component :is="Component" :key="$route.fullPath" v-if="keepStore.show" />
          </keep-alive>
        </transition>
      </router-view>
      <iframe-view />
    </div>
  </a-layout-content>
</template>

<script setup>
  import { useAppStore, useKeepAliveStore } from '@/store'
  import IframeView from './components/iframe-view.vue'
  const appStore = useAppStore()
  const keepStore = useKeepAliveStore()
</script>