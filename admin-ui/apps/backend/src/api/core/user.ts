import type { UserInfo } from '@vben/types';

import { requestClient } from '#/api/request';

export interface SystemUserInfoResult {
  user: {
    avatar: string;
    dashboard?: string;
    email?: string;
    id: number;
    nickname: string;
    phone?: string;
    signed?: string;
    username: string;
    [key: string]: unknown;
  };
  roles: string[];
  codes: string[];
  routers: Array<{
    id: number;
    parent_id: number;
    name: string;
    path: string;
    component: string;
    redirect: string;
    meta: {
      hidden: boolean;
      hiddenBreadcrumb: boolean;
      icon: string;
      title: string;
      type: 'B' | 'I' | 'L' | 'M';
    };
    children: unknown[];
  }>;
}

const DASHBOARD_HOME_PATH_MAP: Record<string, string> = {
  statistics: '/analytics',
  work: '/workspace',
};

function resolveHomePath(dashboard?: string): string {
  if (!dashboard) {
    return '/analytics';
  }

  if (dashboard.startsWith('/')) {
    return dashboard;
  }

  return DASHBOARD_HOME_PATH_MAP[dashboard] || '/analytics';
}

// Module-level cache to avoid duplicated requests in the same event loop.
let _getInfoPromise: null | Promise<SystemUserInfoResult> = null;

export function getSystemInfoApi(): Promise<SystemUserInfoResult> {
  if (!_getInfoPromise) {
    _getInfoPromise = requestClient.get<SystemUserInfoResult>('/system/getInfo');
    _getInfoPromise.finally(() => {
      _getInfoPromise = null;
    });
  }
  return _getInfoPromise;
}

export async function getUserInfoApi(): Promise<UserInfo> {
  const info = await getSystemInfoApi();
  return {
    avatar: info.user.avatar || '',
    desc: info.user.signed || '',
    homePath: resolveHomePath(info.user.dashboard),
    realName: info.user.nickname || info.user.username,
    roles: info.roles,
    token: '',
    userId: String(info.user.id),
    username: info.user.username,
  };
}

export async function getAccessCodesApi(): Promise<string[]> {
  const info = await getSystemInfoApi();
  return info.codes;
}
