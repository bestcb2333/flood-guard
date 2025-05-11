import { defineStore } from 'pinia'
import type { Region, User } from '@/tables'
import { ref } from 'vue'
import { apiAxios } from '@/axios'

export const useSessionStore = defineStore('session', () => {

  const user = ref<User | null>(null)
  async function loadUser() {
    user.value = await apiAxios.get<any, User>('/myinfo')
  }

  const regions = ref<Region[]>([])
  const total = ref(0)
  async function loadMap() {
    const res = await apiAxios.get<any, {
      data: Region[],
      total: number,
    }>('/regions', {params: {pageSize: 100}})
    regions.value = res.data
    total.value = res.total
  }

  return { user, loadUser, regions, loadMap, apiAxios }
})

export default useSessionStore
