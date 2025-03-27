<script setup lang="ts">
import {ref} from 'vue';
import {Comment, DataAnalysis, Notebook, Setting, User, Van} from '@element-plus/icons-vue';

const isCollapse = ref(false)

const menuItems = ([
  ['navbar.monitor.title', '/monitor', DataAnalysis],
  ['navbar.resource.title', '/resource', Van],
  ['navbar.event.title', '/event', Notebook],
  ['navbar.public.title', '/public', Comment],
  ['navbar.user', '/user', User],
  ['navbar.setting', '/setting', Setting],
] as [string, string, object][]).map(([label, path, icon]) => ({label, path, icon}))

</script>

<template>
  <el-card body-class="p-0">
    <el-menu router :collapse="isCollapse" :default-active="'/'+$route.path.split('/')[1]">
      <el-menu-item v-for="item in menuItems" :key="item.label" :index="item.path">
        <el-icon>
          <component :is="item.icon" />
        </el-icon>
        <template #title>
          {{ $t(item.label) }}
        </template>
      </el-menu-item>
      <el-menu-item @click="isCollapse = !isCollapse" class="hide-button">
        <el-icon>
          <Hide v-if="isCollapse" />
          <View v-else />
        </el-icon>
        <template #title>
          {{ $t('navbar.hide') }}
        </template>
      </el-menu-item>
    </el-menu>
  </el-card>
</template>
