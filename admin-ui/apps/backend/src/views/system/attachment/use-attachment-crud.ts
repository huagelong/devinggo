import type { AttachmentApi } from '#/api/system/attachment';

import { useCrudPage } from '#/composables/crud/use-crud-page';
import {
  getAttachmentPageList,
  getRecycleAttachmentList,
} from '#/api/system/attachment';

import type { AttachmentListItem } from './model';
import { createAttachmentSearchForm } from './schemas';

export function useAttachmentCrud() {
  return useCrudPage<
    AttachmentListItem,
    ReturnType<typeof createAttachmentSearchForm>
  >({
    defaultSearchForm: createAttachmentSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin
        ? getRecycleAttachmentList(params)
        : getAttachmentPageList(params),
    buildParams: (form) => {
      const params: AttachmentApi.ListQuery = {};
      if (form.origin_name) params.origin_name = form.origin_name;
      if (form.mime_type) params.mime_type = form.mime_type;
      if (form.storage_mode !== undefined) params.storage_mode = form.storage_mode;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params as Record<string, unknown>;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
