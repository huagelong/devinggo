import { requestClient } from '#/api/request';
import type { BatchIdsPayload, IdType, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace DeptApi {
  export interface ListTreeItem {
    children?: ListTreeItem[];
    created_at?: string;
    id: number;
    leader?: string;
    name: string;
    parent_id?: IdType;
    phone?: string;
    sort?: number;
    status?: StatusValue;
  }

  export interface ListItem {
    created_at?: string;
    id: number;
    leader?: string;
    level?: string;
    name: string;
    parent_id?: IdType;
    phone?: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
    updated_at?: string;
  }

  export interface TreeNode {
    children?: TreeNode[];
    disabled?: boolean;
    id: number;
    label: string;
    leader?: string;
    parent_id?: IdType;
    status?: StatusValue;
    value?: number;
  }

  export interface ListQuery {
    created_at?: string[];
    leader?: string;
    level?: string;
    name?: string;
    phone?: string;
    recycle?: boolean;
    status?: StatusValue;
  }

  export interface LeaderListItem {
    avatar?: string;
    email?: string;
    id: number;
    leader_add_time?: string;
    nickname?: string;
    phone?: string;
    status?: StatusValue;
    user_type?: string;
    username: string;
  }

  export interface LeaderListQuery extends Partial<PageQuery> {
    dept_id: number;
    nickname?: string;
    status?: StatusValue;
    username?: string;
  }

  export interface SubmitPayload {
    id?: number;
    leader?: string;
    level?: string;
    name: string;
    parent_id: number;
    phone?: string;
    remark?: string;
    sort: number;
    status: StatusValue;
  }

  export interface AddLeaderUser {
    nickname?: string;
    user_id: number;
    username?: string;
  }

  export interface AddLeaderPayload {
    id: number;
    users: AddLeaderUser[];
  }

  export interface DeleteLeaderPayload {
    id: number;
    ids: number[];
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
  export type LeaderListResponse = PageResponse<LeaderListItem>;
}

export function getDeptPageList(params?: DeptApi.ListQuery) {
  return requestClient.get<DeptApi.ListTreeItem[]>('/system/dept/index', {
    params,
  });
}

export function getRecycleDeptList(params?: DeptApi.ListQuery) {
  return requestClient.get<DeptApi.ListTreeItem[]>('/system/dept/recycle', {
    params,
  });
}

export function getDeptList() {
  return requestClient.get<DeptApi.ListItem[]>('/system/dept/list');
}

export function getDeptTree() {
  return requestClient.get<DeptApi.TreeNode[]>('/system/dept/tree');
}

export function getDeptLeaderList(params: DeptApi.LeaderListQuery) {
  return requestClient.get<DeptApi.LeaderListResponse>(
    '/system/dept/getLeaderList',
    { params },
  );
}

export function saveDept(data: DeptApi.SubmitPayload) {
  return requestClient.post<void>('/system/dept/save', data);
}

export function updateDept(id: number, data: DeptApi.SubmitPayload) {
  return requestClient.put<void>(`/system/dept/update/${id}`, data);
}

export function deleteDept(ids: number[]) {
  return requestClient.delete<void>('/system/dept/delete', { data: { ids } });
}

export function realDeleteDept(ids: number[]) {
  return requestClient.delete<void>('/system/dept/realDelete', {
    data: { ids },
  });
}

export function recoveryDept(ids: number[]) {
  return requestClient.put<void>('/system/dept/recovery', { ids });
}

export function changeDeptStatus(data: DeptApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/dept/changeStatus', data);
}

export function updateDeptNumber(data: DeptApi.NumberOperationPayload) {
  return requestClient.put<void>('/system/dept/numberOperation', data);
}

export function addDeptLeader(data: DeptApi.AddLeaderPayload) {
  return requestClient.post<void>('/system/dept/addLeader', data);
}

export function deleteDeptLeader(data: DeptApi.DeleteLeaderPayload) {
  return requestClient.delete<void>('/system/dept/delLeader', { data });
}
