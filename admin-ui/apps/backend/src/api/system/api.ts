import { requestClient } from '#/api/request';
import type { BatchIdsPayload, IdType } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace ApiManageApi {
  export interface ListItem {
    auth_mode?: number;
    created_at?: string;
    group_id: IdType;
    group_name?: string;
    id: number;
    name: string;
    access_name: string;
    remark?: string;
    request_mode: number | string;
    status: number;
  }

  export interface ListQuery extends Partial<PageQuery> {
    group_id?: IdType;
    name?: string;
    access_name?: string;
    request_mode?: number | string;
    status?: number;
    created_at?: string[];
  }

  export interface SubmitPayload {
    auth_mode: number;
    group_id: IdType;
    name: string;
    access_name: string;
    remark?: string;
    request_mode: number | string;
    status: number;
    id?: number;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export interface ModuleItem {
    id: IdType;
    name: string;
  }

  export type ListResponse = PageResponse<ListItem>;
  export type BatchPayload = BatchIdsPayload<number>;
}

export function getApiPageList(params: ApiManageApi.ListQuery) {
  return requestClient.get<ApiManageApi.ListResponse>('/system/api/index', {
    params,
  });
}

export function getRecycleApiList(params: ApiManageApi.ListQuery) {
  return requestClient.get<ApiManageApi.ListResponse>('/system/api/recycle', {
    params,
  });
}

export function getApiModuleList() {
  return requestClient.get<ApiManageApi.ModuleItem[]>('/system/api/getModuleList');
}

export function saveApi(data: ApiManageApi.SubmitPayload) {
  return requestClient.post<void>('/system/api/save', data);
}

export function updateApi(id: number, data: ApiManageApi.SubmitPayload) {
  return requestClient.put<void>(`/system/api/update/${id}`, data);
}

export function deleteApi(ids: number[]) {
  return requestClient.delete<void>('/system/api/delete', { data: { ids } });
}

export function realDeleteApi(ids: number[]) {
  return requestClient.delete<void>('/system/api/realDelete', {
    data: { ids },
  });
}

export function recoveryApi(ids: number[]) {
  return requestClient.put<void>('/system/api/recovery', { ids });
}

export function changeApiStatus(data: ApiManageApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/api/changeStatus', data);
}
