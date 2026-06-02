import type { MenuApi } from '#/api/system/menu';

import { reactive, ref } from 'vue';

import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { getMenuTreeList, getRecycleMenuTreeList } from '#/api/system/menu';
import { logger } from '#/utils/logger';

import { createMenuSearchForm } from './schemas';

export function useMenuPage() {
  const searchForm = reactive(createMenuSearchForm());
  const tableData = ref<MenuApi.TreeItem[]>([]);
  const loading = ref(false);
  const isRecycleBin = ref(false);
  const selectedRowKeys = ref<Array<number | string>>([]);
  let fetchRequestId = 0;

  function buildParams() {
    const params: Partial<MenuApi.ListQuery> = {};
    if (searchForm.name) params.name = searchForm.name;
    if (searchForm.code) params.code = searchForm.code;
    if (searchForm.level) params.level = searchForm.level;
    if (searchForm.status !== undefined) params.status = searchForm.status;
    if (searchForm.created_at?.length === 2 && searchForm.created_at[0]) {
      params.created_at = searchForm.created_at;
    }
    return params;
  }

  async function fetchTableData() {
    const requestId = ++fetchRequestId;
    loading.value = true;
    try {
      const params = buildParams();
      const result = isRecycleBin.value
        ? await getRecycleMenuTreeList(params)
        : await getMenuTreeList(params);
      if (requestId !== fetchRequestId) return;
      tableData.value = result;
    } catch (error) {
      if (requestId !== fetchRequestId) return;
      logger.error(error);
      message.error($t('common.listLoadFailed'));
    } finally {
      if (requestId === fetchRequestId) {
        loading.value = false;
      }
    }
  }

  function handleSelectChange(keys: Array<number | string>) {
    selectedRowKeys.value = keys;
  }

  function clearSelectedRowKeys() {
    selectedRowKeys.value = [];
  }

  function handleSearch() {
    void fetchTableData();
  }

  function handleReset() {
    Object.assign(searchForm, createMenuSearchForm());
    void fetchTableData();
  }

  function toggleRecycleBin(next?: boolean) {
    isRecycleBin.value = typeof next === 'boolean' ? next : !isRecycleBin.value;
    clearSelectedRowKeys();
    void fetchTableData();
  }

  return {
    clearSelectedRowKeys,
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
  };
}
