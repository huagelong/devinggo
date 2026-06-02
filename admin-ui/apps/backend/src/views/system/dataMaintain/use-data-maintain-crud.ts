import type { DataMaintainApi } from '#/api/system/data-maintain';

import { getDataMaintainPageList } from '#/api/system/data-maintain';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import type { DataMaintainListItem } from './model';
import { createDataMaintainSearchForm } from './schemas';

export function useDataMaintainCrud() {
  return useCrudPage<
    DataMaintainListItem,
    ReturnType<typeof createDataMaintainSearchForm>
  >({
    defaultSearchForm: createDataMaintainSearchForm,
    fetchList: (params) => getDataMaintainPageList(params),
    buildParams: (form) => {
      const params: DataMaintainApi.ListQuery = {};
      if (form.group_name) params.group_name = form.group_name;
      if (form.name) params.name = form.name;
      return params as Record<string, unknown>;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
