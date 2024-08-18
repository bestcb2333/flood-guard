<template>
  <div class="w-100 h-100 d-flex flex-column">
    <div class="d-flex justify-content-between align-items-center p-2 box-style">
      <h3 class="m-0 ms-2">后台管理</h3>
      <div class="custom-style-panle-edit">
        <el-segmented v-model="activePanel" :options="panelOptions" size="large"/>
      </div>
    </div>
    <div class="flex-grow-1 mt-4 box-style pb-3 d-flex flex-column" style="overflow-y: auto;position: relative;">
      <div style="position: sticky;background: #ffffff;padding: 20px;width: 100%;top: 0px;"></div>
      <table class="w-100 p-2 mb-auto" style="margin-top: -40px;">
        <thead>
          <tr>
            <td v-for="item in fields" :key="item.key" class="text-truncate">{{ item.title }}</td>
          </tr>
        </thead>
        <tbody>
          <div v-if="datas.length===0" style="position: sticky;left: 50%; transform: translateX(-50%);">
            <el-empty class="w-100" description="暂无数据"/>
          </div>
          <tr v-else v-for="(item,id) in datas" :key="id" @contextmenu.prevent="rightClick($event,item)">
            <td v-for="field in fields" :key="field.key">{{ item[field.key] }}</td>
          </tr>
        </tbody>
      </table>
      <div class="edit-el-pagination mt-auto">
        <el-pagination background layout="prev, pager, next" :total="counts" :current-page="nowPage" @current-change="changePage"/>
      </div>
    </div>
    <div v-if="isRightClickMenuShow" class="right-click-menu" :style="{'position':'fixed','left':rightClickMenu.x,'top':rightClickMenu.y}">
      <div class="right-click-menu-item" @click="edit">编辑</div>
      <div class="right-click-menu-item" @click="del">删除</div>
    </div>
    <el-drawer v-model="drawer" :show-close="false" title="编辑数据" size="40%" @close="cancelEdit">
      <el-form :model="nowRightClickChose" label-width="100px" label-position="top">
        <el-form-item v-if="nowRightClickChose" v-for="item in fields" :key="item.key" :label="item.title" :prop="item.key">
          <el-input size="large" v-model="nowRightClickChose[item.key]"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="drawer=false">取消</el-button>
          <el-button type="primary" @click="confirmEdit">确认修改</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script>
import { markRaw } from 'vue'
import http from '../utils/http'

