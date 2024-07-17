<script setup lang="ts">
import { Menu } from '@element-plus/icons-vue';
defineProps({
	title: String,
	minHeight: String,
	center: Boolean,
})
</script>

<template>
	<div class="p-4 mb-3 bg-light rounded-4"	:style="{ minHeight: minHeight }">
	  <h5 class="mb-3">
			<el-icon class="me-1 text-primary"><Menu /></el-icon>
			{{ title }}
		</h5>
		<div :class="['', { 'text-center': center }]">
			<slot></slot>
		</div>
	</div>
</template>
