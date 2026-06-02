<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DictApi } from '#/api/system/dict';

import { reactive, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';
import { AddIcon, DeleteIcon, SearchIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  Input,
  InputNumber,
  Popconfirm,
  Select,
  Space,
  Switch,
  Table,
} from 'tdesign-vue-next';

import {
  changeDictDataStatus,
  deleteDictData,
  getDictDataPageList,
  getRecycleDictDataList,
  realDeleteDictData,
  recoveryDictData,
  updateDictDataNumber,
} from '#/api/system/dict';

import DictDataFormModal from './dict-data-form-modal.vue';
import type { DictDataListItem } from '../model';
import { createDictDataSearchForm, createDictDataTableColumns } from '../schemas';

type DictDataFormInstance = {
  open: (options: {
    data?: DictApi.DictDataSubmitPayload;
    typeInfo: { id: number; code: string; name?: string };
  }) => void;
};

const dictDataModalRef = ref<DictDataFormInstance>();

const searchForm = reactive(createDictDataSearchForm());
const tableData = ref<DictDataListItem[]>([]);
const loading = ref(false);
const isRecycleBin = ref(false);
const selectedRowKeys = ref<Array<number | string>>([]);
const currentDict = ref<{ id: number; code: string; name?: string }>();
const pagination = reactive({
  current: 1,
  pageSize: 20,
  total: 0,
  showJumper: true,
  showPageSize: true,
});

const columns = createDictDataTableColumns();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

function buildParams() {
  const params: DictApi.DictDataQuery = {
    page: pagination.current,
    pageSize: pagination.pageSize,
  };
  if (!currentDict.value) return params;
  params.type_id = currentDict.value.id;
  params.code = currentDict.value.code;
  if (searchForm.label) params.label = searchForm.label;
  if (searchForm.value) params.value = searchForm.value;
  if (searchForm.status !== undefined) params.status = searchForm.status;
  if (searchForm.created_at?.length === 2 && searchForm.created_at[0]) {
    params.created_at = searchForm.created_at;
  }
  return params;
}

async function fetchTableData() {
  if (!currentDict.value) return;
  loading.value = true;
  try {
    const params = buildParams();
    const response = isRecycleBin.value
      ? await getRecycleDictDataList(params)
      : await getDictDataPageList(params);
    tableData.value = response.items ?? [];
    pagination.total = Number(response?.pageInfo?.total || response?.total || 0);
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.dictDataLoadFailed'));
  } finally {
    loading.value = false;
  }
}

function handleSearch() {
  pagination.current = 1;
  void fetchTableData();
}

function handleReset() {
  Object.assign(searchForm, createDictDataSearchForm(), {
    type_id: currentDict.value?.id,
    code: currentDict.value?.code,
  });
  pagination.current = 1;
  void fetchTableData();
}

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  pagination.current = pageInfo.current;
  pagination.pageSize = pageInfo.pageSize;
  void fetchTableData();
}

function handleSelectChange(keys: Array<number | string>) {
  selectedRowKeys.value = keys;
}

function clearSelectedRowKeys() {
  selectedRowKeys.value = [];
}

function handleToggleRecycleBin(next?: boolean) {
  isRecycleBin.value = typeof next === 'boolean' ? next : !isRecycleBin.value;
  pagination.current = 1;
  clearSelectedRowKeys();
  void fetchTableData();
}

function handleAdd() {
  if (!currentDict.value) return;
  dictDataModalRef.value?.open({
    typeInfo: currentDict.value,
  });
}

function handleEdit(row: DictDataListItem) {
  if (!currentDict.value) return;
  dictDataModalRef.value?.open({
    data: {
      ...row,
      sort: Number(row.sort ?? 1),
      status: Number(row.status ?? 1),
      type_id: currentDict.value.id,
      code: currentDict.value.code,
    },
    typeInfo: currentDict.value,
  });
}

async function handleDelete(row: DictDataListItem) {
  try {
    await (isRecycleBin.value ? realDeleteDictData([row.id]) : deleteDictData([row.id]));
    MessagePlugin.success($t('common.operationSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.deleteFailed'));
  }
}

