<script setup lang="ts">
import { onMounted, reactive, provide } from 'vue'
import { routes } from '@/router/index'
import { View, Hide, Ship } from '@element-plus/icons-vue'
import '@/style.css'

const showMenu = reactive({ value: window.innerWidth > 768 })
const collapseMenu = reactive({ value: window.innerWidth < 1200 })

onMounted(() => window.addEventListener('resize', () => {
	showMenu.value = window.innerWidth > 768
	collapseMenu.value = window.innerWidth < 1200
}))

provide('toggleSidebarShow', () => { showMenu.value = !showMenu.value })

</script>

<template>
	<!-- 左侧导航栏 -->
	<el-menu v-if="showMenu.value" default-active="1" :collapse="collapseMenu.value" 
		class="bgc-subtle"
	>

		<!-- 左上角LOGO -->
		<div class="text-center my-4">
			<div v-if="!collapseMenu.value">FloodGuard</div>
			<el-icon v-else="collapseMenu.value"><Ship /></el-icon>
		</div>

		<!-- 导航栏菜单的每个元素 -->
		<router-link v-for="(item, index) in routes" :key="index"
			:to="item.path" custom v-slot="{ navigate, isActive }"
		>
			<el-menu-item :index="index" @click="navigate"
				:class="['rounded-4', {'is-active text-light fw-bold bgc-main': isActive}]"
			>
				<el-icon><component :is="item.icon" /></el-icon>
				<span>{{ item.name }}</span>
			</el-menu-item>
		</router-link>

		<!-- 左下角隐藏菜单按钮 -->
		<el-menu-item class="position-absolute bottom-0 w-100"
			@click="collapseMenu.value = !collapseMenu.value" 
		>
			<el-icon>
				<Hide v-if="collapseMenu.value" />
				<View v-else />
			</el-icon>
			<span>折叠</span>
		</el-menu-item>
	</el-menu>

	<!-- 右侧内容部分 -->
	<div class="flex-grow-1 overflow-y-auto p-4 bgc-bg">
		<RouterView />
	</div>
</template>
