import type { BasicUserInfo } from '@vben-core/typings';

/** 用户信息 */
interface UserInfo extends BasicUserInfo {
  /**
   * 用户描述
   */
  desc: string;
  /**
   * 首页地址
   */
  homePath: string;
  /**
   * 后台首页类型
   */
  dashboard?: string;

  /**
   * accessToken
   */
  token: string;
}

export type { UserInfo };
