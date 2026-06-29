import type { DictApi } from '#/api/system/dict';

import { useCrudPage } from '#/composables/crud/use-crud-page';

import { getDictTypePageList, getRecycleDictTypeList } from '#/api/system/dict';

import type { DictTypeListItem } from './model';
import { createDictTypeSearchForm } from './schemas';

export function useDictTypeCrud() {
  return useCrudPage<DictTypeListItem, ReturnType<typeof createDictTypeSearchForm>>({
    defaultSearchForm: createDictTypeSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleDictTypeList(params) : getDictTypePageList(params),
    buildParams: (form) => {
      const params: DictApi.DictTypeQuery = {};
      if (form.name) params.name = form.name;
      if (form.code) params.code = form.code;
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
