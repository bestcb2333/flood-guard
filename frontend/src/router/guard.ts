import useSessionStore from '@/stores/session'
import { ElMessage } from 'element-plus'
import { type NavigationGuard } from 'vue-router'

export const loginGuard: NavigationGuard = (to) => {
  const session = useSessionStore()
  if (!session.user) {
    ElMessage({ type: 'error', message: '请先登录' })
    return {
      path: '/login',
      query: { redirect: to.fullPath },
    }
  }
}

export const adminGuard: NavigationGuard = () => {
  const session = useSessionStore()
  if (!session.user?.admin) {
    ElMessage({ type: 'error', message: '你不是管理员' })
    return false
  }
}
