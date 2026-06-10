<script lang="ts" setup>
import type { AttachmentListItem, AttachmentTreeItem } from './model';

import { computed, onMounted, onUnmounted, ref, watch } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import {
  deleteAttachment,
  realDeleteAttachment,
  recoveryAttachment,
} from '#/api/system/attachment';
import { getDictList } from '#/api/system/dict';
import type { DictApi } from '#/api/system/dict';
import { logger } from '#/utils/logger';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';

import {
  AppIcon,
  DeleteIcon,
  FolderIcon,
  FolderZipIcon,
  FullscreenExitIcon,
  FullscreenIcon,
  ImageIcon,
  FileIcon,
  MusicIcon,
  SearchIcon,
  VideoIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  ImageViewer,
  Input,
  Popconfirm,
  Select,
  Space,
  Table,
  Tag,
  Tooltip,
  Tree,
} from 'tdesign-vue-next';

import type { AttachmentTableColumn } from './model';
import {
  createAttachmentColumnOptions,
  createAttachmentTableColumns,
} from './schemas';
import { useAttachmentCrud } from './use-attachment-crud';

defineOptions({ name: 'SystemAttachment' });

const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

const viewMode = ref<'list' | 'window'>('list');
const selectedTreeKey = ref<string[]>(['all']);

// Dict data
const storageModeDictOptions = ref<DictApi.Item[]>([]);
const resourceTypeDictOptions = ref<DictApi.Item[]>([]);

// Tree data from dict
const treeData = computed<AttachmentTreeItem[]>(() => {
  const items = resourceTypeDictOptions.value.map((item) => ({
    title: item.title,
    key: String(item.key),
    icon: getResourceTypeIcon(String(item.key)),
  }));
  return [
    { title: $t('system.attachment.filterAll'), key: 'all', icon: 'folder' },
    ...items,
  ];
});

function getResourceTypeIcon(key: string): string {
  const iconMap: Record<string, string> = {
    image: 'image',
    document: 'file',
    audio: 'music',
    video: 'video',
    archive: 'folder-zip',
  };
  return iconMap[key] || 'file';
}

// Tree expand/collapse and search
const isTreeExpanded = ref(true);
const treeSearchKeyword = ref('');
const filteredTreeData = computed(() => {
  if (!treeSearchKeyword.value) return treeData.value;
  const keyword = treeSearchKeyword.value.toLowerCase();
  return treeData.value.filter((item) =>
    item.title.toLowerCase().includes(keyword),
  );
});

function toggleTreeExpand() {
  isTreeExpanded.value = !isTreeExpanded.value;
}

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
  searchForm,
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = useAttachmentCrud();

// Watch tree selection changes
watch(selectedTreeKey, (newValue) => {
  const key = newValue[0];
  if (key === 'all') {
    searchForm.mime_type = undefined;
  } else {
    searchForm.mime_type = key;
  }
  handleSearch();
});

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

function handleTreeClick(context: { node?: { value?: string | number } }) {
  const key = String(context?.node?.value ?? 'all');
  selectedTreeKey.value = [key];
  if (key === 'all') {
    searchForm.mime_type = undefined;
  } else {
    searchForm.mime_type = key;
  }
  handleSearch();
}

function renderTreeIcon(node: AttachmentTreeItem) {
  switch (node.icon) {
    case 'image':
      return ImageIcon;
    case 'video':
      return VideoIcon;
    case 'music':
      return MusicIcon;
    case 'folder-zip':
      return FolderZipIcon;
    case 'file':
      return FileIcon;
    case 'folder':
    default:
      return FolderIcon;
  }
}

