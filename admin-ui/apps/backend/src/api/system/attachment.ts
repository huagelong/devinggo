import { requestClient } from '#/api/request';
import type { BatchIdsPayload } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace AttachmentApi {
  export interface ListItem {
    id: number;
    object_name: string;
    origin_name: string;
    storage_mode: number;
    mime_type: string;
    storage_path: string;
    size_info: string;
    url: string;
    created_at?: string;
    updated_at?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    mime_type?: string;
    origin_name?: string;
    storage_mode?: number;
    created_at?: string[];
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
}

export function getAttachmentPageList(params: AttachmentApi.ListQuery) {
  return requestClient.get<AttachmentApi.ListResponse>(
    '/system/attachment/index',
    { params },
  );
}

export function getRecycleAttachmentList(params: AttachmentApi.ListQuery) {
  return requestClient.get<AttachmentApi.ListResponse>(
    '/system/attachment/recycle',
    { params },
  );
}

export function deleteAttachment(ids: number[]) {
  return requestClient.delete<void>('/system/attachment/delete', {
    data: { ids },
  });
}

export function realDeleteAttachment(ids: number[]) {
  return requestClient.delete<void>('/system/attachment/realDelete', {
    data: { ids },
  });
}

export function recoveryAttachment(ids: number[]) {
  return requestClient.put<void>('/system/attachment/recovery', { ids });
}
