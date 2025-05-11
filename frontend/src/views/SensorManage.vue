<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import RegionMap from '@/components/RegionMap.vue'
import type { Sensor } from '@/tables'
import { reactive, ref, watch } from 'vue'
import { Menu } from '@element-plus/icons-vue'
import { apiAxios } from '@/axios'
import {formatDate} from '@/utils'
import useSessionStore from '@/stores/session'

const region = ref<string|null>(null)

// 获取传感器列表
const sensors = ref<Sensor[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
watch([region, page, pageSize], async ([region, page, pageSize]) => {
  const res = await apiAxios.get<any, {
    data: Sensor[],
    total: number,
  }>('/sensors', {params: {
    region, page, pageSize,
  }})
  sensors.value = res.data
  total.value = res.total
}, {immediate: true})

const isDialogOpen = ref(false)
const addForm = reactive({
  name: '',
  coordinate: [0, 0],
  available: true,
  region: 1,
  description: '',
})

const session = useSessionStore()
</script>

<template>
  <div class="h-full md:grid md:grid-cols-2 md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
      <card-title class="md:col-span-2 p-4" :title="$t('resource.sensor.dashTitle')" :icon="Menu">
      </card-title>

      <el-card
        shadow="never"
        class="grow flex flex-col"
        body-class="grow overflow-y-auto"
        header-class="flex justify-between"
        footer-class="flex justify-end"
      >
        <template #header>
          <div class="font-bold text-xl">
            {{$t('resource.sensor.tableTitle')}}
          </div>
          <el-button round type="primary" @click="isDialogOpen=true">
            {{$t('sensor.add')}}
          </el-button>
        </template>

        <el-table :data="sensors">
          <el-table-column
            :label="$t('data.sensors.createdAt')"
            prop="createdAt"
            :formatter="formatDate"
          />
          <el-table-column :label="$t('data.sensors.name')" prop="name" />
          <el-table-column :label="$t('data.sensors.coordinate')" prop="coordinate" />
          <el-table-column :label="$t('data.sensors.available')" prop="available">
            <template #default="scope">
              <el-tag :type="scope.row.available?'success':'danger'">
                {{$t(scope.row.available?'global.yes':'global.no')}}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="$t('data.sensors.region')" prop="region.name" />
          <el-table-column :label="$t('data.sensors.description')" prop="description" show-overflow-tooltip />
        </el-table>

        <template #footer>
          <el-pagination layout="sizes, prev, pager, next, total"
            :total="total" v-model:current-page="page" v-model:page-size="pageSize"
            :page-sizes="[5, 10, 20, 30]"
          />
        </template>
      </el-card>

      <el-card shadow="never" class="flex flex-col" body-class="grow">
        <template #header>
          <card-title :title="$t('resource.sensor.mapTitle')" :icon="Menu"> </card-title>
        </template>
        <region-map show-markers :markers="sensors" />
      </el-card>
  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('sensor.add')">
    <el-form v-model="addForm">
      <el-form-item :label="$t('sensor.name')">
        <el-input v-model="addForm.name" />
      </el-form-item>
      <el-form-item :label="$t('sensor.coordinate')">
        <el-input-number v-model="addForm.coordinate[0]" />
        <el-input-number v-model="addForm.coordinate[1]" class="ms-4" />
      </el-form-item>
      <el-form-item :label="$t('sensor.available')">
        <el-switch v-model="addForm.available" />
      </el-form-item>
      <el-form-item :label="$t('sensor.region')">
        <el-select v-model="addForm.region">
          <el-option
            v-for="region in session.regions"
            :key="region.id"
            :label="region.name"
            :value="region.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('sensor.description')">
        <el-input type="textarea" v-model="addForm.description" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" round class="ms-auto">
          {{$t('global.confirm')}}
        </el-button>
        <el-button round @click="isDialogOpen=false">
          {{$t('global.cancel')}}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<style scoped></style>
