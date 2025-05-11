<script setup lang="ts">
import { ChromeFilled, Menu, Opportunity, Promotion, Setting } from '@element-plus/icons-vue'
import CardTitle from '@/components/CardTitle.vue'
import { ref } from 'vue'

const input = ref('')
const useBrowse = ref(false)
const useThink = ref(false)
const isDialogShow = ref(false)
</script>

<template>
  <el-card
    shadow="never"
    class="h-full flex flex-col"
    body-class="grow flex flex-col gap-4"
    footer-class="flex max-md:flex-col gap-2"
  >
    <template #header>
      <card-title :title="$t('deepseek.title')" :icon="Menu">
        <el-button :icon="Setting" circle @click="isDialogShow = true" />
      </card-title>
    </template>

    <el-card class="bg-slate-200 self-end max-w-[60%]" shadow="never" body-class="p-4">
      魔镜魔镜告诉我，什么是内涝
    </el-card>

    <el-card class="self-start" shadow="never" body-class="p-4"> 服务器繁忙，请稍后再试。 </el-card>

    <template #footer>
      <el-input
        v-model="input"
        type="textarea"
        :autosize="{ minRows: 1, maxRows: 4 }"
        :placeholder="$t('deepseek.placeholder')"
      />
      <div class="flex flex-nowrap justify-end">
        <el-button :icon="Promotion" circle type="primary" />
        <el-button
          :icon="ChromeFilled"
          circle
          :type="useBrowse ? 'success' : ''"
          @click="useBrowse = !useBrowse"
        />
        <el-button
          :icon="Opportunity"
          circle
          :type="useThink ? 'warning' : ''"
          @click="useThink = !useThink"
        />
      </div>
    </template>
  </el-card>

  <el-dialog v-model="isDialogShow" class="w-4/5 max-w-md" :title="$t('deepseek.setting')">
    <el-form>
      <el-form-item :label="$t('deepseek.apikey')">
        <el-input />
      </el-form-item>
      <el-form-item>
        <el-button type="danger">
          {{ $t('deepseek.clear') }}
        </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
