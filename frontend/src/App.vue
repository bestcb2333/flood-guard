<template>
	<!-- 左侧导航栏 -->
	<el-menu v-if="showMenu" default-active="1" :collapse="collapseMenu" class="nav-galss" style="padding: 10px;">
		<div class="text-center my-4">
			<div class="w-100" style="height: 60px">
				<div class="h-100 w-100 rounded-4 d-flex align-items-center justify-content-center" style="background: #0D6EFD;">
					<svg v-if="collapseMenu" id="图层_1" data-name="图层 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 109.12 90.9" height="25px"><defs></defs><path fill="#fff" d="M110.36,131.75a11.45,11.45,0,0,1,6.9,2.3l3.8,2.9a3.6,3.6,0,0,0,4.4,0l3.8-2.9a11.5,11.5,0,0,1,13.8,0L147,137a3.83,3.83,0,0,1,.7,5.4h0a3.94,3.94,0,0,1-5.4.8l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-3.8,2.9a11.5,11.5,0,0,1-13.8,0l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-3.8,2.9a11.5,11.5,0,0,1-13.8,0l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-3.8,2.9a11.5,11.5,0,0,1-13.8,0l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-5,3.8a3.85,3.85,0,0,1-5.4-.8,4.14,4.14,0,0,1-.8-2.9,3.88,3.88,0,0,1,1.5-2.6l5-3.8a11.5,11.5,0,0,1,13.8,0l3.8,2.9a3.92,3.92,0,0,0,4.4,0l3.8-2.9a11.5,11.5,0,0,1,13.8,0l3.9,2.9a3.75,3.75,0,0,0,4.4,0l3.8-2.9a11,11,0,0,1,6.7-2.2Zm-15.5-77.2.6.2,59.1,21.8-.2.5-13.6,38.2a12,12,0,0,1,3.6,1.7l3.8,2.9a3.88,3.88,0,0,1,1.5,2.6,3.92,3.92,0,0,1-3.3,4.4,4.35,4.35,0,0,1-2.9-.7l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-3.8,2.9a11.5,11.5,0,0,1-13.8,0l-3.8-2.9a3.92,3.92,0,0,0-4.4,0l-3.8,2.9a11.5,11.5,0,0,1-13.8,0l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-3.8,2.9a11.25,11.25,0,0,1-13.7,0l-3.8-2.9a3.6,3.6,0,0,0-4.4,0l-5,3.8a4.15,4.15,0,0,1-4.3.3,3.84,3.84,0,0,1-1.9-3.9,4.11,4.11,0,0,1,1.5-2.6l5-3.8a11.55,11.55,0,0,1,13.7,0l3.9,2.9a3.41,3.41,0,0,0,2.2.7h.4l21.4-66.1Zm4.6,10-15.8,50.5a9.08,9.08,0,0,1,2.3-.3,11.45,11.45,0,0,1,6.9,2.3l3.9,2.9a3.6,3.6,0,0,0,4.4,0l3.8-2.9a11.25,11.25,0,0,1,13.7,0l3.8,2.9a3.6,3.6,0,0,0,4.4,0l3.8-2.9c.3-.2.5-.3.8-.5a.35.35,0,0,0,.2-.1l12.9-35.3-45.1-16.6Zm19,40.5,10.1,3.7-2.7,7.4-10.1-3.7,2.7-7.4Zm-19.1-6.5,10.1,3.7-2.7,7.3-10.1-3.7,2.7-7.3Zm23.9-5.2,10.1,3.7-2.7,7.3-10.1-3.7C120.56,100.75,123.26,93.35,123.26,93.35Zm-19.2-7.1,10.1,3.7-2.7,7.3-10.1-3.7,2.7-7.3Zm23.2-6.4,10.1,3.7-2.7,7.3-10.1-3.7Zm-19.2-7,10.1,3.7-2.7,7.3-10.1-3.7,2.7-7.3Z" transform="translate(-45.44 -54.55)"/></svg>
					<span v-if="!collapseMenu" class="logo-text">FloodGuard</span>
				</div>
			</div>
		</div>

		<!-- 导航栏菜单的每个元素 -->
		<router-link v-for="(item, index) in $router.options.routes" :key="index" :to="item.path" custom v-slot="{ navigate, isActive }">
			<el-menu-item @click="navigate" :class="['rounded-4', {'is-active fw-bold': isActive}]" class="nav-btn" :style="{'color':isActive?'#fff':''}" style="margin-top: 5px;">
				<el-icon><component :is="item.icon" /></el-icon>
				<span>{{ item.name }}</span>
			</el-menu-item>
		</router-link>

		<!-- 左下角隐藏菜单按钮 -->
		<el-menu-item :class="['rounded-4', {'is-active': collapseMenu}]" class="position-absolute btn-collapseMenu" @click="toggleCollapseMenu">
			<el-icon>
				<Hide v-if="collapseMenu" />
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
  
<script>
import '@/style.css'
import './js/css-doodle.min.js'
// import { $router.options.routes } from '@/router/index.js'

export default{
	data() {
		return {
			showMenu: window.innerWidth > 768,
			collapseMenu: window.innerWidth < 1200,
		}
	},
	provide() {
		return {
			toggleSidebarShow: this.toggleSidebarShow
		}
	},
	mounted() {
		
	},
	methods: {
		toggleCollapseMenu() {
			this.collapseMenu = !this.collapseMenu
		},
		toggleSidebarShow(){
			console.log('Toggling sidebar in ancestor component');
      		this.showMenu = !this.showMenu;
		},
		handleResize() {
			this.showMenu = window.innerWidth > 768
			this.collapseMenu = window.innerWidth < 1200
		}
	},
	mounted() {
		window.addEventListener('resize', this.handleResize())
    	// console.log(this.$router.options.routes)
	},
	beforeUnmount() {
		window.removeEventListener('resize', this.handleResize())
	}
}
</script>

<style>
.el-menu--collapse{
	width: calc(var(--el-menu-icon-width) + var(--el-menu-base-level-padding)*2+20);
}

.logo-text{
	margin-left: 4px;
	color:#fff;
	font-size: 1rem;
	transition: 0.2s ease;
	font-weight: bold;
}
</style>