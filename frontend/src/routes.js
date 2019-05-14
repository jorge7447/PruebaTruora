export default [
    {
        path: '/',
        redirect: '/list'
    },
    {
        path: '/list',
        name: 'list',
        component: () => import('./views/List.vue')
    },
    {
        path: '/create',
        name: 'create',
        component: () => import('./views/Create.vue')
    },
    {
      path: "/detail/:id",
      name: "detail",
      component: () => import('./views/Detail.vue')
    },
]