function handleDownload(row: AttachmentListItem) {
  const link = document.createElement('a');
  link.href = row.url;
  link.download = row.origin_name;
  link.target = '_blank';
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
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

function isImageType(mimeType: string, originName?: string): boolean {
  if (mimeType && /^image\//.test(mimeType)) {
    return true;
  }
  // Fallback: check file extension from origin_name
  if (originName) {
    const ext = originName.split('.').pop()?.toLowerCase();
    return ['jpg', 'jpeg', 'png', 'gif', 'bmp', 'webp', 'svg', 'ico'].includes(ext || '');
  }
  return false;
}

// Image preview
const previewVisible = ref(false);
const previewImageUrl = ref('');

function handlePreviewImage(url: string) {
  previewImageUrl.value = url;
  previewVisible.value = true;
}

function handlePreview(row: AttachmentListItem) {
  if (isImageType(row.mime_type, row.origin_name)) {
    handlePreviewImage(row.url);
  } else {
    window.open(row.url, '_blank');
  }
}

function getFileExtension(mimeType: string, originName?: string): string {
  const map: Record<string, string> = {
    'application/pdf': 'PDF',
    'application/zip': 'ZIP',
    'application/x-rar': 'RAR',
    'text/plain': 'TXT',
    'application/vnd.ms-excel': 'XLS',
    'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet': 'XLSX',
  };
  if (mimeType) {
    return map[mimeType] || mimeType.split('/')[1]?.toUpperCase() || 'FILE';
  }
  // Fallback: extract extension from origin_name
  if (originName) {
    const ext = originName.split('.').pop()?.toUpperCase();
    return ext || 'FILE';
  }
  return 'FILE';
}

async function loadDictData() {
  try {
    const [storageModes, resourceTypes] = await Promise.all([
      getDictList('upload_mode'),
      getDictList('attachment_type'),
    ]);
    storageModeDictOptions.value = storageModes;
    resourceTypeDictOptions.value = resourceTypes;
  } catch (error) {
    logger.error(error);
  }
}

function getStorageModeLabel(mode: number): string {
  const item = storageModeDictOptions.value.find(
    (opt) => Number(opt.key) === mode,
  );
  return item?.title || String(mode);
}

onMounted(() => {
  document.addEventListener('fullscreenchange', handleFullscreenChange);
  void loadDictData();
  void fetchTableData();
});

onUnmounted(() => {
  document.removeEventListener('fullscreenchange', handleFullscreenChange);
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full gap-3">
      <!-- Left Tree Slider -->
      <div class="w-48 flex-shrink-0 rounded-md bg-white p-2">
        <div class="mb-2 flex items-center justify-between px-2">
          <span class="text-sm font-medium text-gray-500">{{ $t('system.attachment.resourceType') }}</span>
          <Button
            size="small"
            theme="default"
            variant="text"
            @click="toggleTreeExpand"
          >
            {{ isTreeExpanded ? $t('common.collapse') : $t('common.expand') }}
          </Button>
        </div>
        <div v-show="isTreeExpanded" class="mb-2">
          <Input
            v-model="treeSearchKeyword"
            :placeholder="$t('ui.placeholder.search')"
            clearable
            size="small"
          >
            <template #prefix-icon>
              <SearchIcon />
            </template>
          </Input>
        </div>
        <Tree
          v-show="isTreeExpanded"
          v-model:value="selectedTreeKey"
          :data="filteredTreeData"
          :keys="{ value: 'key', label: 'title', children: 'children' }"
          activable
          hover
          expand-all
          @click="handleTreeClick"
        >
          <template #icon="{ node }">
            <component :is="renderTreeIcon(node.data)" />
          </template>
        </Tree>
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
                <Select
                  v-model="searchForm.storage_mode"
                  :options="storageModeDictOptions.map((item) => ({ label: item.title, value: Number(item.key) }))"
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

          <!-- List View -->
          <div v-if="viewMode === 'list'" class="min-h-0 flex-1 overflow-auto">
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
              @select-change="handleSelectChange"
            >
            <template #url="{ row }">
              <div class="flex items-center justify-center">
                <img
                  v-if="isImageType(row?.mime_type, row?.origin_name)"
                  :src="row?.url"
                  :alt="row?.origin_name"
                  class="h-10 w-10 cursor-zoom-in rounded object-cover transition hover:opacity-80"
                  @click="handlePreviewImage(row?.url)"
                />
                <Tag v-else theme="default">
                  {{ getFileExtension(row?.mime_type, row?.origin_name) }}
                </Tag>
              </div>
            </template>

            <template #storage_mode="{ row }">
              <Tag theme="primary">
                {{ getStorageModeLabel(row?.storage_mode) }}
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
          </div>

          <!-- Window View (Gallery) -->
          <div v-else class="grid grid-cols-4 gap-4">
            <div
              v-for="row in tableData"
              :key="row.id"
              class="group relative rounded-md border border-gray-200 p-2 transition hover:border-blue-400"
            >
              <div class="flex h-32 items-center justify-center overflow-hidden rounded bg-gray-50">
                <img
                  v-if="isImageType(row.mime_type, row.origin_name)"
                  :src="row.url"
                  :alt="row.origin_name"
                  class="max-h-full max-w-full cursor-zoom-in object-contain"
                  @click="handlePreviewImage(row.url)"
                />
                <Tag v-else theme="default" size="large">
                  {{ getFileExtension(row.mime_type, row.origin_name) }}
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
                  theme="default"
                  variant="outline"
                  @click="handlePreview(row)"
                >
                  {{ $t('common.preview') }}
                </Button>
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
