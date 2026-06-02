import type { AppApi } from '#/api/system/app';

import { useCrudPage } from '#/composables/crud/use-crud-page';
import { getAppPageList, getRecycleAppList } from '#/api/system/app';

import type { AppListItem } from './model';
import { createAppSearchForm } from './schemas';

export function useAppCrud() {
  return useCrudPage<AppListItem, ReturnType<typeof createAppSearchForm>>({
    defaultSearchForm: createAppSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleAppList(params) : getAppPageList(params),
    buildParams: (form) => {
      const params: AppApi.ListQuery = {};
      if (form.name) params.name = form.name;
      if (form.status !== undefined) params.status = form.status;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params as Record<string, unknown>;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
