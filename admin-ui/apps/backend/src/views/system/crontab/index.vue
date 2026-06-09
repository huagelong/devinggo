<script lang="ts" setup>
import type { CrontabListItem } from './model';

import { computed, onMounted, onUnmounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import {
  changeCrontabStatus,
  deleteCrontab,
  realDeleteCrontab,
  recoveryCrontab,
  runCrontab,
} from '#/api/system/crontab';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { logger } from '#/utils/logger';

import {
  DeleteIcon,
  EditIcon,
  FullscreenExitIcon,
  FullscreenIcon,
  HistoryIcon,
  PlusIcon,
  PlayIcon,
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
  Tag,
  Tooltip,
} from 'tdesign-vue-next';

import CrontabLogPanel from './components/crontab-log-panel.vue';
import CrontabModal from './components/crontab-modal.vue';
import type { CrontabTableColumn } from './model';
import {
  createCrontabColumnOptions,
  createCrontabSearchForm,
  createCrontabTableColumns,
  crontabStatusOptions,
  crontabTypeOptions,
} from './schemas';
import { useCrontabCrud } from './use-crontab-crud';

defineOptions({ name: 'SystemCrontab' });

const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

type CrontabModalInstance = {
  open: (data?: Partial<CrontabListItem>) => void;
};

type CrontabLogPanelInstance = {
  open: (id: number) => void;
};

const crontabModalRef = ref<CrontabModalInstance>();
const crontabLogPanelRef = ref<CrontabLogPanelInstance>();

const searchForm = ref(createCrontabSearchForm());

const columns: CrontabTableColumn[] = createCrontabTableColumns();
const columnOptions = createCrontabColumnOptions(columns);
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
  handleReset,
  handleSearch,
  handleSelectChange,
  isRecycleBin,
  loading,
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = useCrontabCrud();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

function handleAdd() {
  crontabModalRef.value?.open();
}

function handleEdit(row: CrontabListItem) {
  crontabModalRef.value?.open(row);
}

async function handleDelete(row: CrontabListItem) {
  try {
    await (isRecycleBin.value
      ? realDeleteCrontab([row.id])
      : deleteCrontab([row.id]));
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
    await (isRecycleBin.value ? realDeleteCrontab(ids) : deleteCrontab(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: CrontabListItem) {
  try {
    await recoveryCrontab([row.id]);
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
    await recoveryCrontab(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleRun(row: CrontabListItem) {
  try {
    await runCrontab({ id: row.id });
    message.success($t('common.executeSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.executeFailed'));
  }
}

async function handleChangeStatus(row: CrontabListItem, newStatus: number) {
  const originalStatus = newStatus === 1 ? 2 : 1;
  try {
    await changeCrontabStatus({ id: row.id, status: newStatus });
    message.success($t('common.operationSuccess'));
    await fetchTableData();
  } catch (error) {
    row.status = originalStatus;
    logger.error(error);
    message.error($t('common.operationFailed'));
  }
}

function handleOpenLog(row: CrontabListItem) {
  crontabLogPanelRef.value?.open(row.id);
}

function handleSuccess() {
  void fetchTableData();
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
              <FormItem :label="$t('system.crontab.name')" name="name">
                <Input
                  v-model="searchForm.name"
                  :placeholder="$t('ui.placeholder.input')"
                  clearable
                />
              </FormItem>
              <FormItem :label="$t('system.crontab.taskType')" name="type">
                <Select
                  v-model="searchForm.type"
                  :options="crontabTypeOptions"
                  :placeholder="$t('ui.placeholder.select')"
                  clearable
                />
              </FormItem>
              <FormItem :label="$t('common.status')" name="status">
                <Select
                  v-model="searchForm.status"
                  :options="crontabStatusOptions"
                  :placeholder="$t('ui.placeholder.select')"
                  clearable
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
                <template #icon><PlusIcon /></template>
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
          :selected-row-keys="selectedRowKeys"
          row-key="id"
          hover
          stripe
          @select-change="handleSelectChange"
        >
          <template #type="{ row }">
            <Tag theme="primary">
              <template v-if="row?.type === 1">{{ $t('system.crontab.typeCommand') }}</template>
              <template v-else-if="row?.type === 2">{{ $t('system.crontab.typeClass') }}</template>
              <template v-else-if="row?.type === 3">{{ $t('system.crontab.typeUrl') }}</template>
              <template v-else-if="row?.type === 4">{{ $t('system.crontab.typeEval') }}</template>
              <template v-else>{{ row?.type }}</template>
            </Tag>
          </template>

          <template #status="{ row }">
            <Switch
              v-model="row.status"
              :custom-value="[1, 2]"
              @change="(val) => handleChangeStatus(row, val as number)"
            />
          </template>

          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <template v-if="!isRecycleBin">
                <Button
                  size="small"
                  theme="primary"
                  variant="text"
                  @click="handleRun(row)"
                >
                  <template #icon><PlayIcon /></template>
                  {{ $t('system.crontab.runOnce') }}
                </Button>
                <Button
                  size="small"
                  theme="default"
                  variant="text"
                  @click="handleOpenLog(row)"
                >
                  <template #icon><HistoryIcon /></template>
                  {{ $t('system.crontab.logTitle') }}
                </Button>
                <Button
                  size="small"
                  theme="primary"
                  variant="text"
                  @click="handleEdit(row)"
                >
                  <template #icon><EditIcon /></template>
                  {{ $t('common.edit') }}
                </Button>
                <Popconfirm
                  :content="$t('system.crontab.confirmDelete')"
                  @confirm="handleDelete(row)"
                >
                  <Button size="small" theme="danger" variant="text">
                    <template #icon><DeleteIcon /></template>
                    {{ $t('common.delete') }}
                  </Button>
                </Popconfirm>
              </template>

              <template v-else>
                <Popconfirm
                  :content="$t('system.crontab.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('system.crontab.confirmPermanentDelete')"
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

    <CrontabModal ref="crontabModalRef" @success="handleSuccess" />
    <CrontabLogPanel ref="crontabLogPanelRef" />
  </Page>
</template>
