<!-- MenuItem.vue -->
<template>
  <router-link :to="to" custom v-slot="{ navigate, isActive, href }">
    <el-menu-item :index="index" :class="{ 'is-active': isActive }" @click="navigate">
      <el-icon><component :is="icon" /></el-icon>
      <span>{{ label }}</span>
    </el-menu-item>
  </router-link>
</template>

<script lang="ts" setup>
import { defineProps } from 'vue'

const props = defineProps({
  index: String,
  label: String,
  icon: Object,
  to: String
})
</script>

