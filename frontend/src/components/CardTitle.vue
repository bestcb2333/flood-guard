<script setup lang="ts">
defineProps<{
  title?: string
  icon?: object
  bodyClass?: string
}>()
</script>

<template>
  <div class="flex justify-between items-center flex-wrap gap-2">
    <div v-if="title" class="text-xl font-bold flex items-center gap-2">
      <el-icon v-if="icon">
        <component :is="icon" />
      </el-icon>
      <div>
        {{ title }}
      </div>
    </div>

    <div :class="bodyClass">
      <slot></slot>
    </div>
  </div>
</template>
