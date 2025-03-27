<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue';
import RegionMap from '@/components/RegionMap.vue';
import TableComponent from '@/components/TableComponent.vue';
import useMapStore from '@/stores/map';
import type {TableData} from '@/types';
import apiAxios from '@/utils/axios';
import {ref} from 'vue';

const map = useMapStore()

const events = ref<TableData<Event>|null>(null)
apiAxios.get<TableData<Event>>('/event').then(res => {
  events.value = res.data
}).catch(err => console.log(err))

const tableFields = ['id', 'createdAt', 'startTime', 'endTime', 'user.name', 'severity',
'description']
</script>

<template>
  <article class="flex gap-2 overflow-y-auto">
    <el-card class="basis-20">
      <card-title :title="$t('event.history.statusTitle')">
        <Menu />
      </card-title>
    </el-card>
    <el-card class="basis-40">
      <card-title :title="$t('event.history.mapTitle')">
        <Menu />
      </card-title>
      <region-map :name="$t('event.history.mapTitle')" :geojson="map.map" :data="[]" />
    </el-card>
    <el-card class="flex-grow-1">
      <card-title :title="$t('event.history.tableTitle')">
        <Menu />
      </card-title>
      <table-component v-if="events" :data="events" :fields="tableFields" />
      <el-empty v-else />
    </el-card>
  </article>
</template>

<style scoped>

</style>
