<script lang="ts" setup>
import type { AppApi } from '#/api/system/app';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { computed, onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

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
  EnhancedTable as Table,
  Form,
  FormItem,
  Input,
  Popconfirm,
  Select,
  Space,
  Switch,
  Tooltip,
} from 'tdesign-vue-next';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';
import {
  changeAppStatus,
  deleteApp,
  realDeleteApp,
  recoveryApp,
} from '#/api/system/app';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import AppModal from './components/app-modal.vue';
import type { AppListItem, AppTableColumn } from './model';
import {
  createAppColumnOptions,
  createAppTableColumns,
} from './schemas';
import { useAppCrud } from './use-app-crud';

defineOptions({ name: 'SystemApp' });

const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

type AppModalInstance = {
  open: (data?: Partial<AppApi.SubmitPayload>) => void;
};

const appModalRef = ref<AppModalInstance>();
const fallbackStatusOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

const statusOptions = ref<DictOption[]>([]);

const columns: AppTableColumn[] = createAppTableColumns();
const columnOptions = createAppColumnOptions(columns);
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
} = useAppCrud();

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
  return keys.map((key) => Number(key)).filter((id) => !Number.isNaN(id));
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
  statusOptions.value =
    options.length > 0 ? options : fallbackStatusOptions;
}

function handleAdd() {
  appModalRef.value?.open();
}

function handleEdit(row: AppListItem) {
  appModalRef.value?.open({
    ...row,
    status: Number(row.status ?? 1),
  });
}

async function handleDelete(row: AppListItem) {
  try {
    await (isRecycleBin.value ? realDeleteApp([row.id]) : deleteApp([row.id]));
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
    await (isRecycleBin.value ? realDeleteApp(ids) : deleteApp(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: AppListItem) {
  try {
    await recoveryApp([row.id]);
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
    await recoveryApp(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: AppListItem, checked: boolean) {
  try {
    await changeAppStatus({ id: row.id, status: checked ? 1 : 2 });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

function handleStatusSwitchChange(row: AppListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

function handleSuccess() {
  void fetchTableData();
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
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
        <Form :data="searchForm" label-width="90px" layout="inline" colon>
          <div class="w-full">
            <div class="grid grid-cols-4 gap-x-4 gap-y-3 w-full">
              <FormItem :label="$t('system.app.name')" name="app_name">
                <Input
                  v-model="searchForm.app_name"
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
              <FormItem :label="$t('common.createTime')" name="created_at" class="col-span-2">
                <DateRangePicker
                  v-model="searchForm.created_at"
                  :placeholder="[$t('common.startTime'), $t('common.endTime')]"
                  clearable
                  class="w-full"
                />
              </FormItem>
            </div>
            <div class="flex justify-start gap-2 pt-4">
              <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
              <Button theme="primary" @click="handleSearch">
                <template #icon><SearchIcon /></template>
                {{ $t('common.query') }}
              </Button>
            </div>
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

        <div class="min-h-0 flex-1 overflow-auto">
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
              :value="Number(row?.status) === 1"
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
                  :content="$t('system.app.confirmDelete')"
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
                  :content="$t('system.app.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('system.app.confirmPermanentDelete')"
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
    </div>

    <AppModal ref="appModalRef" @success="handleSuccess" />
  </Page>
</template>
