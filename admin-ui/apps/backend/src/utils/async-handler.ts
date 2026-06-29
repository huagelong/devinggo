import { $t } from '@vben/locales';

import { logger } from './logger';

/**
 * 包装异步操作，提供统一的错误处理和用户提示
 *
 * @param fn - 要执行的异步操作
 * @param options - 配置选项
 * @param options.errorMessage - 错误时显示给用户的提示信息，默认为 '操作失败'
 * @param options.rethrow - 是否在错误处理后重新抛出异常，默认为 false
 * @returns 异步操作的结果，如果出错则返回 null
 *
 * @example
 * // 在 modal open() 中使用
 * async function open(id: string) {
 *   const data = await withAsync(
 *     () => getUserInfoApi(id),
 *     { errorMessage: $t('common.profileUpdateFailed') },
 *   );
 *   if (!data) return;
 *   formApi.setValues(data);
 * }
 *
 * // 包装整个 open 方法
 * const open = withAsyncHandler(async (id: string) => {
 *   modalApi.open();
 *   const data = await getUserInfoApi(id);
 *   formApi.setValues(data);
 * }, { errorMessage: $t('common.operationFailed') });
 */
export async function withAsync<T>(
  fn: () => Promise<T>,
  options: {
    errorMessage?: string;
    rethrow?: boolean;
  } = {},
): Promise<T | null> {
  const { errorMessage, rethrow = false } = options;

  try {
    return await fn();
  } catch (error) {
    logger.error(errorMessage ?? $t('common.operationFailed'), error);

    if (errorMessage) {
      const { MessagePlugin } = await import('tdesign-vue-next');
      MessagePlugin.error(errorMessage);
    }

    if (rethrow) {
      throw error;
    }

    return null;
  }
}