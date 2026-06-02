<script lang="ts" setup>
import { $t } from '@vben/locales';

import { AddIcon, SettingIcon } from 'tdesign-icons-vue-next';
import { Button, Space } from 'tdesign-vue-next';

interface Props {
  activeGroupKey: number | undefined;
}

interface Emits {
  (e: 'add-group'): void;
  (e: 'manage'): void;
  (e: 'delete-group'): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();

function handleAddGroup() {
  emit('add-group');
}

function handleManage() {
  emit('manage');
}

function handleDeleteGroup() {
  emit('delete-group');
}
</script>

<template>
  <div class="config-toolbar flex items-center justify-between">
    <Space>
      <Button
        theme="primary"
        variant="outline"
        @click="handleAddGroup"
      >
        <template #icon><AddIcon /></template>
        {{ $t('system.config.addGroup') }}
      </Button>
      <Button
        theme="default"
        variant="outline"
        :disabled="!activeGroupKey"
        @click="handleManage"
      >
        <template #icon><SettingIcon /></template>
        {{ $t('system.config.manageTitle') }}
      </Button>
      <Button
        v-if="activeGroupKey && activeGroupKey > 2"
        theme="danger"
        variant="outline"
        @click="handleDeleteGroup"
      >
        {{ $t('system.config.deleteGroup') }}
      </Button>
    </Space>
  </div>
</template>

<style scoped>
.config-toolbar {
  margin-bottom: 14px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--td-component-stroke, #f0f1f2);
}

.config-toolbar :deep(.t-space) {
  flex-wrap: wrap;
  row-gap: 8px;
}
</style>