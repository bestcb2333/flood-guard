<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import type { Notice } from '@/tables'
import { computed, reactive, ref } from 'vue'
import RegionMap from '@/components/RegionMap.vue'
import { MapLocation, Plus, Stopwatch } from '@element-plus/icons-vue'
import { apiAxios } from '@/axios'
import type {EChartsOption} from 'echarts'
import VChart from 'vue-echarts'

// 获取公告列表
const notices = ref<Notice[]>([])
const count = ref(0)
const noticeStatus = ref<number | boolean>(false)
apiAxios.get<any, {
  data: Notice[],
  count: number,
}>('/notices', {params: {
  page: 1,
}}).then(res => {
  notices.value = res.data
  count.value = res.count
}).catch(() => {})

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

const isDialogOpen = ref(false)
const addNoticeForm = reactive({
  title: '',
  content: '',
})

async function addNotice() {
  try {
    const res = await apiAxios.post<any, Notice>('/notices', addNoticeForm)
    notices.value.push(res)
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
        <el-button round type="primary" @click="isDialogOpen=true">
          {{$t('dashboard.add')}}
        </el-button>
      </template>
      <el-card v-for="notice in notices" :key="notice.id" shadow="hover">
        <div class="font-bold text-lg">
          {{ notice.title }}
        </div>
        <div>
          {{ notice.content }}
        </div>
      </el-card>
    </el-card>

  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('dashboard.add')">
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
        <el-button round @click="isDialogOpen=false">
          {{$t('global.cancel')}}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
