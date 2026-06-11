<script lang="ts" setup>
import type { LogApi } from '#/api/system/log';

import { onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import { getOperLogPageList } from '#/api/system/log';
import { logger } from '#/utils/logger';

import { SearchIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Dialog,
  Form,
  FormItem,
  Input,
  Select,
  Table,
  Tag,
} from 'tdesign-vue-next';

defineOptions({ name: 'SystemOperLog' });

const loading = ref(false);
const tableData = ref<LogApi.OperLogItem[]>([]);
const searchForm = ref<{ username: string; service_name: string; ip: string; status: number | undefined; created_at: string[] }>({
  username: '',
  service_name: '',
  ip: '',
  status: undefined,
  created_at: [],
});

const pagination = ref({
  current: 1,
  pageSize: 20,
  pageSizeOptions: [10, 20, 50, 100],
  showJumper: true,
  showPageSize: true,
  total: 0,
});

const statusOptions = [
  { label: $t('system.logs.loginSuccess'), value: 1 },
  { label: $t('system.logs.loginFailed'), value: 2 },
];

const detailVisible = ref(false);
const detailRow = ref<LogApi.OperLogItem | null>(null);

const columns = [
  { colKey: 'username', title: $t('system.logs.username'), width: 120 },
  { colKey: 'service_name', title: $t('system.logs.serviceName'), width: 160 },
  { colKey: 'router', title: $t('system.logs.router'), width: 180, ellipsis: true },
  { colKey: 'method', title: $t('system.logs.method'), width: 100 },
  { colKey: 'ip', title: $t('system.logs.ip'), width: 140 },
  { colKey: 'ip_location', title: $t('system.logs.ipLocation'), width: 120 },
  { colKey: 'response_code', title: $t('system.logs.responseCode'), width: 100 },
  { colKey: 'created_at', title: $t('system.logs.createdAt'), width: 180 },
  { colKey: 'operation', title: $t('common.operation'), width: 100 },
];

function handleViewDetail(row: LogApi.OperLogItem) {
  detailRow.value = row;
  detailVisible.value = true;
}

async function fetchTableData() {
  loading.value = true;
  try {
    const params: LogApi.OperLogQuery = {
      page: pagination.value.current,
      pageSize: pagination.value.pageSize,
    };
    if (searchForm.value.username) params.username = searchForm.value.username;
    if (searchForm.value.service_name) params.service_name = searchForm.value.service_name;
    if (searchForm.value.ip) params.ip = searchForm.value.ip;
    if (searchForm.value.status !== undefined && searchForm.value.status !== null) {
      params.status = searchForm.value.status;
    }
    if (searchForm.value.created_at?.length === 2 && searchForm.value.created_at[0]) {
      params.created_at = searchForm.value.created_at;
    }
    const response = await getOperLogPageList(params);
    tableData.value = response.items || [];
    pagination.value.total = Number(response.pageInfo?.total || response.total || 0);
  } catch (error) {
    logger.error(error);
    message.error($t('common.logFetchFailed'));
  } finally {
    loading.value = false;
  }
}

function handleSearch() {
  pagination.value.current = 1;
  void fetchTableData();
}

function handleReset() {
  searchForm.value = { username: '', service_name: '', ip: '', status: undefined, created_at: [] };
  pagination.value.current = 1;
  void fetchTableData();
}

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  pagination.value.current = pageInfo.current;
  pagination.value.pageSize = pageInfo.pageSize;
  void fetchTableData();
}

onMounted(() => {
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-card p-4">
        <Form :data="searchForm" label-width="80px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.logs.username')" name="username">
              <Input
                v-model="searchForm.username"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.logs.serviceName')" name="service_name">
              <Input
                v-model="searchForm.service_name"
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
            <FormItem :label="$t('system.logs.status')" name="status">
              <Select
                v-model="searchForm.status"
                :options="statusOptions"
                :placeholder="$t('ui.placeholder.select')"
                clearable
                class="w-full"
              />
            </FormItem>
            <FormItem :label="$t('system.logs.createdAt')" name="created_at">
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

      <div class="flex min-h-0 flex-1 flex-col rounded-md bg-card p-4">
        <Table
          :columns="columns"
          :data="tableData"
          :loading="loading"
          :pagination="pagination"
          row-key="id"
          hover
          stripe
          @page-change="handlePageChange"
        >
        <template #response_code="{ row }">
          <Tag :theme="String(row.response_code) === '200' ? 'success' : 'danger'" variant="light">
            {{ row.response_code }}
          </Tag>
        </template>
        <template #operation="{ row }">
          <Button theme="primary" variant="text" size="small" @click="handleViewDetail(row)">
            {{ $t('common.detail') }}
          </Button>
        </template>
        </Table>
      </div>
    </div>

    <Dialog
      v-model:visible="detailVisible"
      :header="$t('system.logs.operLogDetail')"
      width="700px"
      :footer="false"
    >
      <div v-if="detailRow" class="space-y-4">
        <div class="grid grid-cols-2 gap-4">
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.router') }}:</span>
            <span class="flex-1 break-all">{{ detailRow.router }}</span>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.username') }}:</span>
            <span>{{ detailRow.username }}</span>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.method') }}:</span>
            <span>{{ detailRow.method }}</span>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.responseCode') }}:</span>
            <Tag :theme="String(detailRow.response_code) === '200' ? 'success' : 'danger'" variant="light">
              {{ detailRow.response_code }}
            </Tag>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.serviceName') }}:</span>
            <span>{{ detailRow.service_name }}</span>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.createdAt') }}:</span>
            <span>{{ detailRow.created_at }}</span>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.ip') }}:</span>
            <span>{{ detailRow.ip }}</span>
          </div>
          <div class="flex">
            <span class="w-24 text-muted-foreground">{{ $t('system.logs.ipLocation') }}:</span>
            <span>{{ detailRow.ip_location }}</span>
          </div>
        </div>
        <div class="border-t pt-4">
          <div class="mb-2 font-medium">{{ $t('system.logs.requestData') }}</div>
          <pre class="max-h-40 overflow-auto rounded bg-muted/50 p-3 text-sm">{{ detailRow.request_data || '-' }}</pre>
        </div>
        <div class="border-t pt-4">
          <div class="mb-2 font-medium">{{ $t('system.logs.responseData') }}</div>
          <pre class="max-h-40 overflow-auto rounded bg-muted/50 p-3 text-sm">{{ detailRow.response_data || '-' }}</pre>
        </div>
      </div>
    </Dialog>
  </Page>
</template>
