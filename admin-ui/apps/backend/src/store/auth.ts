import type { Recordable, UserInfo } from '@vben/types';

import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { LOGIN_PATH } from '@vben/constants';
import { preferences } from '@vben/preferences';
import { resetAllStores, useAccessStore, useUserStore } from '@vben/stores';

import { defineStore } from 'pinia';

import { notification } from '#/adapter/tdesign';
import { getAccessCodesApi, getUserInfoApi, loginApi, logoutApi } from '#/api';
import { $t } from '#/locales';

/**
 * AES-ECB 加密密码（使用 SubtleCrypto 单块 CBC 模拟 ECB）
 * 与后端 secure.AESEncrypt 保持一致：ECB 模式 + PKCS7 填充 + Base64 编码
 */
export async function encryptPassword(password: string): Promise<string> {
  const aesKey = import.meta.env.VITE_APP_AES_KEY as string;
  if (!aesKey) {
    throw new Error('VITE_APP_AES_KEY 未配置');
  }
  const encoder = new TextEncoder();
  const keyBytes = encoder.encode(aesKey);
  const data = encoder.encode(password);

  // PKCS7 填充到 16 字节的倍数
  const blockSize = 16;
  const paddingLen = blockSize - (data.length % blockSize);
  const padded = new Uint8Array(data.length + paddingLen);
  padded.set(data);
  padded.fill(paddingLen, data.length);

  // 导入 AES 密钥
  const cryptoKey = await crypto.subtle.importKey(
    'raw',
    keyBytes,
    { name: 'AES-CBC' },
    false,
    ['encrypt'],
  );

  // ECB 模式：每个 16 字节块使用零 IV 的 AES-CBC 独立加密
  // 单块 AES-CBC(IV=0) == AES-ECB，输出取前 16 字节
  const result = new Uint8Array(padded.length);
  const zeroIv = new Uint8Array(16);
  for (let i = 0; i < padded.length; i += blockSize) {
    const block = padded.slice(i, i + blockSize);
    const encrypted = await crypto.subtle.encrypt(
      { name: 'AES-CBC', iv: zeroIv },
      cryptoKey,
      block,
    );
    result.set(new Uint8Array(encrypted).slice(0, blockSize), i);
  }

  // Base64 编码
  return btoa(String.fromCharCode(...result));
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
