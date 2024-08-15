<template>
	<HEADER right="true">
		<span v-for="(item, index) in head" :key="index">
			<a class="links" :href="item.link" :class="'text-decoration-none mx-1 active-btn ' + item.icon" style="font-weight: normal;">
				&nbsp;{{ item.name }}
			</a>
		</span>
	</HEADER>
  
	<section class="p-2 ps-3 rounded-3 nav-home">
		<h1>FloodGuard</h1>
		<h4>城市内涝隐患管理信息系统</h4>
	</section>
  
	<el-row :gutter="25" class="my-3">
		<el-col :xs="24" :sm="16">
			<CONTAINER title="概览" :center="true" class="nav-home">
				<el-row>
					<el-col v-for="(item, index) in data" :key="index" :span="6">
						<span>{{ item.name }}</span>
						<div class="text-primary fs-4">{{ item.value }}</div>
					</el-col>
				</el-row>
			</CONTAINER>
		
			<CONTAINER class="nav-home" title="内涝地图" minHeight="600px" :center="true">
				<img src="../assets/map.jpg" style="height: 300px;" />
			</CONTAINER>
		</el-col>
  
	  	<el-col :xs="24" :sm="8">
			<CONTAINER class="nav-home" title="公告" minHeight="400px">
				<div v-if="notice.length === 0" class="text-center" style="margin-top: calc(50% - 40px);">
					<div><el-icon size="50px"><Box /></el-icon></div>
					<div>空空如也</div>
				</div>
				<!-- <el-card v-else v-for="(item, index) in notice" :key="index" shadow="hover">
					<h5>{{ item.title }}</h5>
					<div class="text item">{{ item.content }}</div>
				</el-card> -->
				<div class="notice-box" v-for="(item, index) in notice" :key="index">
					<div class="notice-item">
						<h5>{{ item.title }}</h5>
						<div>{{ item.content }}</div>
					</div>
				</div>
			</CONTAINER>
		</el-col>
	</el-row>
</template>
  
<script>
import CONTAINER from '@/components/Container.vue'
import HEADER from '@/components/Header.vue'
// import apiClient from '@/axios';

import '../assets/css/active-btn.css'
import '@/style.css'
  
export default{
	components: { 
		CONTAINER, HEADER
	},
	data() {
		return {
			head: [
				{ name: '下载桌面端', link: 'https://test.mcax.cn', icon: 'bi-display' },
				{ name: '下载移动端', link: 'https://test.mcax.cn', icon: 'bi-phone' },
				{ name: 'Github', link: 'https://github.com/bestcb2333/floodguard', icon: 'bi-github' }
			],
			data: [
				{ name: '已添加区域', value: 0 },
				{ name: '内涝次数', value: 0 },
				{ name: '风险区域', value: 0 },
				{ name: '运行的传感器', value: 0 },
			],
			notice: [
				{ title: '城市排水维护', content: '为防止内涝，请居民协助清理排水口，确保排水畅通，共同维护安全环境。' },
				{ title: '道路封闭通告', content: '因暴雨造成内涝，XX路段临时封闭，请绕行，避免涉水行驶。' },
				{ title: '内涝应急准备', content: '请居民备足饮用水和食物，注意收听天气预报，避免外出，确保安全。' },
			]
		}
	},
	methods: {
		// async getNotice() {
		// 	try {
		// 		this.notice = await apiClient.get('/get/notice').data
		// 	} catch (error) {
		// 		console.error('获取公告失败：', error)
		// 	}
		// }
	},
	mounted() {
	  	// this.getNotice()
	}
}
</script>

<style scoped>
.links{
	text-decoration: none;
	color: #434343;
	background: #f7f7f7;
	transition: 0.2s ease;
	border-radius: 5px;
	padding: 8px;
	font-size: 0.9rem;
}

.notice-item{
	background: #fff;
	padding: 10px;
	border-bottom: 1px dashed #0d6efd;
}

.notice-box:last-child .notice-item {
  border-bottom: none;
}
</style>