async function handleBatchDelete() {
  if (selectedRowKeys.value.length === 0) {
    MessagePlugin.warning($t('common.selectDataFirst'));
    return;
  }
  const ids = toIds(selectedRowKeys.value);
  try {
    await (isRecycleBin.value ? realDeleteDictData(ids) : deleteDictData(ids));
    MessagePlugin.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: DictDataListItem) {
  try {
    await recoveryDictData([row.id]);
    MessagePlugin.success($t('common.recoverySuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.recoveryFailed'));
  }
}

async function handleBatchRecovery() {
  if (selectedRowKeys.value.length === 0) {
    MessagePlugin.warning($t('common.selectDataFirst'));
    return;
  }
  const ids = toIds(selectedRowKeys.value);
  try {
    await recoveryDictData(ids);
    MessagePlugin.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: DictDataListItem, checked: boolean) {
  try {
    await changeDictDataStatus({ id: row.id, status: checked ? 1 : 2 });
    MessagePlugin.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.statusUpdateFailed'));
  }
}

function handleStatusSwitchChange(row: DictDataListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

async function handleSortChange(value: number | string, row: DictDataListItem) {
  const numberValue = Number(value);
  if (Number.isNaN(numberValue)) return;
  try {
    await updateDictDataNumber({
      id: Number(row.id),
      numberName: 'sort',
      numberValue,
    });
    MessagePlugin.success($t('common.sortUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.sortUpdateFailed'));
  }
}

const [PanelModal, panelModalApi] = useVbenModal({
  footer: false,
  class: 'w-[1200px]',
  onClosed: () => {
    currentDict.value = undefined;
    tableData.value = [];
    clearSelectedRowKeys();
  },
});

async function open(row: DictApi.DictTypeItem) {
  currentDict.value = { id: row.id, code: row.code, name: row.name };
  Object.assign(searchForm, createDictDataSearchForm(), {
    type_id: row.id,
    code: row.code,
  });
  pagination.current = 1;
  pagination.total = 0;
  selectedRowKeys.value = [];
  isRecycleBin.value = false;

  panelModalApi.setState({
    title: $t('system.dict.maintainDictData', [row.name]),
  });
  panelModalApi.open();
  await fetchTableData();
}

defineExpose({
  open,
});
</script>

<template>
  <PanelModal>
    <div class="flex flex-col gap-4">
      <div class="rounded-md bg-gray-50 p-3 text-sm text-gray-700">
        {{ $t('system.dict.currentDict') }}：{{ currentDict ? currentDict.name : '-' }}（{{
          currentDict ? currentDict.code : '-'
        }}）
      </div>

      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="90px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.dict.label')" name="label">
              <Input
                v-model="searchForm.label"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.dict.value')" name="value">
              <Input
                v-model="searchForm.value"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('common.status')" name="status">
              <Select
                v-model="searchForm.status"
                :options="[
                  { label: $t('common.statusEnabled'), value: 1 },
                  { label: $t('common.statusDisabled'), value: 2 },
                ]"
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

      <div class="rounded-md bg-white p-4">
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

          <Button variant="outline" @click="handleToggleRecycleBin()">
            {{ isRecycleBin ? $t('common.backToList') : $t('common.viewRecycleBin') }}
          </Button>
        </div>

        <Table
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
          <template #sort="{ row }">
            <InputNumber
              :value="row.sort"
              :min="0"
              :max="1000"
              size="small"
              :disabled="isRecycleBin"
              @change="(value: number | string) => handleSortChange(value, row)"
            />
          </template>
          <template #status="{ row }">
            <Switch
              :disabled="isRecycleBin"
              :value="Number(row.status) === 1"
              @change="(value: unknown) => handleStatusSwitchChange(row, value)"
            />
          </template>
          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <template v-if="!isRecycleBin">
                <Button size="small" theme="primary" variant="outline" @click="handleEdit(row)">
                  {{ $t('common.edit') }}
                </Button>
                <Popconfirm :content="$t('system.dict.confirmDeleteData')" @confirm="handleDelete(row)">
                  <Button size="small" theme="danger" variant="outline">{{ $t('common.delete') }}</Button>
                </Popconfirm>
              </template>
              <template v-else>
                <Popconfirm :content="$t('system.dict.confirmRecovery')" @confirm="handleRecovery(row)">
                  <Button size="small" theme="primary" variant="outline">{{ $t('common.recovery') }}</Button>
                </Popconfirm>
                <Popconfirm :content="$t('system.dict.confirmPermanentDelete')" @confirm="handleDelete(row)">
                  <Button size="small" theme="danger" variant="outline">{{ $t('common.permanentDelete') }}</Button>
                </Popconfirm>
              </template>
            </div>
          </template>
        </Table>
      </div>

      <DictDataFormModal ref="dictDataModalRef" @success="fetchTableData" />
    </div>
  </PanelModal>
</template>
