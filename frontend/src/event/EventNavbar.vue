<script setup lang="ts">
import type {MenuItem} from '@/types';
import {DocumentCopy, Histogram} from '@element-plus/icons-vue';
import SubNavbar from '@/components/SubNavbar.vue';

const menuItems: MenuItem[] = ([
  ['navbar.event.history', '/event/history', Histogram],
  ['navbar.event.optimize', '/event/optimize', DocumentCopy],
] as [string, string, object][]).map(([label, path, icon]) => ({label, path, icon}))
</script>

<template>
  <article class="flex-col h-100 gap-2">
    <sub-navbar title="navbar.event.title" :paths="menuItems" />
    <router-view class="flex-grow-1" />
  </article>
</template>
