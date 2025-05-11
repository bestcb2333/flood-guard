<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import { Setting } from '@element-plus/icons-vue'
import { watch } from 'vue'
import { useI18n } from 'vue-i18n'
import useSessionStore from '@/stores/session'
import usePersistedStore from '@/stores/persisted'
import { apiAxios } from '@/axios'

const session = useSessionStore()
const persisted = usePersistedStore()
const i18n = useI18n()

watch(
  () => persisted.setting.themeColor,
  (value) => {
    document.documentElement.style.setProperty('--el-color-primary', value)
  },
)

watch(
  () => persisted.setting.fontSize,
  (value) => {
    document.documentElement.style.setProperty('--font-size', value.toString() + 'px')
  },
)

watch(
  () => persisted.setting.language,
  (value) => {
    i18n.locale.value = value
  },
)

function UploadRegion() {
  const formData = new FormData()
  /*if (file.value) {
    formData.append('data', file.value)
  } else {
    ElMessage({type: 'error', message: '没有文件'})
  }
  */
  apiAxios.postForm('/upload/region', formData).catch((err) => console.log(err))
}

async function test() {
  apiAxios.get('/ping').catch((err) => console.log(err))
}
</script>

<template>
  <el-card shadow="never" class="h-full" body-class="flex flex-col md:flex-row gap-5">
    <template #header>
      <card-title :title="$t('setting.client.title')" :icon="Setting"> </card-title>
    </template>

    <el-form :model="persisted.setting">
      <el-form-item :label="$t('setting.client.backendAddr.title')">
        <div class="flex flex-grow-1 gap-1">
          <el-input v-model="persisted.setting.apiAddr" />
          <el-button>{{ $t('setting.client.backendAddr.default') }}</el-button>
          <el-button @click="test">{{ $t('setting.client.backendAddr.test') }}</el-button>
        </div>
      </el-form-item>
      <el-form-item :label="$t('setting.client.dark.title')">
        <el-radio-group v-model="persisted.setting.darkMode">
          <el-radio :value="false">{{ $t('setting.client.dark.false') }}</el-radio>
          <el-radio :value="true">{{ $t('setting.client.dark.true') }}</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item :label="$t('setting.client.theme')">
        <el-color-picker v-model="persisted.setting.themeColor" />
      </el-form-item>
      <el-form-item :label="$t('setting.client.fontSize')">
        <el-slider v-model="persisted.setting.fontSize" :min="8" :max="32" />
      </el-form-item>
      <el-form-item :label="$t('setting.client.language')">
        <el-radio-group v-model="persisted.setting.language">
          <el-radio value="zh">简体中文</el-radio>
          <el-radio value="en">English</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>

    <div class="bg-slate-200 h-px w-full md:w-px md:h-full" />

    <el-form>
      <el-form-item :label="$t('setting.server.upload')">
        <el-upload v-if="session.user?.admin" :limit="1" :action="`${persisted.setting.apiAddr}/upload/region`">
          <el-button :color="persisted.setting.themeColor" :dark="persisted.setting.darkMode">
            {{ $t('setting.server.upload') }}
          </el-button>
        </el-upload>
      </el-form-item>
    </el-form>
  </el-card>
</template>
