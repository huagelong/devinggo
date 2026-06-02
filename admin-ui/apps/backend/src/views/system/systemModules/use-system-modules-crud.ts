import type { SystemModulesApi } from '#/api/system/system-modules';

import { useCrudPage } from '#/composables/crud/use-crud-page';
import {
  getRecycleSystemModulesList,
  getSystemModulesPageList,
} from '#/api/system/system-modules';

import type { SystemModulesListItem } from './model';
import { createSystemModulesSearchForm } from './schemas';

export function useSystemModulesCrud() {
  return useCrudPage<
    SystemModulesListItem,
    ReturnType<typeof createSystemModulesSearchForm>
  >({
    defaultSearchForm: createSystemModulesSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin
        ? getRecycleSystemModulesList(params)
        : getSystemModulesPageList(params),
    buildParams: (form) => {
      const params: SystemModulesApi.ListQuery = {};
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
