<script setup lang="ts">
import { reactive, ref, watch } from 'vue'
import md from './markdown'
import CardTitle from '@/components/CardTitle.vue'
import { Menu } from '@element-plus/icons-vue'

const currentRadio = ref('sensor')
const radios = [
  {
    label: 'resource.docs.radio.sensor',
    value: 'sensor',
  },
  {
    label: 'resource.docs.radio.upload',
    value: 'upload',
  },
]

const htmlContent = reactive<Record<string, string>>({})

watch(
  currentRadio,
  async (radio) => {
    if (!htmlContent[radio]) {
      const mdResponse = await fetch(`/md/${radio}.md`)
      const mdContent = await mdResponse.text()
      htmlContent[radio] = md.render(mdContent)
    }
  },
  {
    immediate: true,
  },
)
</script>

<template>
  <el-card class="h-full" body-class="overflow-y-auto">
    <template #header>
      <card-title :title="$t('data.addSensor.title')" :icon="Menu">
        <el-radio-group v-model="currentRadio">
          <el-radio-button
            v-for="radio in radios"
            :key="radio.value"
            :label="$t(radio.label)"
            :value="radio.value"
          />
        </el-radio-group>
      </card-title>
    </template>
    <div class="flex-grow-1 overflow-y-auto" v-html="htmlContent[currentRadio]" />
  </el-card>
</template>
