<script setup lang="ts">
defineProps<{title: string}>()
</script>

<template>
  <div class="flex align-center gap-1">
    <el-icon size="1.5rem">
      <slot></slot>
    </el-icon>
    <b class="fs-2">
      {{ $t(title) }}
    </b>
  </div>
</template>
