// src/stores/counter.js

import { defineStore } from 'pinia';

export const userStore = defineStore('userInfo', {
  state: () => ({
    isLogin: false,
    userInfo: {
      userName:'',
      introduction:'',
      email:'',
      admin:false,
      id:-1
    }
  }),
  getters: {
    
  },
  actions: {
    login(userInfo) {
      this.isLogin=true
      this.userInfo=userInfo
    },
    logOut() {
      this.isLogin=false
      this.initUserInfo()
    },
    initUserInfo(){
      this.userInfo = {
        userName:'',
        introduction:'',
        email:'',
        admin:false,
        id:-1
      }
    }
  },
})
