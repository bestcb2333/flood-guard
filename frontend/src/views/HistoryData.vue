<script setup lang="ts">
import type { History } from '@/tables'
import { Menu } from '@element-plus/icons-vue'
import CardTitle from '@/components/CardTitle.vue'
import { computed, ref, watch } from 'vue'
import { apiAxios } from '@/axios'
import useSessionStore from '@/stores/session'
import {use} from 'echarts/core'
import {LineChart} from 'echarts/charts'
import {GridComponent, LegendComponent, TitleComponent, TooltipComponent} from 'echarts/components'
import {CanvasRenderer} from 'echarts/renderers'
import VChart from 'vue-echarts'
import {formatDate} from '@/utils'
import dayjs from 'dayjs'

use([
  LineChart,
  TitleComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent,
  CanvasRenderer,
])

const session = useSessionStore()

const data = ref<History[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
watch([page, pageSize], async ([page, pageSize]) => {
  const res = await apiAxios.get<null, {
    data: History[],
    total: number,
  }>('/histories', {params: {
    page, pageSize,
  }})
  data.value = res.data
  total.value = res.total
}, {immediate: true})

const option = ref();

interface Trend {
  id: number,
  name: string,
  histories: {
    id: number,
    createdAt: Date,
    rainfall: number,
    waterlevel: number,
  }[],
}

const trends = ref<Trend[]>([])

apiAxios.get<any, Trend[]>('/trends').then(res => {
  trends.value = res
}).catch(() => {})

const trendOptions = computed(() => trends.value.map(trend => ({
  title: {
    text: trend.name,
  },
  tooltip: {
    trigger: 'axis',
  },
  xAxis: {
    type: 'category',
    data: trend.histories.map(history => dayjs(history.createdAt).format('DD')),
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      name: '降水量',
      type: 'line',
      data: trend.histories.map(history => history.rainfall),
    },
    {
      name: '水位',
      type: 'line',
      data: trend.histories.map(history => history.waterlevel),
    },
  ],
})))
</script>

<template>
  <div class="h-full flex flex-col md:flex-row gap-2">

    <el-card shadow="never" class="basis-0 grow flex flex-col" body-class="grow overflow-y-auto" footer-class="flex justify-end">
      <template #header>
        <card-title :title="$t('data.history.title')" :icon="Menu">
          <el-button circle :icon="Menu" />
          <el-button circle :icon="Menu" />
        </card-title>
      </template>

      <el-table :data="data">
        <el-table-column
          :label="$t('data.history.createdAt')"
          prop="createdAt"
          :formatter="formatDate"
        />
        <el-table-column :label="$t('data.history.region')" prop="region.name">
        </el-table-column>
        <el-table-column :label="$t('data.history.rainfall')" prop="rainfall" />
        <el-table-column :label="$t('data.history.waterlevel')" prop="waterlevel" />
        <el-table-column :label="$t('data.history.source')" prop="source" />
        <el-table-column :label="$t('data.history.description')" prop="description" show-overflow-tooltip>
        </el-table-column>
      </el-table>

      <template #footer>
        <el-pagination layout="sizes, prev, pager, next, total"
          :total="total"
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :page-sizes="[5, 10, 20, 30]"
        />
      </template>
    </el-card>

    <el-card shadow="never" class="basis-0 grow flex flex-col"
      body-class="grow grid grid-cols-1 md:grid-cols-2 gap-2 overflow-y-auto"
    >

      <template #header>
        <card-title :title="$t('data.trend')">
        </card-title>
      </template>

      <el-card v-for="option in trendOptions" :key="option.title.text" class="h-64" shadow="hover"
        body-class="h-full">

        <v-chart :option="option" autoresize />

      </el-card>

    </el-card>

  </div>
</template>
