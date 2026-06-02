<script lang="ts" setup>
import type { DataMaintainApi } from '#/api/system/data-maintain';

import { computed, onMounted, ref } from 'vue';

import { useAccess } from '@vben/access';
import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import {
  fragmentDataMaintainTable,
  optimizeDataMaintainTable,
} from '#/api/system/data-maintain';
import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import { logger } from '#/utils/logger';

import {
  InfoCircleFilledIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  Form,
  FormItem,
  Input,
  Popconfirm,
  Popup,
  Space,
  Table,
} from 'tdesign-vue-next';

import DataMaintainDetailPanel from './components/data-maintain-detail-panel.vue';

import type { DataMaintainListItem, DataMaintainTableColumn } from './model';
import {
  createDataMaintainColumnOptions,
  createDataMaintainTableColumns,
} from './schemas';
import { useDataMaintainCrud } from './use-data-maintain-crud';

defineOptions({ name: 'SystemDataMaintain' });

const { hasAccessByCodes } = useAccess();
const canView = computed(() =>
  hasAccessByCodes(['system:dataMaintain:index', 'system:dataMaintain']),
);
const canDetailed = computed(() =>
  hasAccessByCodes(['system:dataMaintain:detailed', 'system:dataMaintain:index']),
);
const canOptimize = computed(() =>
  hasAccessByCodes(['system:dataMaintain:optimize', 'system:dataMaintain:index']),
);
const canFragment = computed(() =>
  hasAccessByCodes(['system:dataMaintain:fragment', 'system:dataMaintain:index']),
);

// Current backend only exposes index; detailed/optimize/fragment are prepared for future rollout.
const hasDetailedApi = false;
const hasOptimizeApi = false;
const hasFragmentApi = false;

type DataMaintainDetailPanelInstance = {
  open: (options: {
    groupName?: string;
    hasDetailedApi: boolean;
    row: DataMaintainApi.ListItem;
  }) => Promise<void>;
};

const detailPanelRef = ref<DataMaintainDetailPanelInstance>();
const optimizingTableName = ref('');
const fragmentingTableName = ref('');

const columns: DataMaintainTableColumn[] = createDataMaintainTableColumns();
const columnOptions = createDataMaintainColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);

const displayColumns = computed({
  get: () => [...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value;
  },
});

const {
  fetchTableData,
  handlePageChange,
  handleReset,
  handleSearch,
  loading,
  pagination,
  searchForm,
  tableData,
} = useDataMaintainCrud();

function handleUnimplementedAction(actionName: string) {
  message.info($t('common.actionNotAvailable', [actionName]));
}

async function handleViewDetail(row: DataMaintainListItem) {
  await detailPanelRef.value?.open({
    groupName: searchForm.group_name,
    hasDetailedApi,
    row,
  });
}

async function handleOptimize(row: DataMaintainListItem) {
  if (!hasOptimizeApi) {
    handleUnimplementedAction($t('system.dataMaintain.optimizeTable'));
    return;
  }

  optimizingTableName.value = row.name;
  try {
    await optimizeDataMaintainTable({
      group_name: searchForm.group_name,
      table_name: row.name,
    });
    message.success($t('common.optimizeSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.optimizeFailed'));
  } finally {
    optimizingTableName.value = '';
  }
}

async function handleFragment(row: DataMaintainListItem) {
  if (!hasFragmentApi) {
    handleUnimplementedAction($t('system.dataMaintain.cleanFragment'));
    return;
  }

  fragmentingTableName.value = row.name;
  try {
    await fragmentDataMaintainTable({
      group_name: searchForm.group_name,
      table_name: row.name,
    });
    message.success($t('common.cleanSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    message.error($t('common.cleanFailed'));
  } finally {
    fragmentingTableName.value = '';
  }
}

onMounted(() => {
  if (!canView.value) {
    return;
  }
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="90px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.dataMaintain.dbGroup')" name="group_name">
              <Input
                v-model="searchForm.group_name"
                :placeholder="$t('system.dataMaintain.defaultGroup')"
                clearable
              />
            </FormItem>
            <FormItem :label="$t('system.dataMaintain.tableName')" name="name">
              <Input v-model="searchForm.name" :placeholder="$t('ui.placeholder.input')" clearable />
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
            <Button variant="outline" @click="handleUnimplementedAction($t('system.dataMaintain.viewFields'))">
              {{ $t('system.dataMaintain.viewFields') }}
            </Button>
            <Button variant="outline" @click="handleUnimplementedAction($t('system.dataMaintain.optimizeTable'))">
              {{ $t('system.dataMaintain.optimizeTable') }}
            </Button>
            <Button variant="outline" @click="handleUnimplementedAction($t('system.dataMaintain.cleanFragment'))">
              {{ $t('system.dataMaintain.cleanFragment') }}
            </Button>
            <Popup placement="bottom" trigger="hover" :content="$t('system.dataMaintain.firstVersionTooltip')">
              <InfoCircleFilledIcon class="cursor-help text-gray-400" />
            </Popup>
          </Space>

          <CrudToolbar
            v-model="displayColumns"
            :column-options="columnOptions"
            :is-recycle-bin="false"
            @refresh="fetchTableData"
          />
        </div>

        <div v-if="!canView" class="rounded-md border border-dashed border-gray-300 p-6 text-center text-gray-500">
          {{ $t('common.noPermission', ['system:dataMaintain:index']) }}
        </div>

        <div v-else class="min-h-0 flex-1">
          <Table
            row-key="name"
            hover
            stripe
            :columns="columns"
            :column-controller-visible="false"
            :display-columns="displayColumns"
            :data="tableData"
            :loading="loading"
            :pagination="pagination"
            @page-change="handlePageChange"
          >
            <template #rows="{ row }">
              {{ row.rows ?? '-' }}
            </template>
            <template #comment="{ row }">
              <span :title="row.comment || '-'">{{ row.comment || '-' }}</span>
            </template>
            <template #create_time="{ row }">
              {{ row.create_time || '-' }}
            </template>
            <template #action="{ row }">
              <Space>
                <Button
                  v-if="canDetailed"
                  size="small"
                  variant="text"
                  @click="handleViewDetail(row)"
                >
                  {{ $t('common.detail') }}
                </Button>
                <Popconfirm
                  v-if="canOptimize"
                  :content="$t('system.dataMaintain.confirmOptimize')"
                  placement="bottom"
                  @confirm="handleOptimize(row)"
                >
                  <Button
                    size="small"
                    theme="warning"
                    variant="text"
                    :loading="optimizingTableName === row.name"
                  >
                    {{ $t('system.dataMaintain.optimizeTable') }}
                  </Button>
                </Popconfirm>
                <Popconfirm
                  v-if="canFragment"
                  :content="$t('system.dataMaintain.confirmFragment')"
                  placement="bottom"
                  @confirm="handleFragment(row)"
                >
                  <Button
                    size="small"
                    theme="danger"
                    variant="text"
                    :loading="fragmentingTableName === row.name"
                  >
                    {{ $t('system.dataMaintain.fragmentClean') }}
                  </Button>
                </Popconfirm>
              </Space>
            </template>
          </Table>

          <DataMaintainDetailPanel ref="detailPanelRef" />
        </div>
      </div>
    </div>
  </Page>
</template>
