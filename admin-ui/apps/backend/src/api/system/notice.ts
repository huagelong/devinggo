import { requestClient } from '#/api/request';
import type { BatchIdsPayload } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace NoticeApi {
  export interface ListItem {
    content: string;
    created_at?: string;
    id: number;
    remark?: string;
    title: string;
    type: number;
    users?: number[];
  }

  export interface ListQuery extends Partial<PageQuery> {
    created_at?: string[];
    title?: string;
    type?: number;
  }

  export interface SubmitPayload {
    content: string;
    id?: number;
    remark?: string;
    title: string;
    type: number;
    users?: number[];
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
}

export function getNoticePageList(params: NoticeApi.ListQuery) {
  return requestClient.get<NoticeApi.ListResponse>('/system/notice/index', {
    params,
  });
}

export function getRecycleNoticeList(params: NoticeApi.ListQuery) {
  return requestClient.get<NoticeApi.ListResponse>('/system/notice/recycle', {
    params,
  });
}

export function saveNotice(data: NoticeApi.SubmitPayload) {
  return requestClient.post<void>('/system/notice/save', data);
}

export function updateNotice(id: number, data: NoticeApi.SubmitPayload) {
  return requestClient.put<void>(`/system/notice/update/${id}`, data);
}

export function deleteNotice(ids: number[]) {
  return requestClient.delete<void>('/system/notice/delete', { data: { ids } });
}

export function realDeleteNotice(ids: number[]) {
  return requestClient.delete<void>('/system/notice/realDelete', {
    data: { ids },
  });
}

export function recoveryNotice(ids: number[]) {
  return requestClient.put<void>('/system/notice/recovery', { ids });
}
