import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace AppApi {
  export interface ListItem {
    app_id?: string;
    app_secret?: string;
    created_at?: string;
    id: number;
    intro?: string;
    name: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
    updated_at?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    created_at?: string[];
    name?: string;
    status?: StatusValue;
  }

  export interface SubmitPayload {
    app_id?: string;
    id?: number;
    intro?: string;
    name: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
  }

  export interface BindApiPayload {
    api_ids: number[];
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export interface NumberOperationPayload {
    id: number;
    numberName: string;
    numberValue: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
}

export function getAppPageList(params: AppApi.ListQuery) {
  return requestClient.get<AppApi.ListResponse>('/system/app/index', {
    params,
  });
}

export function getRecycleAppList(params: AppApi.ListQuery) {
  return requestClient.get<AppApi.ListResponse>('/system/app/recycle', {
    params,
  });
}

export function getAppDetail(id: number) {
  return requestClient.get<AppApi.ListItem>(`/system/app/read/${id}`);
}

export function saveApp(data: AppApi.SubmitPayload) {
  return requestClient.post<void>('/system/app/save', data);
}

export function updateApp(id: number, data: AppApi.SubmitPayload) {
  return requestClient.put<void>(`/system/app/update/${id}`, data);
}

export function deleteApp(ids: number[]) {
  return requestClient.delete<void>('/system/app/delete', { data: { ids } });
}

export function realDeleteApp(ids: number[]) {
  return requestClient.delete<void>('/system/app/realDelete', {
    data: { ids },
  });
}

export function recoveryApp(ids: number[]) {
  return requestClient.put<void>('/system/app/recovery', { ids });
}

export function changeAppStatus(data: AppApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/app/changeStatus', data);
}

export function updateAppNumber(data: AppApi.NumberOperationPayload) {
  return requestClient.put<void>('/system/app/numberOperation', data);
}

export function getAppId() {
  return requestClient.get<{ app_id: string }>('/system/app/getAppId');
}

export function getAppSecret() {
  return requestClient.get<{ app_secret: string }>('/system/app/getAppSecret');
}

export function bindAppApi(id: number, data: AppApi.BindApiPayload) {
  return requestClient.put<void>(`/system/app/bind/${id}`, data);
}

export function getAppBindApiList(id: number) {
  return requestClient.get<number[]>('/system/app/getApiList', {
    params: { id },
  });
}
