<script setup lang="ts">
import { ref } from 'vue'
import CardTitle from '@/components/CardTitle.vue'
import { Close } from '@element-plus/icons-vue'
import {wikiGroups} from '@/wiki/wikiGroup';
import md from '@/wiki/markdown';

const currentGroup = ref<number|null>(null)
const currentDoc = ref<number|null>(null)

const content = `
城市内涝是指由于降水量过大或排水系统不畅，导致城市区域内积水无法及时排出，从而影响正常的生活、交通和生产活动。它是现代城市中普遍存在的一种环境问题，尤其是在暴雨等极端气候事件发生时。以下是城市内涝的主要成因及分类。

### 一、城市内涝的成因

1. **降水量过大或极端降雨**
   短时间内大规模的暴雨或持续性降水会超过城市排水系统的承载能力，造成积水。随着气候变化，极端天气事件愈发频繁，降水量的增加是内涝发生的重要因素。

2. **城市排水系统设计不合理或老化**
   许多城市的排水系统设计不合理，缺乏有效的分流系统，或者在建设时考虑的排水标准较低，随着城市化进程的加速，这些系统无法适应现代城市的需求。此外，排水管网老化、破损，也使得排水效率大大降低。

3. **城市化进程加速**
   城市化导致大量的土地被硬化，建筑物、道路等人工表面增加，原本能够吸收水分的土壤被覆盖，导致降水无法迅速渗透到地下，增加了地表径流，进而加剧了内涝现象。

4. **绿地面积减少**
   随着城市化和土地开发，绿地面积大幅减少，树木和植被的覆盖作用减弱，水分蒸发和渗透作用减少，从而使降水容易积聚在地面，导致内涝。

5. **地势低洼或地下水位高**
   城市某些地区的地势较低，或地下水位较高，这些地方在降雨时容易积水，排水不畅时容易形成内涝。

6. **人为因素**
   排水管道被垃圾或杂物堵塞，或者市民将污水排入雨水管道，导致管道堵塞，使排水能力减弱，进一步加剧内涝。

### 二、城市内涝的分类

根据内涝的发生特点和表现形式，城市内涝可以分为以下几类：

1. **暴雨型内涝**
   暴雨型内涝是指由于降水量极大，超过了城市排水系统的处理能力，导致短时间内的积水。此类内涝一般持续时间较短，降雨过后水位会逐渐回落，但可能会影响交通和生活秩序。

2. **排水系统故障型内涝**
   排水系统故障型内涝是由于排水管道堵塞、损坏或设计不合理，导致水无法及时排出，从而发生积水。此类内涝可能在短时间内发生，但有时会因排水系统修复不及时而持续较长时间。

3. **持续性降雨型内涝**
   持续性降雨型内涝是指由于长期降水或连续性的中小雨天气，导致城市排水系统负荷过大，出现积水现象。与暴雨型内涝不同，持续性降雨型内涝的特点是水位上升缓慢，但持续时间较长。

4. **地势低洼型内涝**
   地势低洼型内涝通常发生在城市内较低的地区，或者地下水位较高的地方。这类地区的排水系统可能无法完全排出积水，导致积水持续存在，形成内涝。

5. **城市热岛效应引发的内涝**
   城市热岛效应是指城市区域比周围地区温度更高，导致空气湿度增加，降水更多。城市热岛效应的加剧可能会导致降水量增加，从而引发内涝。

### 三、应对城市内涝的措施

1. **完善排水设施建设**
   改进城市排水系统的设计，扩大排水管网的容量，提高雨水排放效率，建设更加先进的雨水分流系统。

2. **增加绿色基础设施**
   提高绿化率，建设更多的透水性道路和绿地，增强城市的渗水能力，减轻地表径流。

3. **提升城市规划和管理**
   在城市规划时考虑防洪排水问题，避免在易涝区进行过度开发。同时加强对排水管网的维护和管理，定期检查和清理管道，确保其畅通。

4. **加强应急预案和公众教育**
   制定合理的内涝应急预案，提升市民应对内涝的意识，加强灾后快速恢复能力。

城市内涝是一个复杂的多因素问题，需要从多个方面进行综合治理才能有效缓解其带来的影响。
`
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
          @click="currentGroup=groupIndex;currentDoc=docIndex;"
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
      <div v-html="md.render(content)" />
    </el-card>
  </div>
</template>
