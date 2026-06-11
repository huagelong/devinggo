import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace <%.EntityName%>Api {
  export interface ListItem {
<%range .ListFields%>    <%.Name%>?: <%.Type%>;
<%end%>    id: number;
  }

  export interface ListQuery extends Partial<PageQuery> {
<%range .SearchFields%>    <%.Name%>?: <%.Type%>;
<%end%>  }

  export interface SubmitPayload {
<%range .EditableFields%>    <%.Name%><%if .Optional%>?<%end%>: <%.Type%>;
<%end%>    id?: number;
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export interface UpdateNumberPayload {
    id: number;
    numberName: string;
    numberValue: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
  export type OptionListResponse = ListItem[] | ListResponse;
}

export function get<%.EntityName%>PageList(params: <%.EntityName%>Api.ListQuery) {
  return requestClient.get<<%.EntityName%>Api.ListResponse>('<%.ApiPrefix%>/index', { params });
}

export function get<%.EntityName%>List(params?: <%.EntityName%>Api.ListQuery) {
  return requestClient.get<<%.EntityName%>Api.OptionListResponse>('<%.ApiPrefix%>/list', { params });
}

export function getRecycle<%.EntityName%>List(params: <%.EntityName%>Api.ListQuery) {
  return requestClient.get<<%.EntityName%>Api.ListResponse>('<%.ApiPrefix%>/recycle', { params });
}

export function save<%.EntityName%>(data: <%.EntityName%>Api.SubmitPayload) {
  return requestClient.post<void>('<%.ApiPrefix%>/save', data);
}

export function update<%.EntityName%>(id: number, data: <%.EntityName%>Api.SubmitPayload) {
  return requestClient.put<void>(`<%.ApiPrefix%>/update/${id}`, data);
}

export function delete<%.EntityName%>(ids: number[]) {
  return requestClient.delete<void>('<%.ApiPrefix%>/delete', { data: { ids } });
}

export function realDelete<%.EntityName%>(ids: number[]) {
  return requestClient.delete<void>('<%.ApiPrefix%>/realDelete', { data: { ids } });
}

export function recovery<%.EntityName%>(ids: number[]) {
  return requestClient.put<void>('<%.ApiPrefix%>/recovery', { ids });
}

export function change<%.EntityName%>Status(data: <%.EntityName%>Api.ChangeStatusPayload) {
  return requestClient.put<void>('<%.ApiPrefix%>/changeStatus', data);
}

export function update<%.EntityName%>Number(data: <%.EntityName%>Api.UpdateNumberPayload) {
  return requestClient.put<void>('<%.ApiPrefix%>/updateNumber', data);
}
