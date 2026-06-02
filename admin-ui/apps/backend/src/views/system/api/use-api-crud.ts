import type { ApiListItem } from './model';

import { getApiPageList, getRecycleApiList } from '#/api/system/api';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { createApiSearchForm } from './schemas';

export function useApiCrud() {
  return useCrudPage<ApiListItem, ReturnType<typeof createApiSearchForm>>({
    defaultSearchForm: createApiSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleApiList(params) : getApiPageList(params),
    buildParams: (form) => {
      const params: Record<string, unknown> = {};
      if (form.group_id) params.group_id = form.group_id;
      if (form.name) params.name = form.name;
      if (form.access_name) params.access_name = form.access_name;
      if (form.request_mode) params.request_mode = form.request_mode;
      if (typeof form.status === 'number') params.status = form.status;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params;
    },
  });
}
