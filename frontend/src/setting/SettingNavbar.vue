<script setup lang="ts">
import SubNavbar from '@/components/SubNavbar.vue';
import type {MenuItem} from '@/types';
import {Menu} from '@element-plus/icons-vue';

const menuItems: MenuItem[] = ([
  ['/setting/client', 'setting.client.title', Menu],
  ['/setting/server', 'setting.server.title', Menu],
] as [string, string, object][]).map(([path, label, icon]) => ({path, label, icon}))
</script>

<template>
  <article class="flex-col h-100 gap-2">
    <sub-navbar title="setting.title" :paths="menuItems" />
    <router-view class="flex-grow-1" />
  </article>
</template>
