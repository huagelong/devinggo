import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace AppGroupApi {
  export interface ListItem {
    created_at?: string;
    id: number;
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
    id?: number;
    name: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
}

export function getAppGroupPageList(params: AppGroupApi.ListQuery) {
  return requestClient.get<AppGroupApi.ListResponse>('/system/appGroup/index', {
    params,
  });
}

export function getRecycleAppGroupList(params: AppGroupApi.ListQuery) {
  return requestClient.get<AppGroupApi.ListResponse>('/system/appGroup/recycle', {
    params,
  });
}

export function saveAppGroup(data: AppGroupApi.SubmitPayload) {
  return requestClient.post<void>('/system/appGroup/save', data);
}

export function updateAppGroup(id: number, data: AppGroupApi.SubmitPayload) {
  return requestClient.put<void>(`/system/appGroup/update/${id}`, data);
}

export function deleteAppGroup(ids: number[]) {
  return requestClient.delete<void>('/system/appGroup/delete', { data: { ids } });
}

export function realDeleteAppGroup(ids: number[]) {
  return requestClient.delete<void>('/system/appGroup/realDelete', {
    data: { ids },
  });
}

export function recoveryAppGroup(ids: number[]) {
  return requestClient.put<void>('/system/appGroup/recovery', { ids });
}

export function changeAppGroupStatus(data: AppGroupApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/appGroup/changeStatus', data);
}
