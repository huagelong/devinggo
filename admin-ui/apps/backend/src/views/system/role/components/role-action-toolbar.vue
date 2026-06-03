<script lang="ts" setup>
import { $t } from '@vben/locales';

import {
  AddIcon,
  DeleteIcon,
  FullscreenExitIcon,
  FullscreenIcon,
} from 'tdesign-icons-vue-next';
import { Button, Space, Tooltip } from 'tdesign-vue-next';

import CrudToolbar from '#/components/crud/crud-toolbar.vue';

interface ColumnOption {
  label: string;
  value: string;
}

interface Props {
  isRecycleBin: boolean;
  isFullscreen: boolean;
  visibleColumns: string[];
  columnOptions: ColumnOption[];
}

interface Emits {
  (e: 'add'): void;
  (e: 'batch-delete'): void;
  (e: 'batch-recovery'): void;
  (e: 'toggle-fullscreen'): void;
  (e: 'refresh'): void;
  (e: 'toggle-recycle'): void;
  (e: 'update:visibleColumns', value: string[]): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();

function handleAdd() {
  emit('add');
}

function handleBatchDelete() {
  emit('batch-delete');
}

function handleBatchRecovery() {
  emit('batch-recovery');
}

function handleToggleFullscreen() {
  emit('toggle-fullscreen');
}

function handleRefresh() {
  emit('refresh');
}

function handleToggleRecycle() {
  emit('toggle-recycle');
}

function handleUpdateVisibleColumns(value: string[]) {
  emit('update:visibleColumns', value);
}
</script>

<template>
  <div class="mb-3 flex items-center justify-between">
    <Space>
      <template v-if="!isRecycleBin">
        <Button theme="primary" @click="handleAdd">
          <template #icon><AddIcon /></template>
          {{ $t('common.create') }}
        </Button>
        <Button theme="danger" variant="outline" @click="handleBatchDelete">
          <template #icon><DeleteIcon /></template>
          {{ $t('common.delete') }}
        </Button>
      </template>
      <template v-else>
        <Button theme="success" @click="handleBatchRecovery">{{ $t('common.recovery') }}</Button>
        <Button theme="danger" @click="handleBatchDelete">{{ $t('common.permanentDelete') }}</Button>
      </template>
    </Space>

    <div class="flex items-center gap-2">
      <Tooltip :content="isFullscreen ? $t('common.exitFullscreen') : $t('common.fullscreen')">
        <Button shape="square" variant="outline" @click="handleToggleFullscreen">
          <template #icon>
            <FullscreenExitIcon v-if="isFullscreen" />
            <FullscreenIcon v-else />
          </template>
        </Button>
      </Tooltip>

      <CrudToolbar
        :model-value="visibleColumns"
        :column-options="columnOptions"
        :is-recycle-bin="isRecycleBin"
        @update:model-value="handleUpdateVisibleColumns"
        @refresh="handleRefresh"
        @toggle-recycle="handleToggleRecycle"
      />
    </div>
  </div>
</template>