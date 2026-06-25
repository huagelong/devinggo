<script lang="ts" setup>
import { computed, onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
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
  Table,
  Tooltip,
} from 'tdesign-vue-next';

import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import {
  deleteNotice,
  realDeleteNotice,
  recoveryNotice,
} from '#/api/system/notice';
import type { DictOption } from '#/composables/crud/use-dict-options';
import { useDictOptions } from '#/composables/crud/use-dict-options';
import { sanitizeHtml } from '#/utils/sanitize';
import { logger } from '#/utils/logger';

import NoticeModal from './components/notice-modal.vue';
import type { NoticeListItem, NoticeTableColumn } from './model';
import {
  createNoticeColumnOptions,
  createNoticeTableColumns,
} from './schemas';
import { useNoticeCrud } from './use-notice-crud';

defineOptions({ name: 'SystemNotice' });

const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

type NoticeModalInstance = {
  open: (data?: NoticeListItem) => void;
};

const noticeModalRef = ref<NoticeModalInstance>();
const fallbackNoticeTypeOptions: DictOption[] = [
  { label: $t('system.notice.noticeType'), value: 1 },
  { label: $t('system.notice.announcementType'), value: 2 },
];

function normalizeNoticeTypeOptions(options: DictOption[]) {
  return (options || []).map((item) => {
    const numericValue = Number(item.value);
    return Number.isNaN(numericValue) ? { ...item } : { ...item, value: numericValue };
  });
}

const noticeTypeOptions = ref<DictOption[]>([]);

const columns: NoticeTableColumn[] = createNoticeTableColumns();
const columnOptions = createNoticeColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);
const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value.filter((item) => item !== 'row-select');
  },
});

const noticeTypeMap = computed(() => {
  const map = new Map<string, string>();
  noticeTypeOptions.value.forEach((item) => {
    map.set(String(item.value), item.label ?? '');
  });
  return map;
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
} = useNoticeCrud();

const { getDictOptions } = useDictOptions();

async function fetchNoticeTypeDict() {
  const options = await getDictOptions('backend_notice_type');
  noticeTypeOptions.value = normalizeNoticeTypeOptions(
    options.length > 0 ? options : fallbackNoticeTypeOptions,
  );
}

function resolveNoticeTypeLabel(value?: number | string) {
  return noticeTypeMap.value.get(String(value ?? '')) || '-';
}

function normalizeId(value: number | string) {
  const numericValue = Number(value);
  return Number.isNaN(numericValue) ? undefined : numericValue;
}

function toIds(keys: Array<number | string>) {
  return keys
    .map((key) => normalizeId(key))
    .filter((id): id is number => typeof id === 'number');
}

function handleAdd() {
  noticeModalRef.value?.open();
}

function handleEdit(row: NoticeListItem) {
  noticeModalRef.value?.open(row);
}

async function handleDelete(row: NoticeListItem) {
  try {
    await (isRecycleBin.value ? realDeleteNotice([row.id]) : deleteNotice([row.id]));
    message.success($t('common.operationSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error(isRecycleBin.value ? $t('common.permanentDeleteFailed') : $t('common.deleteFailed'));
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
    await (isRecycleBin.value ? realDeleteNotice(ids) : deleteNotice(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error(isRecycleBin.value ? $t('common.batchPermanentDeleteFailed') : $t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: NoticeListItem) {
  try {
    await recoveryNotice([row.id]);
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
    await recoveryNotice(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
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
  void fetchNoticeTypeDict();
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
      <div class="rounded-md bg-card p-4">
        <Form :data="searchForm" label-width="80px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.notice.title')" name="title">
              <Input
                v-model="searchForm.title"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.notice.type')" name="type">
              <Select
                v-model="searchForm.type"
                :options="noticeTypeOptions"
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

      <div ref="tableContainerRef" class="flex min-h-0 flex-1 flex-col rounded-md bg-card p-4">
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
          <template #type="{ row }">
            <span>{{ resolveNoticeTypeLabel(row?.type) }}</span>
          </template>

          <template #remark="{ row }">
            <span class="block max-w-[320px] truncate">
              {{ row?.remark || '-' }}
            </span>
          </template>

          <template #content="{ row }">
            <span
              class="block max-w-[320px] truncate text-muted-foreground text-sm"
              v-html="sanitizeHtml(row?.content)"
            />
          </template>

          <template #created_at="{ row }">
            <span>{{ row?.created_at || '-' }}</span>
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
                  :content="$t('system.notice.confirmDelete')"
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
                  :content="$t('system.notice.confirmRecovery')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="success" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('system.notice.confirmPermanentDelete')"
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

    <NoticeModal ref="noticeModalRef" @success="handleSuccess" />
  </Page>
</template>
