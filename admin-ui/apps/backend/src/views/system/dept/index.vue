<script lang="ts" setup>
import type { DictOption } from '#/composables/crud/use-dict-options';

import { computed, onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';
import {
  changeDeptStatus,
  deleteDept,
  realDeleteDept,
  recoveryDept,
  updateDeptNumber,
} from '#/api/system/dept';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import { DeleteIcon, EditIcon, PlusIcon, SearchIcon, UserIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  EnhancedTable as Table,
  Form,
  FormItem,
  Input,
  InputNumber,
  Popconfirm,
  Select,
  Space,
  Switch,
} from 'tdesign-vue-next';

import DeptLeaderModal from './components/dept-leader-modal.vue';
import DeptModal from './components/dept-modal.vue';
import type { DeptTableColumn, DeptTreeItem } from './model';
import {
  createDeptColumnOptions,
  createDeptTableColumns,
} from './schemas';
import { useDeptPage } from './use-dept-page';

defineOptions({ name: 'SystemDept' });

type DeptModalInstance = {
  open: (data?: Record<string, unknown>) => void;
};

type DeptLeaderModalInstance = {
  open: (data: { id: number; name?: string }) => void;
};

const deptModalRef = ref<DeptModalInstance>();
const deptLeaderModalRef = ref<DeptLeaderModalInstance>();
const statusOptions = ref<DictOption[]>([]);

const columns: DeptTableColumn[] = createDeptTableColumns();
const columnOptions = createDeptColumnOptions(columns);
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
  expandedTreeNodes,
  fetchTableData,
  handleReset,
  handleSearch,
  handleSelectChange,
  isRecycleBin,
  loading,
  searchForm,
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = useDeptPage();

const { getDictOptions } = useDictOptions();

function handleExpandedTreeNodesChange(value: Array<number | string>) {
  expandedTreeNodes.value = value;
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
    statusOptions.value =
      options.length > 0
        ? options
        : [
            { label: $t('common.statusEnabled'), value: 1 },
            { label: $t('common.statusDisabled'), value: 2 },
          ];
}

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

function handleAdd(parentId = 0) {
  deptModalRef.value?.open({ parent_id: parentId });
}

function handleEdit(row: DeptTreeItem) {
  deptModalRef.value?.open({
    ...row,
    status: Number(row.status ?? 1),
  });
}

function handleOpenLeaderList(row: DeptTreeItem) {
  deptLeaderModalRef.value?.open({ id: row.id, name: row.name });
}

async function handleDelete(row: DeptTreeItem) {
  try {
    await (isRecycleBin.value ? realDeleteDept([row.id]) : deleteDept([row.id]));
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
    await (isRecycleBin.value ? realDeleteDept(ids) : deleteDept(ids));
    message.success($t('common.operationSuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

async function handleRecovery(row: DeptTreeItem) {
  try {
    await recoveryDept([row.id]);
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
    await recoveryDept(ids);
    message.success($t('common.recoverySuccess'));
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchRecoveryFailed'));
  }
}

async function handleStatusChange(row: DeptTreeItem, checked: boolean) {
  try {
    await changeDeptStatus({ id: row.id, status: checked ? 1 : 2 });
    message.success($t('common.statusUpdateSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.statusUpdateFailed'));
  }
}

async function handleSortChange(value: number | string, row: DeptTreeItem) {
  const numberValue = Number(value);
  if (Number.isNaN(numberValue)) return;

  try {
    await updateDeptNumber({
      id: row.id,
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

function handleSuccess() {
  void fetchTableData();
}

function handleStatusSwitchChange(row: DeptTreeItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

onMounted(() => {
  void fetchStatusOptions();
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="80px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.dept.name')" name="name">
              <Input
                v-model="searchForm.name"
                :placeholder="$t('ui.placeholder.input', [$t('system.dept.name')])"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.dept.leader')" name="leader">
              <Input
                v-model="searchForm.leader"
                :placeholder="$t('ui.placeholder.input', [$t('system.dept.leader')])"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.user.phone')" name="phone">
              <Input
                v-model="searchForm.phone"
                :placeholder="$t('ui.placeholder.input', [$t('system.user.phone')])"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('common.status')" name="status">
              <Select
                v-model="searchForm.status"
                :options="statusOptions"
                :placeholder="$t('ui.placeholder.select', [$t('common.status')])"
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
              <Button theme="primary" @click="handleAdd()">
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
          :expanded-tree-nodes="expandedTreeNodes"
          :loading="loading"
          :selected-row-keys="selectedRowKeys"
          :tree="{
            childrenKey: 'children',
            defaultExpandAll: true,
            treeNodeColumnIndex: 1,
          }"
          row-key="id"
          hover
          stripe
          @expanded-tree-nodes-change="handleExpandedTreeNodesChange"
          @select-change="handleSelectChange"
          >
          <template #sort="{ row }">
            <InputNumber
              :value="row?.sort"
              :min="0"
              :max="1000"
              size="small"
              @change="(value: number | string) => handleSortChange(value, row)"
            />
          </template>

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
                  theme="default"
                  variant="outline"
                  @click="handleOpenLeaderList(row)"
                >
                  <template #icon><UserIcon /></template>
                  {{ $t('system.dept.leaderList') }}
                </Button>
                <Button
                  size="small"
                  theme="primary"
                  variant="outline"
                  @click="handleAdd(row.id)"
                >
                  <template #icon><PlusIcon /></template>
                  {{ $t('common.create') }}
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
                  :content="$t('system.dept.confirmDelete')"
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
                  :content="$t('common.confirmRecoveryDept')"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    {{ $t('common.recovery') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  :content="$t('common.confirmPermanentDeleteDept')"
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

    <DeptModal ref="deptModalRef" @success="handleSuccess" />
    <DeptLeaderModal ref="deptLeaderModalRef" />
  </Page>
</template>
