<script setup lang="ts">
import type {GeoJSON} from 'echarts/types/src/coord/geo/geoTypes.js';
import NavBar from './NavBar.vue';
import {useUserStore, type User} from './stores/user';
import apiAxios from './utils/axios';
import useMapStore from './stores/map';

// 自动登录逻辑
const user = useUserStore()
apiAxios.get<User>('/myinfo').then(res => {
  user.val = res.data
}).catch(err => console.log(err))

// 获取地图
const map = useMapStore()
apiAxios.get<GeoJSON>('/map').then(res => {
  map.map = res.data
}).catch(err => console.log(err))
</script>

<template>
  <div class="app-layout flex gap-2">
    <nav-bar />
    <div class="flex-grow-1 h-100">
      <router-view />
    </div>
  </div>
</template>

<style scoped>
.app-layout {
  height: calc(100vh - 2rem);
}
</style>
