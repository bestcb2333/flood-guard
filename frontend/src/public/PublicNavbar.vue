<script setup lang="ts">
import type {MenuItem} from '@/types';
import {DocumentCopy, Histogram} from '@element-plus/icons-vue';
import SubNavbar from '@/components/SubNavbar.vue';

const menuItems: MenuItem[] = ([
  ['navbar.public.report', '/public/report', Histogram],
  ['navbar.public.volunteer', '/public/volunteer', DocumentCopy],
] as [string, string, object][]).map(([label, path, icon]) => ({label, path, icon}))
</script>

<template>
  <sub-navbar title="navbar.public.title" :paths="menuItems" />
  <router-view />
</template>
