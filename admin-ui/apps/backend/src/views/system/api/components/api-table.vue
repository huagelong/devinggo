<script lang="ts" setup>
import { computed } from 'vue';

import { $t } from '@vben/locales';

import { DeleteIcon, EditIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  Popconfirm,
  Switch,
  Table,
} from 'tdesign-vue-next';

import type { ApiListItem, ApiTableColumn } from '../model';

interface Props {
  columns: ApiTableColumn[];
  data: ApiListItem[];
  loading: boolean;
  pagination: {
    current: number;
    pageSize: number;
    total: number;
    pageSizeOptions: number[];
    showJumper: boolean;
    showPageSize: boolean;
  };
  selectedRowKeys: Array<number | string>;
  displayColumns: string[];
  isRecycleBin: boolean;
  resolveGroupLabel: (id?: number | string) => string;
  resolveRequestModeLabel: (value?: number | string) => string;
  resolveAuthModeLabel: (value?: number | string) => string;
}

interface Emits {
  (e: 'page-change', pageInfo: { current: number; pageSize: number }): void;
  (e: 'select-change', keys: Array<number | string>): void;
  (e: 'edit', row: ApiListItem): void;
  (e: 'delete', row: ApiListItem): void;
  (e: 'status-change', row: ApiListItem, value: boolean): void;
  (e: 'recovery', row: ApiListItem): void;
  (e: 'update:displayColumns', value: string[]): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

// 虚拟滚动配置：当数据量大于100条时启用虚拟滚动
const scrollConfig = computed(() => {
  if (props.data.length > 100) {
    return {
      type: 'virtual' as const,
      rowHeight: 48,
    };
  }
  return undefined;
});

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  emit('page-change', pageInfo);
}

function handleSelectChange(keys: Array<number | string>) {
  emit('select-change', keys);
}

function handleEdit(row: ApiListItem) {
  emit('edit', row);
}

function handleDelete(row: ApiListItem) {
  emit('delete', row);
}

function handleStatusChange(row: ApiListItem, value: unknown) {
  emit('status-change', row, Boolean(value));
}

function handleRecovery(row: ApiListItem) {
  emit('recovery', row);
}

function handleUpdateDisplayColumns(value: string[]) {
  emit('update:displayColumns', value);
}
</script>

<template>
  <Table
    :display-columns="displayColumns"
    :columns="columns"
    :data="data"
    :loading="loading"
    :pagination="pagination"
    :selected-row-keys="selectedRowKeys"
    :scroll="scrollConfig"
    height="100%"
    row-key="id"
    hover
    stripe
    @page-change="handlePageChange"
    @select-change="handleSelectChange"
    @update:display-columns="handleUpdateDisplayColumns"
  >
    <template #group_name="{ row }">
      {{ resolveGroupLabel(row?.group_id) }}
    </template>

    <template #request_mode="{ row }">
      {{ resolveRequestModeLabel(row?.request_mode) }}
    </template>

    <template #auth_mode="{ row }">
      {{ resolveAuthModeLabel(row?.auth_mode) }}
    </template>

    <template #status="{ row }">
      <div class="flex items-center justify-center">
        <Switch
          :disabled="isRecycleBin"
          :value="row?.status === 1"
          @change="(value: unknown) => handleStatusChange(row, value)"
        />
      </div>
    </template>

    <template #remark="{ row }">
      <span class="block max-w-[240px] truncate">
        {{ row?.remark || '-' }}
      </span>
    </template>

    <template #action="{ row }">
      <div class="flex items-center justify-center gap-1">
        <template v-if="!isRecycleBin">
          <Button
            size="small"
            theme="primary"
            variant="outline"
            @click="handleEdit(row)"
          >
            <template #icon><EditIcon /></template>
            {{ $t('common.edit') }}
          </Button>
          <Popconfirm
            :content="$t('system.api.confirmDelete')"
            @confirm="handleDelete(row)"
          >
            <Button size="small" theme="danger" variant="outline">
              <template #icon><DeleteIcon /></template>
              {{ $t('common.delete') }}
            </Button>
          </Popconfirm>
        </template>

        <template v-else>
          <Popconfirm
            :content="$t('system.api.confirmRecovery')"
            @confirm="handleRecovery(row)"
          >
            <Button size="small" theme="primary" variant="outline">
              {{ $t('common.recovery') }}
            </Button>
          </Popconfirm>
          <Popconfirm
            :content="$t('system.api.confirmPermanentDelete')"
            @confirm="handleDelete(row)"
          >
            <Button size="small" theme="danger" variant="outline">
              {{ $t('common.permanentDelete') }}
            </Button>
          </Popconfirm>
        </template>
      </div>
    </template>
  </Table>
</template>