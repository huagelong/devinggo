<script lang="ts" setup>
import type { CrontabListItem } from './model';

import { computed, onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import {
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
  Table,
  Tag,
} from 'tdesign-vue-next';

import CrontabLogPanel from './components/crontab-log-panel.vue';
import CrontabModal from './components/crontab-modal.vue';
import type { CrontabTableColumn } from './model';
import {
  createCrontabColumnOptions,
  createCrontabSearchForm,
  createCrontabTableColumns,
  crontabFinallyOptions,
  crontabTypeOptions,
} from './schemas';
import { useCrontabCrud } from './use-crontab-crud';

defineOptions({ name: 'SystemCrontab' });

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

function handleOpenLog(row: CrontabListItem) {
  crontabLogPanelRef.value?.open(row.id);
}

function handleSuccess() {
  void fetchTableData();
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
            <FormItem :label="$t('system.crontab.isFinally')" name="is_finally">
              <Select
                v-model="searchForm.is_finally"
                :options="crontabFinallyOptions"
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
                <template #icon><PlusIcon /></template>
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
          :selected-row-keys="selectedRowKeys"
          row-key="id"
          hover
          stripe
          @select-change="handleSelectChange"
        >
          <template #type="{ row }">
            <Tag theme="primary">
              {{ row?.type === 1 ? $t('system.crontab.typeInterval') : $t('system.crontab.typeCron') }}
            </Tag>
          </template>

          <template #is_finally="{ row }">
            <Tag :theme="row?.is_finally === 1 ? 'success' : 'default'">
              {{ row?.is_finally === 1 ? $t('common.yes') : $t('common.no') }}
            </Tag>
          </template>

          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <template v-if="!isRecycleBin">
                <Button
                  size="small"
                  theme="default"
                  variant="outline"
                  @click="handleOpenLog(row)"
                >
                  <template #icon><HistoryIcon /></template>
                  {{ $t('system.crontab.logTitle') }}
                </Button>
                <Button
                  size="small"
                  theme="warning"
                  variant="outline"
                  @click="handleRun(row)"
                >
                  <template #icon><PlayIcon /></template>
                  {{ $t('common.execute') }}
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
                  :content="$t('system.crontab.confirmDelete')"
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
