<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import RegionMap from '@/components/RegionMap.vue'
import type { Sensor } from '@/tables'
import { reactive, ref, watch } from 'vue'
import { Menu } from '@element-plus/icons-vue'
import { apiAxios } from '@/axios'
import {formatDate} from '@/utils'
import useSessionStore from '@/stores/session'
import {useRouteQuery} from '@vueuse/router'
import {useI18n} from 'vue-i18n'

// 获取传感器列表
const sensors = ref<Sensor[]>([])
const total = ref(0)
const page = useRouteQuery('page', 1, {transform: Number})
const pageSize = useRouteQuery('page_size', 10, {transform: Number})
const regionId = useRouteQuery('region_id', 0, {transform: Number})
const available = useRouteQuery<any, boolean>('available', false, {transform: Boolean})
const isDialogOpen = ref(false)
const isSelectMode = ref(false)
const selected = ref<number[]>([])

watch([page, pageSize, regionId, available], loadTable, {immediate: true})

async function loadTable() {
  const res = await apiAxios.get<any, {
    data: Sensor[],
    total: number,
  }>('/sensors', {params: {
    region_id: regionId.value,
    page: page.value,
    page_size: pageSize.value,
    available: available.value,
  }})
  sensors.value = res.data
  total.value = res.total
}

const addForm = reactive({
  name: '',
  coordinate: [0, 0],
  available: true,
  regionId: 1,
  description: '',
})

const session = useSessionStore()

const {t} = useI18n({messages: {
  zh: {
    title: '传感器管理系统',
    allRegions: '所有区域',
    allSensors: '所有传感器',
    availableOnly: '仅可用传感器',
    selectMode: '选择模式',
    viewMode: '查看模式',
    deleteSelected: '删除所选项',
  },
}})

async function deleteSelected() {
  try {
    await apiAxios.delete('/sensors', {params: {id: selected.value}})
    await loadTable()
    isSelectMode.value = false
  } catch {}
}

async function addSensor() {
  try {
    await apiAxios.post('/sensors', addForm)
    await loadTable()
    isDialogOpen.value = false
  } catch {}
}
</script>

<template>
  <div class="h-full md:grid md:grid-cols-2 md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
      <div class="md:col-span-2 p-4 flex justify-between">
        <div class="text-xl font-bold">{{t('title')}}</div>
        <el-button round type="primary" @click="isDialogOpen=true">
          {{$t('sensor.add')}}
        </el-button>
      </div>

      <el-card
        shadow="never"
        class="grow flex flex-col"
        body-class="grow overflow-y-auto"
        header-class="flex"
        footer-class="flex justify-end"
      >
        <template #header>
          <div class="font-bold text-xl">
            {{$t('resource.sensor.tableTitle')}}
          </div>
          <el-select v-model="regionId" class="ms-auto w-32">
            <el-option :label="t('allRegions')" :value="0" />
            <el-option
              v-for="region in session.regions" :key="region.id"
              :label="region.name" :value="region.id"
            />
          </el-select>
          <el-button v-if="isSelectMode" type="danger" round @click="deleteSelected" class="ms-2">
            {{t('deleteSelected')}}
          </el-button>
          <el-switch v-model="isSelectMode" inline-prompt class="ms-2"
            :active-text="t('selectMode')" :inactive-text="t('viewMode')"
          />
          <el-switch v-model="available" inline-prompt class="ms-2"
            :active-text="t('availableOnly')" :inactive-text="t('allSensors')"
          />
        </template>

        <el-table :data="sensors" @selection-change="(vals:Sensor[])=>selected=vals.map(val=>val.id)">
          <el-table-column type="selection" v-if="isSelectMode" />
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
        <el-select v-model="addForm.regionId">
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
        <el-button type="primary" round class="ms-auto" @click="addSensor">
          {{$t('global.confirm')}}
        </el-button>
        <el-button round @click="isDialogOpen=false">
          {{$t('global.cancel')}}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
