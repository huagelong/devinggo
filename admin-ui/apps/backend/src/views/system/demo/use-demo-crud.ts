import type { DemoListItem } from './model';
import type { DemoApi } from '#/api/system/demo';

import { getDemoPageList, getRecycleDemoList } from '#/api/system/demo';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { createDemoSearchForm } from './schemas';

export function useDemoCrud() {
  return useCrudPage<DemoListItem, ReturnType<typeof createDemoSearchForm>>({
    defaultSearchForm: createDemoSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleDemoList(params) : getDemoPageList(params),
    buildParams: (form) => {
      const params: Partial<DemoApi.ListQuery> = {};
      if (form.name) params.name = form.name;
      if (form.code) params.code = form.code;
      if (form.status !== undefined) params.status = form.status;
      if (form.birthday) params.birthday = form.birthday;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
