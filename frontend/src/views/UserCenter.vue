<script setup lang="ts">
import useSessionStore from './stores/session';


const session = useSessionStore()
</script>

<template>
  <el-card>
    <el-descriptions v-if="session.user" :title="$t('user.myinfo')">
      <el-descriptions-item :label="$t('user.username')">
        {{ session.user.name }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.userId')">
        {{ session.user.id }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.email')">
        {{ session.user.email }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.createdAt')">
        {{ session.user.createdAt.toString() }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.updatedAt')">
        {{ session.user.updatedAt.toString() }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.admin')">
        <el-tag type="success" v-if="session.user.admin">
          {{ $t('global.yes') }}
        </el-tag>
        <el-tag type="info" v-else>
          {{ $t('global.no') }}
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.regionName')">
        {{ session.user.region?.name }}
      </el-descriptions-item>
      <el-descriptions-item :label="$t('user.profile')">
        {{ session.user.profile }}
      </el-descriptions-item>
    </el-descriptions>

    <el-empty v-else :description="$t('user.notlogin')" />
  </el-card>
</template>
