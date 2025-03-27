<script setup lang="ts">
import type {MenuItem} from '@/types';
import {Cpu, DocumentCopy, Histogram} from '@element-plus/icons-vue';
import SubNavbar from '@/components/SubNavbar.vue';

const menuItems: MenuItem[] = ([
  ['navbar.monitor.dashboard', '/monitor/dashboard', Histogram],
  ['navbar.monitor.history', '/monitor/history', DocumentCopy],
  ['navbar.monitor.analyze', '/monitor/analyze', Cpu],
] as [string, string, object][]).map(([label, path, icon]) => ({label, path, icon}))
</script>

<template>
  <article class="flex-col h-100 gap-2">
    <sub-navbar title="navbar.monitor.title" :paths="menuItems" />
    <router-view class="flex-grow-1" />
  </article>
</template>
