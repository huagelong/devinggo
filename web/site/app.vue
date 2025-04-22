<script setup lang="ts">
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn'

if (useHelper.isClient()) {
  const colorMode = useColorMode()
  document.body.setAttribute('arco-theme', colorMode.value)
  const htmlElement = document.querySelector('html')
  if (htmlElement) {
    htmlElement.classList.add(colorMode.value)
  }
}
useHead({
  title: 'Deving Hub',
})
</script>

<template>
  <a-config-provider :locale="zhCN">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </a-config-provider>
</template>
