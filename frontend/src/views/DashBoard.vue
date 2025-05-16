<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import type { Notice } from '@/tables'
import { computed, reactive, ref } from 'vue'
import RegionMap from '@/components/RegionMap.vue'
import { MapLocation, Stopwatch } from '@element-plus/icons-vue'
import { apiAxios } from '@/axios'
import type {EChartsOption} from 'echarts'
import VChart from 'vue-echarts'
import {useI18n} from 'vue-i18n'
import {use} from 'echarts/core'
import {GridComponent, TooltipComponent} from 'echarts/components'
import {BarChart} from 'echarts/charts'
import {CanvasRenderer} from 'echarts/renderers'
import {watch} from 'vue'

use([
  GridComponent,
  BarChart,
  CanvasRenderer,
  TooltipComponent,
])

// 获取公告列表
const notices = ref<Notice[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const noticeStatus = ref<number | boolean>(false)

watch([page, pageSize], loadNotice, {immediate: true})

async function loadNotice() {
  try {
    const res = await apiAxios.get<any, {
      data: Notice[],
      total: number,
    }>('/notices', {params: {
      page: page.value,
      page_size: pageSize.value,
    }})
    notices.value = res.data
    total.value = res.total
  } catch {}
}

interface Factor {
  name: string,
  value: number,
}

const factors = ref<Factor[]>([])
apiAxios.get<any, Factor[]>('/factors').then(res => {
  factors.value = res
}).catch(() => {})

const barOption = computed((): EChartsOption => ({
  xAxis: {
    type: 'category',
    data: factors.value.map(factor => factor.name),
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow',
    },
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      data: factors.value.map(factor => factor.value),
      type: 'bar',
    },
  ],
}))

const isAddDialogOpen = ref(false)
const addNoticeForm = reactive({
  title: '',
  content: '',
})

async function addNotice() {
  try {
    await apiAxios.post<any, Notice>('/notices', addNoticeForm)
    await loadNotice()
    isAddDialogOpen.value = false
  } catch {}
}

const isViewDialogOpen = ref(false)
const {t} = useI18n({messages: {
  zh: {
    viewDialogTitle: '查看公告',
  },
}})

function viewNotice(notice: Notice) {
  currentNotice.value = notice
  isViewDialogOpen.value = true
}

const currentNotice = ref<Notice|null>(null)

async function deleteNotice() {
  try {
    await apiAxios.delete('/notices', {params: {id: currentNotice.value?.id}})
    await loadNotice()
    isViewDialogOpen.value = false
  } catch {}
}
</script>

<template>
  <div
    class="md:h-full md:grid md:grid-cols-3 md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2"
  >
    <!-- 仪表盘 -->
    <el-card shadow="never" class="flex flex-col"
      body-class="grow flex gap-3"
    >
      <template #header>
        <card-title :title="$t('dashboard.title')" :icon="Stopwatch"> </card-title>
      </template>
      <el-card shadow="hover" class="grow">
        <el-statistic :value="1" :title="$t('dashboard.riskRegions')" />
      </el-card>
      <el-card shadow="hover" class="grow">
        <el-statistic :value="1" :title="$t('dashboard.availableResource')" />
      </el-card>
      <el-card shadow="hover" class="grow">
        <el-statistic :value="1" :title="$t('dashboard.rainPercent')" />
      </el-card>
      <el-card shadow="hover" class="grow">
        <el-statistic :value="1" :title="$t('dashboard.availableSensor')" />
      </el-card>
      <el-card shadow="hover" class="grow">
        <el-statistic :value="1" :title="$t('dashboard.events')" />
      </el-card>
    </el-card>

    <el-card shadow="never" body-class="h-full">
      <template #header>
        {{$t('dashboard.bar')}}
      </template>
      <v-chart :option="barOption" class="w-full h-full" autoresize />
    </el-card>

    <!-- 实时地图 -->
    <el-card class="row-span-2 flex flex-col" shadow="never" body-class="grow min-h-0 p-0">
      <template #header>
        <card-title :title="$t('dashboard.mapTitle')" :icon="MapLocation">
        </card-title>
      </template>
      <region-map :series="factors" show-region-colors />
    </el-card>

    <!-- 公告列表 -->
    <el-card
      v-if="noticeStatus == false"
      shadow="never"
      class="md:col-span-2 flex flex-col"
      body-class="grow space-y-2 overflow-y-auto"
      header-class="flex justify-between"
    >
      <template #header>
        <div class="text-xl font-bold">
          {{$t('dashboard.noticeTitle')}}
        </div>
        <el-button round type="primary" @click="isAddDialogOpen=true">
          {{$t('dashboard.add')}}
        </el-button>
      </template>
      <el-card v-for="notice in notices" :key="notice.id" shadow="hover" @click="viewNotice(notice)">
        <div class="font-bold text-lg">
          {{ notice.title }}
        </div>
        <div>
          {{ notice.content }}
        </div>
      </el-card>

      <template #footer>
        <el-pagination layout="sizes, prev, pager, next, total" :total="total"
          v-model:current-page="page" v-model:page-size="pageSize"
        />
      </template>
    </el-card>

  </div>

  <el-dialog v-model="isAddDialogOpen" :title="$t('dashboard.add')">
    <el-form :model="addNoticeForm">
      <el-form-item :label="$t('dashboard.title')">
        <el-input v-model="addNoticeForm.title" />
      </el-form-item>
      <el-form-item :label="$t('dashboard.content')">
        <el-input v-model="addNoticeForm.content" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" round class="ms-auto" @click="addNotice">
          {{$t('global.confirm')}}
        </el-button>
        <el-button round @click="isAddDialogOpen=false">
          {{$t('global.cancel')}}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
  <el-dialog v-model="isViewDialogOpen" :title="t('viewDialogTitle')">
    <div v-if="currentNotice">
      <div class="text-lg font-bold">
        {{currentNotice.title}}
      </div>
      <div v-if="currentNotice.user">
        作者：{{currentNotice.user.name}}
      </div>
      <div>
        {{currentNotice.content}}
      </div>
    </div>
    <el-button type="danger" @click="deleteNotice">
      删除公告
    </el-button>
  </el-dialog>
</template>
