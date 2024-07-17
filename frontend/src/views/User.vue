<script setup>
import {reactive} from 'vue';


const hasToken = reactive({value: false})
const noToken = reactive({action: 'login'})
</script>

<template>

	<!-- 已登录的用户中心 -->
	<div v-if="hasToken.value">
		个人中心
	</div>

	<!-- 未登录的用户中心 -->
	<div v-else>
		<div v-if="noToken.action === 'login'">
			登录
		</div>

		<div v-else-if="noToken.action === 'signup'">
			注册
		</div>

		<div v-else>
			找回密码
		</div>
	</div>
</template>
