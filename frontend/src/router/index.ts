import { createRouter, createWebHistory } from 'vue-router'
import {adminGuard, loginGuard} from './guard'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'root',
      redirect: '/monitor',
    },
    {
      path: '/monitor',
      name: 'monitor',
      component: () => import('@/monitor/MonitorNavbar.vue'),
      redirect: '/monitor/dashboard',
      children: [
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('@/monitor/DashBoard.vue'),
        },
        {
          path: 'history',
          name: 'history-data',
          component: () => import('@/monitor/HistoryData.vue'),
        },
        {
          path: 'analyze',
          name: 'ai-analyze',
          component: () => import('@/monitor/AiAnalyze.vue'),
        },
      ],
    },
    {
      path: '/resource',
      name: 'resource',
      component: () => import('@/resource/ResourceNavbar.vue'),
      redirect: '/resource/sensor',
      children: [
        {
          path: 'sensor',
          name: 'sensor-manage',
          component: () => import('@/resource/SensorManage.vue'),
        },
        {
          path: 'resource',
          name: 'resource-manage',
          component: () => import('@/resource/ResourceManage.vue'),
        },
        {
          path: 'docs',
          name: 'sensor-docs',
          component: () => import('@/resource/SensorDocs.vue'),
        },
      ],
    },
    {
      path: '/event',
      name: 'event',
      component: () => import('@/event/EventNavbar.vue'),
      redirect: '/event/history',
      children: [
        {
          path: 'history',
          name: 'history-events',
          component: () => import('@/event/HistoryEvents.vue'),
        },
        {
          path: 'optimize',
          name: 'opti-analyze',
          component: () => import('@/event/OptiAnalyze.vue'),
        },
      ],
    },
    {
      path: '/public',
      name: 'public',
      component: () => import('@/public/PublicNavbar.vue'),
      children: [
        {
          path: 'report',
          name: 'event-report',
          component: () => import('@/public/EventReport.vue'),
        },
        {
          path: 'volunteer',
          name: 'volunteer',
          component: () => import('@/public/VolunteerRegister.vue'),
        },
      ],
    },
    {
      path: '/user',
      name: 'user',
      component: () => import('@/UserCenter.vue'),
      beforeEnter: loginGuard,
    },
    {
      path: '/setting',
      name: 'app-setting',
      component: () => import('@/setting/SettingNavbar.vue'),
      redirect: '/setting/client',
      children: [
        {
          path: 'client',
          name: 'setting-client',
          component: () => import('@/setting/ClientSetting.vue'),
        },
        {
          path: 'server',
          name: 'setting-server',
          component: () => import('@/setting/ServerSetting.vue'),
          beforeEnter: [loginGuard, adminGuard],
        },
      ],
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/LoginSignup.vue'),
    },
    {
      path: '/test',
      component: () => import('@/TestComponent.vue'),
    },
  ],
})

export default router
