<script lang="ts" setup>
import type { LogApi } from '#/api/system/log';

import { onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import {
  deleteApiLog,
  getApiLogPageList,
} from '#/api/system/log';
import { logger } from '#/utils/logger';

import {
  DeleteIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  Input,
  Popconfirm,
  Space,
  Table,
  Tag,
} from 'tdesign-vue-next';

defineOptions({ name: 'SystemApiLog' });

const loading = ref(false);
const tableData = ref<LogApi.ApiLogItem[]>([]);
const selectedRowKeys = ref<number[]>([]);
const searchForm = ref<{ api_name: string; access_name: string; ip: string; access_time: string[] }>({
  api_name: '',
  access_name: '',
  ip: '',
  access_time: [],
});

const pagination = ref({
  current: 1,
  pageSize: 20,
  pageSizeOptions: [10, 20, 50, 100],
  showJumper: true,
  showPageSize: true,
  total: 0,
});

const columns = [
  { colKey: 'row-select', type: 'multiple' as const, width: 50 },
  { colKey: 'api_name', title: $t('system.logs.apiName'), width: 180 },
  { colKey: 'access_name', title: $t('system.logs.accessName'), width: 140 },
  { colKey: 'response_code', title: $t('system.logs.responseCode'), width: 100 },
  { colKey: 'ip', title: $t('system.logs.ip'), width: 140 },
  { colKey: 'ip_location', title: $t('system.logs.ipLocation'), width: 120 },
  { colKey: 'access_time', title: $t('system.logs.accessTime'), width: 180 },
  { colKey: 'remark', title: $t('system.logs.remark'), ellipsis: true },
  { colKey: 'action', title: $t('common.action'), width: 100, fixed: 'right' as const },
];

async function fetchTableData() {
  loading.value = true;
  try {
    const params: LogApi.ApiLogQuery = {
      page: pagination.value.current,
      pageSize: pagination.value.pageSize,
    };
    if (searchForm.value.api_name) params.api_name = searchForm.value.api_name;
    if (searchForm.value.access_name) params.access_name = searchForm.value.access_name;
    if (searchForm.value.ip) params.ip = searchForm.value.ip;
    if (searchForm.value.access_time?.length === 2 && searchForm.value.access_time[0]) {
      params.access_time = searchForm.value.access_time;
    }
    const response = await getApiLogPageList(params);
    tableData.value = response.items || [];
    pagination.value.total = Number(response.pageInfo?.total || response.total || 0);
  } catch (error) {
    logger.error(error);
    message.error($t('common.logFetchFailed'));
  } finally {
    loading.value = false;
  }
}

async function handleDelete(row: LogApi.ApiLogItem) {
  try {
    await deleteApiLog([row.id]);
    message.success($t('common.deleteSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.deleteFailed'));
  }
}

async function handleBatchDelete() {
  if (selectedRowKeys.value.length === 0) {
    message.warning($t('common.selectLogFirst'));
    return;
  }
  try {
    await deleteApiLog(selectedRowKeys.value);
    message.success($t('common.batchDeleteSuccess'));
    selectedRowKeys.value = [];
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

function handleSearch() {
  pagination.value.current = 1;
  void fetchTableData();
}

function handleReset() {
  searchForm.value = { api_name: '', access_name: '', ip: '', access_time: [] };
  pagination.value.current = 1;
  void fetchTableData();
}

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  pagination.value.current = pageInfo.current;
  pagination.value.pageSize = pageInfo.pageSize;
  void fetchTableData();
}

function handleSelectChange(keys: Array<string | number>) {
  selectedRowKeys.value = keys as number[];
}

onMounted(() => {
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="80px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.logs.apiName')" name="api_name">
              <Input
                v-model="searchForm.api_name"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.logs.accessName')" name="access_name">
              <Input
                v-model="searchForm.access_name"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.logs.ip')" name="ip">
              <Input
                v-model="searchForm.ip"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.logs.accessTime')" name="access_time">
              <DateRangePicker
                v-model="searchForm.access_time"
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
        <div class="mb-3 flex items-center">
          <Space>
            <Button theme="danger" variant="outline" @click="handleBatchDelete">
              <template #icon><DeleteIcon /></template>
              {{ $t('common.batchDelete') }}
            </Button>
          </Space>
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
          @page-change="handlePageChange"
          @select-change="handleSelectChange"
        >
          <template #response_code="{ row }">
            <Tag :theme="String(row.response_code) === '200' ? 'success' : 'danger'" variant="light">
              {{ row.response_code }}
            </Tag>
          </template>

          <template #remark="{ row }">
            <span class="block max-w-[320px] truncate">
              {{ row?.remark || '-' }}
            </span>
          </template>

          <template #action="{ row }">
            <Popconfirm
              :content="$t('system.logs.confirmDeleteApiLogSingle')"
              @confirm="handleDelete(row)"
            >
              <Button size="small" theme="danger" variant="outline">
                <template #icon><DeleteIcon /></template>
                {{ $t('common.delete') }}
              </Button>
            </Popconfirm>
          </template>
        </Table>
      </div>
    </div>
  </Page>
</template>
