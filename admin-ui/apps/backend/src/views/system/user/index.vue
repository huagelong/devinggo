<script lang="ts" setup>
import { computed, onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

import { getDeptTree } from '#/api/system/dept';
import { getPostList } from '#/api/system/post';
import { getRoleList } from '#/api/system/role';
import type { DeptApi } from '#/api/system/dept';
import type { PostApi } from '#/api/system/post';
import type { RoleApi } from '#/api/system/role';
import type { DictOption } from '#/composables/crud/use-dict-options';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import DeptTree from './components/dept-tree.vue';
import ImportDialog from './components/import-dialog.vue';
import SetHomePageDialog from './components/set-home-page-dialog.vue';
import UserActionToolbar from './components/user-action-toolbar.vue';
import UserModal from './components/user-modal.vue';
import UserSearchForm from './components/user-search-form.vue';
import UserTable from './components/user-table.vue';
import type {
  UserActionDropdownItem,
  UserListItem,
  UserTableColumn,
} from './model';
import { createUserColumnOptions, createUserTableColumns } from './schemas';
import { useUserActions } from './use-user-actions';
import { useUserCrud } from './use-user-crud';

const currentDeptId = ref<number | string>('');
type UserModalInstance = {
  open: (data?: Partial<UserListItem>) => void;
};

const userModalRef = ref<UserModalInstance>();
const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);
const importDialogVisible = ref(false);

const roleOptions = ref<RoleApi.ListItem[]>([]);
const postOptions = ref<PostApi.ListItem[]>([]);
const deptTreeData = ref<DeptApi.TreeNode[]>([]);
const statusOptions = ref<DictOption[]>([]);
const userTypeOptions = ref<DictOption[]>([]);
const homePageOptions = ref<DictOption[]>([]);

const columns: UserTableColumn[] = createUserTableColumns();
const columnOptions = createUserColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);

const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value.filter((item) => item !== 'row-select');
  },
});

