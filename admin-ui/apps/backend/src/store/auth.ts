import type { Recordable, UserInfo } from '@vben/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import CryptoJS from 'crypto-js';

import { LOGIN_PATH } from '@vben/constants';
import { preferences } from '@vben/preferences';
import { resetAllStores, useAccessStore, useUserStore } from '@vben/stores';

import { defineStore } from 'pinia';

import { notification } from '#/adapter/tdesign';
import { getAccessCodesApi, getUserInfoApi, loginApi, logoutApi } from '#/api';
import { $t } from '#/locales';

/**
 * AES-ECB 加密密码
 * 与后端 secure.AESEncrypt/AESDecrypt 保持一致：ECB 模式 + PKCS7 填充 + Base64 编码
 * 注意：不使用 crypto.subtle（Web Crypto API），因其仅在安全上下文（HTTPS/localhost）可用，
 * 通过 HTTP 访问时 crypto.subtle 为 undefined 会导致登录失败。
 */
export function encryptPassword(password: string): string {
  const aesKey = import.meta.env.VITE_APP_AES_KEY as string;
  if (!aesKey) {
    throw new Error('VITE_APP_AES_KEY 未配置');
  }
  return CryptoJS.AES.encrypt(password, CryptoJS.enc.Utf8.parse(aesKey), {
    mode: CryptoJS.mode.ECB,
    padding: CryptoJS.pad.Pkcs7,
  }).toString();
}

export const useAuthStore = defineStore('auth', () => {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const router = useRouter();

  const loginLoading = ref(false);

  /**
   * 异步处理登录操作
   * Asynchronously handle the login process
   * @param params 登录表单数据
   */
  async function authLogin(
    params: Recordable<any>,
    onSuccess?: () => Promise<void> | void,
  ) {
    // 异步处理用户登录操作并获取 accessToken
    let userInfo: null | UserInfo = null;
    try {
      loginLoading.value = true;
      // 加密密码后再提交
      const encryptedParams = { ...params };
      if (encryptedParams.password) {
        encryptedParams.password = await encryptPassword(
          encryptedParams.password,
        );
      }
      const { token } = await loginApi(encryptedParams);

      // 如果成功获取到 token
      if (token) {
        accessStore.setAccessToken(token);
        // 获取用户信息并存储到 accessStore 中
        const [fetchUserInfoResult, accessCodes] = await Promise.all([
          fetchUserInfo(),
          getAccessCodesApi(),
        ]);

        userInfo = fetchUserInfoResult;

        userStore.setUserInfo(userInfo);
        accessStore.setAccessCodes(accessCodes);

        if (accessStore.loginExpired) {
          accessStore.setLoginExpired(false);
        } else {
          onSuccess
            ? await onSuccess?.()
            : await router.push(
                userInfo.homePath || preferences.app.defaultHomePath,
              );
        }

        if (userInfo?.realName) {
          notification.success({
            title: $t('authentication.loginSuccess'),
            content: `${$t('authentication.loginSuccessDesc')}:${userInfo?.realName}`,
            duration: 3000,
          });
        }
      }
    } catch (error) {
      // 登录失败时清理状态并提示用户
      accessStore.setAccessToken(null);
      accessStore.setRefreshToken(null);
      const message =
        error instanceof Error ? error.message : $t('common.loginFailed');
      notification.error({
        title: $t('authentication.loginFailed'),
        content: message,
        duration: 5000,
      });
      throw error;
    } finally {
      loginLoading.value = false;
    }

    return {
      userInfo,
    };
  }

  async function logout(redirect: boolean = true) {
    try {
      await logoutApi();
    } catch {
      // 不做任何处理
    }
    resetAllStores();
    accessStore.setLoginExpired(false);

    // 回登录页带上当前路由地址
    await router.replace({
      path: LOGIN_PATH,
      query: redirect
        ? {
            redirect: encodeURIComponent(router.currentRoute.value.fullPath),
          }
        : {},
    });
  }

  async function fetchUserInfo() {
    try {
      const userInfo = await getUserInfoApi();
      userStore.setUserInfo(userInfo);
      return userInfo;
    } catch (error) {
      // 获取用户信息失败时清理状态并向上传播错误
      userStore.setUserInfo(null as unknown as UserInfo);
      throw error;
    }
  }

  function $reset() {
    loginLoading.value = false;
  }

  return {
    $reset,
    authLogin,
    fetchUserInfo,
    loginLoading,
    logout,
  };
});
