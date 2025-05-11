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

const currentRow = ref<Event|undefined>()

const isCurrent = ref(false)
const isDialogOpen = ref(false)
const region  = ref<string|null>(null)

const events = ref<Event[]>([])
const total = ref(0)
watch(region, async () => {
  const res = await apiAxios.get<any, {
    data: Event[],
    total: number,
  }>('/events')
  events.value = res.data
  total.value = res.total
}, {immediate: true})

const severities: {
  type: 'success'|'primary'|'warning'|'danger',
  content: string,
}[] = [
  {type: 'success', content: '无风险'},
  {type: 'primary', content: '低风险'},
  {type: 'warning', content: '中风险'},
  {type: 'danger', content: '高风险'},
]

const uploadForm = reactive({
  startTime: new Date(),
  endTime: new Date(),
  name: '',
  desctiption: '',
  region: 1,
  severity: 0,
})

const session = useSessionStore()
</script>

<template>
  <div class="h-full md:grid md:grid-cols-[2fr_2fr_1fr] md:grid-rows-[auto_1fr] md:gap-2 max-md:space-y-2 max-md:overflow-y-auto">
    <card-title class="md:col-span-3 p-4" :title="$t('event.title')" :icon="Menu" body-class="flex gap-2">
      <el-button round type="primary" @click="isDialogOpen=true">
        {{ $t('event.report') }}
      </el-button>
      <el-button round :type="isCurrent ? 'warning' : ''" @click="isCurrent = !isCurrent">
        {{ $t('event.onlyCurrent') }}
      </el-button>
    </card-title>

      <el-card
        shadow="never"
        class="grow flex flex-col"
        body-class="grow overflow-y-auto"
        footer-class="flex justify-end"
      >
        <template #header>
          <card-title :title="$t('event.history.tableTitle')" :icon="Menu"> </card-title>
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
            <template #default="scope">
              <el-tag :type="severities[scope.row.severity].type">
                {{severities[scope.row.severity].content}}
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
            <el-tag :type="severities[currentRow.severity].type">
              {{severities[currentRow.severity].content}}
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
            v-for="severity, index in severities"
            :key="severity.content"
            :label="severity.content"
            :value="index"
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
