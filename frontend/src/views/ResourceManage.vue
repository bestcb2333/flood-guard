<script setup lang="ts">
import { reactive, ref, watch } from 'vue'
import CardTitle from '@/components/CardTitle.vue'
import type { Resource } from '@/tables'
import { MapLocation, Menu } from '@element-plus/icons-vue'
import RegionMap from '@/components/RegionMap.vue'
import { apiAxios } from '@/axios'
import { formatDate } from '@/utils'
import useSessionStore from '@/stores/session'
import {useRouteQuery} from '@vueuse/router'
import {useI18n} from 'vue-i18n'

const types = [
  'human', 'transport', 'equipment', 'medical', 'supply', 'others',
]

const resources = ref<Resource[]>([])
const batchManage = ref(false)
const selected = ref<number[]>([])
const total = ref(0)
const page = useRouteQuery('page', 1, {transform: Number})
const pageSize = useRouteQuery('page_size', 10, {transform: Number})
const type = useRouteQuery('type', 'all')
const regionId = useRouteQuery('region_id', 0, {transform: Number})
const available = useRouteQuery<any, boolean>('available', false, {transform: Boolean})

watch([page, pageSize, type, regionId, available], loadTable, {immediate: true})

const isDialogOpen = ref(false)
const addForm = reactive({
  type: '',
  name: '',
  quantity: 0,
  regionId: 1,
  coordinate: [0, 0],
  available: true,
})

async function loadTable() {
  const res = await apiAxios.get<any, {
    data: Resource[],
    total: number,
  }>('/resources', {params: {
    page: page.value,
    page_size: pageSize.value,
    type: type.value,
    region_id: regionId.value,
    available: available.value,
  }})
  resources.value = res.data
  total.value = res.total
}

async function deleteSelected() {
  try {
    await apiAxios.delete('/resources', {params: {id: selected.value}})
    await loadTable()
    batchManage.value = false
  } catch {}
}

async function addResource() {
 console.log(1)
  try {
    await apiAxios.post('/resources', addForm)
    await loadTable()
    isDialogOpen.value = false
  } catch {}
}

const session = useSessionStore()

const {t} = useI18n({messages: {
  zh: {
    tableTitle: '资源列表',
    updatedAt: '更新时间',
    name: '名称',
    type: '类型',
    quantity: '数量',
    region: '区域',
    coordinate: '坐标',
    available: '是否可用',
    availableOnly: '仅可用资源',
    all: '所有资源',
    human: '人力资源',
    transport: '交通运输资源',
    equipment: '设备工具资源',
    medical: '医疗物资资源',
    supply: '生活保障资源',
    others: '其他资源',
    allRegion: '所有区域',
    editingMode: '编辑模式',
    viewMode: '查看模式',
    deleteSelected: '删除所选项',
  },
}})
</script>

<template>
  <div class="h-full md:grid md:grid-cols-2 md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
    <card-title class="md:col-span-2 p-4" :title="$t('resource.title')" :icon="Menu">
      <el-segmented :options="['all',...types].map(type=>({label:t(type),value:type}))" v-model="type" />
    </card-title>

    <el-card
      shadow="never"
      class="grow flex flex-col"
      header-class="flex"
      body-class="grow overflow-y-auto"
      footer-class="flex justify-end"
    >
      <template #header>
        <div>
          {{t('tableTitle')}}
        </div>
        <el-select v-model="regionId" class="w-32 ms-auto">
          <el-option :label="t('allRegion')" :value="0"  />
          <el-option
            v-for="region in session.regions" :key="region.id"
            :label="region.name" :value="region.id"
          />
        </el-select>
        <el-button type="danger" v-if="batchManage" round class="ms-2" @click="deleteSelected">
          {{t('deleteSelected')}}
        </el-button>
        <el-switch class="ms-2" v-model="batchManage" inline-prompt
          :active-text="t('editingMode')" :inactive-text="t('viewMode')"
        />
        <el-switch v-model="available" class="ms-2" inline-prompt
          :active-text="t('availableOnly')" :inactive-text="t('all')"
        />
        <el-button round @click="isDialogOpen=true" class="ms-2">
          {{$t('resource.add')}}
        </el-button>
      </template>

      <el-table :data="resources" @selection-change="(vals:Resource[])=>selected=vals.map(val=>val.id)">

        <el-table-column v-if="batchManage" type="selection" />

        <el-table-column
          :label="$t('resource.updatedAt')"
          prop="updatedAt"
          :formatter="formatDate"
        />

        <el-table-column :label="t('name')" prop="name" />
        <el-table-column :label="t('type')" prop="type" :formatter="(_1,_2,val)=>t(val)" />
        <el-table-column :label="t('quantity')" prop="quantity" />
        <el-table-column :label="t('region')" prop="region.name" />
        <el-table-column :label="t('coordinate')" prop="coordinate" />
        <el-table-column :label="t('available')" prop="available">
          <template #default="scope">
            <el-tag :type="scope.row.available?'success':'danger'">
              {{$t(scope.row.available?'global.yes':'global.no')}}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>

      <template #footer>
        <el-pagination layout="sizes, prev, pager, next, total"
          :total="total" v-model:current-page="page" v-model:page-size="pageSize"
        />
      </template>
    </el-card>

    <el-card shadow="never" class="grow flex flex-col" body-class="grow p-0">
      <template #header>
        <card-title :title="$t('resource.map.title')" :icon="MapLocation">
        </card-title>
      </template>
      <region-map name="资源分布图"
        show-markers :markers="resources"
      />
    </el-card>
  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('resource.add')">
    <el-form v-model="addForm">
      <el-form-item :label="$t('resource.name')">
        <el-input v-model="addForm.name" />
      </el-form-item>
      <el-form-item :label="$t('resource.type')">
        <el-select v-model="addForm.type">
          <el-option
            v-for="type in types" :key="type"
            :label="t(type)" :value="type"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('resource.quantity')">
        <el-input-number v-model="addForm.quantity" />
      </el-form-item>
      <el-form-item>
        <el-select v-model="addForm.regionId">
          <el-option
            v-for="region in session.regions"
            :key="region.id"
            :label="region.name"
            :value="region.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('resource.coordinate')">
        <el-input-number v-model="addForm.coordinate[0]" />
        <el-input-number class="ms-4" v-model="addForm.coordinate[1]" />
      </el-form-item>
      <el-form-item :label="$t('resource.available')">
        <el-switch v-model="addForm.available" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" round class="ms-auto" @click="addResource">
          {{$t('global.confirm')}}
        </el-button>
        <el-button @click="isDialogOpen=false" round>
          {{$t('global.cancel')}}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
