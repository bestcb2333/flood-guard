import { createRouter, createWebHistory } from 'vue-router'
import {
  Cpu,
  DataBoard,
  FirstAidKit,
  Warning,
  Notebook,
  Search,
  Setting,
  User,
  Histogram,
} from '@element-plus/icons-vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: () => import('@/views/DashBoard.vue'),
      meta: {
        name: 'dashboard',
        icon: DataBoard,
      },
    },
    {
      path: '/resources',
      name: 'resources',
      component: () => import('@/views/ResourceManage.vue'),
      meta: {
        name: 'resources',
        icon: FirstAidKit,
      },
    },
    {
      path: '/histories',
      name: 'histories',
      component: () => import('@/views/HistoryData.vue'),
      meta: {
        name: 'histories',
        icon: Histogram,
      },
    },
    {
      path: '/sensors',
      name: 'sensors',
      component: () => import('@/views/SensorManage.vue'),
      meta: {
        name: 'sensors',
        icon: Cpu,
      },
    },
    {
      path: '/events',
      name: 'events',
      component: () => import('@/views/FloodEvents.vue'),
      meta: {
        name: 'events',
        icon: Warning,
      },
    },
    {
      path: '/users',
      name: 'users',
      component: () => import('@/views/UserManage.vue'),
      meta: {
        name: 'users',
        icon: User,
      },
    },
    {
      path: '/wikis',
      name: 'wikis',
      component: () => import('@/views/FloodWiki.vue'),
      meta: {
        name: 'wikis',
        icon: Notebook,
      },
    },
    {
      path: '/deepseek',
      name: 'deepseek',
      component: () => import('@/views/DeepSeek.vue'),
      meta: {
        name: 'deepseek',
        icon: Search,
      },
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('@/views/AppSetting.vue'),
      meta: {
        name: 'settings',
        icon: Setting,
      },
    },
  ],
})

export default router
