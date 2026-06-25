import { requestClient } from '#/api/request';
import type { BatchIdsPayload, IdType } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace ApiGroupApi {
  export interface ListItem {
    created_at?: string;
    id: IdType;
    name: string;
    remark?: string;
    status?: number;
    updated_at?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    created_at?: string[];
    name?: string;
    status?: number;
  }

  export interface SubmitPayload {
    id?: number;
    name: string;
    remark?: string;
    status: number;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export type ListResponse = PageResponse<ListItem>;
  export type BatchPayload = BatchIdsPayload<number>;
}

export function getApiGroupList() {
  return requestClient.get<ApiGroupApi.ListItem[]>('/system/apiGroup/list');
}

export function getApiGroupPageList(params: ApiGroupApi.ListQuery) {
  return requestClient.get<ApiGroupApi.ListResponse>('/system/apiGroup/index', {
    params,
  });
}

export function getRecycleApiGroupList(params: ApiGroupApi.ListQuery) {
  return requestClient.get<ApiGroupApi.ListResponse>('/system/apiGroup/recycle', {
    params,
  });
}

export function saveApiGroup(data: ApiGroupApi.SubmitPayload) {
  return requestClient.post<void>('/system/apiGroup/save', data);
}

export function updateApiGroup(id: number, data: ApiGroupApi.SubmitPayload) {
  return requestClient.put<void>(`/system/apiGroup/update/${id}`, data);
}

export function deleteApiGroup(ids: number[]) {
  return requestClient.delete<void>('/system/apiGroup/delete', { data: { ids } });
}

export function realDeleteApiGroup(ids: number[]) {
  return requestClient.delete<void>('/system/apiGroup/realDelete', {
    data: { ids },
  });
}

export function recoveryApiGroup(ids: number[]) {
  return requestClient.put<void>('/system/apiGroup/recovery', { ids });
}

export function changeApiGroupStatus(data: ApiGroupApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/apiGroup/changeStatus', data);
}
