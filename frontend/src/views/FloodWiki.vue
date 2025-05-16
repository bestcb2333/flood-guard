<script setup lang="ts">
import { ref } from 'vue'
import CardTitle from '@/components/CardTitle.vue'
import { Close } from '@element-plus/icons-vue'
import {wikiGroups} from '@/wiki/wikiGroup';
import md from '@/wiki/markdown';

const currentGroup = ref<number|null>(null)
const currentDoc = ref<number|null>(null)

const wikiDoc = ref<{
  path: string,
  title: string,
  content: string,
}>({
  path: '',
  title: '',
  content: '',
})
</script>

<template>
  <div class="h-full flex gap-2 px-auto">
    <el-card shadow="never" class="basis-0 grow flex flex-col"
      body-class="grow flex flex-wrap gap-3 justify-center items-start overflow-auto"
    >
      <template #header>
        <card-title :title="$t('wiki.menu')">
        </card-title>
      </template>
      <el-card shadow="never" v-for="group, groupIndex in wikiGroups" :key="group.path"
        body-class="flex flex-col gap-2 items-start"
      >
        <template #header>
          {{group.title}}
        </template>
        <el-button v-for="doc, docIndex in group.docs" :key="doc.path" class="!ml-0"
          @click="currentGroup=groupIndex;currentDoc=docIndex;wikiDoc=doc"
        >
          {{doc.title}}
        </el-button>
      </el-card>
    </el-card>

    <el-card v-if="currentGroup!==null&&currentDoc!==null"
      shadow="never" class="basis-0 grow flex flex-col"
      body-class="grow overflow-y-auto"
    >
      <template #header>
        <card-title :title="wikiGroups[currentGroup].docs[currentDoc].title">
          <el-button :icon="Close" circle @click="currentDoc, currentGroup=null, null" />
        </card-title>
      </template>
      <div v-html="md.render(wikiDoc.content)" />
    </el-card>
  </div>
</template>
