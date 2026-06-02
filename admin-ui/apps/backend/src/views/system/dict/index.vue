<script lang="ts" setup>
import type { DictApi } from '#/api/system/dict';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { computed, onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';
import {
  changeDictTypeStatus,
  deleteDictType,
  realDeleteDictType,
  recoveryDictType,
} from '#/api/system/dict';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import {
  AddIcon,
  DeleteIcon,
  EditIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  Input,
  Popconfirm,
  Select,
  Space,
  Switch,
  Table,
} from 'tdesign-vue-next';

import DictDataPanel from './components/dict-data-panel.vue';
import DictTypeModal from './components/dict-type-modal.vue';
import type { DictTableColumn, DictTypeListItem } from './model';
import {
  createDictTypeColumnOptions,
  createDictTypeTableColumns,
} from './schemas';
import { useDictTypeCrud } from './use-dict-type-crud';

defineOptions({ name: 'SystemDict' });

type DictTypeModalInstance = {
  open: (data?: DictApi.DictTypeSubmitPayload) => void;
};

type DictDataPanelInstance = {
  open: (row: DictApi.DictTypeItem) => void;
};

const dictTypeModalRef = ref<DictTypeModalInstance>();
const dictDataPanelRef = ref<DictDataPanelInstance>();
const statusOptions = ref<DictOption[]>([]);

const columns: DictTableColumn[] = createDictTypeTableColumns();
const columnOptions = createDictTypeColumnOptions(columns);
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
} = useDictTypeCrud();

const { getDictOptions } = useDictOptions();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

async function fetchStatusOptions() {
  statusOptions.value =
    (await getDictOptions('data_status')) || [
      { label: $t('common.statusEnabled'), value: 1 },
      { label: $t('common.statusDisabled'), value: 2 },
    ];
}

function handleAdd() {
  dictTypeModalRef.value?.open();
}

function handleEdit(row: DictTypeListItem) {
  dictTypeModalRef.value?.open({
    ...row,
    status: Number(row.status ?? 1),
  });
}

function handleOpenData(row: DictTypeListItem) {
  dictDataPanelRef.value?.open({
    code: row.code,
    id: row.id,
    name: row.name,
  } as DictApi.DictTypeItem);
}

async function handleDelete(row: DictTypeListItem) {
  try {
    await (isRecycleBin.value ? realDeleteDictType([row.id]) : deleteDictType([row.id]));
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
  try {
    await (isRecycleBin.value ? realDeleteDictType(ids) : deleteDictType(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: DictTypeListItem) {
  try {
    await recoveryDictType([row.id]);
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
    await recoveryDictType(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: DictTypeListItem, checked: boolean) {
  try {
    await changeDictTypeStatus({ id: row.id, status: checked ? 1 : 2 });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

function handleStatusSwitchChange(row: DictTypeListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

function handleSuccess() {
  void fetchTableData();
}

onMounted(() => {
  void fetchStatusOptions();
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="90px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.dict.name')" name="name">
              <Input
                v-model="searchForm.name"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.dict.code')" name="code">
              <Input
                v-model="searchForm.code"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('common.status')" name="status">
              <Select
                v-model="searchForm.status"
                :options="statusOptions"
                :placeholder="$t('ui.placeholder.select')"
                clearable
                class="w-full"
              />
            </FormItem>
            <FormItem :label="$t('common.createTime')" name="created_at">
              <DateRangePicker
                v-model="searchForm.created_at"
                :placeholder="[$t('common.startTime'), $t('common.endTime')]"
                clearable
                class="w-full"
              />
            </FormItem>
          </div>
          <div class="flex justify-end gap-2 pt-2">
            <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
            <Button theme="primary" @click="handleSearch">
              <template #icon><SearchIcon /></template>
              {{ $t('common.query') }}
            </Button>
          </div>
        </Form>
      </div>

      <div class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
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

          <CrudToolbar
            v-model="visibleColumns"
            :column-options="columnOptions"
            :is-recycle-bin="isRecycleBin"
            @refresh="fetchTableData"
            @toggle-recycle="toggleRecycleBin"
          />
        </div>

        <Table
          v-model:display-columns="displayColumns"
          :columns="columns"
          :data="tableData"
          :loading="loading"
          :pagination="pagination"
          :selected-row-keys="selectedRowKeys"
          row-key="id"
          hover
          stripe
          table-layout="fixed"
          @page-change="handlePageChange"
          @select-change="handleSelectChange"
        >
          <template #code="{ row }">
            <Button v-if="!isRecycleBin" theme="primary" variant="text" @click="handleOpenData(row)">
              {{ row?.code }}
            </Button>
            <span v-else>{{ row?.code }}</span>
          </template>

          <template #status="{ row }">
            <Switch
              :disabled="isRecycleBin"
              :value="Number(row?.status) === 1"
              @change="(value: unknown) => handleStatusSwitchChange(row, value)"
            />
          </template>

          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <template v-if="!isRecycleBin">
                <Button
                  size="small"
                  theme="default"
                  variant="outline"
                  @click="handleOpenData(row)"
                >
                  {{ $t('system.dict.dictData') }}
                </Button>
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
                  :content="$t('system.dict.confirmDelete')"
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
                  :content="$t('system.dict.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('system.dict.confirmPermanentDelete')"
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
      </div>
    </div>

    <DictTypeModal ref="dictTypeModalRef" @success="handleSuccess" />
    <DictDataPanel ref="dictDataPanelRef" />
  </Page>
</template>
