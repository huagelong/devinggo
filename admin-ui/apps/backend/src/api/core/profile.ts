import type { LogApi } from '#/api/system/log';

import { requestClient } from '#/api/request';

export namespace ProfileApi {
  export interface UpdateUserInfoPayload {
    avatar?: string;
    email?: string;
    nickname?: string;
    phone?: string;
    remark?: string;
    signed?: string;
  }

  export interface ModifyPasswordPayload {
    oldPassword?: string;
    newPassword?: string;
    newPasswordConfirmation?: string;
  }
}

// User profile APIs
export async function updateUserInfoApi(data: ProfileApi.UpdateUserInfoPayload) {
  return requestClient.post<void>('/system/user/updateInfo', data);
}

export async function modifyPasswordApi(data: ProfileApi.ModifyPasswordPayload) {
  return requestClient.post<void>('/system/user/modifyPassword', data);
}

// Log APIs (simplified endpoints for dashboard profile page)
export async function getLoginLogListApi(params: LogApi.LoginLogQuery) {
  return requestClient.get<LogApi.LoginLogResponse>(
    '/system/common/getLoginLogList',
    { params },
  );
}

export async function getOperationLogListApi(params: LogApi.OperLogQuery) {
  return requestClient.get<LogApi.OperLogResponse>(
    '/system/common/getOperationLogList',
    { params },
  );
}
