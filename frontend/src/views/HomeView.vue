<template>
  <div class="w-100 h-100 d-flex justify-content-between">
    <NAVLEFT style="width: 200px;"/>
    <div class="flex-grow-1">主体内容</div>
  </div>
</template>

<script>
import NAVLEFT from '../components/navLeft.vue'
export default {
  name: 'HomeView',
  components: {
    NAVLEFT
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

<style scoped lang="scss">

</style>