import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';
import { validateRoleListItem, validateRoleDetail, logValidationWarning } from '#/api/validation';

export namespace RoleApi {
  export interface ListItem {
    code: string;
    created_at?: string;
    data_scope?: number;
    id: number;
    name: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
    updated_at?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    code?: string;
    created_at?: string[];
    name?: string;
    status?: StatusValue;
  }

  export interface SubmitPayload {
    code: string;
    data_scope: number;
    dept_ids?: number[];
    id?: number;
    menu_ids?: number[];
    name: string;
    remark?: string;
    sort: number;
    status: number;
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

  export interface MenuRelationItem {
    id: number;
    pivot: {
      menu_id: number;
      role_id: number;
    };
  }

  export interface RoleMenuPermission {
    id: number;
    menus: MenuRelationItem[];
  }

  export interface DeptRelationItem {
    id: number;
    pivot: {
      dept_id: number;
      role_id: number;
    };
  }

  export interface RoleDeptPermission {
    id: number;
    depts: DeptRelationItem[];
  }

  export interface PermissionPayload extends Partial<SubmitPayload> {
    data_scope?: number;
    dept_ids?: number[];
    menu_ids?: number[];
  }

  export interface MenuPermissionPayload {
    menu_ids?: number[];
  }

  export interface DataPermissionPayload {
    data_scope: number;
    dept_ids?: number[];
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
  export type OptionListResponse = ListItem[] | PageResponse<ListItem>;
}

export function getRolePageList(params: RoleApi.ListQuery) {
  return requestClient.get<RoleApi.ListResponse>('/system/role/index', { params }).then((response) => {
    if (response?.items) {
      response.items.forEach((item, index) => {
        if (!validateRoleListItem(item)) {
          logValidationWarning(`getRolePageList item[${index}]`, item);
        }
      });
    }
    return response;
  });
}

export function getRecycleRoleList(params: RoleApi.ListQuery) {
  return requestClient.get<RoleApi.ListResponse>('/system/role/recycle', {
    params,
  });
}

export function getRoleList(params?: RoleApi.ListQuery) {
  return requestClient.get<RoleApi.OptionListResponse>('/system/role/list', {
    params,
  }).then((response) => {
    if (Array.isArray(response)) {
      response.forEach((item, index) => {
        if (!validateRoleListItem(item)) {
          logValidationWarning(`getRoleList item[${index}]`, item);
        }
      });
    } else if (response?.items) {
      response.items.forEach((item, index) => {
        if (!validateRoleListItem(item)) {
          logValidationWarning(`getRoleList item[${index}]`, item);
        }
      });
    }
    return response;
  });
}

export function saveRole(data: RoleApi.SubmitPayload) {
  return requestClient.post<void>('/system/role/save', data);
}

export function updateRole(id: number, data: RoleApi.SubmitPayload) {
  return requestClient.put<void>(`/system/role/update/${id}`, data);
}

export function deleteRole(ids: number[]) {
  return requestClient.delete<void>('/system/role/delete', { data: { ids } });
}

export function realDeleteRole(ids: number[]) {
  return requestClient.delete<void>('/system/role/realDelete', {
    data: { ids },
  });
}

export function recoveryRole(ids: number[]) {
  return requestClient.put<void>('/system/role/recovery', { ids });
}

export function changeRoleStatus(data: RoleApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/role/changeStatus', data);
}

export function updateRoleNumber(data: RoleApi.NumberOperationPayload) {
  return requestClient.put<void>('/system/role/numberOperation', data);
}

export function updateRoleMenuPermission(
  id: number,
  data: RoleApi.MenuPermissionPayload,
) {
  return requestClient.put<void>(`/system/role/menuPermission/${id}`, data);
}

export function updateRoleDataPermission(
  id: number,
  data: RoleApi.DataPermissionPayload,
) {
  return requestClient.put<void>(`/system/role/dataPermission/${id}`, data);
}

export function getMenuByRole(id: number) {
  return requestClient.get<RoleApi.RoleMenuPermission[]>(
    `/system/role/getMenuByRole/${id}`,
  );
}

export function getDeptByRole(id: number) {
  return requestClient.get<RoleApi.RoleDeptPermission[]>(
    `/system/role/getDeptByRole/${id}`,
  );
}
