import type { ApiGroupListItem } from './model';

import {
  getApiGroupPageList,
  getRecycleApiGroupList,
} from '#/api/system/api-group';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { createApiGroupSearchForm } from './schemas';

export function useApiGroupCrud() {
  return useCrudPage<ApiGroupListItem, ReturnType<typeof createApiGroupSearchForm>>({
    defaultSearchForm: createApiGroupSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleApiGroupList(params) : getApiGroupPageList(params),
    buildParams: (form) => {
      const params: Record<string, unknown> = {};
      if (form.name) params.name = form.name;
      if (typeof form.status === 'number') params.status = form.status;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params;
    },
  });
}
