import type { Ref } from 'vue';
import type { UserApi } from '#/api/system/user';
import type { IdType } from '#/types/common';

import { useCrudPage } from '#/composables/crud/use-crud-page';

import { getRecycleUserList, getUserList } from '#/api/system/user';

import type { UserListItem } from './model';
import { createUserSearchForm } from './schemas';

export function useUserCrud(currentDeptId: Ref<IdType>) {
  const crud = useCrudPage<UserListItem, ReturnType<typeof createUserSearchForm>>({
    defaultSearchForm: createUserSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleUserList(params) : getUserList(params),
    buildParams: (form) => {
      const params: Partial<UserApi.ListQuery> = {};

      if (form.username) params.username = form.username;
      if (form.role_id !== undefined) params.role_id = form.role_id;
      if (form.phone) params.phone = form.phone;
      if (form.post_id !== undefined) params.post_id = form.post_id;
      if (form.email) params.email = form.email;
      if (form.status !== undefined) params.status = form.status;
      if (form.user_type) params.user_type = form.user_type;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }

      if (currentDeptId.value) {
        params.dept_id = currentDeptId.value;
      } else if (form.dept_ids?.length) {
        const ids = form.dept_ids
          .map((item) => Number(item))
          .filter((item) => !Number.isNaN(item));
        if (ids.length > 0) {
          params.dept_ids = ids;
        }
      }

      return params;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });

  function handleDeptSelect(deptId: IdType) {
    currentDeptId.value = deptId;
    crud.pagination.current = 1;
    void crud.fetchTableData();
  }

  function handleResetWithDept() {
    currentDeptId.value = '';
    crud.handleReset();
  }

  return {
    ...crud,
    handleDeptSelect,
    handleResetWithDept,
  };
}
