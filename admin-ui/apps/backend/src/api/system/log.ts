import type { PageResponse } from '#/types/paging';

import { requestClient } from '#/api/request';

export namespace LogApi {
  export interface LoginLogItem {
    id: number;
    username: string;
    status: number;
    ip: string;
    ip_location: string;
    os: string;
    browser: string;
    message: string;
    login_time: string;
    created_at?: string;
  }

  export interface LoginLogQuery {
    page?: number;
    pageSize?: number;
    username?: string;
    status?: number;
    ip?: string;
    login_time?: string[];
  }

  export interface OperLogItem {
    id: number;
    username: string;
    service_name: string;
    router: string;
    method: string;
    ip: string;
    ip_location: string;
    response_code: string;
    request_data: string;
    response_data: string;
    created_at: string;
    status?: number;
    remark?: string;
  }

  export interface OperLogQuery {
    page?: number;
    pageSize?: number;
    username?: string;
    service_name?: string;
    ip?: string;
    status?: number;
    created_at?: string[];
  }

  export interface ApiLogItem {
    id: number;
    api_id?: number;
    api_name: string;
    access_name: string;
    response_code: string;
    access_time: string;
    ip: string;
    ip_location: string;
    request_data: string;
    response_data: string;
    remark?: string;
  }

  export interface ApiLogQuery {
    page?: number;
    pageSize?: number;
    api_name?: string;
    access_name?: string;
    ip?: string;
    access_time?: string[];
  }

  export type LoginLogResponse = PageResponse<LoginLogItem>;
  export type OperLogResponse = PageResponse<OperLogItem>;
  export type ApiLogResponse = PageResponse<ApiLogItem>;
}

// Login Log APIs
export function getLoginLogPageList(params: LogApi.LoginLogQuery) {
  return requestClient.get<LogApi.LoginLogResponse>(
    '/system/logs/getLoginLogPageList',
    { params },
  );
}

export function deleteLoginLog(ids: number[]) {
  return requestClient.delete<void>('/system/logs/deleteLoginLog', {
    data: { ids },
  });
}

// Operation Log APIs
export function getOperLogPageList(params: LogApi.OperLogQuery) {
  return requestClient.get<LogApi.OperLogResponse>(
    '/system/logs/getOperLogPageList',
    {
    params,
    },
  );
}

export function deleteOperLog(ids: number[]) {
  return requestClient.delete<void>('/system/logs/deleteOperLog', {
    data: { ids },
  });
}

// API Log APIs
export function getApiLogPageList(params: LogApi.ApiLogQuery) {
  return requestClient.get<LogApi.ApiLogResponse>(
    '/system/logs/getApiLogPageList',
    {
    params,
    },
  );
}

export function deleteApiLog(ids: number[]) {
  return requestClient.delete<void>('/system/logs/deleteApiLog', {
    data: { ids },
  });
}
