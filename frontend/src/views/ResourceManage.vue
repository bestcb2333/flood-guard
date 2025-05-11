<script setup lang="ts">
import { ref, watch } from 'vue'
import CardTitle from '@/components/CardTitle.vue'
import type { Resource } from '@/tables'
import { MapLocation, Menu } from '@element-plus/icons-vue'
import RegionMap from '@/components/RegionMap.vue'
import { apiAxios } from '@/axios'
import { formatDate } from '@/utils'
import useSessionStore from '@/stores/session'

const activeRadio = ref<string|undefined>(undefined)

const radios = [
  ['resource.resource.radio.all', undefined],
  ['resource.resource.radio.staff', 'staff'],
  ['resource.resource.radio.sandbox', 'sandbox'],
  ['resource.resource.radio.vehicle', 'vehicle'],
  ['resource.resource.radio.pump', 'pump'],
].map(([label, id]) => ({ label, id }))

const resources = ref<Resource[]>([])
const isAddResource = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
watch([activeRadio, page, pageSize], async ([type, page, pageSize]) => {
  const res = await apiAxios.get<any, {
    data: Resource[],
    total: number,
  }>('/resources', {params: {
    type, page, pageSize,
  }})
  resources.value = res.data
  total.value = res.total
}, {immediate: true})

const isDialogOpen = ref(false)
const addForm = ref({
  type: '',
  name: '',
  quantity: 0,
  regionId: 1,
  coordinate: [0, 0],
  available: true,
})

async function addResource() {
  try {
    const res = await apiAxios.post<any, Resource>('/resources', addForm)
    resources.value.push(res)
  } catch {}
}

const session = useSessionStore()
</script>

<template>
  <div class="h-full md:grid md:grid-cols-2 md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
    <card-title class="md:col-span-2 p-4" :title="$t('resource.title')" :icon="Menu">
      <el-radio-group v-model="activeRadio">
        <el-radio-button
          v-for="radio in radios"
          :key="radio.id"
          :value="radio.id"
        >
          {{$t(radio.label??'')}}
        </el-radio-button>
      </el-radio-group>
    </card-title>

    <el-card shadow="never" v-if="isAddResource">
      <template #header>
        <card-title :title="$t('resource.add')">
          <el-button type=success round @click="isAddResource=false">
            {{$t('resource.table.title')}}
          </el-button>
        </card-title>
      </template>
      <el-form>
        <el-form-item :label="$t('resource.name')">
          <el-input />
        </el-form-item>
        <el-form-item :label="$t('resource.quantity')">
        </el-form-item>
        <el-form-item :label="$t('resource.region')">

        </el-form-item>
        <el-form-item :label="$t('resource.coordinate')">

        </el-form-item>
        <el-form-item :label="$t('resource.available')">

        </el-form-item>
      </el-form>
    </el-card>

    <el-card
      v-else
      shadow="never"
      class="grow flex flex-col"
      body-class="grow overflow-y-auto"
      footer-class="flex justify-end"
    >
      <template #header>
        <card-title :title="$t('resource.table.title')" :icon="Menu">
          <el-button round @click="isDialogOpen=true">
            {{$t('resource.add')}}
          </el-button>
        </card-title>
      </template>

      <el-table :data="resources">

        <el-table-column
          :label="$t('resource.updatedAt')"
          prop="updatedAt"
          :formatter="formatDate"
        />

        <el-table-column :label="$t('resource.name')" prop="name" />
        <el-table-column :label="$t('resource.quantity')" prop="quantity" />
        <el-table-column :label="$t('resource.region')" prop="region.name" />
        <el-table-column :label="$t('resource.coordinate')" prop="coordinate" />
        <el-table-column :label="$t('resource.available')" prop="available">
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
          :page-sizes="[5, 10, 20, 30]"
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
