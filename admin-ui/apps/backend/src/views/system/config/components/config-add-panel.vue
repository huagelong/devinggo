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
  <div class="config-add-panel">
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
.config-add-panel {
  width: 100%;
}

.config-right-head {
  margin-bottom: 12px;
  padding-bottom: 16px;
  border-bottom: 1px solid hsl(var(--border));
}

.config-right-head h3 {
  color: var(--td-text-color-primary, #1f2937);
  font-size: 16px;
}

.config-add-tip {
  margin: 0 0 20px 0;
  color: var(--td-text-color-secondary, #6b7280);
  line-height: 1.7;
  font-size: 13px;
  padding: 10px 14px;
  background: hsl(var(--muted) / 0.3);
  border-radius: 8px;
  border-left: 3px solid hsl(var(--primary));
}

.config-empty-state {
  min-height: 360px;
  border: 1px dashed hsl(var(--primary) / 0.2);
  border-radius: 12px;
  background: linear-gradient(180deg, hsl(var(--primary) / 0.05) 0%, hsl(var(--card)) 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 14px;
  color: hsl(var(--muted-foreground));
  transition: all 0.3s ease;
}

.config-empty-state:hover {
  border-color: hsl(var(--primary) / 0.4);
  background: linear-gradient(180deg, hsl(var(--primary) / 0.08) 0%, hsl(var(--card)) 100%);
}

.config-empty-icon {
  font-size: 36px;
  color: hsl(var(--primary));
}

@media (max-width: 1023px) {
  .config-empty-state {
    min-height: 200px;
  }

  .config-add-tip {
    margin-bottom: 16px;
  }
}
</style>