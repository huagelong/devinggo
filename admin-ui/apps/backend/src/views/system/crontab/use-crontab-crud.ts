import type { CrontabApi } from '#/api/system/crontab';

import { useCrudPage } from '#/composables/crud/use-crud-page';
import {
  getCrontabPageList,
  getRecycleCrontabList,
} from '#/api/system/crontab';

import type { CrontabListItem } from './model';
import { createCrontabSearchForm } from './schemas';

export function useCrontabCrud() {
  return useCrudPage<
    CrontabListItem,
    ReturnType<typeof createCrontabSearchForm>
  >({
    defaultSearchForm: createCrontabSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin
        ? getRecycleCrontabList(params)
        : getCrontabPageList(params),
    buildParams: (form) => {
      const params: CrontabApi.ListQuery = {};
      if (form.name) params.name = form.name;
      if (form.type !== undefined) params.type = form.type;
      if (form.is_finally !== undefined) params.is_finally = form.is_finally;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params as Record<string, unknown>;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