export default{
  data(){
    return{
      rightClickMenu:{},
      isRightClickMenuShow:false,
      activePanel:'/floodevent',
      panelOptions:[
        {"label": "内涝事件", "value": "/floodevent"},
        {"label": "历史信息", "value": "/historydata"},
        {"label": "公告", "value": "/notice"},
        {"label": "地区", "value": "/region"},
        {"label": "传感器状态", "value": "/sensorstatus"},
        {"label": "传感器", "value": "/sensor"},
        {"label": "用户", "value": "/user"}
      ],
      fields:[],
      datas: [],
      counts:0,
      nowPage:1,
      orignData:markRaw({}),
      nowRightClickChose:null,
      drawer:false,
    }
  },
  mounted() {
    this.handleDatas('get','get')
    // 鼠标左键点击
    window.addEventListener('click', this.rightClickClose)
    // 鼠标滚轮
    window.addEventListener('wheel', this.rightClickClose)
  },
  methods: {
    rightClick(e,item){//展示右键菜单
      // console.log(e.clientY,window.innerHeight,window.innerWidth-e.clientX)
      this.isRightClickMenuShow=true
      this.rightClickMenu={
        x:`${window.innerWidth-e.clientX<=200?e.clientX-160:e.clientX}px`,
        y:`${window.innerHeight-e.clientY<=120?e.clientY-100:e.clientY}px`
      }
      this.nowRightClickChose=item
      this.orignData=markRaw({...item})
    },
    rightClickClose(){//关闭右键菜单
      this.isRightClickMenuShow=false
      this.rightClickMenu={
        x:`0px`,
        y:`0px`
      }
    },
    handleDatas(type,link,data){
      let url
      if(type==='get'){
        url=`/${link+this.activePanel}?limit=10&page=${this.nowPage}`
      }else if(type==='post'){
        url=`/${link+this.activePanel}`
      }
      // console.log(data)
      http[type](url,data)
        .then(res=>{
          console.log(res.data)
          if(type==='get'){
            this.fields=res.data.field
            this.datas=res.data.data
            this.counts=res.data.count
          }else if(type==='post'){
            console.log(res)
            if(link==='delete'){
              this.handleDatas('get','get')
            }else if(link==='edit'){
              this.handleDatas('get','get')
            }
            this.nowRightClickChose=null
          }
        })
    },
    changePage(page){
      this.nowPage=page
      this.handleDatas('get','get')
    },
    edit(){
      this.drawer=true
    },
    cancelEdit(){
      this.datas.forEach(item=>{
        if(item.ID==this.orignData.ID){
          Object.assign(item, this.orignData)
        }
      })
      this.nowRightClickChose=null
      this.orignData=markRaw({})
    },
    confirmEdit(){
      this.handleDatas('post','edit',this.nowRightClickChose)
    },
    del(){
      console.log('del',this.nowRightClickChose)
      this.$msgbox.confirm(
        `你确定要删除当前条目的数据吗？此操作不可撤回`,
        '警告',
        {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(() => {
          console.log([this.nowRightClickChose.ID])
          this.handleDatas('post','delete',[this.nowRightClickChose.ID])
        })
        .catch((err)=>{
          // console.log(err)
        })
    }
  },
  watch:{
    activePanel(newVal){
      this.handleDatas('get','get')
    },
  },
}
</script>

<style scoped>
.box-style{
  background: #fff;
  border-radius: 20px;
  box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
}

table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 6px;
}

td {
  padding: 10px;
  /* background: #fff; */
}

tr {
  border-radius: 12px;
  transition: 0.2s ease;
  overflow: hidden;
  &:hover{
    box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
    cursor: pointer;
  }
  &:active{
    /* box-shadow: inset 2px 2px 12px 1px rgba(0, 0, 0, 0.1); */
    background: #efefef;
    box-shadow:none;
  }
}

thead{
  position: sticky;
  font-weight: bold;
  background: #ffffff50;
  backdrop-filter: blur(8px);
  box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
  border-radius: 15px;
  top: 15px;
}

thead tr{
  &:hover{
    box-shadow: none;
    cursor:auto;
  }
  &:active{
    background: none;
  }
}

tr td:first-child {
  border-top-left-radius: 12px;
  border-bottom-left-radius: 12px;
}

tr td:last-child {
    border-top-right-radius: 12px;
    border-bottom-right-radius: 12px;
}

.right-click-menu{
  background: #ffffff57;
  backdrop-filter: blur(6px);
  border-radius: 15px;
  box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
  padding: 5px;
  width: 200px;
  .right-click-menu-item{
    padding: 8px;
    padding-left: 12px;
    margin: 4px;
    border-radius: 15px;
    transition: 0.2s ease;
    &:hover{
      cursor: pointer;
      box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
    }
    &:active{
      box-shadow: inset 2px 2px 12px 1px rgba(0, 0, 0, 0.1);
      /* background: #0d6dfd10;
      &:last-child{
        background: #ff000010;
      } */
    }
  }
}
</style>

<style>
.custom-style-panle-edit{
  transition: 0.2s ease !important;
  .el-segmented {
    --el-segmented-item-selected-color: #0d6dfd;
    --el-segmented-item-selected-bg-color: #0d6dfd00;
    --el-segmented-bg-color:#fff;
    --el-border-radius-base: 16px;
    --el-segmented-padding:1px;
    --el-segmented-item-hover-color:#434343;
    --el-segmented-item-hover-bg-color:rgba(0,0,0,0);
    --el-segmented-item-active-bg-color:rgba(0,0,0,0);
    /* box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1); */
    .el-segmented__item{
      transition: 0.2s ease;
      height: 40px;
      margin: 0 3px;
      &:hover{
        box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
      }
      &:active{
        box-shadow: inset 2px 2px 12px 1px rgba(0, 0, 0, 0.1);
      }
    }
    .el-segmented__item-selected{
      box-shadow: 5px 2px 10px 2px rgba(0, 0, 0, 0.1);
      /* box-shadow: inset 5px 2px 12px 1px rgba(0, 0, 0, 0.1); */
    }
  }
}

.edit-el-pagination{
  position: sticky;
  bottom: 0;
  transition: 0.2s ease;
  display: flex;
  justify-content: end;
  padding-right: 10px;
  .el-pagination{
    padding: 5px;
    background: #ffffff50;
    backdrop-filter: blur(8px);
    box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
    border-radius: 15px;
  }
  .number{
    transition: 0.2s ease;
    border-radius: 10px;
    color: #434343;
    background: #ffffff00 !important;
    margin: 0 5px !important;
    &:hover{
      box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
      color: #434343;
    }
    &:active{
      box-shadow: inset 2px 2px 12px 1px rgba(0, 0, 0, 0.1);
    }
  }
  .more{
    transition: 0.2s ease;
    border-radius: 10px;
    background: #ffffff00 !important;
    margin: 0 5px !important;
    &:hover{
      box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
      color: #0d6dfd;
    }
    &:active{
      box-shadow: inset 2px 2px 12px 1px rgba(0, 0, 0, 0.1);
    }
  }
  button{
    transition: 0.2s ease;
    border-radius: 10px;
    background: #ffffff00 !important;
    margin: 0 5px !important;
    &:hover{
      box-shadow: 5px 2px 12px 1px rgba(0, 0, 0, 0.1);
      color: #0d6dfd;
    }
    &:active{
      box-shadow: inset 2px 2px 12px 1px rgba(0, 0, 0, 0.1);
    }
  }
  .is-active{
    background: #fff !important;
    border-radius: 10px;
  }
}
</style>