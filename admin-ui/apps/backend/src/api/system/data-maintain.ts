import type { PageQuery, PageResponse } from '#/types/paging';

import { requestClient } from '#/api/request';

export namespace DataMaintainApi {
  export interface ListItem {
    name: string;
    collation?: string;
    comment?: string;
    engine?: string;
    create_time?: string;
    rows?: number;
  }

  export interface ListQuery extends Partial<PageQuery> {
    group_name?: string;
    name?: string;
  }

  export interface DetailQuery {
    group_name?: string;
    table_name: string;
  }

  export interface ColumnItem {
    field: string;
    type?: string;
    key?: string;
    null?: string;
    default?: string;
    extra?: string;
    comment?: string;
  }

  export interface TableActionPayload {
    group_name?: string;
    table_name: string;
  }

  export type ListResponse = PageResponse<ListItem>;
  export type DetailResponse = Record<string, ColumnItem>;
}

export function getDataMaintainPageList(params: DataMaintainApi.ListQuery) {
  return requestClient.get<DataMaintainApi.ListResponse>(
    '/system/dataMaintain/index',
    { params },
  );
}

export function getDataMaintainDetailed(params: DataMaintainApi.DetailQuery) {
  return requestClient.get<DataMaintainApi.DetailResponse>(
    '/system/dataMaintain/detailed',
    { params },
  );
}

export function optimizeDataMaintainTable(
  data: DataMaintainApi.TableActionPayload,
) {
  return requestClient.post('/system/dataMaintain/optimize', data);
}

export function fragmentDataMaintainTable(
  data: DataMaintainApi.TableActionPayload,
) {
  return requestClient.post('/system/dataMaintain/fragment', data);
}
