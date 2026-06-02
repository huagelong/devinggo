import type {
  ComponentRecordType,
  GenerateMenuAndRoutesOptions,
} from '@vben/types';

import { generateAccessible } from '@vben/access';
import type { RouteRecordRaw } from 'vue-router';
import type { Component } from 'vue';

import { message } from '#/adapter/tdesign';
import { getAllMenusApi } from '#/api';
import { BasicLayout, IFrameView } from '#/layouts';
import { $t } from '#/locales';

const forbiddenComponent = (): Promise<Component> => import('#/views/_core/fallback/forbidden.vue');

async function generateAccess(options: GenerateMenuAndRoutesOptions) {
  const globMap = import.meta.glob('../views/**/*.vue');

  const pageMap: ComponentRecordType = {};
  for (const [key, value] of Object.entries(globMap)) {
    pageMap[key] = value as ComponentRecordType[string];
    // 为目录结构下的 index.vue 注册 .vue 别名
    // 使后端返回的 component 值（如 views/system/logs/apiLog）能正确解析
    if (key.endsWith('/index.vue')) {
      const alias = key.replace(/\/index\.vue$/, '.vue');
      if (!(alias in pageMap)) {
        pageMap[alias] = value as ComponentRecordType[string];
      }
    }
  }

  // 修复 Vben5 generateRoutesByBackend 中 fallback 组件路径没有带 views/ 导致的 undefined
  // 导致在加载未实现的后端路由组件时 vue-router 崩溃，菜单整个空白
  const fallbackNotFound = pageMap['../views/_core/fallback/not-found.vue'];
  if (fallbackNotFound) {
    pageMap['/_core/fallback/not-found.vue'] = fallbackNotFound;
  }

  const layoutMap: ComponentRecordType = {
    BasicLayout,
    IFrameView,
  };

  return (await generateAccessible('mixed', {
    ...options,
    fetchMenuListAsync: async () => {
      message.loading({
        content: `${$t('common.loadingMenu')}...`,
        duration: 1500,
      });
      return await getAllMenusApi();
    },
    forbiddenComponent,
    layoutMap,
    pageMap,
  })) as { accessibleMenus: RouteRecordRaw[]; accessibleRoutes: RouteRecordRaw[] };
}

export { generateAccess };
