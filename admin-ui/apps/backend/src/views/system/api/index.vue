<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

import {
  changeApiStatus,
  deleteApi,
  realDeleteApi,
  recoveryApi,
} from '#/api/system/api';
import { getApiGroupList } from '#/api/system/api-group';
import type { ApiGroupApi } from '#/api/system/api-group';
import type { OptionItem, IdType } from '#/types/common';
import type { DictOption } from '#/composables/crud/use-dict-options';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import ApiActionToolbar from './components/api-action-toolbar.vue';
import ApiModal from './components/api-modal.vue';
import ApiSearchForm from './components/api-search-form.vue';
import ApiTable from './components/api-table.vue';
import type { ApiListItem, ApiTableColumn } from './model';
import {
  createApiTableColumnOptions,
  createApiTableColumns,
} from './schemas';
import { useApiCrud } from './use-api-crud';

defineOptions({ name: 'SystemApi' });

type ApiModalInstance = {
  open: (data?: Partial<ApiListItem>) => void;
};

const apiModalRef = ref<ApiModalInstance>();

const columns: ApiTableColumn[] = createApiTableColumns();
const columnOptions = createApiTableColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);

const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value.filter((item) => item !== 'row-select');
  },
});

const groupOptions = ref<OptionItem<IdType>[]>([]);
const requestModeOptions = ref<DictOption[]>([]);
const statusOptions = ref<DictOption[]>([]);

const authModeOptions: DictOption[] = [
  { label: $t('system.api.simpleMode'), value: 1 },
  { label: $t('system.api.complexMode'), value: 2 },
];

const fallbackRequestModes: DictOption[] = [
  { label: 'GET', value: 'GET' },
  { label: 'POST', value: 'POST' },
  { label: 'PUT', value: 'PUT' },
  { label: 'DELETE', value: 'DELETE' },
];

const fallbackStatusOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

const groupMap = computed(() => {
  const map = new Map<string, string>();
  groupOptions.value.forEach((item: OptionItem<IdType>) => {
    map.set(String(item.value), item.label ?? '');
  });
  return map;
});

const requestModeMap = computed(() => {
  const map = new Map<string, string>();
  requestModeOptions.value.forEach((item: DictOption) => {
    map.set(String(item.value ?? item.label), item.label ?? '');
  });
  return map;
});

const authModeMap = computed(() => {
  const map = new Map<number, string>();
  authModeOptions.forEach((item: DictOption) => {
    map.set(Number(item.value), item.label ?? '');
  });
  return map;
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
} = useApiCrud();

const { getDictOptions } = useDictOptions();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key)).filter((id) => !Number.isNaN(id));
}

async function fetchFilterOptions() {
  try {
    const [groups, requestModes, statuses] = await Promise.all([
      getApiGroupList().catch(() => []),
      getDictOptions('request_mode'),
      getDictOptions('data_status'),
    ]);
    const groupList = Array.isArray(groups) ? (groups as ApiGroupApi.ListItem[]) : [];
    groupOptions.value = groupList.map((item) => ({
      label: item.name,
      value: item.id,
    }));
    requestModeOptions.value =
      requestModes && requestModes.length > 0 ? requestModes : fallbackRequestModes;
    statusOptions.value =
      statuses && statuses.length > 0 ? statuses : fallbackStatusOptions;
  } catch (error) {
    logger.error(error);
    message.error($t('common.filterLoadFailed'));
    groupOptions.value = [];
    requestModeOptions.value = fallbackRequestModes;
    statusOptions.value = fallbackStatusOptions;
  }
}

function handleAdd() {
  apiModalRef.value?.open();
}

function handleEdit(row: ApiListItem) {
  apiModalRef.value?.open(row);
}

async function handleDelete(row: ApiListItem) {
  try {
    await (isRecycleBin.value ? realDeleteApi([row.id]) : deleteApi([row.id]));
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
  if (ids.length === 0) {
    message.warning($t('common.invalidDataFormat'));
    return;
  }
  try {
    await (isRecycleBin.value ? realDeleteApi(ids) : deleteApi(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: ApiListItem) {
  try {
    await recoveryApi([row.id]);
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
  if (ids.length === 0) {
    message.warning($t('common.invalidDataFormat'));
    return;
  }
  try {
    await recoveryApi(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: ApiListItem, checked: boolean) {
  const status = checked ? 1 : 2;
  try {
    await changeApiStatus({ id: row.id, status });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
}

function handleStatusSwitchChange(row: ApiListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

function resolveGroupLabel(id?: IdType) {
  if (!id) return '-';
  return groupMap.value.get(String(id)) || '-';
}

function resolveRequestModeLabel(value?: number | string) {
  if (value === undefined || value === null) return '-';
  return requestModeMap.value.get(String(value)) || String(value);
}

function resolveAuthModeLabel(value?: number | string) {
  const numberValue = Number(value);
  if (Number.isNaN(numberValue)) return '-';
  return authModeMap.value.get(numberValue) || '-';
}

onMounted(() => {
  void fetchFilterOptions();
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <ApiSearchForm
        :form-data="searchForm"
        :group-options="groupOptions"
        :request-mode-options="requestModeOptions"
        :status-options="statusOptions"
        @search="handleSearch"
        @reset="handleReset"
        @update:form-data="(val) => Object.assign(searchForm, val)"
      />

      <div class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
        <ApiActionToolbar
          :is-recycle-bin="isRecycleBin"
          :visible-columns="visibleColumns"
          :column-options="columnOptions"
          @add="handleAdd"
          @batch-delete="handleBatchDelete"
          @batch-recovery="handleBatchRecovery"
          @refresh="fetchTableData"
          @toggle-recycle="toggleRecycleBin"
          @update:visible-columns="(val) => (visibleColumns = val)"
        />

        <div class="min-h-0 flex-1 overflow-auto">
          <ApiTable
          :columns="columns"
          :data="tableData"
          :loading="loading"
          :pagination="pagination"
          :selected-row-keys="selectedRowKeys"
          :display-columns="displayColumns"
          :is-recycle-bin="isRecycleBin"
          :resolve-group-label="resolveGroupLabel"
          :resolve-request-mode-label="resolveRequestModeLabel"
          :resolve-auth-mode-label="resolveAuthModeLabel"
          @page-change="handlePageChange"
          @select-change="handleTableSelectChange"
          @edit="handleEdit"
          @delete="handleDelete"
          @status-change="handleStatusSwitchChange"
          @recovery="handleRecovery"
          @update:display-columns="(val) => (displayColumns = val)"
        />
        </div>
      </div>
    </div>

    <ApiModal ref="apiModalRef" @success="fetchTableData" />
  </Page>
</template>