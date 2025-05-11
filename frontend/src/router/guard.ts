import { useDataStore } from '@/stores/data'
import { ElMessage } from 'element-plus'
import { type NavigationGuard } from 'vue-router'

export const loginGuard: NavigationGuard = (to) => {
  const data = useDataStore()
  if (!data.user) {
    ElMessage({ type: 'error', message: '请先登录' })
    return {
      path: '/login',
      query: { redirect: to.fullPath },
    }
  }
}

export const adminGuard: NavigationGuard = () => {
  const data = useUserStore()
  if (!data.user?.admin) {
    ElMessage({ type: 'error', message: '你不是管理员' })
    return false
  }
}
