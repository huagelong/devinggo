import homePageRoutes from './homePageRoutes';
//系统路由
const routes = [
  {
    name: 'layout',
    path: '/',
    component: () => import('@/layout/index.vue'),
    redirect: { name: 'dashboard' },
    children: homePageRoutes,
  },
  {
    name: 'login',
    path:'/login',
    component: () => import('@/views/login.vue'),
    meta: { title: '登录' },
  },
  {
    name: 'formLayout',
    path:'/formLayout',
    component: () => import('@/layout/index.vue'),
    redirect: { name: 'openForm' },
    children: [
      {
        name: 'openForm',
        path: '/openForm/:id',
        meta: {
          title: '公共表单',
          type: 'M',
        },
        component: () => import('@/layout/form.vue'),
      },
    ],
  },
  {
    path:'/:pathMatch(.*)*',
    hidden: true,
    meta: { title: '访问的页面不存在' },
    component: () => import('@/layout/404.vue'),
  },
];

export default routes;
