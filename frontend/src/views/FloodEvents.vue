<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import RegionMap from '@/components/RegionMap.vue'
import { Menu } from '@element-plus/icons-vue'
import { reactive, ref } from 'vue'
import type { Event } from '@/tables'
import { apiAxios } from '@/axios'
import {watch} from 'vue'
import {formatDate} from '@/utils'
import dayjs from 'dayjs'
import useSessionStore from '@/stores/session'
import {useRouteQuery} from '@vueuse/router'
import {useI18n} from 'vue-i18n'

const currentRow = ref<Event|undefined>()

const isDialogOpen = ref(false)
const regionId = useRouteQuery('region_id', 0, {transform: Number})
const severity = useRouteQuery('severity', '')
const page = useRouteQuery('page', 1, {transform: Number})
const pageSize = useRouteQuery('page_size', 10, {transform: Number})
const current = useRouteQuery<any, boolean>('current', false, {transform: Boolean})

const events = ref<Event[]>([])
const total = ref(0)
watch([regionId, severity, page, pageSize, current], loadTable, {immediate: true})

async function loadTable() {
  const res = await apiAxios.get<any, {
    data: Event[],
    total: number,
  }>('/events', {params: {
    page: page.value,
    page_size: pageSize.value,
    region_id: regionId.value,
    severity: severity.value,
    current: current.value,
  }})
  events.value = res.data
  total.value = res.total
}

const severities = [
  'safe', 'low', 'medium', 'high',
]

const severityMap: Record<string, {
  type: 'success'|'primary'|'warning'|'danger',
  color: string,
}> = {
  safe: {type: 'success', color: 'green'},
  low: {type: 'primary', color: 'yellow'},
  medium: {type: 'warning', color: 'orange'},
  high: {type: 'danger', color: 'red'},
}

const uploadForm = reactive({
  startTime: new Date(),
  endTime: new Date(),
  name: '',
  desctiption: '',
  region: 1,
  severity: 0,
})

const session = useSessionStore()

const {t} = useI18n({messages: {
  zh: {
    tableTitle: '内涝事件列表',
    startTime: '开始时间',
    endTime: '结束时间',
    region: '区域名',
    severity: '严重性',
    name: '名称',
    descripion: '描述',
    uploader: '上传者',
    safe: '无风险',
    low: '低风险',
    medium: '中风险',
    high: '高风险',
    current: '仅当前事件',
    allEvents: '所有事件',
  },
}})
</script>

<template>
  <div class="h-full md:grid md:grid-cols-[2fr_2fr_1fr] md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
    <card-title class="md:col-span-3 p-4" :title="$t('event.title')" :icon="Menu" body-class="flex gap-2">
      <el-button round type="primary" @click="isDialogOpen=true">
        {{ $t('event.report') }}
      </el-button>
    </card-title>

      <el-card
        shadow="never"
        class="grow flex flex-col"
        header-class="flex"
        body-class="grow overflow-y-auto"
        footer-class="flex justify-end"
      >
        <template #header>
          <div class="text-xl font-bold">
            {{t('tableTitle')}}
          </div>
          <el-switch v-model="current" inline-prompt :active-text="t('current')"
            :inactive-text="t('allEvents')" size="large" class="ms-auto"
          />
          <el-select class="w-32 ms-2" v-model="severity">
            <el-option :label="t('allEvents')" value="" />
            <el-option v-for="severity in severities" :key="severity"
              :label="t(severity)" :value="severity"
            />
          </el-select>
        </template>

        <el-table :data="events" highlight-current-row @current-change="val => currentRow = val">
          <el-table-column
            :label="$t('event.startTime')"
            prop="startTime"
            :formatter="formatDate"
          />
          <el-table-column
            :label="$t('event.endTime')"
            prop="endTime"
            :formatter="formatDate"
          />
          <el-table-column :label="$t('event.region')" prop="region.name" />
          <el-table-column :label="$t('event.severity')" prop="severity">
            <template #default="{row}">
              <el-tag :type="severityMap[row.severity].type">
                {{t(row.severity)}}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="$t('event.user')" prop="user.name" />
          <el-table-column :label="$t('event.name')" prop="name" show-overflow-tooltip />
        </el-table>

        <template #footer>
          <el-pagination layout="sizes, prev, pager, next, total" :total="total" />
        </template>
      </el-card>

      <el-card shadow="never" class="flex flex-col" body-class="grow">
        <template #header>
          <card-title :title="$t('event.history.mapTitle')" :icon="Menu"> </card-title>
        </template>
        <region-map :name="$t('event.history.mapTitle')" :markers="events" />
      </el-card>

      <el-card shadow="never">
        <template #header>
          <card-title :title="$t('event.history.statusTitle')" :icon="Menu"> </card-title>
        </template>
        <el-descriptions v-if="currentRow" :column="1">
          <el-descriptions-item :label="$t('event.startTime')">
            {{dayjs(currentRow.startTime).format('MM月DD日 HH:mm')}}
          </el-descriptions-item>
          <el-descriptions-item :label="$t('event.endTime')">
            {{dayjs(currentRow.endTime).format('MM月DD日 HH:mm')}}
          </el-descriptions-item>
          <el-descriptions-item :label="$t('event.severity')">
            <el-tag :type="severityMap[currentRow.severity].type">
              {{t(currentRow.severity)}}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item :label="$t('event.user')">
            {{currentRow.user?.name}}
          </el-descriptions-item>
          <el-descriptions-item :label="$t('event.name')">
            {{currentRow.name}}
          </el-descriptions-item>
          <el-descriptions-item :label="$t('event.description')">
            {{currentRow.description}}
          </el-descriptions-item>
        </el-descriptions>
        <el-empty v-else :description="$t('event.pleaseSelect')" />
      </el-card>
  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('event.report')">
    <el-form>
      <el-form-item :label="$t('event.startTime')">
        <el-date-picker v-model="uploadForm.startTime" type="datetime" />
      </el-form-item>
      <el-form-item :label="$t('event.endTime')">
        <el-date-picker v-model="uploadForm.endTime" type="datetime" />
      </el-form-item>
      <el-form-item :label="$t('event.name')">
        <el-input v-model="uploadForm.name" />
      </el-form-item>
      <el-form-item :label="$t('event.region')">
        <el-select v-model="uploadForm.region">
          <el-option
            v-for="region in session.regions"
            :key="region.id"
            :label="region.name"
            :value="region.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('event.severity')">
        <el-select v-model="uploadForm.severity">
          <el-option
            v-for="severity in severities"
            :key="severity"
            :label="t('severity')"
            :value="severity"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('event.description')">
        <el-input v-model="uploadForm.desctiption" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button round type="primary" class="ms-auto">
          {{$t('global.confirm')}}
        </el-button>
        <el-button round @click="isDialogOpen=false">
          {{$t('global.cancel')}}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
