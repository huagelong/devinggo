import type { RouteRecordStringComponent } from '@vben/types';

import { getSystemInfoApi } from './user';

/** 后端返回的菜单节点原始结构 */
interface BackendRouter {
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
    /** M=菜单 B=按钮 L=链接 I=iframe */
    type: 'B' | 'I' | 'L' | 'M';
  };
  children: BackendRouter[];
}

/**
 * 递归转换后端路由数据为 vben RouteRecordStringComponent 格式
 * - type=B（按钮）：过滤掉，不生成路由（仅用于权限码控制）
 * - type=M 且有 component：页面路由，component 字符串交由 generateRoutesByBackend 解析
 * - type=M 且无 component：目录/分组路由，使用 BasicLayout 作为布局
 * - type=I：IFrame 路由
 */
function transformBackendRouters(
  routers: BackendRouter[],
): RouteRecordStringComponent[] {
  const result: RouteRecordStringComponent[] = [];

  for (const router of routers) {
    // 按钮类型仅作为权限码，不生成路由
    if (router.meta.type === 'B') {
      continue;
    }

    const children = router.children?.length
      ? transformBackendRouters(router.children)
      : undefined;

    let component: string;
    if (router.meta.type === 'I') {
      component = 'IFrameView';
    } else if (router.component) {
      // 页面组件路径，如 "system/user/index"
      // Vben 要求 component 匹配 import.meta.glob('../views/**/*.vue')
      // 因此 key 需要以 views/ 开头
      component = router.component.startsWith('views/')
        ? router.component
        : `views/${router.component}`;
    } else {
      // 无 component 的目录/分组节点，作为布局容器
      component = 'BasicLayout';
    }

    const route: RouteRecordStringComponent = {
      name: router.name,
      path: router.path,
      component,
      meta: {
        title: router.meta.title,
        icon: router.meta.icon || undefined,
        hideInMenu: router.meta.hidden,
        hideInBreadcrumb: router.meta.hiddenBreadcrumb,
      },
    };

    if (router.redirect) {
      route.redirect = router.redirect;
    }
    if (children) {
      route.children = children;
    }

    result.push(route);
  }

  return result;
}

/**
 * 获取用户所有菜单路由（从 /system/getInfo 的 routers 字段提取并转换）
 */
export async function getAllMenusApi() {
  const info = await getSystemInfoApi();
  return transformBackendRouters(info.routers as BackendRouter[]);
}
