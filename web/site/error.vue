<script setup>
import hljs from 'highlight.js/lib/core'
import json from 'highlight.js/lib/languages/json'
import 'highlight.js/styles/atom-one-dark.css'

defineProps({
  error: {
    type: Object,
    default: null,
  },
})

hljs.registerLanguage('json', json)
hljs.configure({ ignoreUnescapedHTML: true })

const vHighlight = {
  mounted(el) {
    hljs.highlightElement(el)
  },
  updated(el) {
    hljs.highlightElement(el)
  },
}

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
      <a-button
        key="back"
        type="primary" size="small"
        @click="handleError"
      >
        返回首页
      </a-button>
      <div v-if="error.message && isDev" class="error-details text-center">
        <div v-if="error.stack" class="error-section text-center">
          <div v-highlight style="padding: 10px;">
            <code class="language-javascript">{{ error.stack }}</code>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
