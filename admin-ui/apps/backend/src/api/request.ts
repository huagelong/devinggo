/**
 * 该文件可自行根据业务逻辑进行调整
 */
import type { RequestClientOptions } from '@vben/request';

import { useAppConfig } from '@vben/hooks';
import { preferences } from '@vben/preferences';
import {
  authenticateResponseInterceptor,
  defaultResponseInterceptor,
  errorMessageResponseInterceptor,
  RequestClient,
} from '@vben/request';
import { useAccessStore } from '@vben/stores';

import { message } from '#/adapter/tdesign';
import { useAuthStore } from '#/store';
import { validateApiResponse, logValidationWarning } from '#/api/validation';

const { apiURL } = useAppConfig(import.meta.env, import.meta.env.PROD);

const BUSINESS_CODE_UNAUTHORIZED = 1000;

let isReAuthenticating = false;

function createRequestClient(baseURL: string, options?: RequestClientOptions) {
  const client = new RequestClient({
    ...options,
    baseURL,
  });

  async function doReAuthenticate() {
    if (isReAuthenticating) return;
    isReAuthenticating = true;
    try {
      const accessStore = useAccessStore();
      const authStore = useAuthStore();
      accessStore.setAccessToken(null);
      if (
        preferences.app.loginExpiredMode === 'modal' &&
        accessStore.isAccessChecked
      ) {
        accessStore.setLoginExpired(true);
      } else {
        await authStore.logout();
      }
    } finally {
      isReAuthenticating = false;
    }
  }

  /**
   * 刷新token逻辑（enableRefreshToken 默认关闭，此函数暂不使用）
   */
  async function doRefreshToken(): Promise<string> {
    // enableRefreshToken 默认为 false，此处不会被调用
    throw new Error('Refresh token is not supported');
  }

  function formatToken(token: null | string) {
    return token ? `Bearer ${token}` : null;
  }

  // 请求头处理
  client.addRequestInterceptor({
    fulfilled: async (config) => {
      const accessStore = useAccessStore();

      config.headers.Authorization = formatToken(accessStore.accessToken);
      config.headers['Accept-Language'] = preferences.app.locale;
      config.headers['X-App-Id'] = import.meta.env.VITE_APP_ID || 'system';

      // Prevent FormData from being transformed into JSON by Axios when a
      // default JSON content-type header is present on the client.
      if (config.data instanceof FormData && config.headers) {
        if (typeof config.headers.delete === 'function') {
          config.headers.delete('Content-Type');
          config.headers.delete('content-type');
        } else {
          delete (config.headers as Record<string, unknown>)['Content-Type'];
          delete (config.headers as Record<string, unknown>)['content-type'];
        }
      }
      return config;
    },
  });

  // 处理返回的响应数据格式
  client.addResponseInterceptor(
    defaultResponseInterceptor({
      codeField: 'code',
      dataField: 'data',
      successCode: 0,
    }),
  );

  client.addResponseInterceptor({
    fulfilled: async (response) => {
      const responseData = response?.data;
      if (responseData && !validateApiResponse(responseData)) {
        logValidationWarning('Invalid API response format', responseData);
      }
      return response;
    },
  });

  // 处理业务 code 1000：未登录或 token 过期
  client.addResponseInterceptor({
    rejected: async (error) => {
      const responseData = error?.response?.data ?? {};
      if (responseData?.code === BUSINESS_CODE_UNAUTHORIZED) {
        await doReAuthenticate();
      }
      throw error;
    },
  });

  // token过期的处理
  client.addResponseInterceptor(
    authenticateResponseInterceptor({
      client,
      doReAuthenticate,
      doRefreshToken,
      enableRefreshToken: preferences.app.enableRefreshToken,
      formatToken,
    }),
  );

  // 通用的错误处理,如果没有进入上面的错误处理逻辑，就会进入这里
  client.addResponseInterceptor(
    errorMessageResponseInterceptor((msg: string, error) => {
      // 这里可以根据业务进行定制,你可以拿到 error 内的信息进行定制化处理，根据不同的 code 做不同的提示，而不是直接使用 message.error 提示 msg
      // 当前mock接口返回的错误字段是 error 或者 message
      const responseData = error?.response?.data ?? {};
      const errorMessage = responseData?.error ?? responseData?.message ?? '';
      // 如果没有错误信息，则会根据状态码进行提示
      message.error(errorMessage || msg);
    }),
  );

  return client;
}

export const requestClient = createRequestClient(apiURL, {
  responseReturn: 'data',
});

export const baseRequestClient = new RequestClient({ baseURL: apiURL });
