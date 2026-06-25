import { requestClient } from '#/api/request';
import type { BatchIdsPayload, IdType, StatusValue } from '#/types/common';

export namespace MenuApi {
  export interface TreeItem {
    children?: TreeItem[];
    code?: string;
    component?: string;
    created_at?: string;
    deleted_at?: string;
    icon?: string;
    id: number;
    is_hidden?: number;
    level?: string;
    name: string;
    parent_id?: IdType;
    redirect?: string;
    remark?: string;
    restful?: string;
    route?: string;
    sort?: number;
    status?: StatusValue;
    type?: string;
    updated_at?: string;
  }

  export interface TreeOptionItem {
    children?: TreeOptionItem[];
    id: number;
    label: string;
    parent_id?: IdType;
    value: number;
  }

  export interface ListQuery {
    code?: string;
    created_at?: string[];
    level?: string;
    name?: string;
    no_buttons?: boolean;
    recycle?: boolean;
    status?: StatusValue;
  }

  export interface TreeQuery {
    onlyMenu?: boolean;
    scope?: boolean;
  }

  export interface SubmitPayload {
    code?: string;
    component?: string;
    icon?: string;
    id?: number;
    is_hidden: number;
    level?: string;
    name: string;
    parent_id: number;
    redirect?: string;
    remark?: string;
    restful?: string;
    route?: string;
    sort: number;
    status: number;
    type: string;
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
}

export function getMenuTreeList(params?: MenuApi.ListQuery) {
  return requestClient.get<MenuApi.TreeItem[]>('/system/menu/index', { params });
}

export function getRecycleMenuTreeList(params?: MenuApi.ListQuery) {
  return requestClient.get<MenuApi.TreeItem[]>('/system/menu/recycle', {
    params,
  });
}

export function getMenuTreeOptions(params?: MenuApi.TreeQuery) {
  return requestClient.get<MenuApi.TreeOptionItem[]>('/system/menu/tree', {
    params,
  });
}

export function saveMenu(data: MenuApi.SubmitPayload) {
  return requestClient.post<void>('/system/menu/save', data);
}

export function updateMenu(id: number, data: MenuApi.SubmitPayload) {
  return requestClient.put<void>(`/system/menu/update/${id}`, data);
}

export function deleteMenu(ids: number[]) {
  return requestClient.delete<void>('/system/menu/delete', { data: { ids } });
}

export function realDeleteMenu(ids: number[]) {
  return requestClient.delete<void>('/system/menu/realDelete', {
    data: { ids },
  });
}

export function recoveryMenu(ids: number[]) {
  return requestClient.put<void>('/system/menu/recovery', { ids });
}

export function changeMenuStatus(data: MenuApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/menu/changeStatus', data);
}

export function updateMenuNumber(data: MenuApi.NumberOperationPayload) {
  return requestClient.put<void>('/system/menu/numberOperation', data);
}
