<script setup lang="ts">
import { ChromeFilled, Menu, Opportunity, Promotion, Setting } from '@element-plus/icons-vue'
import CardTitle from '@/components/CardTitle.vue'
import { ref } from 'vue'
import {type Message} from '@/tables'
import {apiAxios} from '@/axios'
import {dayjs, ElNotification} from 'element-plus'

const input = ref('')
const useBrowse = ref(false)
const useThink = ref(false)
const isDialogShow = ref(false)

const messages = ref<Message[]>([])

const questions = [
  '常见的内涝因素有哪些？',
  '如何有效地防范内涝？',
  '什么是城市内涝？',
  '哪些城市或区域容易发生内涝？',
]

async function loadMessages() {
  try {
    const res = await apiAxios.get<any, Message[]>('/deepseek')
    messages.value = res
  } catch {}
}

loadMessages()

async function sendMessage() {
  try {
    ElNotification({'type': 'success', 'message': '发送成功，请稍候，不要重复发送'})
    const res = await apiAxios.post<any, Message[]>('/deepseek', {
      message: input.value,
    }, {
      timeout: 60000,
    })
    messages.value = res
    input.value = ''
    console.log(messages.value)
  } catch {}
}

async function deleteMessages() {
  await apiAxios.delete('/deepseek')
  await loadMessages()
}
</script>

<template>
  <el-card
    shadow="never"
    class="h-full flex flex-col"
    body-class="grow flex flex-col gap-4 min-h-0 overflow-y-auto"
    footer-class="space-y-4"
  >
    <template #header>
      <card-title :title="$t('deepseek.title')" :icon="Menu">
        <el-button round @click="deleteMessages">
          清空历史记录
        </el-button>
        <el-button :icon="Setting" circle @click="isDialogShow = true" />
      </card-title>
    </template>

    <el-card v-for="message in messages" :key="message.id" shadow="hover" body-class="p-4"
      :class="message.role==='user'?'bg-slate-200 self-end max-w-[60%] shrink-0':'self-start shrink-0'">
      <div class="text-sm text-slate-600">
        {{dayjs(message.createdAt).format('MM月DD日 HH:mm')}}
      </div>
      <div>
        {{message.content}}
      </div>
    </el-card>

    <template #footer>
      <div>
        <el-button v-for="question in questions" :key="question" round @click="input=question">
          {{question}}
        </el-button>
      </div>
      <div class="flex max-md:flex-col gap-2">
        <el-input
          v-model="input"
          type="textarea"
          :autosize="{ minRows: 1, maxRows: 4 }"
          :placeholder="$t('deepseek.placeholder')"
        />
        <div class="flex flex-nowrap justify-end">
          <el-button :icon="Promotion" circle type="primary" @click="sendMessage" />
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
