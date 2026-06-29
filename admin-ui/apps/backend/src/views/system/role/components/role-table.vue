<script lang="ts" setup>
import { computed } from 'vue';

import { $t } from '@vben/locales';

import { DeleteIcon, EditIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  InputNumber,
  Popconfirm,
  Switch,
  Table,
} from 'tdesign-vue-next';

import type { RoleListItem, RoleTableColumn } from '../model';
import { getRoleDataScopeLabel } from '../schemas';

interface Props {
  columns: RoleTableColumn[];
  data: RoleListItem[];
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
}

interface Emits {
  (e: 'page-change', pageInfo: { current: number; pageSize: number }): void;
  (e: 'select-change', keys: Array<number | string>): void;
  (e: 'edit', row: RoleListItem): void;
  (e: 'delete', row: RoleListItem): void;
  (e: 'status-change', row: RoleListItem, value: boolean): void;
  (e: 'sort-change', value: number | string, row: RoleListItem): void;
  (e: 'menu-permission', row: RoleListItem): void;
  (e: 'data-permission', row: RoleListItem): void;
  (e: 'recovery', row: RoleListItem): void;
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

function handleEdit(row: RoleListItem) {
  emit('edit', row);
}

function handleDelete(row: RoleListItem) {
  emit('delete', row);
}

function handleStatusChange(row: RoleListItem, value: unknown) {
  emit('status-change', row, Boolean(value));
}

function handleSortChange(value: number | string, row: RoleListItem) {
  emit('sort-change', value, row);
}

function handleMenuPermission(row: RoleListItem) {
  emit('menu-permission', row);
}

function handleDataPermission(row: RoleListItem) {
  emit('data-permission', row);
}

function handleRecovery(row: RoleListItem) {
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
    row-key="id"
    hover
    stripe
    @page-change="handlePageChange"
    @select-change="handleSelectChange"
    @update:display-columns="handleUpdateDisplayColumns"
  >
    <template #data_scope="{ row }">
      <span>{{ getRoleDataScopeLabel(row?.data_scope) }}</span>
    </template>

    <template #sort="{ row }">
      <InputNumber
        :value="row?.sort"
        :min="0"
        :max="1000"
        size="small"
        @change="(value: number | string) => handleSortChange(value, row)"
      />
    </template>

    <template #status="{ row }">
      <Switch
        :disabled="isRecycleBin || row?.code === 'superAdmin'"
        :value="Number(row?.status) === 1"
        @change="(value: unknown) => handleStatusChange(row, value)"
      />
    </template>

    <template #action="{ row }">
      <div class="flex items-center justify-center gap-1">
        <template v-if="!isRecycleBin">
          <template v-if="row?.code !== 'superAdmin'">
            <Button
              size="small"
              theme="primary"
              variant="outline"
              @click="handleEdit(row)"
            >
              <template #icon><EditIcon /></template>
              {{ $t('common.edit') }}
            </Button>
            <Button
              size="small"
              theme="default"
              variant="outline"
              @click="handleMenuPermission(row)"
            >
              {{ $t('system.role.menuPermission') }}
            </Button>
            <Button
              size="small"
              theme="default"
              variant="outline"
              @click="handleDataPermission(row)"
            >
              {{ $t('system.role.dataPermission') }}
            </Button>
            <Popconfirm
              :content="$t('system.role.confirmDelete')"
              @confirm="handleDelete(row)"
            >
              <Button size="small" theme="danger" variant="outline">
                <template #icon><DeleteIcon /></template>
                {{ $t('common.delete') }}
              </Button>
            </Popconfirm>
          </template>
        </template>

        <template v-else>
          <Popconfirm
            :content="$t('system.role.confirmRecovery')"
            @confirm="handleRecovery(row)"
          >
            <Button size="small" theme="primary" variant="outline">
              {{ $t('common.recovery') }}
            </Button>
          </Popconfirm>
          <Popconfirm
            :content="$t('system.role.confirmPermanentDelete')"
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