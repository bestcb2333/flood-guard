<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import { MapLocation } from '@element-plus/icons-vue'
import {computed, ref, watch} from 'vue';
import type {User} from '@/tables';
import {apiAxios} from '@/axios';
import {formatDate} from '@/utils';
import {type EChartsOption} from 'echarts';
import VChart from 'vue-echarts'
import {use} from 'echarts/core';
import {PieChart} from 'echarts/charts';
import {LegendComponent, TitleComponent, TooltipComponent} from 'echarts/components';
import {CanvasRenderer} from 'echarts/renderers';

use([
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  CanvasRenderer,
])

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const users = ref<User[]>([])
const region = ref<string|null>(null)

watch([region, page, pageSize], async ([_, page, pageSize]) => {
  const res = await apiAxios.get<any, {
    data: User[],
    total: number,
  }>('/users', {params: {page, pageSize}})
  users.value = res.data
  total.value = res.total
}, {immediate: true})

const roseData = computed(() => {
  const regionMap: Record<string, number> = {}
  users.value.forEach(user => {
    if (user.region) {
      regionMap[user.region.name] = (regionMap[user.region.name] || 0) + 1;
    }
  })
  return Object.entries(regionMap).map(([region, count]) => ({
    name: region,
    value: count,
  }))
})

const roseOption = computed((): EChartsOption => {
  return {
    series: [
      {
        name: 'Rose Chart',
        type: 'pie',
        radius: [50, 250],
        center: ['50%', '50%'],
        roseType: 'area',
        itemStyle: {
          borderRadius: 8,
        },
        data: roseData.value,
      },
    ],
  }
})
</script>

<template>
  <div class="h-full md:grid md:grid-cols-2 md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
    <card-title class="md:col-span-2 p-4" :title="$t('staff.title')">
    </card-title>

    <el-card shadow="never" class="flex flex-col" body-class="grow">
      <template #header>
        <card-title :title="$t('staff.volunteersMap')" :icon="MapLocation">
          <el-button>
            {{ $t('staff.becomeVolunteer') }}
          </el-button>
        </card-title>
      </template>
      <v-chart v-if="roseData.length > 0" :option="roseOption" />
    </el-card>

    <el-card
      shadow="never"
      class="grow flex flex-col"
      body-class="grow overflow-y-auto"
      footer-class="flex justify-end"
    >
      <template #header>
        <card-title :title="$t('staff.list')">
          <el-button>
            {{ $t('staff.onlyVolunteers') }}
          </el-button>
        </card-title>
      </template>
      <el-table :data="users">
        <el-table-column
          :label="$t('staff.createdAt')"
          prop="createdAt"
          :formatter="formatDate"
        />
        <el-table-column :label="$t('staff.name')" prop="name">
        </el-table-column>
        <el-table-column :label="$t('staff.region')" prop="region.name">
        </el-table-column>
      </el-table>
      <template #footer>
        <el-pagination layout="sizes, prev, pager, next, total" :total="total"
          v-model:current-page="page" v-model:page-size="pageSize"
        />
      </template>
    </el-card>
  </div>
</template>
