import type { PageQuery, PageResponse } from '#/types/paging';

import { requestClient } from '#/api/request';

export namespace MessageApi {
  export interface SendUser {
    id: number;
    username: string;
    nickname: string;
  }

  export interface QueueMessageItem {
    id: number;
    content_type: string;
    title: string;
    send_by: number;
    content: string;
    created_by: number;
    updated_by: number;
    created_at: string;
    updated_at: string;
    remark: string;
    send_user: SendUser;
    read_status?: number; // 1 未读, 2 已读
  }

  export interface QueueMessageQuery extends Partial<PageQuery> {
    title?: string;
    read_status?: string; // 'all' | '1' | '2'
    content_type?: string;
    created_at?: string[];
  }

  export interface QueueMessageResponse extends PageResponse<QueueMessageItem> {}

  export interface UpdateReadStatusPayload {
    ids: number[];
  }

  export interface DeleteMessagesPayload {
    ids: number[];
  }

  // Data Dict
  export interface DataDictItem {
    key: string;
    title: string;
    value?: string;
  }

  export interface DataDictQuery {
    code: string;
  }

  export type DataDictListResponse = DataDictItem[];
}

// Queue Message APIs
export function getQueueMessageReceiveListApi(
  params: MessageApi.QueueMessageQuery,
) {
  return requestClient.get<MessageApi.QueueMessageResponse>(
    '/system/queueMessage/receiveList',
    { params },
  );
}

export function updateQueueMessageReadStatusApi(
  data: MessageApi.UpdateReadStatusPayload,
) {
  return requestClient.put<void>('/system/queueMessage/updateReadStatus', data);
}

export function deleteQueueMessageApi(data: MessageApi.DeleteMessagesPayload) {
  return requestClient.delete<void>('/system/queueMessage/deletes', { data });
}

// Data Dict APIs
export function getDataDictListApi(params: MessageApi.DataDictQuery) {
  return requestClient.get<MessageApi.DataDictListResponse>(
    '/system/dataDict/list',
    { params },
  );
}
