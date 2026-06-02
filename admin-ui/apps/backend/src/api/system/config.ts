import { requestClient } from '#/api/request';
import type { BatchIdsPayload, IdType, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace ConfigApi {
  export type InputType =
    | 'input'
    | 'textarea'
    | 'select'
    | 'radio'
    | 'checkbox'
    | 'switch'
    | 'upload'
    | 'key-value'
    | 'editor';

  export interface ConfigGroupItem {
    code: string;
    created_at?: string;
    id: number;
    name: string;
    remark?: string;
    sort?: number;
    updated_at?: string;
  }

  export interface ConfigGroupSubmitPayload {
    code: string;
    id?: number;
    name: string;
    remark?: string;
    sort?: number;
  }

  export interface ConfigSelectOption {
    label: string;
    value: string;
  }

  export interface ConfigItem {
    config_select_data?: ConfigSelectOption[] | string;
    group_id: number;
    id: number;
    input_type: InputType;
    key: string;
    name: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
    value: string | number | boolean | string[] | Record<string, unknown>;
  }

  export interface ConfigListQuery extends Partial<PageQuery> {
    group_id?: number;
    key?: string;
    name?: string;
    orderBy?: string;
    orderType?: string;
  }

  export interface ConfigSubmitPayload {
    config_select_data?: ConfigSelectOption[] | string;
    group_id: number;
    id?: number;
    input_type: InputType;
    key: string;
    name: string;
    remark?: string;
    sort?: number;
    status?: StatusValue;
    value?: string | number | boolean | string[] | Record<string, unknown>;
  }

  export interface DeletePayload extends Partial<BatchIdsPayload<IdType>> {
    id?: number;
  }

  export interface UpdateByKeysPayload {
    [key: string]:
      | string
      | number
      | boolean
      | string[]
      | Record<string, unknown>
      | undefined;
  }

  export interface ConfigListResponse extends PageResponse<ConfigItem> {
    data?: ConfigItem[];
  }
}

export function getConfigGroupList() {
  return requestClient.get<ConfigApi.ConfigGroupItem[]>(
    '/system/setting/configGroup/index',
  );
}

export function saveConfigGroup(data: ConfigApi.ConfigGroupSubmitPayload) {
  return requestClient.post<void>('/system/setting/configGroup/save', data);
}

export function deleteConfigGroup(data: ConfigApi.DeletePayload) {
  return requestClient.delete<void>('/system/setting/configGroup/delete', {
    data,
  });
}

export function getConfigList(params: ConfigApi.ConfigListQuery) {
  return requestClient.get<ConfigApi.ConfigItem[]>(
    '/system/setting/config/index',
    { params },
  );
}

export function saveConfig(data: ConfigApi.ConfigSubmitPayload) {
  return requestClient.post<void>('/system/setting/config/save', data);
}

export function updateConfig(data: ConfigApi.ConfigSubmitPayload) {
  return requestClient.post<void>('/system/setting/config/update', data);
}

export function deleteConfig(data: ConfigApi.DeletePayload) {
  return requestClient.delete<void>('/system/setting/config/delete', {
    data,
  });
}

export function updateConfigByKeys(data: ConfigApi.UpdateByKeysPayload) {
  return requestClient.post<void>('/system/setting/config/updateByKeys', data);
}

export function clearConfigCache() {
  return requestClient.post<void>('/system/setting/config/clearCache');
}
