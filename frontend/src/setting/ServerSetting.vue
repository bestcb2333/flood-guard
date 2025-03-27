<script setup lang="ts">
import apiAxios from '@/utils/axios';
import {ElMessage} from 'element-plus';
import {ref} from 'vue';
import useSystemStore from '@/stores/system';
import useUserStore from '@/stores/user';
import {getStorage} from '@/utils/utils';

const system = useSystemStore()
const user = useUserStore()

const file = ref<Blob | null>(null)

function UploadRegion() {
  const formData = new FormData()
  if (file.value) {
    formData.append('data', file.value)
  } else {
    ElMessage({type: 'error', message: '没有文件'})
  }
  apiAxios.postForm('/upload/region', formData).catch(err => console.log(err))
}
</script>

<template>
  <el-card>
    <el-form>
      <el-form-item :label="$t('setting.server.upload')">
        <el-upload v-if="user.val?.admin" :limit="1"
          :action="`${system.backendAddr}/upload/region`"
          :headers="{Authorization: `Bearer ${getStorage('token')}`}">
          <el-button :color="system.themeColor" :dark="system.darkMode">
            {{ $t('setting.server.upload') }}
          </el-button>
        </el-upload>
      </el-form-item>
    </el-form>
  </el-card>
</template>
