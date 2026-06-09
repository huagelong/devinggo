<script lang="ts" setup>
import type { DemoListItem } from './model';

import { computed, onMounted, onUnmounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';
import { message } from '#/adapter/tdesign';

import {
  AddIcon,
  DeleteIcon,
  EditIcon,
  FullscreenExitIcon,
  FullscreenIcon,
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
  Tooltip,
} from 'tdesign-vue-next';

import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import {
  changeDemoStatus,
  deleteDemo,
  realDeleteDemo,
  recoveryDemo,
} from '#/api/system/demo';
import { logger } from '#/utils/logger';
import type { DictOption } from '#/composables/crud/use-dict-options';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import DemoModal from './components/demo-modal.vue';
import type { DemoTableColumn } from './model';
import { createDemoColumnOptions, createDemoTableColumns } from './schemas';
import { useDemoCrud } from './use-demo-crud';

defineOptions({ name: 'SystemDemo' });

const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

type DemoModalInstance = {
  open: (data?: Partial<DemoListItem>) => void;
};

const demoModalRef = ref<DemoModalInstance>();
const statusOptions = ref<DictOption[]>([{label: $t('common.statusEnabled'), value: "1"}, {label: $t('common.disabled'), value: "2"}]);

const columns: DemoTableColumn[] = createDemoTableColumns();
const columnOptions = createDemoColumnOptions(columns);
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
} = useDemoCrud();

const { getDictOptions } = useDictOptions();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
  if (options.length > 0) {
    statusOptions.value = options;
  }
}

function handleAdd() {
  demoModalRef.value?.open();
}

function handleEdit(row: DemoListItem) {
  demoModalRef.value?.open(row);
}

async function handleDelete(row: DemoListItem) {
  try {
    await (isRecycleBin.value ? realDeleteDemo([row.id]) : deleteDemo([row.id]));
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
    await (isRecycleBin.value ? realDeleteDemo(ids) : deleteDemo(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: DemoListItem) {
  try {
    await recoveryDemo([row.id]);
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
    await recoveryDemo(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: DemoListItem, checked: boolean) {
  const status = checked ? 1 : 2;
  try {
    await changeDemoStatus({ id: row.id, status });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

function handleSuccess() {
  void fetchTableData();
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
}

function handleStatusSwitchChange(row: DemoListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

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
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="80px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('common.name')" name="name">
              <Input
                v-model="searchForm.name"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.demo.code')" name="code">
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
            <FormItem :label="$t('system.demo.birthday')" name="birthday">
              <DatePicker
                v-model="searchForm.birthday"
                :placeholder="$t('system.demo.selectDate')"
                clearable
                class="w-full"
              />
            </FormItem>
            <FormItem :label="$t('common.createTime')" name="created_at" class="col-span-2">
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

      <div ref="tableContainerRef" class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
        <div class="mb-3 flex items-center justify-between">
          <Space>
            <template v-if="!isRecycleBin">
              <Button theme="primary" @click="handleAdd">
                <template #icon><AddIcon /></template>
                {{ $t('common.create') }}
              </Button>
              <Popconfirm
                :content="$t('common.batchDeleteDataConfirm')"
                @confirm="handleBatchDelete"
              >
                <Button theme="danger" variant="outline">
                  <template #icon><DeleteIcon /></template>
                  {{ $t('common.delete') }}
                </Button>
              </Popconfirm>
            </template>
            <template v-else>
              <Popconfirm
                :content="$t('common.batchRecoveryDataConfirm')"
                @confirm="handleBatchRecovery"
              >
                <Button theme="success">{{ $t('common.recovery') }}</Button>
              </Popconfirm>
              <Popconfirm
                :content="$t('common.batchPermanentDeleteDataConfirm')"
                @confirm="handleBatchDelete"
              >
                <Button theme="danger">{{ $t('common.permanentDelete') }}</Button>
              </Popconfirm>
            </template>
          </Space>

          <div class="flex items-center gap-2">
            <Tooltip :content="isFullscreen ? $t('common.exitFullscreen') : $t('common.fullscreen')">
              <Button shape="square" variant="outline" @click="toggleFullscreen">
                <template #icon>
                  <FullscreenExitIcon v-if="isFullscreen" />
                  <FullscreenIcon v-else />
                </template>
              </Button>
            </Tooltip>

            <CrudToolbar
              v-model="visibleColumns"
              :column-options="columnOptions"
              :is-recycle-bin="isRecycleBin"
              @refresh="fetchTableData"
              @toggle-recycle="toggleRecycleBin"
            />
          </div>
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
          @page-change="handlePageChange"
          @select-change="handleTableSelectChange"
        >
          <template #status="{ row }">
            <Switch
              :disabled="isRecycleBin"
              :value="row?.status === 1"
              @change="(value: unknown) => handleStatusSwitchChange(row, value)"
            />
          </template>

          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <template v-if="!isRecycleBin">
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
                  :content="$t('system.demo.confirmDelete')"
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
                  :content="$t('system.demo.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('system.demo.confirmPermanentDelete')"
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

    <DemoModal ref="demoModalRef" @success="handleSuccess" />
  </Page>
</template>
