<script lang="ts" setup>
import type { AttachmentListItem, AttachmentTreeItem } from './model';

import { computed, onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import {
  deleteAttachment,
  realDeleteAttachment,
  recoveryAttachment,
} from '#/api/system/attachment';
import { logger } from '#/utils/logger';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';

import {
  AppIcon,
  DeleteIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  ImageViewer,
  Input,
  Space,
  Table,
  Tag,
  Tree,
} from 'tdesign-vue-next';

import type { AttachmentTableColumn } from './model';
import {
  createAttachmentColumnOptions,
  createAttachmentSearchForm,
  createAttachmentTableColumns,
  defaultAttachmentTreeData,
} from './schemas';
import { useAttachmentCrud } from './use-attachment-crud';

defineOptions({ name: 'SystemAttachment' });

const viewMode = ref<'list' | 'window'>('list');
const selectedTreeKey = ref<string[]>(['all']);
const treeData = ref<AttachmentTreeItem[]>(defaultAttachmentTreeData);

const searchForm = ref(createAttachmentSearchForm());

const columns: AttachmentTableColumn[] = createAttachmentTableColumns();
const columnOptions = createAttachmentColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>(
  allColumnKeys.filter((key) => key !== 'url')
);

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
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = useAttachmentCrud();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

function handleTreeChange(value: Array<string | number>) {
  const keys = value.map((item) => String(item));
  selectedTreeKey.value = keys.length > 0 ? keys : ['all'];
  const key = selectedTreeKey.value[0];
  if (key === 'all') {
    searchForm.value.mime_type = undefined;
  } else {
    searchForm.value.mime_type = key;
  }
}

function handleDownload(row: AttachmentListItem) {
  // In a real implementation, this would trigger file download
  message.success($t('system.attachment.downloadFile', [row.origin_name]));
}

async function handleDelete(row: AttachmentListItem) {
  try {
    await (isRecycleBin.value
      ? realDeleteAttachment([row.id])
      : deleteAttachment([row.id]));
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
    await (isRecycleBin.value
      ? realDeleteAttachment(ids)
      : deleteAttachment(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: AttachmentListItem) {
  try {
    await recoveryAttachment([row.id]);
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
    await recoveryAttachment(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

function switchViewMode() {
  viewMode.value = viewMode.value === 'list' ? 'window' : 'list';
}

function isImageType(mimeType: string): boolean {
  return /^image\//.test(mimeType);
}

// Image preview
const previewVisible = ref(false);
const previewImageUrl = ref('');

function handlePreviewImage(url: string) {
  previewImageUrl.value = url;
  previewVisible.value = true;
}

function getFileExtension(mimeType: string): string {
  const map: Record<string, string> = {
    'application/pdf': 'PDF',
    'application/zip': 'ZIP',
    'application/x-rar': 'RAR',
    'text/plain': 'TXT',
    'application/vnd.ms-excel': 'XLS',
    'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet': 'XLSX',
  };
  return map[mimeType] || mimeType.split('/')[1]?.toUpperCase() || 'FILE';
}

onMounted(() => {
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full gap-3">
      <!-- Left Tree Slider -->
      <div class="w-48 flex-shrink-0 rounded-md bg-white p-2">
        <div class="mb-2 px-2 text-sm font-medium text-gray-500">{{ $t('system.upload.fileType') }}</div>
        <Tree
          v-model="selectedTreeKey"
          :data="treeData"
          hover
          expand-all
          @change="handleTreeChange"
        />
      </div>

      <!-- Main Content -->
      <div class="flex min-h-0 flex-1 flex-col gap-3">
        <div class="rounded-md bg-white p-4">
          <Form :data="searchForm" label-width="80px" layout="inline" colon>
            <div class="grid grid-cols-4 gap-x-4 gap-y-3">
              <FormItem :label="$t('system.attachment.originName')" name="origin_name">
                <Input
                  v-model="searchForm.origin_name"
                  :placeholder="$t('ui.placeholder.input')"
                  clearable
                />
              </FormItem>
              <FormItem :label="$t('system.attachment.storageMode')" name="storage_mode">
                <Input
                  v-model="searchForm.storage_mode"
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

            <Space>
              <Button
                theme="default"
                variant="outline"
                @click="switchViewMode"
              >
                <template #icon>
                  <AppIcon />
                </template>
                {{ viewMode === 'list' ? $t('system.attachment.windowMode') : $t('system.attachment.listMode') }}
              </Button>

              <CrudToolbar
                v-model="visibleColumns"
                :column-options="columnOptions"
                :is-recycle-bin="isRecycleBin"
                @refresh="fetchTableData"
                @toggle-recycle="toggleRecycleBin"
              />
            </Space>
          </div>

          <!-- List View -->
          <Table
            v-if="viewMode === 'list'"
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
            @select-change="handleSelectChange"
          >
            <template #url="{ row }">
              <div class="flex items-center justify-center">
                <img
                  v-if="isImageType(row?.mime_type)"
                  :src="row?.url"
                  :alt="row?.origin_name"
                  class="h-10 w-10 cursor-zoom-in rounded object-cover transition hover:opacity-80"
                  @click="handlePreviewImage(row?.url)"
                />
                <Tag v-else theme="default">
                  {{ getFileExtension(row?.mime_type) }}
                </Tag>
              </div>
            </template>

            <template #storage_mode="{ row }">
              <Tag :theme="row?.storage_mode === 1 ? 'primary' : 'warning'">
                {{ row?.storage_mode === 1 ? $t('common.local') : $t('common.cloudStorage') }}
              </Tag>
            </template>

            <template #action="{ row }">
              <div class="flex items-center justify-center gap-1">
                <template v-if="!isRecycleBin">
                  <Button
                    size="small"
                    theme="primary"
                    variant="outline"
                    @click="handleDownload(row)"
                  >
                    {{ $t('common.download') }}
                  </Button>
                  <Button
                    size="small"
                    theme="danger"
                    variant="outline"
                    @click="handleDelete(row)"
                  >
                    {{ $t('common.delete') }}
                  </Button>
                </template>
                <template v-else>
                  <Button
                    size="small"
                    theme="primary"
                    variant="outline"
                    @click="handleRecovery(row)"
                  >
                    {{ $t('common.recovery') }}
                  </Button>
                  <Button
                    size="small"
                    theme="danger"
                    variant="outline"
                    @click="handleDelete(row)"
                  >
                    {{ $t('common.permanentDelete') }}
                  </Button>
                </template>
              </div>
            </template>
          </Table>

          <!-- Window View (Gallery) -->
          <div v-else class="grid grid-cols-4 gap-4">
            <div
              v-for="row in tableData"
              :key="row.id"
              class="group relative rounded-md border border-gray-200 p-2 transition hover:border-blue-400"
            >
              <div class="flex h-32 items-center justify-center overflow-hidden rounded bg-gray-50">
                <img
                  v-if="isImageType(row.mime_type)"
                  :src="row.url"
                  :alt="row.origin_name"
                  class="max-h-full max-w-full cursor-zoom-in object-contain"
                  @click="handlePreviewImage(row.url)"
                />
                <Tag v-else theme="default" size="large">
                  {{ getFileExtension(row.mime_type) }}
                </Tag>
              </div>
              <div class="mt-2 text-sm">
                <div class="truncate text-gray-700" :title="row.origin_name">
                  {{ row.origin_name }}
                </div>
                <div class="text-xs text-gray-400">{{ row.size_info }}</div>
              </div>
              <div
                class="absolute left-0 top-0 flex h-full w-full items-center justify-center gap-2 rounded bg-black/50 opacity-0 transition group-hover:opacity-100"
              >
                <Button
                  size="small"
                  theme="primary"
                  @click="handleDownload(row)"
                >
                  {{ $t('common.download') }}
                </Button>
                <Button
                  size="small"
                  theme="danger"
                  @click="handleDelete(row)"
                >
                  {{ $t('common.delete') }}
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Image Preview -->
    <ImageViewer
      v-model:visible="previewVisible"
      :images="[previewImageUrl]"
      :close-on-overlay="true"
    />
  </Page>
</template>
