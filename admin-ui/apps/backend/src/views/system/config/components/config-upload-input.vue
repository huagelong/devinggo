<script lang="ts" setup>
import { logger } from '#/utils/logger';
import { ref } from 'vue';

import { $t } from '@vben/locales';

import { UploadIcon } from 'tdesign-icons-vue-next';
import { Button, Input, MessagePlugin, Space } from 'tdesign-vue-next';

import { uploadImageFileApi } from '#/api/system/upload';

const props = defineProps<{
  modelValue?: string;
  placeholder?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const uploading = ref(false);
const fileInputRef = ref<HTMLInputElement>();

function handleInput(value: string) {
  emit('update:modelValue', value);
}

function triggerUpload() {
  fileInputRef.value?.click();
}

async function handleFileChange(event: Event) {
  const input = event.target as HTMLInputElement;
  const file = input.files?.[0];
  if (!file) return;
  uploading.value = true;
  try {
    const response = (await uploadImageFileApi(file)) as { url?: string };
    if (response?.url) {
      emit('update:modelValue', response.url);
      MessagePlugin.success($t('common.uploadSuccess2'));
    } else {
      MessagePlugin.error($t('common.uploadFailed2'));
    }
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.uploadFailed2'));
  } finally {
    uploading.value = false;
    input.value = '';
  }
}
</script>

<template>
  <div class="config-upload-input flex flex-col gap-2">
    <Space class="upload-row">
      <Input
        :model-value="modelValue"
        :placeholder="placeholder ?? $t('common.uploadLinkPlaceholder')"
        class="upload-url-input"
        @change="(value) => handleInput(value as string)"
      />
      <Button :loading="uploading" variant="outline" @click="triggerUpload">
        <template #icon><UploadIcon /></template>
        {{ $t('common.uploadFile') }}
      </Button>
    </Space>
    <div v-if="modelValue" class="rounded-md border border-gray-100 p-2">
      <img
        :src="modelValue"
        alt="config upload preview"
        class="h-24 max-w-full rounded-md object-contain"
      />
    </div>
    <input
      ref="fileInputRef"
      type="file"
      accept="image/*"
      class="hidden"
      @change="handleFileChange"
    />
  </div>
</template>

<style scoped>
.upload-row :deep(.t-space-item:first-child) {
  flex: 1;
  min-width: 0;
}

.upload-url-input {
  width: 100%;
}

.config-upload-input :deep(.t-space) {
  width: 100%;
}

@media (max-width: 767px) {
  .upload-row :deep(.t-space) {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: stretch;
    gap: 8px;
  }

  .upload-row :deep(.t-space-item:last-child .t-button) {
    width: 100%;
  }
}
</style>
