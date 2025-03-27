<script setup lang="ts">
import type {MenuItem} from '@/types';

defineProps<{
  title: string,
  paths: MenuItem[],
}>()
</script>

<template>
  <el-card body-class="p-0">
    <el-menu router mode="horizontal" :default-active="$route.path">
      <el-menu-item class="title-text">
        {{ $t(title) }}
      </el-menu-item>
      <el-menu-item v-for="path in paths" :key="path.label" :index="path.path">
        <el-icon>
          <component :is="path.icon" />
        </el-icon>
        {{ $t(path.label) }}
      </el-menu-item>
    </el-menu>
  </el-card>
</template>
