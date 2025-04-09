<template>
  <a-drawer v-bind="$attrs" v-model:visible="visible" @close="handleClose">
    <a-watermark
      v-if="watermarkEnabled"
      :content="[userStore.user.nickname, currentDate]"
      :font-size="14"
      :line-height="14"
      :gap="[80, 80]"
      :rotate="-22"
      class="h-full"
    >
      <slot></slot>
    </a-watermark>
    <div v-else class="h-full">
      <slot></slot>
    </div>
  </a-drawer>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useUserStore } from '@/store'
import dayjs from 'dayjs'

const userStore = useUserStore()
const currentDate = computed(() => dayjs().format('YYYY-MM-DD'))

// 从环境变量中读取水印开关配置
const watermarkEnabled = computed(() => {
  const enabled = import.meta.env.VITE_APP_WATERMARK_ENABLED
  console.log(enabled)
  return enabled === 'true' || enabled === true
})

const visible = ref(false)
const emit = defineEmits(['update:visible'])

const handleClose = () => {
  emit('update:visible', false)
}

defineExpose({
  visible
})
</script>

<style scoped>
.h-full {
  height: 100%;
}
</style>