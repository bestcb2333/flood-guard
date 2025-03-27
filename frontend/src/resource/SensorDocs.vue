<script setup lang="ts">
import {reactive, ref, watch} from 'vue';
import md from './markdown';

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

watch(currentRadio, async (radio) => {
  if (!htmlContent[radio]) {
    const mdResponse = await fetch(`/md/${radio}.md`)
    const mdContent = await mdResponse.text()
    htmlContent[radio] = md.render(mdContent)
  }
}, {
  immediate: true,
})
</script>

<template>
  <el-card class="h-100" body-class="flex-col gap-2">
    <el-radio-group v-model="currentRadio">
      <el-radio-button v-for="radio in radios" :key="radio.value"
        :label="$t(radio.label)" :value="radio.value" />
    </el-radio-group>
    <div class="flex-grow-1 overflow-y-auto" v-html="htmlContent[currentRadio]" />
  </el-card>
</template>
