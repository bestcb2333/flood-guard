import { createRouter, createWebHistory } from 'vue-router'
import {HomeFilled, Histogram, Search, Document, Cpu, Files, UserFilled } from '@element-plus/icons-vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '主页',
      icon: HomeFilled,
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/visualization',
      name: '数据大屏',
      icon: Histogram,
      component: () => import('../views/Visualization.vue')
    },
    {
      path: '/query',
      name: '数据查询',
      icon: Search,
      component: () => import('../views/Query.vue')
    },
    {
      path: '/documentation',
      name: '接口文档',
      icon: Document,
      component: () => import('../views/Documention.vue')
    },
    {
      path: '/analysis',
      name: '智能分析',
      icon: Cpu,
      component: () => import('../views/Analysis.vue')
    },
    {
      path: '/edit',
      name: '后台管理',
      icon: Files,
      component: () => import('../views/Edit.vue')
    },
    {
      path: '/user',
      name: '用户中心',
      icon: UserFilled,
      component: () => import('../views/User.vue')
    }
  ]
})

export default router
