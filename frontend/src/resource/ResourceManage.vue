<script setup lang="ts">
import apiAxios from '@/utils/axios';
import {ref, watch} from 'vue';
import CardTitle from '@/components/CardTitle.vue';
import TableComponent from '@/components/TableComponent.vue';
import RegionMap from '@/components/RegionMap.vue';
import useMapStore from '@/stores/map';
import type {Resource, TableData} from '@/types';

const map = useMapStore()

const activeRadio = ref('')
const radios = [
  ['resource.resource.radio.all', ''],
  ['resource.resource.radio.staff', 'staff'],
  ['resource.resource.radio.sandbox', 'sandbox'],
  ['resource.resource.radio.vehicle', 'vehicle'],
  ['resource.resource.radio.pump', 'pump'],
].map(([label, id]) => ({label, id}))

const tableEnums = { type: {
  sandbox: '沙袋',
} }

const resource = ref<TableData<Resource>|null>(null)
watch(activeRadio, (radio) => {
  apiAxios.get<TableData<Resource>>(`/resource?type=${radio}`).then(res => {
    resource.value = res.data
  }).catch(err => console.log(err))
}, { immediate: true })

const tableFields = ['id', 'createdAt', 'type', 'name', 'quantity', 'region.name', 'coordinate', 'available']
</script>

<template>
  <el-card class="h-100" body-class="flex-col gap-2">
    <el-radio-group v-model="activeRadio">
      <el-radio-button
        v-for="radio in radios" :key="radio.id"
        :label="$t(radio.label)" :value="radio.id" />
    </el-radio-group>
    <div class="flex-grow-1 flex gap-2 overflow-y-auto">
      <el-card class="resource-manage-map">
        <card-title title="resource.resource.map">
          <MapLocation />
        </card-title>
        <region-map name="资源分布图" :geojson="map.map" :points="[]" />
      </el-card>
      <el-card class="flex-grow-1">
        <card-title title="resource.resource.table">
          <Grid />
        </card-title>
        <table-component v-if="resource" :data="resource" :enums="tableEnums" :fields="tableFields" />
        <el-empty v-else />
      </el-card>
    </div>
  </el-card>
</template>

<style>
.resource-manage-map {
  flex-basis: 45%;
}
</style>
