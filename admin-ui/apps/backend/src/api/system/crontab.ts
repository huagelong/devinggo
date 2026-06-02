import { requestClient } from '#/api/request';
import type { BatchIdsPayload } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace CrontabApi {
  export interface ListItem {
    created_at?: string;
    id: number;
    is_finally: number;
    name: string;
    remark?: string;
    rule: string;
    target: string;
    type: number;
    updated_at?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    created_at?: string[];
    is_finally?: number;
    name?: string;
    type?: number;
  }

  export interface SubmitPayload {
    id?: number;
    is_finally?: number;
    name: string;
    remark?: string;
    rule: string;
    target: string;
    type: number;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export interface RunPayload {
    id: number;
  }

  export interface LogItem {
    created_at?: string;
    end_time?: string;
    error?: string;
    id: number;
    crontab_id?: number;
    name?: string;
    output?: string;
    start_time?: string;
    status?: number;
  }

  export interface LogQuery extends Partial<PageQuery> {
    crontab_id?: number;
    created_at?: string[];
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
  export type LogResponse = PageResponse<LogItem>;
}

export function getCrontabPageList(params: CrontabApi.ListQuery) {
  return requestClient.get<CrontabApi.ListResponse>(
    '/system/setting/crontab/index',
    { params },
  );
}

export function getRecycleCrontabList(params: CrontabApi.ListQuery) {
  return requestClient.get<CrontabApi.ListResponse>(
    '/system/setting/crontab/recycle',
    { params },
  );
}

export function saveCrontab(data: CrontabApi.SubmitPayload) {
  return requestClient.post<void>('/system/setting/crontab/save', data);
}

export function updateCrontab(id: number, data: CrontabApi.SubmitPayload) {
  return requestClient.put<void>(`/system/setting/crontab/update/${id}`, data);
}

export function deleteCrontab(ids: number[]) {
  return requestClient.delete<void>('/system/setting/crontab/delete', {
    data: { ids },
  });
}

export function realDeleteCrontab(ids: number[]) {
  return requestClient.delete<void>('/system/setting/crontab/realDelete', {
    data: { ids },
  });
}

export function recoveryCrontab(ids: number[]) {
  return requestClient.put<void>('/system/setting/crontab/recovery', { ids });
}

export function changeCrontabStatus(data: CrontabApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/setting/crontab/changeStatus', data);
}

export function runCrontab(data: CrontabApi.RunPayload) {
  return requestClient.post<void>('/system/setting/crontab/run', data);
}

export function getCrontabLogPageList(params: CrontabApi.LogQuery) {
  return requestClient.get<CrontabApi.LogResponse>(
    '/system/setting/crontab/logPageList',
    { params },
  );
}

export function deleteCrontabLog(ids: number[]) {
  return requestClient.delete<void>('/system/setting/crontab/deleteCrontabLog', {
    data: { ids },
  });
}
