import type { GenerateApi } from '#/api/system/generate';

import { useCrudPage } from '#/composables/crud/use-crud-page';
import { getCodePageList } from '#/api/system/generate';

import type { CodeListItem } from './model';
import { createCodeSearchForm } from './schemas';

export function useCodeCrud() {
  return useCrudPage<
    CodeListItem,
    ReturnType<typeof createCodeSearchForm>
  >({
    defaultSearchForm: createCodeSearchForm,
    fetchList: (params) => getCodePageList(params as GenerateApi.ListQuery),
    buildParams: (form) => {
      const params: GenerateApi.ListQuery = {};
      if (form.table_name) params.table_name = form.table_name;
      if (form.type) params.type = form.type as GenerateApi.GenerateType;
      return params as Record<string, unknown>;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
