import type { UploadListItem, UploadSearchFormModel } from './model';

import { reactive, ref } from 'vue';

import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';

import {
  deleteUploadFiles,
  downloadFileApi,
  getFileInfoApi,
  getUploadPageList,
} from '#/api/system/upload';
import { logger } from '#/utils/logger';

export function useUploadCrud() {
  const loading = ref(false);
  const tableData = ref<UploadListItem[]>([]);
  const selectedRowKeys = ref<Array<number | string>>([]);
  let fetchRequestId = 0;

  const searchForm = reactive<UploadSearchFormModel>({
    origin_name: '',
    mime_type: '',
    storage_mode: undefined,
    created_at: [],
  });

  const pagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
  });

  const fetchTableData = async () => {
    const requestId = ++fetchRequestId;
    loading.value = true;
    try {
      const response = await getUploadPageList({
        page: pagination.current,
        pageSize: pagination.pageSize,
        ...searchForm,
      });
      if (requestId !== fetchRequestId) return;
      tableData.value = response.items ?? [];
      pagination.total = Number(response.pageInfo?.total ?? response.total ?? 0);
    } catch (error) {
      if (requestId !== fetchRequestId) return;
      if (import.meta.env.DEV) {
        logger.error(error);
      }
      message.error($t('common.listLoadFailed'));
    } finally {
      if (requestId === fetchRequestId) {
        loading.value = false;
      }
    }
  };

  const handleSearch = () => {
    pagination.current = 1;
    fetchTableData();
  };

  const handleReset = () => {
    searchForm.origin_name = '';
    searchForm.mime_type = '';
    searchForm.storage_mode = undefined;
    searchForm.created_at = [];
    handleSearch();
  };

  const handlePageChange = (pageInfo: { current: number; pageSize: number }) => {
    pagination.current = pageInfo.current;
    pagination.pageSize = pageInfo.pageSize;
    fetchTableData();
  };

  const handleSelectChange = (val: Array<number | string>) => {
    selectedRowKeys.value = val;
  };

  const clearSelectedRowKeys = () => {
    selectedRowKeys.value = [];
  };

  const handleDownload = async (row: UploadListItem) => {
    try {
      const blob = await downloadFileApi({ id: row.id });
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.download = row.origin_name;
      link.click();
      window.URL.revokeObjectURL(url);
      message.success($t('common.downloadSuccess'));
    } catch (error) {
      logger.error(error);
      message.error($t('common.downloadFailed'));
    }
  };

  const handleView = async (row: UploadListItem) => {
    try {
      await getFileInfoApi({ id: row.id });
    } catch (error) {
      logger.error(error);
      message.error($t('common.listLoadFailed'));
    }
  };

  const handleBatchDelete = async () => {
    if (selectedRowKeys.value.length === 0) {
      message.warning($t('common.selectDataFirst'));
      return;
    }

    try {
      await deleteUploadFiles(selectedRowKeys.value.map((key) => Number(key)));
      message.success($t('common.operationSuccess'));
      clearSelectedRowKeys();
      await fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.batchDeleteFailed'));
    }
  };

  return {
    clearSelectedRowKeys,
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
    handleView,
  };
}
