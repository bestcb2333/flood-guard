import { defineStore } from "pinia";

const useSystemStore = defineStore('system', {
  state: () => ({
    backendAddr: 'http://axtl.cn:8700',
    darkMode: false,
    themeColor: '#28abce',
    fontSize: 14,
    language: 'zh',
  }),
})

export default useSystemStore
