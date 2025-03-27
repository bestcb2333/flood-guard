<script setup lang="ts">
import useUserStore from './stores/user';

const user = useUserStore()
</script>

<template>
  <el-card>
    <el-descriptions v-if="user.val" :title="$t('user.myinfo')">
      <el-descriptions-item :label="$t('user.username')">
        {{ user.val.name }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.userId')">
        {{ user.val.id }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.email')">
        {{ user.val.email }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.createdAt')">
        {{ user.val.createdAt.toString() }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.updatedAt')">
        {{ user.val.updatedAt.toString() }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.admin')">
        <el-tag type="success" v-if="user.val.admin">
          {{ $t('global.yes') }}
        </el-tag>
        <el-tag type="info" v-else>
          {{ $t('global.no') }}
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.regionName')">
        {{ user.val.region?.name }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.profile')">
        {{ user.val.profile }}
      </el-descriptions-item>
    </el-descriptions>

    <el-empty v-else :description="$t('user.notlogin')" />
  </el-card>
</template>
