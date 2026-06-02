<script lang="ts" setup>
import type { CodeListItem } from './model';

import { computed, onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';
import { deleteCode, generateCode, syncTable } from '#/api/system/generate';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { downloadResponseBlob } from '#/utils/download';

import {
  BrowseIcon,
  CodeIcon,
  DeleteIcon,
  EditIcon,
  ExportIcon,
  RefreshIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  Form,
  FormItem,
  Input,
  Popconfirm,
  Select,
  Space,
  Table,
  Tag,
} from 'tdesign-vue-next';

import EditInfo from './components/edit-info.vue';
import LoadTable from './components/load-table.vue';
import Preview from './components/preview.vue';
import type { CodeTableColumn } from './model';
import {
  createCodeColumnOptions,
  createCodeSearchForm,
  createCodeTableColumns,
  generateTypeOptions,
} from './schemas';
import { useCodeCrud } from './use-code-crud';

defineOptions({ name: 'SystemCode' });

type LoadTableInstance = {
  open: () => void;
};

type PreviewInstance = {
  open: (id: number) => void;
};

type EditInfoInstance = {
  open: (id: number) => void;
};

const loadTableRef = ref<LoadTableInstance>();
const previewRef = ref<PreviewInstance>();
const editInfoRef = ref<EditInfoInstance>();

const searchForm = ref(createCodeSearchForm());

const columns: CodeTableColumn[] = createCodeTableColumns();
const columnOptions = createCodeColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);

const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value.filter((item) => item !== 'row-select');
  },
});

const selectedTables = ref<number[]>([]);

const {
  fetchTableData,
  handleReset,
  handleSearch,
  handleSelectChange,
  loading,
  pagination,
  selectedRowKeys,
  tableData,
} = useCodeCrud();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key)).filter((id) => !Number.isNaN(id));
}

function handleOpenLoadTable() {
  loadTableRef.value?.open();
}

function handleOpenPreview(row: CodeListItem) {
  previewRef.value?.open(row.id);
}

function handleOpenEdit(row: CodeListItem) {
  editInfoRef.value?.open(row.id);
}

async function handleSync(row: CodeListItem) {
  try {
    await syncTable(row.id);
    message.success($t('common.syncSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.syncFailed'));
  }
}

async function handleGenerate(row: CodeListItem) {
  try {
    message.info($t('common.generating'));
    const response = await generateCode({ ids: String(row.id) });
    downloadResponseBlob({ data: response as unknown as Blob }, `code_${row.table_name}.zip`);
    message.success($t('common.generateSuccess'));
  } catch (error) {
    logger.error(error);
    message.error($t('common.generateFailed'));
  }
}

async function handleBatchGenerate() {
  if (selectedTables.value.length === 0) {
    message.warning($t('common.selectGenerateData'));
    return;
  }

  try {
    message.info($t('common.generating'));
    const response = await generateCode({
      ids: selectedTables.value.join(','),
    });
    downloadResponseBlob({ data: response as unknown as Blob }, `code_batch_${Date.now()}.zip`);
    message.success($t('common.generateSuccess'));
  } catch (error) {
    logger.error(error);
    message.error($t('common.generateFailed'));
  }
}

async function handleDelete(row: CodeListItem) {
  try {
    await deleteCode([row.id]);
    message.success($t('common.deleteSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.deleteFailed'));
  }
}

function handleSelectChangeFn(keys: Array<number | string>) {
  handleSelectChange(keys);
  selectedTables.value = toIds(keys);
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
            <FormItem :label="$t('system.code.tableName')" name="table_name">
              <Input
                v-model="searchForm.table_name"
                :placeholder="$t('ui.placeholder.input')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.code.genType')" name="type">
              <Select
                v-model="searchForm.type"
                :options="generateTypeOptions"
                :placeholder="$t('ui.placeholder.select')"
                clearable
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
            <Button theme="primary" @click="handleOpenLoadTable">
              <template #icon><ExportIcon /></template>
              {{ $t('system.code.loadTable') }}
            </Button>
            <Button
              theme="primary"
              variant="outline"
              :disabled="selectedTables.length === 0"
              @click="handleBatchGenerate"
            >
              <template #icon><CodeIcon /></template>
              {{ $t('system.code.batchGenerate') }}
            </Button>
          </Space>

          <CrudToolbar
            v-model="visibleColumns"
            :column-options="columnOptions"
            @refresh="fetchTableData"
          />
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
          @select-change="handleSelectChangeFn"
        >
          <template #type="{ row }">
            <Tag :theme="row?.type === 'single' ? 'primary' : 'warning'">
              {{ row?.type === 'single' ? $t('system.code.singleCrud') : $t('system.code.treeCrud') }}
            </Tag>
          </template>

          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <Button
                size="small"
                theme="default"
                variant="outline"
                @click="handleOpenPreview(row)"
              >
                <template #icon><BrowseIcon /></template>
                {{ $t('common.preview') }}
              </Button>
              <Button
                size="small"
                theme="default"
                variant="outline"
                @click="handleSync(row)"
              >
                <template #icon><RefreshIcon /></template>
                {{ $t('common.sync') }}
              </Button>
              <Button
                size="small"
                theme="primary"
                variant="outline"
                @click="handleOpenEdit(row)"
              >
                <template #icon><EditIcon /></template>
                {{ $t('common.edit') }}
              </Button>
              <Button
                size="small"
                theme="primary"
                variant="outline"
                @click="handleGenerate(row)"
              >
                <template #icon><CodeIcon /></template>
                {{ $t('common.generate') }}
              </Button>
              <Popconfirm
                :content="$t('system.code.confirmDelete')"
                @confirm="handleDelete(row)"
              >
                <Button size="small" theme="danger" variant="outline">
                  <template #icon><DeleteIcon /></template>
                  {{ $t('common.delete') }}
                </Button>
              </Popconfirm>
            </div>
          </template>
        </Table>
      </div>
    </div>

    <LoadTable ref="loadTableRef" @success="handleSuccess" />
    <Preview ref="previewRef" />
    <EditInfo ref="editInfoRef" @success="handleSuccess" />
  </Page>
</template>
