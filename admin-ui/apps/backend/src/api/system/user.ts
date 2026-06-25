import { requestClient } from '#/api/request';
import type { RequestClientConfig } from '@vben/request';
import type { BatchIdsPayload, IdType, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';
import { validateUserListItem, validateUserDetail, logValidationWarning } from '#/api/validation';

export namespace UserApi {
  export interface ListItem {
    avatar?: string;
    dashboard?: string;
    dept_name?: string;
    email?: string;
    id: number;
    nickname?: string;
    phone?: string;
    post_name?: string;
    role_name?: string;
    status?: number;
    user_type?: string;
    username: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    created_at?: string[];
    dept_id?: IdType;
    dept_ids?: IdType[];
    email?: string;
    nickname?: string;
    phone?: string;
    post_id?: number;
    role_id?: number;
    status?: StatusValue;
    user_type?: string;
    username?: string;
  }

  export interface RelatedEntity {
    id: number;
    name?: string;
  }

  export interface Detail extends SubmitPayload {
    deptList?: RelatedEntity[];
    dept_ids?: number[];
    postList?: RelatedEntity[];
    post_ids?: number[];
    roleList?: RelatedEntity[];
    role_ids?: number[];
  }

  export interface SubmitPayload {
    avatar?: string;
    dashboard?: string;
    dept_ids?: number[];
    email?: string;
    id?: number;
    nickname?: string;
    password?: string;
    phone?: string;
    post_ids?: number[];
    remark?: string;
    role_ids?: number[];
    status?: number;
    user_type?: string;
    username: string;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export interface ResetPasswordPayload {
    id: number;
  }

  export interface SetHomePagePayload {
    dashboard: string;
    id: number;
  }

  export interface DownloadResponse {
    data: Blob;
    headers?:
      | Record<string, string | undefined>
      | {
          get?: (name: string) => null | string;
          [key: string]: unknown;
        };
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
}

export function getUserList(
  params: UserApi.ListQuery,
  config?: RequestClientConfig,
) {
  return requestClient.get<UserApi.ListResponse>('/system/user/index', {
    ...config,
    params,
  }).then((response) => {
    if (response?.items) {
      response.items.forEach((item, index) => {
        if (!validateUserListItem(item)) {
          logValidationWarning(`getUserList item[${index}]`, item);
        }
      });
    }
    return response;
  });
}

export function getRecycleUserList(params: UserApi.ListQuery) {
  return requestClient.get<UserApi.ListResponse>('/system/user/recycle', { params });
}

export function getUserDetail(id: number) {
  return requestClient.get<UserApi.Detail>(`/system/user/read/${id}`).then((response) => {
    if (!validateUserDetail(response)) {
      logValidationWarning('getUserDetail', response);
    }
    return response;
  });
}

export function saveUser(data: UserApi.SubmitPayload) {
  return requestClient.post<void>('/system/user/save', data);
}

export function updateUser(id: number, data: UserApi.SubmitPayload) {
  return requestClient.put<void>(`/system/user/update/${id}`, data);
}

export function deleteUser(ids: number[]) {
  return requestClient.delete<void>('/system/user/delete', { data: { ids } });
}

export function realDeleteUser(ids: number[]) {
  return requestClient.delete<void>('/system/user/realDelete', { data: { ids } });
}

export function recoveryUser(ids: number[]) {
  return requestClient.put<void>('/system/user/recovery', { ids });
}

export function changeUserStatus(data: UserApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/user/changeStatus', data);
}

export function resetPassword(data: UserApi.ResetPasswordPayload) {
  return requestClient.put<void>('/system/user/initUserPassword', data);
}

export function clearUserCache(data: UserApi.ResetPasswordPayload) {
  return requestClient.post<void>('/system/user/clearCache', data);
}

// Set home page
export function setHomePage(data: UserApi.SetHomePagePayload) {
  return requestClient.post<void>('/system/user/setHomePage', data);
}

export function importUserFile(file: File) {
  const formData = new FormData();
  formData.append('file', file);
  return requestClient.post<void>('/system/user/import', formData);
}

export function exportUserList(data: Omit<UserApi.ListQuery, keyof PageQuery>) {
  return requestClient.download<UserApi.DownloadResponse>('/system/user/export', {
    data,
    method: 'POST',
    responseReturn: 'raw',
  });
}

export function downloadUserImportTemplate() {
  return requestClient.download<UserApi.DownloadResponse>(
    '/system/user/downloadTemplate',
    {
    method: 'GET',
    responseReturn: 'raw',
    },
  );
}
