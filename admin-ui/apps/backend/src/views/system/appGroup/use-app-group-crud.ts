import type { AppGroupApi } from '#/api/system/app-group';

import { useCrudPage } from '#/composables/crud/use-crud-page';
import { getAppGroupPageList, getRecycleAppGroupList } from '#/api/system/app-group';

import type { AppGroupListItem } from './model';
import { createAppGroupSearchForm } from './schemas';

export function useAppGroupCrud() {
  return useCrudPage<
    AppGroupListItem,
    ReturnType<typeof createAppGroupSearchForm>
  >({
    defaultSearchForm: createAppGroupSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin
        ? getRecycleAppGroupList(params)
        : getAppGroupPageList(params),
    buildParams: (form) => {
      const params: AppGroupApi.ListQuery = {};
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
