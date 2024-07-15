<template>
    <div>
        <el-menu
            default-active="1"
            class="left-menu"
        >
            <div class="text-center" style="padding: 15px 20px;">LOGO--</div>
            <el-menu-item index="1">
                <el-icon><HomeFilled /></el-icon>
                <span>主页</span>
            </el-menu-item>
            <el-menu-item index="2">
                <el-icon><Histogram /></el-icon>
                <span>数据可视化</span>
            </el-menu-item>
            <el-menu-item index="3">
                <el-icon><Search /></el-icon>
                <span>数据查询</span>
            </el-menu-item>
            <el-menu-item index="4">
                <el-icon><Document /></el-icon>
                <span>接口文档</span>
            </el-menu-item>
            <el-menu-item index="5">
                <el-icon><Cpu /></el-icon>
                <span>智能分析</span>
            </el-menu-item>
            <el-menu-item index="6">
                <el-icon><Files /></el-icon>
                <span>数据处理</span>
            </el-menu-item>
            <el-menu-item index="7" class="setting-box">
                <el-icon><UserFilled /></el-icon>
                <span>用户中心</span>
            </el-menu-item>
        </el-menu>
    </div>
</template>

<script>
// import cookie from 'js-cookie'

export default {
    name: 'NavLeft',
    components: {
    },
    props: {
        
    },
    data(){
        return{
            
        }
    },
    mounted(){
        
    },
    methods:{
        
    },
    watch: {
        
    },
    beforeUnmount() {
        
    }
}
</script>

<style scoped>
.left-menu{
    height: 100%;
    position: relative;
    .setting-box{
        position: absolute;
        bottom: 0;
    }
}
</style>