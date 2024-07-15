<script setup>
import NavBar from '@/components/NavBar.vue'
</script>

<template>
  <div class="w-100 h-100 d-flex justify-content-between">
    <NavBar style="width: 200px;"/>
		<div class="flex-grow-1"><RouterView /></div>
  </div>
</template>

<style>
#app {
  height: 100vh;
}
</style>
