<script setup lang="ts">
import type {MenuItem} from '@/types';
import {Cpu, DocumentCopy, Histogram} from '@element-plus/icons-vue';
import SubNavbar from '@/components/SubNavbar.vue';

const menuItems: MenuItem[] = ([
  ['navbar.resource.sensor', '/resource/sensor', Histogram],
  ['navbar.resource.resource', '/resource/resource', DocumentCopy],
  ['navbar.resource.docs', '/resource/docs', Cpu],
] as [string, string, object][]).map(([label, path, icon]) => ({label, path, icon}))
</script>

<template>
  <div class="flex-col gap-2 h-100">
    <div>
      <sub-navbar title="navbar.resource.title" :paths="menuItems" />
    </div>
    <div class="flex-grow-1 overflow-y-auto">
      <router-view />
    </div>
  </div>
</template>
