<script lang="ts" setup>
import type { GlobalConfigProvider } from 'tdesign-vue-next';

import { onErrorCaptured, watch } from 'vue';

import { usePreferences } from '@vben/preferences';

import { merge } from 'es-toolkit/compat';
import { ConfigProvider } from 'tdesign-vue-next';
import zhConfig from 'tdesign-vue-next/es/locale/zh_CN';

defineOptions({ name: 'App' });
const { isDark } = usePreferences();

// 全局错误捕获：防止未处理的组件错误导致白屏
onErrorCaptured((error, instance, info) => {
  // 开发环境输出详细错误信息
  if (import.meta.env.DEV) {
    console.error('[GlobalErrorBoundary]', error, { component: instance?.$options?.name, info });
  }
  // 返回 false 阻止错误继续向上传播，避免白屏
  // 返回 true 则允许传播到父级错误边界
  return false;
});

watch(
  () => isDark.value,
  (dark) => {
    document.documentElement.setAttribute('theme-mode', dark ? 'dark' : '');
  },
  { immediate: true },
);

const customConfig: GlobalConfigProvider = {
  // 可以在此处定义更多自定义配置，具体可配置内容参看 API 文档
  calendar: {},
  table: {},
  pagination: {},
};
const globalConfig = merge(zhConfig, customConfig);
</script>

<template>
  <ConfigProvider :global-config="globalConfig">
    <RouterView />
  </ConfigProvider>
</template>
