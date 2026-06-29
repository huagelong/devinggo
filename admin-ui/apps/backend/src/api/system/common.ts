import { requestClient } from '#/api/request';
import type { NoticeApi } from '#/api/system/notice';
import type { IdType } from '#/types/common';
import type { PageQuery } from '#/types/paging';

export namespace SystemCommonApi {
  export interface UserInfoItem {
    id: IdType;
    nickname?: string;
    username: string;
  }

  export interface UserInfoByIdsPayload {
    ids: IdType[];
  }
}

export function getUserInfoByIds(data: SystemCommonApi.UserInfoByIdsPayload) {
  return requestClient.post<SystemCommonApi.UserInfoItem[]>(
    '/system/common/getUserInfoByIds',
    data,
  );
}

export function getNoticeList(params: Partial<PageQuery>) {
  return requestClient.get<NoticeApi.ListResponse>('/system/common/getNoticeList', {
    params,
  });
}
