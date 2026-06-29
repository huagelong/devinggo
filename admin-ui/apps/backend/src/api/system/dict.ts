import { requestClient } from '#/api/request';
import type { BatchIdsPayload, IdType, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace DictApi {
  export interface Item<TKey extends IdType = IdType> {
    [key: string]: unknown;
    code?: string;
    id?: number;
    key: TKey;
    title: string;
  }

  export interface DictTypeItem {
    code: string;
    created_at?: string;
    id: number;
    name: string;
    remark?: string;
    status?: StatusValue;
    updated_at?: string;
  }

  export interface DictTypeQuery extends Partial<PageQuery> {
    code?: string;
    created_at?: string[];
    name?: string;
    status?: StatusValue;
  }

  export interface DictTypeSubmitPayload {
    code: string;
    id?: number;
    name: string;
    remark?: string;
    status: number;
  }

  export interface DictDataItem {
    code: string;
    created_at?: string;
    id: number;
    label: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
    type_id?: number;
    typeId?: number;
    updated_at?: string;
    value: string;
  }

  export interface DictDataQuery extends Partial<PageQuery> {
    code?: string;
    codes?: string;
    created_at?: string[];
    label?: string;
    status?: StatusValue;
    type_id?: number;
    value?: string;
  }

  export interface DictDataSubmitPayload {
    code: string;
    id?: number;
    label: string;
    remark?: string;
    sort: number;
    status: number;
    type_id?: number;
    value: string;
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
  export type DictTypeResponse = PageResponse<DictTypeItem>;
  export type DictDataResponse = PageResponse<DictDataItem>;
  export type DictTypeListResponse = DictTypeItem[] | DictTypeResponse;
  export type DictLookupMap = Record<string, Item>;
}

export function getDictTypePageList(params: DictApi.DictTypeQuery) {
  return requestClient.get<DictApi.DictTypeResponse>('/system/dictType/index', {
    params,
  });
}

export function getDictTypeList(params?: DictApi.DictTypeQuery) {
  return requestClient.get<DictApi.DictTypeListResponse>(
    '/system/dictType/list',
    {
      params,
    },
  );
}

export function getRecycleDictTypeList(params: DictApi.DictTypeQuery) {
  return requestClient.get<DictApi.DictTypeResponse>(
    '/system/dictType/recycle',
    {
      params,
    },
  );
}

export function getDictTypeDetail(id: number) {
  return requestClient.get<DictApi.DictTypeItem>(`/system/dictType/read/${id}`);
}

export function saveDictType(data: DictApi.DictTypeSubmitPayload) {
  return requestClient.post<void>('/system/dictType/save', data);
}

export function updateDictType(id: number, data: DictApi.DictTypeSubmitPayload) {
  return requestClient.put<void>(`/system/dictType/update/${id}`, data);
}

export function deleteDictType(ids: number[]) {
  return requestClient.delete<void>('/system/dictType/delete', {
    data: { ids },
  });
}

export function realDeleteDictType(ids: number[]) {
  return requestClient.delete<void>('/system/dictType/realDelete', {
    data: { ids },
  });
}

export function recoveryDictType(ids: number[]) {
  return requestClient.put<void>('/system/dictType/recovery', { ids });
}

export function changeDictTypeStatus(data: DictApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/dictType/changeStatus', data);
}

export function getDictDataPageList(params: DictApi.DictDataQuery) {
  return requestClient.get<DictApi.DictDataResponse>('/system/dataDict/index', {
    params,
  });
}

export function getDictDataList(params?: DictApi.DictDataQuery) {
  return requestClient.get<DictApi.Item[]>('/system/dataDict/list', {
    params,
  });
}

export function getDictLists(codes: string[] | string) {
  return requestClient.get<DictApi.DictLookupMap>('/system/dataDict/lists', {
    params: { codes: Array.isArray(codes) ? codes.join(',') : codes },
  });
}

export function clearDictDataCache() {
  return requestClient.post<void>('/system/dataDict/clearCache');
}

export function getRecycleDictDataList(params: DictApi.DictDataQuery) {
  return requestClient.get<DictApi.DictDataResponse>(
    '/system/dataDict/recycle',
    {
      params,
    },
  );
}

export function getDictDataDetail(id: number) {
  return requestClient.get<DictApi.DictDataItem>(`/system/dataDict/read/${id}`);
}

export function saveDictData(data: DictApi.DictDataSubmitPayload) {
  return requestClient.post<void>('/system/dataDict/save', data);
}

export function updateDictData(id: number, data: DictApi.DictDataSubmitPayload) {
  return requestClient.put<void>(`/system/dataDict/update/${id}`, data);
}

export function deleteDictData(ids: number[]) {
  return requestClient.delete<void>('/system/dataDict/delete', {
    data: { ids },
  });
}

export function realDeleteDictData(ids: number[]) {
  return requestClient.delete<void>('/system/dataDict/realDelete', {
    data: { ids },
  });
}

export function recoveryDictData(ids: number[]) {
  return requestClient.put<void>('/system/dataDict/recovery', { ids });
}

export function changeDictDataStatus(data: DictApi.ChangeStatusPayload) {
  return requestClient.put<void>('/system/dataDict/changeStatus', data);
}

export function updateDictDataNumber(data: DictApi.NumberOperationPayload) {
  return requestClient.put<void>('/system/dataDict/numberOperation', data);
}

export function getDictList(code: string) {
  return getDictDataList({ code });
}
