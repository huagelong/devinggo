<script lang="ts" setup>
import type { UploadTreeItem } from './model';

import { computed, onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { uploadFileFromFileApi } from '#/api/system/upload';
import { logger } from '#/utils/logger';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';

import {
  DeleteIcon,
  FullscreenExitIcon,
  FullscreenIcon,
  SearchIcon,
  UploadIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  Input,
  MessagePlugin,
  Popconfirm,
  Radio,
  RadioGroup,
  Space,
  Table,
  Tag,
  Tooltip,
  Tree,
  Upload,
} from 'tdesign-vue-next';

import type { UploadTableColumn } from './model';
import {
  createUploadColumnOptions,
  createUploadTableColumns,
  defaultUploadTreeData,
} from './schemas';
import { useUploadCrud } from './use-upload-crud';

defineOptions({ name: 'SystemUpload' });

const selectedTreeKey = ref<string[]>(['all']);
const treeData = ref<UploadTreeItem[]>(defaultUploadTreeData);
const uploadingFiles = ref(0);
const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

const columns: UploadTableColumn[] = createUploadTableColumns();
const columnOptions = createUploadColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>(
  allColumnKeys.filter((key) => key !== 'storage_path')
);

const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value, 'action'],
  set: (value: string[]) => {
    visibleColumns.value = value.filter(
      (item) => item !== 'row-select' && item !== 'action',
    );
  },
});

const {
  fetchTableData,
  handleBatchDelete,
  handleDownload,
  handlePageChange,
  handleReset,
  handleSearch,
  handleSelectChange,
  loading,
  pagination,
  searchForm,
  selectedRowKeys,
  tableData,
} = useUploadCrud();

function handleTreeChange(value: Array<string | number>) {
  const keys = value.map((item) => String(item));
  selectedTreeKey.value = keys.length > 0 ? keys : ['all'];
  const key = selectedTreeKey.value[0];
  if (key === 'all') {
    searchForm.mime_type = '';
  } else {
    searchForm.mime_type = key ?? '';
  }
  handleSearch();
}

async function handleUpload(file: File) {
  uploadingFiles.value++;
  try {
    await uploadFileFromFileApi(file);
    MessagePlugin.success($t('common.uploadSuccess'));
    await fetchTableData();
  } catch (error) {
    if (import.meta.env.DEV) {
      logger.error(error);
    }
    MessagePlugin.error($t('common.uploadFailed'));
  } finally {
    uploadingFiles.value--;
  }
}

function getStorageModeText(mode?: number) {
  return mode === 1 ? $t('common.local') : $t('common.cloudStorage');
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
    <div class="flex h-full gap-3">
      <!-- Left Tree Slider -->
      <div class="w-48 flex-shrink-0 rounded-md bg-card p-2">
        <div class="mb-2 px-2 text-sm font-medium text-muted-foreground">{{ $t('system.upload.fileType') }}</div>
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
        <!-- Upload Area -->
        <div class="rounded-md bg-card p-4">
          <div class="flex items-center gap-4">
            <Upload
              :auto="false"
              :show-upload-progress="true"
              accept="*"
              multiple
              theme="file-input"
@select-files="(files: any[]) => {
                 files.forEach((f: any) => {
                   handleUpload(f.raw ?? f);
                 });
               }"
            >
              <template #trigger>
                <Button theme="primary">
                  <template #icon>
                    <UploadIcon />
                  </template>
                  {{ $t('common.uploadFile') }}
                </Button>
              </template>
            </Upload>
            <div v-if="uploadingFiles > 0" class="text-sm text-muted-foreground">
              {{ $t('common.uploadingCount', [uploadingFiles]) }}
            </div>
          </div>
        </div>

        <!-- Search Form -->
        <div class="rounded-md bg-card p-4">
          <Form :data="searchForm" label-width="80px" layout="inline" colon>
            <div class="grid grid-cols-4 gap-x-4 gap-y-3">
              <FormItem :label="$t('system.upload.fileName')" name="origin_name">
                <Input
                  v-model="searchForm.origin_name"
                  :placeholder="$t('ui.placeholder.input')"
                  clearable
                />
              </FormItem>
              <FormItem :label="$t('system.upload.storageMode')" name="storage_mode">
                <RadioGroup v-model="searchForm.storage_mode">
                  <Radio :value="1">{{ $t('common.local') }}</Radio>
                  <Radio :value="2">{{ $t('common.cloudStorage') }}</Radio>
                </RadioGroup>
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

        <!-- Table Area -->
        <div ref="tableContainerRef" class="flex min-h-0 flex-1 flex-col rounded-md bg-card p-4">
          <div class="mb-3 flex items-center justify-between">
            <Space>
                <Popconfirm
                  :content="$t('common.batchDeleteDataConfirm')"
                  @confirm="handleBatchDelete"
                >
                  <Button theme="danger" variant="outline">
                    <template #icon><DeleteIcon /></template>
                    {{ $t('common.delete') }}
                  </Button>
                </Popconfirm>
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
                v-model="displayColumns"
                :column-options="columnOptions"
                :is-recycle-bin="false"
                @refresh="fetchTableData"
              />
            </div>
          </div>

          <div class="min-h-0 flex-1 overflow-hidden">
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
              <template #storage_mode="{ row }">
                <Tag :theme="row?.storage_mode === 1 ? 'primary' : 'warning'">
                  {{ getStorageModeText(row?.storage_mode) }}
                </Tag>
              </template>

              <template #action="{ row }">
                <Button
                  size="small"
                  theme="primary"
                  variant="outline"
                  @click="handleDownload(row)"
                >
                  {{ $t('common.download') }}
                </Button>
              </template>
            </Table>
          </div>
        </div>
      </div>
    </div>
  </Page>
</template>

<style scoped>
.upload-area {
  border: 2px dashed hsl(var(--border));
  border-radius: 0.5rem;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s;
}

.upload-area:hover {
  border-color: hsl(var(--primary));
  background-color: hsl(var(--muted) / 0.3);
}
</style>
