<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue';
import RegionMap, {type Data, type MapPoint} from '@/components/RegionMap.vue';
import TableComponent from '@/components/TableComponent.vue';
import useMapStore from '@/stores/map';
import type {Sensor, TableData} from '@/types';
import apiAxios from '@/utils/axios';
import {ref} from 'vue';

const map = useMapStore()

// 获取传感器列表
const sensors = ref<TableData<Sensor>|null>(null)
const points = ref<MapPoint[]>([])
apiAxios.get<TableData<Sensor>>('/sensor').then(res => {
  sensors.value = res.data
  points.value = res.data.data.map(({name, coordinate}) => ({name: name, coordinate: coordinate}))
}).catch(err => console.log(err))


// 获取风险指数
const factors = ref<Data[]>([])
apiAxios.get<Data[]>('/factor').then(res => {
  factors.value = res.data
}).catch(err => console.log(err))

const tableFields = ['id', 'createdAt', 'name', 'description', 'coordinate', 'available', 'region.name']
</script>

<template>
  <article class="flex-col overflow-y-auto gap-2 h-100">
    <el-card>
      <card-title :title="$t('resource.sensor.dashTitle')">
        <Menu />
      </card-title>
    </el-card>

    <section class="flex overflow-y-auto gap-2 flex-grow-1">
      <el-card class="basis-40">
        <card-title :title="$t('resource.sensor.mapTitle')">
          <Menu />
        </card-title>
        <region-map :name="$t('resource.sensor.mapTitle')" :geojson="map.map" :data="factors"
          :points="points" />
      </el-card>

      <el-card class="flex-grow-1">
        <card-title :title="$t('resource.sensor.tableTitle')">
          <Menu />
        </card-title>
        <table-component v-if="sensors" :data="sensors" :fields="tableFields" />
        <el-empty v-else />
      </el-card>
    </section>
  </article>
</template>

<style scoped>

</style>
