import type { NoticeListItem } from './model';
import type { NoticeApi } from '#/api/system/notice';

import { getNoticePageList, getRecycleNoticeList } from '#/api/system/notice';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { createNoticeSearchForm } from './schemas';

export function useNoticeCrud() {
  return useCrudPage<NoticeListItem, ReturnType<typeof createNoticeSearchForm>>({
    defaultSearchForm: createNoticeSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycleNoticeList(params) : getNoticePageList(params),
    buildParams: (form) => {
      const params: NoticeApi.ListQuery = {};
      if (form.title) params.title = form.title;
      if (form.type !== undefined && form.type !== null && form.type !== '') {
        const typeValue = Number(form.type);
        if (!Number.isNaN(typeValue)) {
          params.type = typeValue;
        }
      }
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params as Record<string, unknown>;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
