<!-- navLeft.vue -->
<template>
  <div>
    <el-menu
      default-active="1"
      class="left-menu"
      :collapse="isCollapse"
      @open="handleOpen"
      @close="handleClose"
    >
      <div class="text-center" style="padding: 15px 20px;">FloodGuard</div>
      <MenuItem v-for="item in menuItems" :key="item.index" :index="item.index" :label="item.label" :icon="item.icon" :to="item.to" />
      <el-menu-item index="8" @click="toggleCollapse" class="setting-box">
        <el-icon><View /></el-icon>
        <span>折叠侧边栏</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import MenuItem from './NavItem.vue'
import {
  HomeFilled,
  Histogram,
  Search,
  Document,
  Cpu,
  Files,
  UserFilled,
  View
} from '@element-plus/icons-vue'

const menuItems = [
  { index: '1', label: '主页', icon: HomeFilled, to: '/' },
  { index: '2', label: '数据可视化', icon: Histogram, to: '/data-visualization' },
  { index: '3', label: '数据查询', icon: Search, to: '/data-query' },
  { index: '4', label: '接口文档', icon: Document, to: '/documentation' },
  { index: '5', label: '智能分析', icon: Cpu, to: '/smart-analysis' },
  { index: '6', label: '数据处理', icon: Files, to: '/data-processing' },
  { index: '7', label: '用户中心', icon: UserFilled, to: '/user-center' }
]

const isCollapse = ref(false)

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}
</script>

<style scoped>
.left-menu {
  height: 100%;
  position: relative;
  .setting-box {
    position: absolute;
    bottom: 0;
  }
}
</style>

