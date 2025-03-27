<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue';
import type {Notice, TableData} from '@/types';
import apiAxios from '@/utils/axios';
import {ref} from 'vue';
import RegionMap, {type Data} from '@/components/RegionMap.vue';
import useMapStore from '@/stores/map';

// 获取地图
const map = useMapStore()
map.load()

// 获取公告列表
const notices = ref<Notice[] | null>(null)
apiAxios.get<TableData<Notice>>('/notice?preload=User').then(res => {
  notices.value = res.data.data
}).catch(err => console.log(err))

// 获取风险指数
const factors = ref<Data[]>([])
apiAxios.get<Data[]>('/factor').then(res => {
  factors.value = res.data
}).catch(err => console.log(err))

</script>

<template>
  <article class="flex gap-2 overflow-y-auto">
    <section class="dash-left flex-col gap-2">

      <!-- 仪表盘 -->
      <el-card>
        <card-title title="monitor.dash.dashboard.title">
          <Stopwatch />
        </card-title>
        <div class="dash-dashboard">
          <el-statistic :title="$t('monitor.dash.dashboard.sensor')" :value="10" />
        </div>
      </el-card>

      <!-- 实时地图 -->
      <el-card class="flex-grow-1">
        <card-title title="monitor.dash.map.title">
          <MapLocation />
        </card-title>
          <region-map name="实时地图" :geojson="map.map" :points="[]" :data="factors" />
      </el-card>
    </section>

    <!-- 公告区 -->
    <el-card class="flex-grow-1">
      <card-title title="monitor.dash.notice.title">
        <DataLine />
      </card-title>
      <div v-if="notices">
        <div v-for="notice in notices" :key="notice.title">
          <hr>
          <b>{{ notice.title }}</b>
          <div>{{ notice.content }}</div>
        </div>
      </div>
      <el-empty v-else />
    </el-card>
  </article>
</template>

<style scoped>
.dash-left {
  flex-basis: 70%;
}
.dash-dashboard {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
