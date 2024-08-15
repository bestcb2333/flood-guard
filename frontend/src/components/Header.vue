<template>
	<div :class="['d-flex align-items-center rounded-3 mb-4 p-2 nav-home', {'justify-content-between': right}]">
		<button @click="toggleSidebarShow" class="bi bi-list active-btn" text bg>菜单</button>
		<div>
			<slot></slot>
		</div>
	</div>
</template>

<script>

import '../assets/css/active-btn.css'

export default{
    data(){
		return{

		}
    },
	props:{
		right:{},
		type: {},
		default: {},
	},
	inject: ['toggleSidebarShow'],
	methods:{
		handleToggleSidebarShow(){
			this.toggleSidebarShow()
		}
	}
}
</script>