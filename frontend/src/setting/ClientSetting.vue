<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue';
import useSystemStore from '@/stores/system';
import apiAxios from '@/utils/axios';
import {watch} from 'vue';
import { useI18n } from 'vue-i18n';
const system = useSystemStore()
const i18n = useI18n()

watch(() => system.darkMode, (value) => {
  if (value) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
})

watch(() => system.themeColor, (value) => {
  document.documentElement.style.setProperty('--el-color-primary', value)
})

watch(() => system.fontSize, (value) => {
  document.documentElement.style.setProperty('--font-size', value.toString() + 'px')
})

watch(() => system.language, (value) => {
  i18n.locale.value = value
})

async function test() {
  apiAxios.get('/ping').catch(err => console.log(err))
}
</script>

<template>
  <el-card class="100vh-2rem">
    <card-title title="setting.client.title">
      <Setting />
    </card-title>
    <el-form :model="system.$state">
      <el-form-item :label="$t('setting.client.backendAddr.title')">
        <div class="flex flex-grow-1 gap-1">
          <el-input v-model="system.backendAddr" />
          <el-button>{{ $t('setting.client.backendAddr.default') }}</el-button>
          <el-button @click="test">{{ $t('setting.client.backendAddr.test') }}</el-button>
        </div>
      </el-form-item>
      <el-form-item :label="$t('setting.client.dark.title')">
        <el-radio-group v-model="system.darkMode">
          <el-radio :value="false">{{ $t('setting.client.dark.false') }}</el-radio>
          <el-radio :value="true">{{ $t('setting.client.dark.true') }}</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item :label="$t('setting.client.theme')">
        <el-color-picker v-model="system.themeColor" />
      </el-form-item>
      <el-form-item :label="$t('setting.client.fontSize')">
        <el-slider v-model="system.fontSize" :min="8" :max="32" />
      </el-form-item>
      <el-form-item :label="$t('setting.client.language')">
        <el-radio-group v-model="system.language">
          <el-radio value="zh">简体中文</el-radio>
          <el-radio value="en">English</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
  </el-card>
</template>
