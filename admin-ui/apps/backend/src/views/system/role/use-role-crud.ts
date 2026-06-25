import type { RoleApi } from '#/api/system/role';

import { getRecycleRoleList, getRolePageList } from '#/api/system/role';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import type { RoleListItem } from './model';
import { createRoleSearchForm } from './schemas';

export function useRoleCrud() {
  return useCrudPage<RoleListItem, ReturnType<typeof createRoleSearchForm>>({
    defaultSearchForm: createRoleSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleRoleList(params) : getRolePageList(params),
    buildParams: (form) => {
      const params: Partial<RoleApi.ListQuery> = {};
      if (form.name) params.name = form.name;
      if (form.code) params.code = form.code;
      if (form.status !== undefined) params.status = form.status;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
