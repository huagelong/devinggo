<script lang="ts" setup>
import type { RoleApi } from '#/api/system/role';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { computed, onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';
import {
  changeRoleStatus,
  deleteRole,
  realDeleteRole,
  recoveryRole,
  updateRoleNumber,
} from '#/api/system/role';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import RoleActionToolbar from './components/role-action-toolbar.vue';
import RoleDataPermissionModal from './components/role-data-permission-modal.vue';
import RoleMenuPermissionModal from './components/role-menu-permission-modal.vue';
import RoleModal from './components/role-modal.vue';
import RoleSearchForm from './components/role-search-form.vue';
import RoleTable from './components/role-table.vue';
import type { RoleListItem, RoleTableColumn } from './model';
import {
  createRoleColumnOptions,
  createRoleTableColumns,
} from './schemas';
import { useRoleCrud } from './use-role-crud';

defineOptions({ name: 'SystemRole' });

type RoleModalInstance = {
  open: (data?: Partial<RoleApi.SubmitPayload>) => void;
};

type RolePermissionModalInstance = {
  open: (data: RoleApi.ListItem) => void;
};

const roleModalRef = ref<RoleModalInstance>();
const roleMenuPermissionModalRef = ref<RolePermissionModalInstance>();
const roleDataPermissionModalRef = ref<RolePermissionModalInstance>();
const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);
const statusOptions = ref<DictOption[]>([]);

const columns: RoleTableColumn[] = createRoleTableColumns();
const columnOptions = createRoleColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);

const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value.filter((item) => item !== 'row-select');
  },
});

const {
  clearSelectedRowKeys,
  fetchTableData,
  handlePageChange,
  handleReset,
  handleSearch,
  handleSelectChange,
  isRecycleBin,
  loading,
  pagination,
  searchForm,
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = useRoleCrud();

const { getDictOptions } = useDictOptions();

function handleFullscreenChange() {
  isFullscreen.value = !!document.fullscreenElement;
}

function toggleFullscreen() {
  if (document.fullscreenElement) {
    document.exitFullscreen();
    return;
  }
  tableContainerRef.value?.requestFullscreen();
}

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
    statusOptions.value =
      options.length > 0
        ? options
        : [
            { label: $t('common.statusEnabled'), value: 1 },
            { label: $t('common.statusDisabled'), value: 2 },
          ];
}

function handleAdd() {
  roleModalRef.value?.open();
}

function handleEdit(row: RoleListItem) {
  if (row.id === 1 || row.code === 'superAdmin') {
    message.error($t('common.superAdminRoleCannotEdit'));
    return;
  }
  roleModalRef.value?.open({
    ...row,
    data_scope: Number(row.data_scope ?? 1),
    status: Number(row.status ?? 1),
  });
}

function handleMenuPermission(row: RoleListItem) {
  roleMenuPermissionModalRef.value?.open(row);
}

function handleDataPermission(row: RoleListItem) {
  roleDataPermissionModalRef.value?.open(row);
}

async function handleDelete(row: RoleListItem) {
  if (row.id === 1 || row.code === 'superAdmin') {
    message.error($t('common.superAdminRoleCannotDelete'));
    return;
  }
  try {
    await (isRecycleBin.value ? realDeleteRole([row.id]) : deleteRole([row.id]));
    message.success($t('common.operationSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.deleteFailed'));
  }
}

async function handleBatchDelete() {
  if (selectedRowKeys.value.length === 0) {
    message.warning($t('common.selectDataFirst'));
    return;
  }

  const ids = toIds(selectedRowKeys.value);
  if (ids.includes(1)) {
    message.error($t('common.superAdminRoleCannotDelete'));
    return;
  }
  try {
    await (isRecycleBin.value ? realDeleteRole(ids) : deleteRole(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: RoleListItem) {
  try {
    await recoveryRole([row.id]);
    message.success($t('common.recoverySuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.recoveryFailed'));
  }
}

async function handleBatchRecovery() {
  if (selectedRowKeys.value.length === 0) {
    message.warning($t('common.selectDataFirst'));
    return;
  }

  const ids = toIds(selectedRowKeys.value);
  try {
    await recoveryRole(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: RoleListItem, checked: boolean) {
  if (row.code === 'superAdmin') {
    message.info($t('common.superAdminCannotDisable'));
    return;
  }
  const status = checked ? 1 : 2;
  try {
    await changeRoleStatus({ id: row.id, status });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

async function handleSortChange(value: number | string, row: RoleListItem) {
  const numberValue = Number(value);
  if (Number.isNaN(numberValue)) return;
  if (row.id === 1) {
    message.info($t('common.superAdminCannotModify'));
    return;
  }

  try {
    await updateRoleNumber({
      id: Number(row.id),
      numberName: 'sort',
      numberValue,
    });
    message.success($t('common.sortUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.sortUpdateFailed'));
  }
}

function handleSuccess() {
  void fetchTableData();
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
}

function handleStatusSwitchChange(row: RoleListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

onMounted(() => {
  void fetchStatusOptions();
  void fetchTableData();
  document.addEventListener('fullscreenchange', handleFullscreenChange);
});

onUnmounted(() => {
  document.removeEventListener('fullscreenchange', handleFullscreenChange);
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <RoleSearchForm
        :form-data="searchForm"
        :status-options="statusOptions"
        @search="handleSearch"
        @reset="handleReset"
        @update:form-data="(val) => Object.assign(searchForm, val)"
      />

      <div ref="tableContainerRef" class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
        <RoleActionToolbar
          :is-recycle-bin="isRecycleBin"
          :is-fullscreen="isFullscreen"
          :visible-columns="visibleColumns"
          :column-options="columnOptions"
          @add="handleAdd"
          @batch-delete="handleBatchDelete"
          @batch-recovery="handleBatchRecovery"
          @toggle-fullscreen="toggleFullscreen"
          @refresh="fetchTableData"
          @toggle-recycle="toggleRecycleBin"
          @update:visible-columns="(val) => (visibleColumns = val)"
        />

        <RoleTable
          :columns="columns"
          :data="tableData"
          :loading="loading"
          :pagination="pagination"
          :selected-row-keys="selectedRowKeys"
          :display-columns="displayColumns"
          :is-recycle-bin="isRecycleBin"
          @page-change="handlePageChange"
          @select-change="handleTableSelectChange"
          @edit="handleEdit"
          @delete="handleDelete"
          @status-change="handleStatusSwitchChange"
          @sort-change="handleSortChange"
          @menu-permission="handleMenuPermission"
          @data-permission="handleDataPermission"
          @recovery="handleRecovery"
          @update:display-columns="(val) => (displayColumns = val)"
        />
      </div>
    </div>

    <RoleModal ref="roleModalRef" @success="handleSuccess" />
    <RoleMenuPermissionModal
      ref="roleMenuPermissionModalRef"
      @success="handleSuccess"
    />
    <RoleDataPermissionModal
      ref="roleDataPermissionModalRef"
      @success="handleSuccess"
    />
  </Page>
</template>