<script setup lang="ts" generic="T">
import { type TableData } from '@/types';

defineProps<{
  data: TableData<T>,
  fields: string[],
  enums?: Record<string, Record<string, string>>,
}>()
</script>

<template>
  <el-table :data="data.data">
    <el-table-column v-for="field in fields" :key="data.fields[field].comment" :prop="field"
      :label="data.fields[field].comment">
      <template #default="scope">
        <el-tag v-if="data.fields[field].type === 'bool'" :type="scope.row[field] === true ? 'success' : 'danger'">
          {{ scope.row[field] === true ? $t('table.true') : $t('table.false') }}
        </el-tag>
        <div v-if="enums?.[field]">
          {{ enums[field][scope.row[field]] }}
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>
