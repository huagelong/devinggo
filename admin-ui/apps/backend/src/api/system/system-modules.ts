import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace SystemModulesApi {
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

export function getSystemModulesPageList(
  params: SystemModulesApi.ListQuery,
) {
  return requestClient.get<SystemModulesApi.ListResponse>(
    '/system/systemModules/index',
    { params },
  );
}

export function getRecycleSystemModulesList(
  params: SystemModulesApi.ListQuery,
) {
  return requestClient.get<SystemModulesApi.ListResponse>(
    '/system/systemModules/recycle',
    { params },
  );
}

export function saveSystemModules(data: SystemModulesApi.SubmitPayload) {
  return requestClient.post<void>('/system/systemModules/save', data);
}

export function updateSystemModules(
  id: number,
  data: SystemModulesApi.SubmitPayload,
) {
  return requestClient.put<void>(`/system/systemModules/update/${id}`, data);
}

export function deleteSystemModules(ids: number[]) {
  return requestClient.delete<void>('/system/systemModules/delete', {
    data: { ids },
  });
}

export function realDeleteSystemModules(ids: number[]) {
  return requestClient.delete<void>('/system/systemModules/realDelete', {
    data: { ids },
  });
}

export function recoverySystemModules(ids: number[]) {
  return requestClient.put<void>('/system/systemModules/recovery', { ids });
}

export function changeSystemModulesStatus(
  data: SystemModulesApi.ChangeStatusPayload,
) {
  return requestClient.put<void>('/system/systemModules/changeStatus', data);
}
