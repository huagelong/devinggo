<script lang="ts" setup>
import { $t } from '@vben/locales';

import { AddIcon, InfoCircleIcon } from 'tdesign-icons-vue-next';
import { Button } from 'tdesign-vue-next';

interface Emits {
  (e: 'create'): void;
}

const emit = defineEmits<Emits>();

function handleCreate() {
  emit('create');
}
</script>

<template>
  <div class="config-card lg:w-1/2">
    <div class="config-right-head flex items-center justify-between">
      <h3 class="text-lg font-semibold">{{ $t('system.config.addConfigTitle') }}</h3>
      <Button theme="primary" @click="handleCreate">
        <template #icon><AddIcon /></template>
        {{ $t('common.create') }}
      </Button>
    </div>
    <p class="config-add-tip text-sm">
      {{ $t('system.config.addConfigTip') }}
    </p>
    <div class="config-empty-state">
      <InfoCircleIcon class="config-empty-icon" />
      <span>{{ $t('system.config.addConfigTitle') }}</span>
    </div>
  </div>
</template>

<style scoped>
.config-card {
  border: 1px solid var(--td-component-border, #e7e7e7);
  border-radius: 10px;
  background: #fff;
  padding: 16px 18px;
  box-shadow: 0 6px 18px rgb(15 23 42 / 4%);
}

.config-right-head {
  margin-bottom: 16px;
}

.config-add-tip {
  margin: 0;
  color: var(--td-text-color-secondary, #6b7280);
  line-height: 1.7;
}

.config-empty-state {
  margin-top: 20px;
  min-height: 320px;
  border: 1px dashed #d9e3f3;
  border-radius: 12px;
  background: linear-gradient(180deg, #f8fbff 0%, #ffffff 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
  color: #5f6f85;
}

.config-empty-icon {
  font-size: 24px;
  color: #3b82f6;
}

@media (max-width: 1023px) {
  .config-card {
    padding: 14px;
  }

  .config-empty-state {
    min-height: 200px;
  }
}
</style>