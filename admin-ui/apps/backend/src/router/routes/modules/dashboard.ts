import type { RouteRecordRaw } from 'vue-router';

import { $t } from '#/locales';

const routes: RouteRecordRaw[] = [
  {
    meta: {
      icon: 'lucide:layout-dashboard',
      order: -1,
      title: $t('page.dashboard.title'),
    },
    name: 'Dashboard',
    path: '/dashboard',
    children: [
      {
        name: 'Analytics',
        path: '/analytics',
        component: () => import('#/views/dashboard/analytics/index.vue'),
        meta: {
          affixTab: true,
          icon: 'lucide:area-chart',
          title: $t('page.dashboard.analytics'),
        },
      },
      {
        name: 'Workspace',
        path: '/workspace',
        component: () => import('#/views/dashboard/workspace/index.vue'),
        meta: {
          icon: 'carbon:workspace',
          title: $t('page.dashboard.workspace'),
        },
      },
      {
        name: 'DashboardProfile',
        path: '/dashboard/profile',
        component: () => import('#/views/dashboard/profile/index.vue'),
        meta: {
          icon: 'lucide:user',
          title: $t('page.dashboard.profile'),
        },
      },
      {
        name: 'DashboardMessage',
        path: '/dashboard/message',
        component: () => import('#/views/dashboard/message/index.vue'),
        meta: {
          icon: 'lucide:bell',
          title: $t('page.dashboard.message'),
        },
      },
    ],
  },
];

export default routes;
