import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace PostApi {
  export interface ListItem {
    code: string;
    id: number;
    name: string;
    remark?: string;
    sort?: number;
    status?: number;
  }

  export interface ListQuery extends Partial<PageQuery> {
    code?: string;
    created_at?: string[];
    name?: string;
    status?: StatusValue;
  }

  export interface SubmitPayload {
    code: string;
    id?: number;
    name: string;
    remark?: string;
    sort: number;
    status: number;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export interface SortPayload {
    id: number;
    numberName: string;
    numberValue: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
  export type OptionListResponse = ListItem[] | ListResponse;
}

export function getPostPageList(params: PostApi.ListQuery) {
  return requestClient.get<PostApi.ListResponse>('/system/post/index', { params });
}

export function getPostList(params?: PostApi.ListQuery) {
  return requestClient.get<PostApi.OptionListResponse>('/system/post/list', { params });
}

export function getRecyclePostList(params: PostApi.ListQuery) {
  return requestClient.get<PostApi.ListResponse>('/system/post/recycle', { params });
}

export function savePost(data: PostApi.SubmitPayload) {
  return requestClient.post<void>('/system/post/save', data);
}

export function updatePost(id: number, data: PostApi.SubmitPayload) {
  return requestClient.put<void>(`/system/post/update/${id}`, data);
}

export function deletePost(ids: number[]) {
  return requestClient.delete<void>('/system/post/delete', { data: { ids } });
}

export function realDeletePost(ids: number[]) {
  return requestClient.delete<void>('/system/post/realDelete', { data: { ids } });
}

export function recoveryPost(ids: number[]) {
  return requestClient.put<void>('/system/post/recovery', { ids });
}

export function changePostStatus(data: PostApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/post/changeStatus', data);
}

export function updatePostSort(data: PostApi.SortPayload) {
  return requestClient.put<void>('/system/post/numberOperation', data);
}
