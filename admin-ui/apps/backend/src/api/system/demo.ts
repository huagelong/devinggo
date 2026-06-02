import { requestClient } from '#/api/request';
import type { BatchIdsPayload } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace DemoApi {
  export interface ListItem {
    id: number;
    name: string;
    code?: string;
    status: number;
    sort: number;
    price: number;
    cover: string;
    email: string;
    phone: string;
    birthday: string;
    remark?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    name?: string;
    code?: string;
    status?: number;
    birthday?: string;
    created_at?: string[];
  }

  export interface SubmitPayload {
    name: string;
    code: string;
    status: number;
    sort: number;
    price: number;
    cover: string;
    email: string;
    phone: string;
    birthday: string;
    remark: string;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
  export type OptionListResponse = ListItem[] | ListResponse;
}

export function getDemoPageList(params: DemoApi.ListQuery) {
  return requestClient.get<DemoApi.ListResponse>('/system/demo/index', { params });
}

export function getDemoList(params?: DemoApi.ListQuery) {
  return requestClient.get<DemoApi.OptionListResponse>('/system/demo/list', { params });
}

export function getRecycleDemoList(params: DemoApi.ListQuery) {
  return requestClient.get<DemoApi.ListResponse>('/system/demo/recycle', { params });
}

export function saveDemo(data: DemoApi.SubmitPayload) {
  return requestClient.post<void>('/system/demo/save', data);
}

export function updateDemo(id: number, data: DemoApi.SubmitPayload) {
  return requestClient.put<void>(`/system/demo/update/${id}`, data);
}

export function deleteDemo(ids: number[]) {
  return requestClient.delete<void>('/system/demo/delete', { data: { ids } });
}

export function realDeleteDemo(ids: number[]) {
  return requestClient.delete<void>('/system/demo/realDelete', { data: { ids } });
}

export function recoveryDemo(ids: number[]) {
  return requestClient.put<void>('/system/demo/recovery', { ids });
}

export function changeDemoStatus(data: DemoApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/demo/changeStatus', data);
}
