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
        icon: DataBoard,
      },
    },
    {
      path: '/resource',
      name: 'resource',
      component: () => import('@/views/ResourceManage.vue'),
      meta: {
        icon: FirstAidKit,
      },
    },
    {
      path: '/history',
      name: 'history',
      component: () => import('@/views/HistoryData.vue'),
      meta: {
        icon: Histogram,
      },
    },
    {
      path: '/sensors',
      name: 'sensors',
      component: () => import('@/views/SensorManage.vue'),
      meta: {
        icon: Cpu,
      },
    },
    {
      path: '/event',
      name: 'event',
      component: () => import('@/views/FloodEvents.vue'),
      meta: {
        icon: Warning,
      },
    },
    {
      path: '/staff',
      name: 'staff',
      component: () => import('@/views/AllStaff.vue'),
      meta: {
        icon: User,
      },
    },
    {
      path: '/wiki',
      name: 'wiki',
      component: () => import('@/views/FloodWiki.vue'),
      meta: {
        icon: Notebook,
      },
    },
    {
      path: '/deepseek',
      name: 'deepseek',
      component: () => import('@/views/DeepSeek.vue'),
      meta: {
        icon: Search,
      },
    },
    {
      path: '/setting',
      name: 'setting',
      component: () => import('@/views/AppSetting.vue'),
      meta: {
        icon: Setting,
      },
    },
  ],
})

export default router
