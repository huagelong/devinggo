<script setup>
defineProps({
  error: {
    type: Object,
    default: null,
  },
})

definePageMeta({
  layout: 'default',
})

function handleError() {
  clearError({ redirect: '/' })
}

const isDev = useHelper.isDev()
</script>

<template>
  <div class="grid h-screen place-content-center bg-white px-4">
    <div class="text-center">
      <a-result v-if="error.statusCode === 404" class="result" status="404" subtitle="页面没找到" />
      <a-result v-if="error.statusCode === 403" class="result" status="403" subtitle="Forbidden" />
      <a-result v-if="error.statusCode === 500" class="result" status="500" subtitle="服务器错误" />
      <!-- <a-result v-else class="result" status="error" subtitle="异常">
        <div v-if="error.message">
          {{ error.message }}
        </div>
      </a-result> -->
      <a-button
        key="back"
        type="primary" size="small"
        @click="handleError"
      >
        返回首页
      </a-button>
      <div v-if="error.message && isDev">
        {{ error }}
      </div>
    </div>
  </div>
</template>
