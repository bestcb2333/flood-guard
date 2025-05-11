<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const routes = useRouter().options.routes

const isCollapse = ref(false)
</script>

<template>
  <el-menu
    class="min-h-full flex flex-col bg-transparent border-0 gap-4 px-2"
    router
    :collapse="isCollapse"
    :default-active="'/' + $route.path.split('/')[1]"
  >

    <el-menu-item class="font-bold text-2xl mt-4">
      <span>
        FloodGuard
      </span>
    </el-menu-item>

    <template v-for="route in routes" :key="route.path">
      <el-menu-item v-if="route.meta?.icon" :index="route.path" class="rounded-full">
        <el-icon>
          <component :is="route.meta.icon" />
        </el-icon>
        <span>{{ $t(`navbar.${route.name?.toString()}`) }}</span>
      </el-menu-item>
    </template>

    <el-menu-item @click="isCollapse = !isCollapse" class="mt-auto">
      <el-icon>
        <Hide v-if="isCollapse" />
        <View v-else />
      </el-icon>
      <template #title>
        {{ $t('navbar.hide') }}
      </template>
    </el-menu-item>
  </el-menu>
</template>