const {
  buildRequestParams,
  clearSelectedRowKeys,
  fetchTableData,
  handleDeptSelect,
  handlePageChange,
  handleResetWithDept,
  handleSearch,
  handleSelectChange,
  isRecycleBin,
  loading,
  pagination,
  searchForm,
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = useUserCrud(currentDeptId);

const {
  exportLoading,
  handleActionDropdownClick,
  handleBatchDelete,
  handleBatchRecovery,
  handleClearCache,
  handleDelete,
  handleDownloadTemplate,
  handleExport,
  handleImportChange,
  handleRecovery,
  handleSetHomePage,
  handleStatusChange,
  importInputRef,
  importLoading,
  isSuperAdmin,
  selectedHomePage,
  setHomePageLoading,
  setHomePageVisible,
  templateLoading,
  triggerImport,
} = useUserActions({
  buildRequestParams,
  clearSelectedRowKeys,
  fetchTableData,
  isRecycleBin,
  selectedRowKeys,
});

void importInputRef;

const { getDictOptions } = useDictOptions();

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

function normalizeListData<T>(data: T[] | { items?: T[] } | null | undefined): T[] {
  if (Array.isArray(data)) {
    return data;
  }
  return Array.isArray(data?.items) ? data.items : [];
}

async function fetchOptions() {
  try {
    const [roleRes, postRes, deptRes, statusDict, userTypeDict, dashboardDict] = await Promise.all([
      getRoleList().catch(() => null),
      getPostList().catch(() => null),
      getDeptTree().catch(() => null),
      getDictOptions('data_status'),
      getDictOptions('user_type'),
      getDictOptions('dashboard'),
    ]);

    roleOptions.value = normalizeListData(roleRes);
    postOptions.value = normalizeListData(postRes);
    deptTreeData.value = deptRes || [];
    statusOptions.value = statusDict || [];
    userTypeOptions.value = userTypeDict || [];
    homePageOptions.value = dashboardDict || [];
  } catch (error) {
    logger.error(error);
    message.error($t('common.filterLoadFailed'));
  }
}

function handleAdd() {
  userModalRef.value?.open();
}

function openImportDialog() {
  importDialogVisible.value = true;
}

async function handleImportChangeWithClose(event: Event) {
  const success = await handleImportChange(event);
  if (success) {
    importDialogVisible.value = false;
  }
}

function handleEdit(row: UserListItem) {
  if (isSuperAdmin(row)) {
    message.warning($t('common.superAdminCannotEdit'));
    return;
  }
  userModalRef.value?.open(row);
}

function handleSuccess() {
  void fetchTableData();
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
}

function handleStatusSwitchChange(row: UserListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

function handleActionDropdownItemClick(
  item: unknown,
  row: UserListItem,
) {
  handleActionDropdownClick(item as UserActionDropdownItem, row);
}

onMounted(() => {
  void fetchOptions();
  void fetchTableData();
  document.addEventListener('fullscreenchange', handleFullscreenChange);
});

onUnmounted(() => {
  document.removeEventListener('fullscreenchange', handleFullscreenChange);
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-row gap-4">
      <div class="h-full rounded-md bg-background p-2">
        <DeptTree @select="handleDeptSelect" />
      </div>

      <div class="flex h-full min-w-0 flex-1 flex-col gap-3 overflow-hidden">
        <UserSearchForm
          :form-data="searchForm"
          :role-options="roleOptions"
          :post-options="postOptions"
          :dept-tree-data="deptTreeData"
          :status-options="statusOptions"
          :user-type-options="userTypeOptions"
          @search="handleSearch"
          @reset="handleResetWithDept"
          @update:form-data="(val) => Object.assign(searchForm, val)"
        />

        <div ref="tableContainerRef" class="flex min-h-0 flex-1 flex-col rounded-md bg-card p-4">
          <UserActionToolbar
            :is-recycle-bin="isRecycleBin"
            :import-loading="importLoading"
            :export-loading="exportLoading"
            :is-fullscreen="isFullscreen"
            :visible-columns="visibleColumns"
            :column-options="columnOptions"
            @add="handleAdd"
            @batch-delete="handleBatchDelete"
            @import="openImportDialog"
            @export="handleExport"
            @batch-recovery="handleBatchRecovery"
            @toggle-fullscreen="toggleFullscreen"
            @refresh="fetchTableData"
            @toggle-recycle="toggleRecycleBin"
            @update:visible-columns="(val) => (visibleColumns = val)"
          />

          <UserTable
            :columns="columns"
            :data="tableData"
            :loading="loading"
            :pagination="pagination"
            :selected-row-keys="selectedRowKeys"
            :display-columns="displayColumns"
            :is-recycle-bin="isRecycleBin"
            :is-super-admin="isSuperAdmin"
            @page-change="handlePageChange"
            @select-change="handleTableSelectChange"
            @edit="handleEdit"
            @delete="handleDelete"
            @status-change="handleStatusSwitchChange"
            @action-dropdown-click="handleActionDropdownItemClick"
            @clear-cache="handleClearCache"
            @recovery="handleRecovery"
            @update:display-columns="(val) => (displayColumns = val)"
          />
        </div>
      </div>
    </div>

    <input
      ref="importInputRef"
      type="file"
      accept=".xlsx,.xls,.csv"
      class="hidden"
      @change="handleImportChangeWithClose"
    />

    <ImportDialog
      v-model:visible="importDialogVisible"
      :import-loading="importLoading"
      :template-loading="templateLoading"
      @download-template="handleDownloadTemplate"
      @import="triggerImport"
      @cancel="importDialogVisible = false"
    />

    <SetHomePageDialog
      v-model:visible="setHomePageVisible"
      :selected-home-page="selectedHomePage"
      :home-page-options="homePageOptions"
      :loading="setHomePageLoading"
      @update:selected-home-page="(val) => (selectedHomePage = val ?? '')"
      @save="handleSetHomePage"
      @cancel="setHomePageVisible = false"
    />

    <UserModal ref="userModalRef" @success="handleSuccess" />
  </Page>
</template>