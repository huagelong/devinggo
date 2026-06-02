import type { DeptApi } from '#/api/system/dept';

import { reactive, ref } from 'vue';

import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { getDeptPageList, getRecycleDeptList } from '#/api/system/dept';
import { logger } from '#/utils/logger';

import { createDeptSearchForm } from './schemas';

export function useDeptPage() {
  const searchForm = reactive(createDeptSearchForm());
  const tableData = ref<DeptApi.ListTreeItem[]>([]);
  const expandedTreeNodes = ref<Array<number | string>>([]);
  const loading = ref(false);
  const selectedRowKeys = ref<Array<number | string>>([]);
  const isRecycleBin = ref(false);
  let fetchRequestId = 0;

  function normalizeDeptTree(items: DeptApi.ListTreeItem[] = []): DeptApi.ListTreeItem[] {
    return items.map((item) => {
      const children = Array.isArray(item.children)
        ? normalizeDeptTree(item.children)
        : undefined;
      return { ...item, children };
    });
  }

  function collectExpandedTreeNodes(items: DeptApi.ListTreeItem[] = []): Array<number | string> {
    const nodes: Array<number | string> = [];
    for (const item of items) {
      nodes.push(item.id);
      if (Array.isArray(item.children) && item.children.length > 0) {
        nodes.push(...collectExpandedTreeNodes(item.children));
      }
    }
    return nodes;
  }

  function buildParams() {
    const params: Partial<DeptApi.ListQuery> = {};
    if (searchForm.name) params.name = searchForm.name;
    if (searchForm.leader) params.leader = searchForm.leader;
    if (searchForm.phone) params.phone = searchForm.phone;
    if (searchForm.status !== undefined) params.status = searchForm.status;
    if (searchForm.created_at?.length === 2 && searchForm.created_at[0]) {
      params.created_at = searchForm.created_at;
    }
    return params;
  }

  function clearSelectedRowKeys() {
    selectedRowKeys.value = [];
  }

  function handleSelectChange(keys: Array<number | string>) {
    selectedRowKeys.value = keys;
  }

  function resetSearchForm() {
    Object.assign(searchForm, createDeptSearchForm());
  }

  async function fetchTableData() {
    const requestId = ++fetchRequestId;
    loading.value = true;
    try {
      const params = buildParams();
      const result = isRecycleBin.value
        ? await getRecycleDeptList(params)
        : await getDeptPageList(params);
      if (requestId !== fetchRequestId) return;
      const treeData = normalizeDeptTree(result);
      tableData.value = treeData;
      expandedTreeNodes.value = collectExpandedTreeNodes(treeData);
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

  function handleSearch() {
    void fetchTableData();
  }

  function handleReset() {
    resetSearchForm();
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
    expandedTreeNodes,
    searchForm,
    selectedRowKeys,
    tableData,
    toggleRecycleBin,
  };
}
