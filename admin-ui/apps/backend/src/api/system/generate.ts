import { requestClient } from '#/api/request';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace GenerateApi {
  export type GenerateType = 'single' | 'tree';
  export type ComponentType = 1 | 2 | 3; // 模态框/抽屉/Tag页

  export interface ListItem {
    id: number;
    table_name: string;
    table_comment: string;
    type: GenerateType;
    module_name?: string;
    menu_name?: string;
    created_at?: string;
    updated_at?: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    table_name?: string;
    type?: GenerateType;
  }

  export interface ListResponse extends PageResponse<ListItem> {}

  // 装载数据表
  export interface LoadTableItem {
    name: string;
    comment: string;
    sourceName: string;
  }

  export interface LoadTablePayload {
    source: string;
    names: LoadTableItem[];
  }

  // 读取表信息
  export interface ReadTableResponse {
    table_name: string;
    table_comment: string;
    columns: TableColumn[];
  }

  export interface TableColumn {
    column_name: string;
    column_comment: string;
    column_type: string;
    is_nullable: string;
    data_type: string;
  }

  // 预览代码
  export interface PreviewCodeItem {
    name: string;
    tab_name: string;
    code: string;
    lang: string;
  }

  // 更新表信息
  export interface UpdatePayload {
    id: number;
    table_name?: string;
    table_comment?: string;
    remark?: string;
    module_name?: string;
    belong_menu_id?: number;
    type: GenerateType;
    menu_name?: string;
    component_type?: ComponentType;
    tpl_type?: string;
    // 树表配置
    tree_id?: string;
    tree_parent_id?: string;
    tree_name?: string;
    // Tag页配置
    tag_id?: string;
    tag_name?: string;
    tag_view_name?: string;
    // 字段配置
    fields?: FieldConfig[];
    // 菜单配置
    menu_buttons?: string[];
  }

  export interface FieldConfig {
    id?: number;
    column_name: string;
    column_comment: string;
    column_type: string;
    sort: number;
    is_required: number;
    is_insert: number;
    is_edit: number;
    is_list: number;
    is_query: number;
    is_sort: number;
    query_type: string;
    view_type: string;
    dict_type?: string;
    allow_roles?: string;
  }

  // 同步
  export interface SyncPayload {
    id: number;
  }

  // 生成代码
  export interface GeneratePayload {
    ids: string;
  }
}

// 获取代码列表
export function getCodePageList(params: GenerateApi.ListQuery) {
  return requestClient.get<GenerateApi.ListResponse>(
    '/system/code/index',
    { params },
  );
}

// 删除代码记录
export function deleteCode(ids: number[]) {
  return requestClient.delete<void>('/system/code/delete', {
    data: { ids },
  });
}

// 更新代码信息
export function updateCode(data: GenerateApi.UpdatePayload) {
  return requestClient.post<void>('/system/code/update', data);
}

// 读取表信息
export function readTable(id: number) {
  return requestClient.get<GenerateApi.ReadTableResponse>(
    `/system/code/readTable/${id}`,
  );
}

// 装载数据表
export function loadTable(data: GenerateApi.LoadTablePayload) {
  return requestClient.post<void>('/system/code/loadTable', data);
}

// 同步数据表
export function syncTable(id: number) {
  return requestClient.put<void>(`/system/code/sync/${id}`);
}

// 生成代码（返回 blob）
export function generateCode(data: GenerateApi.GeneratePayload) {
  return requestClient.post<string>('/system/code/generate', data, {
    responseType: 'blob',
  } as any);
}

// 预览代码
export function previewCode(id: number) {
  return requestClient.get<{ data: GenerateApi.PreviewCodeItem[] }>(
    '/system/code/preview',
    { params: { id } },
  );
}
