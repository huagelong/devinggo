import { requestClient } from '#/api/request';
import type { IdType } from '#/types/common';

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
