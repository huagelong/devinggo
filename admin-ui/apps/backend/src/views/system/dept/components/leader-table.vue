<script lang="ts" setup>
import { computed } from 'vue';

import { $t } from '@vben/locales';

import { Button, Popconfirm, Table } from 'tdesign-vue-next';
import type { PrimaryTableCol, TableRowData } from 'tdesign-vue-next/es/table/type';

interface LeaderItem {
  id: number;
  username: string;
  nickname?: string;
  phone?: string;
  email?: string;
  status?: number | string;
  leader_add_time?: string;
}

interface Props {
  data: LeaderItem[];
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
}

interface Emits {
  (e: 'page-change', pageInfo: { current: number; pageSize: number }): void;
  (e: 'select-change', keys: Array<number | string>): void;
  (e: 'delete', id: number): void;
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

const columns = [
  { colKey: 'row-select', type: 'multiple', width: 50 },
  { colKey: 'username', title: $t('system.dept.username'), minWidth: 140 },
  { colKey: 'nickname', title: $t('system.dept.nickname'), minWidth: 140 },
  { colKey: 'phone', title: $t('system.dept.phone'), minWidth: 140 },
  { colKey: 'email', title: $t('system.dept.email'), minWidth: 180 },
  { colKey: 'status', title: $t('common.status'), width: 100, align: 'center' as const },
  {
    colKey: 'leader_add_time',
    title: $t('system.dept.leaderAddTime'),
    width: 180,
    align: 'center' as const,
  },
  { colKey: 'action', title: $t('common.action'), width: 120, align: 'center' as const },
] satisfies PrimaryTableCol<TableRowData>[];

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  emit('page-change', pageInfo);
}

function handleSelectChange(keys: Array<number | string>) {
  emit('select-change', keys);
}

function handleDelete(id: number) {
  emit('delete', id);
}
</script>

<template>
  <div class="rounded-md border border-gray-100 bg-white p-4">
    <Table
      :columns="columns"
      :data="data"
      :loading="loading"
      :pagination="pagination"
      :selected-row-keys="selectedRowKeys"
      :scroll="scrollConfig"
      row-key="id"
      hover
      stripe
      size="small"
      @page-change="handlePageChange"
      @select-change="handleSelectChange"
    >
      <template #status="{ row }">
        <span>{{ Number(row.status) === 1 ? $t('common.statusEnabled') : $t('common.statusDisabled') }}</span>
      </template>

      <template #action="{ row }">
        <Popconfirm
          :content="$t('system.dept.confirmDeleteLeader')"
          @confirm="handleDelete(Number(row.id))"
        >
          <Button size="small" theme="danger" variant="outline">
            {{ $t('common.delete') }}
          </Button>
        </Popconfirm>
      </template>
    </Table>
  </div>
</template>