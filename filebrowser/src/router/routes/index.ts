import type { RouteRecordRaw } from 'vue-router'

/**
 * 路由配置
 * @description 所有路由都在这里集中管理
 */
const routes: RouteRecordRaw[] = [
  /**
   * 首页
   */
  {
    path: '/',
    name: 'home',
    component: () => import(/* webpackChunkName: "home" */ '@views/home.vue'),
    meta: {
      title: 'Home',
    },
  },
  /**
   * 子路由示例
   */
  {
    path: '/browser',
    name: 'browser',
    component: () =>
      import(/* webpackChunkName: "foo" */ '@views/browser/browser.vue'),
    meta: {
      title: 'browser',
    },
  },
  {
    path: '/show_content',
    name: 'show_content',
    component: () =>
      import(/* webpackChunkName: "foo" */ '@views/browser/show_content.vue'),
    meta: {
      title: 'show_content',
    },
  },
  {
    path: '/foo',
    name: 'foo',
    component: () =>
      import(/* webpackChunkName: "foo" */ '@cp/TransferStation.vue'),
    meta: {
      title: 'Foo',
    },
    redirect: {
      name: 'bar',
    },
    children: [
      {
        path: 'bar',
        name: 'bar',
        component: () =>
          import(/* webpackChunkName: "bar" */ '@views/foo/bar.vue'),
        meta: {
          title: 'Bar',
        },
      },
    ],
  },
]

export default routes
