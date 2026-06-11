<script lang="ts" setup>
import { computed, onMounted, onUnmounted, ref } from 'vue';
import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

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
  InputNumber,
  Popconfirm,
  Select,
  Space,
  Switch,
  Table,
  Textarea,
  Tooltip,
} from 'tdesign-vue-next';

import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import {
  change<%.EntityName%>Status,
  delete<%.EntityName%>,
  realDelete<%.EntityName%>,
  recovery<%.EntityName%>,
  update<%.EntityName%>Number,
} from '#/api/<%.ModuleName%>/<%.VarName%>';
import type { DictOption } from '#/composables/crud/use-dict-options';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import <%.EntityName%>Modal from './components/<%.VarName%>-modal.vue';
import type { <%.EntityName%>ListItem, <%.EntityName%>TableColumn } from './model';
import { create<%.EntityName%>ColumnOptions, create<%.EntityName%>TableColumns } from './schemas';
import { use<%.EntityName%>Crud } from './use-<%.VarName%>-crud';

defineOptions({ name: '<%.ComponentName%>' });

type <%.EntityName%>ModalInstance = {
  open: (data?: Partial<<%.EntityName%>ListItem>) => void;
};

const <%.VarName%>ModalRef = ref<<%.EntityName%>ModalInstance>();
const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);
const statusOptions = ref<DictOption[]>([]);

const columns: <%.EntityName%>TableColumn[] = create<%.EntityName%>TableColumns();
const columnOptions = create<%.EntityName%>ColumnOptions(columns);
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
} = use<%.EntityName%>Crud();

const { getDictOptions } = useDictOptions();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
  statusOptions.value =
    options.length > 0
      ? options
      : [
          { label: '已上架', value: 1 },
          { label: '未上架', value: 2 },
        ];
}

function handleAdd() {
  <%.VarName%>ModalRef.value?.open();
}

function handleEdit(row: <%.EntityName%>ListItem) {
  <%.VarName%>ModalRef.value?.open(row);
}

async function handleDelete(row: <%.EntityName%>ListItem) {
  try {
    await (isRecycleBin.value ? realDelete<%.EntityName%>([row.id]) : delete<%.EntityName%>([row.id]));
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
    await (isRecycleBin.value ? realDelete<%.EntityName%>(ids) : delete<%.EntityName%>(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: <%.EntityName%>ListItem) {
  try {
    await recovery<%.EntityName%>([row.id]);
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
    await recovery<%.EntityName%>(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: <%.EntityName%>ListItem, checked: boolean) {
  const status = checked ? 1 : 2;
  try {
    await change<%.EntityName%>Status({ id: row.id, status });
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

function handleStatusSwitchChange(row: <%.EntityName%>ListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

async function handleSortChange(value: number | string, row: <%.EntityName%>ListItem) {
  const numberValue = Number(value);
  if (Number.isNaN(numberValue)) return;

  try {
    await update<%.EntityName%>Number({
      id: Number(row.id),
      numberName: 'sort',
      numberValue,
    });
    message.success($t('common.sortUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.sortUpdateFailed'));
  }
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
<%range .SearchFormItems%>            <FormItem :label="$t('<%.LabelKey%>')" name="<%.Name%>">
              <<%.SearchComponent%>
                v-model="searchForm.<%.Name%>"
                :placeholder="<%.Placeholder%>"
                clearable
                <%if eq .SearchComponent "Select"%>class="w-full" :options="statusOptions"<%end%>
              />
            </FormItem>
<%end%>          </div>
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

          <template #sort="{ row }">
            <InputNumber
              :value="row?.sort"
              :min="0"
              :max="9999"
              size="small"
              :disabled="isRecycleBin"
              @change="(value: number | string) => handleSortChange(value, row)"
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
                  :content="$t('<%.ModuleName%>.<%.VarName%>.confirmDelete')"
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
                  :content="$t('<%.ModuleName%>.<%.VarName%>.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('<%.ModuleName%>.<%.VarName%>.confirmPermanentDelete')"
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

    <<%.EntityName%>Modal ref="<%.VarName%>ModalRef" @success="handleSuccess" />
  </Page>
</template>
