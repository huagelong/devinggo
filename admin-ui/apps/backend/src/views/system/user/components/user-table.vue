<script lang="ts" setup>
import { computed } from 'vue';

import { $t } from '@vben/locales';

import {
  DeleteIcon,
  EditIcon,
  MoreIcon,
  RefreshIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  Dropdown,
  Popconfirm,
  Switch,
  Table,
} from 'tdesign-vue-next';

import type { UserActionDropdownItem, UserListItem, UserTableColumn } from '../model';
import { userActionDropdownOptions } from '../schemas';

interface Props {
  columns: UserTableColumn[];
  data: UserListItem[];
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
  isSuperAdmin: (row: UserListItem) => boolean;
}

interface Emits {
  (e: 'page-change', pageInfo: { current: number; pageSize: number }): void;
  (e: 'select-change', keys: Array<number | string>): void;
  (e: 'edit', row: UserListItem): void;
  (e: 'delete', row: UserListItem): void;
  (e: 'status-change', row: UserListItem, value: boolean): void;
  (e: 'action-dropdown-click', item: UserActionDropdownItem, row: UserListItem): void;
  (e: 'clear-cache', row: UserListItem): void;
  (e: 'recovery', row: UserListItem): void;
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

function handleEdit(row: UserListItem) {
  emit('edit', row);
}

function handleDelete(row: UserListItem) {
  emit('delete', row);
}

function handleStatusChange(row: UserListItem, value: unknown) {
  emit('status-change', row, Boolean(value));
}

function handleActionDropdownClick(item: unknown, row: UserListItem) {
  emit('action-dropdown-click', item as UserActionDropdownItem, row);
}

function handleClearCache(row: UserListItem) {
  emit('clear-cache', row);
}

function handleRecovery(row: UserListItem) {
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
    <template #avatar="{ row }">
      <img
        v-if="row?.avatar"
        :src="row?.avatar"
        class="mx-auto h-8 w-8 rounded-full object-cover"
        alt="avatar"
      />
      <span v-else class="text-gray-400">-</span>
    </template>

    <template #status="{ row }">
      <Switch
        :disabled="isRecycleBin || isSuperAdmin(row)"
        :value="row?.status === 1"
        @change="(value: unknown) => handleStatusChange(row, value)"
      />
    </template>

    <template #action="{ row }">
      <div class="flex items-center justify-center gap-1">
        <template v-if="!isRecycleBin">
          <template v-if="isSuperAdmin(row)">
            <Button
              size="small"
              theme="default"
              variant="outline"
              @click="handleClearCache(row)"
            >
              <template #icon><RefreshIcon /></template>
              {{ $t('system.user.clearCache') }}
            </Button>
          </template>
          <template v-else>
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
              :content="$t('system.user.confirmDelete')"
              @confirm="handleDelete(row)"
            >
              <Button size="small" theme="danger" variant="outline">
                <template #icon><DeleteIcon /></template>
                {{ $t('common.delete') }}
              </Button>
            </Popconfirm>
            <Dropdown
              :options="userActionDropdownOptions"
              trigger="click"
              @click="(item: unknown) => handleActionDropdownClick(item, row)"
            >
              <Button size="small" theme="default" variant="outline">
                <template #icon><MoreIcon /></template>
                {{ $t('common.more') }}
              </Button>
            </Dropdown>
          </template>
        </template>

        <template v-else>
          <Popconfirm
            :content="$t('system.user.confirmRecovery')"
            @confirm="handleRecovery(row)"
          >
            <Button size="small" theme="primary" variant="outline">
              {{ $t('common.recovery') }}
            </Button>
          </Popconfirm>
          <Popconfirm
            :content="$t('system.user.confirmPermanentDelete')"
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