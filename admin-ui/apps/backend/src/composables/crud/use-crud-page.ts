import type { IdType } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

import { reactive, ref } from 'vue';

import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

type RowKey = IdType;
type CrudRequestParams = Partial<PageQuery> & Record<string, unknown>;

interface CrudPageContext {
  isRecycleBin: boolean;
  pagination: {
    current: number;
    pageSize: number;
    total: number;
  };
}

interface UseCrudPageOptions<
  TItem,
  TSearchForm extends object,
  TResponse extends PageResponse<TItem> = PageResponse<TItem>,
> {
  buildParams?: (
    searchForm: TSearchForm,
    context: CrudPageContext,
  ) => Record<string, unknown>;
  defaultSearchForm: () => TSearchForm;
  errorMessage?: string;
  fetchList: (
    params: CrudRequestParams,
    context: CrudPageContext,
  ) => Promise<TResponse>;
  onFetchError?: (error: unknown) => void;
  pageSize?: number;
  resolveItems?: (response: TResponse) => TItem[];
  resolveTotal?: (response: TResponse) => number;
}

export function useCrudPage<
  TItem = Record<string, unknown>,
  TSearchForm extends object = Record<string, unknown>,
  TResponse extends PageResponse<TItem> = PageResponse<TItem>,
>(options: UseCrudPageOptions<TItem, TSearchForm, TResponse>) {
  const searchForm = reactive<TSearchForm>(options.defaultSearchForm());
  const tableData = ref<TItem[]>([]);
  const loading = ref(false);
  const selectedRowKeys = ref<RowKey[]>([]);
  const isRecycleBin = ref(false);
  let fetchRequestId = 0;

  const pagination = reactive({
    current: 1,
    pageSize: options.pageSize ?? 20,
    pageSizeOptions: [10, 20, 50, 100],
    showJumper: true,
    showPageSize: true,
    total: 0,
  });

  const resolveItems =
    options.resolveItems ??
    ((response: TResponse) =>
      Array.isArray(response?.items) ? response.items : []);
  const resolveTotal =
    options.resolveTotal ??
    ((response: TResponse) =>
      Number(response?.pageInfo?.total ?? response?.total ?? 0));

  function getContext(): CrudPageContext {
    return {
      isRecycleBin: isRecycleBin.value,
      pagination: {
        current: pagination.current,
        pageSize: pagination.pageSize,
        total: pagination.total,
      },
    };
  }

  function buildRequestParams(includePagination = true): CrudRequestParams {
    const context = getContext();
    const businessParams = options.buildParams
      ? options.buildParams(searchForm as TSearchForm, context)
      : { ...(searchForm as Record<string, unknown>) };
    if (!includePagination) {
      return businessParams;
    }
    return {
      page: pagination.current,
      pageSize: pagination.pageSize,
      ...businessParams,
    };
  }

  function resetSearchForm() {
    Object.assign(searchForm, options.defaultSearchForm());
  }

  async function fetchTableData() {
    const requestId = ++fetchRequestId;
    loading.value = true;
    try {
      const response = await options.fetchList(buildRequestParams(true), getContext());
      if (requestId !== fetchRequestId) return;
      tableData.value = resolveItems(response);
      pagination.total = resolveTotal(response);
    } catch (error) {
      if (requestId !== fetchRequestId) return;
      logger.error(error);
      message.error(options.errorMessage ?? $t('common.listLoadFailed'));
      options.onFetchError?.(error);
    } finally {
      if (requestId === fetchRequestId) {
        loading.value = false;
      }
    }
  }

  function clearSelectedRowKeys() {
    selectedRowKeys.value = [];
  }

  function handleSelectChange(keys: RowKey[]) {
    selectedRowKeys.value = keys;
  }

  function handleSearch() {
    pagination.current = 1;
    void fetchTableData();
  }

  function handleReset() {
    resetSearchForm();
    pagination.current = 1;
    void fetchTableData();
  }

  function handlePageChange(pageInfo: { current: number; pageSize: number }) {
    pagination.current = pageInfo.current;
    pagination.pageSize = pageInfo.pageSize;
    clearSelectedRowKeys();
    void fetchTableData();
  }

  function toggleRecycleBin(next?: boolean) {
    isRecycleBin.value = typeof next === 'boolean' ? next : !isRecycleBin.value;
    clearSelectedRowKeys();
    pagination.current = 1;
    void fetchTableData();
  }

  return {
    buildRequestParams,
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
  };
}
