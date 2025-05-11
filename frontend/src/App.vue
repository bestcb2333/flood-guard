<script setup lang="ts">
import NavBar from './NavBar.vue'
import useSessionStore from './stores/session';

const session = useSessionStore()
session.loadUser()
session.loadMap()
</script>

<template>
  <div class="h-screen p-3 flex gap-3">
    <nav-bar class="shrink-0" />
    <div class="grow no-scrollbar overflow-y-auto rounded-xl">
      <router-view />
    </div>
  </div>
</template>
