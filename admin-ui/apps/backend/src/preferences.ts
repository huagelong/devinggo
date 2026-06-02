import { defineOverridesPreferences } from '@vben/preferences';

/**
 * @description 项目配置文件
 * 只需要覆盖项目中的一部分配置，不需要的配置不用覆盖，会自动使用默认配置
 * !!! 更改配置后请清空缓存，否则可能不生效
 */
export const overridesPreferences = defineOverridesPreferences({
  // overrides
  app: {
    name: import.meta.env.VITE_APP_TITLE,
    // 使用 mixed 模式，既保留前端内置路由（/analytics），也加载后端动态路由
    accessMode: 'mixed',
    // 登录后默认跳转首页（与后端第一个可访问菜单路径保持一致）
    defaultHomePath: '/analytics',
    enablePreferences: import.meta.env.VITE_SHOW_PREFERENCES_BUTTON !== 'false',
    enableStickyPreferencesNavigationBar: false,
    layout: 'sidebar-nav',
    watermark: true,
  },
  sidebar: {
    collapsed: false,
  },
  tabbar: {
    middleClickToClose: true,
  },
  breadcrumb: {
    showHome: true,
  },
  copyright: {
    companyName: 'devinggo',
    companySiteLink: 'https://devinggo.devinghub.com',
    icp: 'A1111aaaa',
    icpLink: 'https://devinggo.devinghub.com',
  },
  footer: {
    enable: true,
  },
  logo: {
    source: '/logo.png',
  },
  theme: {
    mode: 'light',
    semiDarkHeader: false,
    semiDarkSidebar: true,
    semiDarkSidebarSub: false,
  },
  transition: {
    enable: false,
  },
  widget: {
    timezone: false,
  },
});
