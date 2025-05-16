import { defineStore } from 'pinia'
import {reactive, ref} from 'vue'

const usePersistedStore = defineStore('persisted', () => {

  const setting = reactive({
    apiAddr: 'https://axogc.net:7899/flood',
    darkMode: false,
    themeColor: '#28abce',
    fontSize: 14,
    language: 'zh',
  })

  const token = ref<string|null>(null)

  return {setting, token}
}, {persist: true})

export default usePersistedStore
