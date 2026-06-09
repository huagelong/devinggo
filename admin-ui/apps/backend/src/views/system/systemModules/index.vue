<script lang="ts" setup>
import type { SystemModulesApi } from '#/api/system/system-modules';
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
  RefreshIcon,
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
  changeSystemModulesStatus,
  deleteSystemModules,
  realDeleteSystemModules,
  recoverySystemModules,
} from '#/api/system/system-modules';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import SystemModulesModal from './components/system-modules-modal.vue';
import type {
  SystemModulesListItem,
  SystemModulesTableColumn,
} from './model';
import {
  createSystemModulesColumnOptions,
  createSystemModulesTableColumns,
} from './schemas';
import { useSystemModulesCrud } from './use-system-modules-crud';

defineOptions({ name: 'SystemModules' });

type SystemModulesModalInstance = {
  open: (data?: Partial<SystemModulesApi.SubmitPayload>) => void;
};

const systemModulesModalRef = ref<SystemModulesModalInstance>();
const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);
const fallbackStatusOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

const installedOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

const statusOptions = ref<DictOption[]>([]);

const columns: SystemModulesTableColumn[] = createSystemModulesTableColumns();
const columnOptions = createSystemModulesColumnOptions(columns);
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
} = useSystemModulesCrud();

const { getDictOptions } = useDictOptions();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key)).filter((id) => !Number.isNaN(id));
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
  statusOptions.value = options.length > 0 ? options : fallbackStatusOptions;
}

function handleAdd() {
  systemModulesModalRef.value?.open();
}

function handleEdit(row: SystemModulesListItem) {
  systemModulesModalRef.value?.open({
    ...row,
    installed: Number(row.installed ?? 1),
    status: Number(row.status ?? 1),
  });
}

async function handleDelete(row: SystemModulesListItem) {
  try {
    await (isRecycleBin.value
      ? realDeleteSystemModules([row.id])
      : deleteSystemModules([row.id]));
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
    await (isRecycleBin.value
      ? realDeleteSystemModules(ids)
      : deleteSystemModules(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: SystemModulesListItem) {
  try {
    await recoverySystemModules([row.id]);
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
    await recoverySystemModules(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(
  row: SystemModulesListItem,
  checked: boolean,
) {
  try {
    await changeSystemModulesStatus({ id: row.id, status: checked ? 1 : 2 });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

function handleStatusSwitchChange(row: SystemModulesListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

function isDefaultModule(row: SystemModulesListItem) {
  return row.name === 'system';
}

function handleSuccess() {
  void fetchTableData();
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
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
          <div class="w-full">
            <div class="grid grid-cols-4 gap-x-4 gap-y-3 w-full">
              <FormItem label="ID" name="id">
                <Input
                  v-model="searchForm.id"
                  :placeholder="$t('ui.placeholder.input')"
                  clearable
                  type="number"
                />
              </FormItem>
              <FormItem :label="$t('system.systemModules.name')" name="name">
                <Input
                  v-model="searchForm.name"
                  :placeholder="$t('ui.placeholder.input')"
                  clearable
                />
              </FormItem>
              <FormItem :label="$t('system.systemModules.label')" name="label">
                <Input
                  v-model="searchForm.label"
                  :placeholder="$t('ui.placeholder.input')"
                  clearable
                />
              </FormItem>
              <FormItem :label="$t('system.systemModules.installed')" name="installed">
                <Select
                  v-model="searchForm.installed"
                  :options="installedOptions"
                  :placeholder="$t('ui.placeholder.select')"
                  clearable
                  class="w-full"
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
            <div class="flex justify-center gap-2 pt-4">
              <Button theme="primary" @click="handleSearch">
                <template #icon><SearchIcon /></template>
                {{ $t('common.query') }}
              </Button>
              <Button theme="default" @click="handleReset">
                <template #icon><RefreshIcon /></template>
                {{ $t('common.reset') }}
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
                <Button
                  theme="danger"
                  variant="outline"
                >
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
              :disabled="isRecycleBin || isDefaultModule(row)"
              :value="Number(row.status) === 1"
              @change="(value: unknown) => handleStatusSwitchChange(row, value)"
            />
          </template>

          <template #installed="{ row }">
            <span :class="Number(row.installed) === 1 ? 'text-primary' : 'text-danger'"
              >{{ Number(row.installed) === 1 ? $t('common.statusEnabled') : $t('common.statusDisabled') }}</span
            >
          </template>

          <template #action="{ row }">
            <div class="flex items-center justify-center gap-3">
              <template v-if="isDefaultModule(row)">
                <span class="text-gray-400">-</span>
              </template>
              <template v-else-if="!isRecycleBin">
                <a class="inline-flex items-center gap-1 text-primary hover:underline cursor-pointer" @click="handleEdit(row)">
                  <EditIcon />
                  {{ $t('common.edit') }}
                </a>
                <Popconfirm
                  :content="$t('system.systemModules.confirmDelete')"
                  @confirm="handleDelete(row)"
                >
                  <a class="inline-flex items-center gap-1 text-danger hover:underline cursor-pointer">
                    <DeleteIcon />
                    {{ $t('common.delete') }}
                  </a>
                </Popconfirm>
              </template>
              <template v-else>
                <Popconfirm
                  :content="$t('system.systemModules.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <a class="inline-flex items-center gap-1 text-primary hover:underline cursor-pointer">
                    {{ $t('common.recovery') }}
                  </a>
                </Popconfirm>
                <Popconfirm
                  :content="$t('system.systemModules.confirmPermanentDelete')"
                  @confirm="handleDelete(row)"
                >
                  <a class="inline-flex items-center gap-1 text-danger hover:underline cursor-pointer">
                    {{ $t('common.permanentDelete') }}
                  </a>
                </Popconfirm>
              </template>
            </div>
          </template>
        </Table>
      </div>
    </div>

    <SystemModulesModal
      ref="systemModulesModalRef"
      @success="handleSuccess"
    />
  </Page>
</template>
