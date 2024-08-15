<template>
	<div class="p-4 mb-3 rounded-3 d-flex flex-column" :style="[{minHeight: minHeight}]">
		<h5 class="mb-3">
			<div class="d-flex align-items-center">
				<el-icon class="me-1 color-main"><Menu /></el-icon>
				<span>{{ title }}</span>
			</div>
		</h5>
		<div :class="[{ 'text-center': center }]" class="flex-grow-1">
			<slot></slot>
		</div>
	</div>
</template>

<script>
export default{
    data(){
		return{

		}
    },
	props:{
		title: {},
		minHeight: {},
		center: {},
	}
}
</script>