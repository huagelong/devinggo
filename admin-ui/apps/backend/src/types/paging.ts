export interface PageQuery {
  page: number;
  pageSize: number;
}

export interface PageInfo {
  current?: number;
  page?: number;
  pageSize?: number;
  total: number;
}

export interface PageResponse<TItem> {
  items?: TItem[];
  pageInfo?: PageInfo;
  total?: number;
  [key: string]: unknown;
}
