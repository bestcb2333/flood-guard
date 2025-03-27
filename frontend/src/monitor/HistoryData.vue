<script setup lang="ts">
import type {TableData} from '@/types';
import apiAxios from '@/utils/axios';
import {ref} from 'vue';

const response = ref<TableData<History>|null>(null)

apiAxios.get<TableData<History>>(`/history?region=`).then(res => {
  response.value = res.data
}).catch(err => console.log(err))
</script>

<template>
  <el-card>
    <section v-if="response">

    </section>
    <el-skeleton v-else />

  </el-card>
</template>

<style scoped>

</style>
