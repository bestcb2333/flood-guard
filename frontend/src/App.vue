<script setup lang="ts">
import { reactive } from 'vue'
import { routes } from '@/router/index'
import { View, Hide } from '@element-plus/icons-vue'

const isCollapsed = reactive({ value: false })

</script>

<template>
  <div class="w-100 h-100 d-flex">

		<!-- 左侧导航栏 -->
		<el-menu default-active="1" :collapse="isCollapsed.value"
			 class="bg-primary-subtle" style="user-select: none;"
			>

			<!-- 左上角LOGO -->
			<div class="text-center p-3">
				<div v-if="!isCollapsed.value">FloodGuard</div>
				<div v-else="isCollapsed.value">FG</div>
			</div>

			<!-- 导航栏菜单的每个元素 -->
			<router-link v-for="(item, index) in routes" :key="index"
				:to="item.path" custom v-slot="{ navigate, isActive }"
			>
				<el-menu-item :index="index" @click="navigate"
					:class="{'is-active bg-primary text-light': isActive}"
					style="border-radius: 10px; margin: 5px 5px;"
				>  
				  <el-icon><component :is="item.icon" /></el-icon>
				  <span>{{ item.name }}</span>
				</el-menu-item>
			</router-link>

			<!-- 左下角隐藏菜单按钮 -->
			<el-menu-item class="position-absolute bottom-0 w-100"
				@click="isCollapsed.value = !isCollapsed.value" 
			>
				<el-icon>
					<Hide v-if="isCollapsed.value" />
					<View v-else />
				</el-icon>
				<span>折叠</span>
			</el-menu-item>
		</el-menu>

		<!-- 右侧内容部分 -->
    <div class="flex-grow-1 p-5 bg-body-tertiary">
      <RouterView />
    </div>
  </div>
</template>
