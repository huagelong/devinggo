<script lang="ts" setup>
import { $t } from '@vben/locales';

import { Table } from 'tdesign-vue-next';
import type { PrimaryTableCol, TableRowData } from 'tdesign-vue-next/es/table/type';

interface CandidateUser {
  id: number;
  username: string;
  nickname?: string;
  phone?: string;
  email?: string;
  dept_name?: string;
  role_name?: string;
  post_name?: string;
}

interface Props {
  data: CandidateUser[];
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
}

defineProps<Props>();
const emit = defineEmits<Emits>();

const columns = [
  { colKey: 'row-select', type: 'multiple', width: 50 },
  { colKey: 'username', title: $t('system.user.username'), minWidth: 120 },
  { colKey: 'nickname', title: $t('system.user.nickname'), minWidth: 120 },
  { colKey: 'phone', title: $t('system.user.phone'), minWidth: 140 },
  { colKey: 'email', title: $t('system.user.email'), minWidth: 180 },
  { colKey: 'dept_name', title: $t('system.user.dept'), minWidth: 160 },
  { colKey: 'role_name', title: $t('system.user.role'), minWidth: 140 },
  { colKey: 'post_name', title: $t('system.user.post'), minWidth: 140 },
] satisfies PrimaryTableCol<TableRowData>[];

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  emit('page-change', pageInfo);
}

function handleSelectChange(keys: Array<number | string>) {
  emit('select-change', keys);
}
</script>

<template>
  <div class="rounded-md border border-border bg-card p-3">
    <Table
      :columns="columns"
      :data="data"
      :loading="loading"
      :pagination="pagination"
      :selected-row-keys="selectedRowKeys"
      row-key="id"
      hover
      stripe
      size="small"
      @page-change="handlePageChange"
      @select-change="handleSelectChange"
    />
  </div>
</template>