<script setup>
import Container from '@/components/Container.vue';
import {reactive} from 'vue';
import { Box } from '@element-plus/icons-vue';
import apiClient from '@/axios';

const head = reactive([
	{ name: '下载桌面端', link: 'https://test.mcax.cn', icon: 'bi-display' },
	{ name: '下载移动端', link: 'https://test.mcax.cn', icon: 'bi-phone' },
	{ name: 'Github', link: 'https://github.com/bestcb2333/floodguard', icon: 'bi-github' },
])

const data = reactive([
	{ name: '已添加区域', value: 0 },
	{ name: '内涝次数', value: 0 },
	{ name: '风险区域', value: 0 },
	{ name: '运行的传感器', value: 0 },
])

const notice = reactive([
	{ title: '城市排水维护', content: '为防止内涝，请居民协助清理排水口，确保排水畅通，共同维护安全环境。' },
	{ title: '道路封闭通告', content: '因暴雨造成内涝，XX路段临时封闭，请绕行，避免涉水行驶。' },
	{ title: '内涝应急准备', content: '请居民备足饮用水和食物，注意收听天气预报，避免外出，确保安全。' },
])

async function getNotice() {
	try {
		notice = await apiClient.get('/get/notice').data
	} catch (error) {
		console.error('获取公告失败：', error)
	}
}

</script>

<template>

	<el-row :gutter="20" class="mb-4 mx-4 justify-content-end bg-light py-2 rounded">
		<el-col v-for="(item, index) in head" :key="index" :span="3">
			<a :href="item.link" :class="'bi text-decoration-none fw-medium ' + item.icon">
				&nbsp;{{ item.name }}
			</a>
		</el-col>
	</el-row>

	<section class="text-center">
		<h1>FloodGuard</h1>
		<h4>城市内涝隐患管理信息系统</h4>
	</section>

	<el-row :gutter="25" class="my-4">
		<el-col :span="16">
			<Container title="概览" :center="true">
				<el-row>
					<el-col v-for="(item, index) in data" :key="index" :span="6">
						{{ item.name }}
						<div class="text-primary fs-4">{{ item.value }}</div>
					</el-col>
				</el-row>
			</Container>
		
			<Container title="内涝地图" minHeight="100px" :center="true">
				<img src="../assets/map.jpg" style="height: 300px;" />
			</Container>
		</el-col>

		<el-col :span="8">
			<Container title="公告" minHeight="400px">
				<div v-if="notice.length === 0" class="text-center pt-5">
					<el-icon size="100px"><Box /></el-icon>
					<div>空空如也</div>
				</div>
				<el-card v-else v-for="(item, index) in notice" :key="index" shadow="hover">
					<h5>{{ item.title }}</h5>
    			<div class="text item">{{ item.content }}</div>
				</el-card>
			</Container>
		</el-col>
	</el-row>
</template>
