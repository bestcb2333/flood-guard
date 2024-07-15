<script setup lang="ts">
import NavBar from '@/components/NavBar.vue'
import { ref } from 'vue'

const isCollapsed = ref(false)
const toggleCollapse = () => { isCollapsed.value = !isCollapsed.value }
</script>

<template>
  <div class="w-100 h-100 d-flex">
    <NavBar :is-collapsed="isCollapsed" @toggle-collapse="toggleCollapse" />
    <div :class="['content', { 'content-expanded': isCollapsed }]">
			<RouterView />
		</div>
  </div>
</template>

<style>
#app {
  height: 100vh;
}

.content {
  flex-grow: 1;
  transition: margin-left 0.3s ease;
}

.content-expanded {
  margin-left: 80px; /* 折叠后的宽度 */
}

.content {
  margin-left: 200px; /* 未折叠的宽度 */
}
</style>

