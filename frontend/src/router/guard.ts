import useUserStore from "@/stores/user";
import {ElMessage} from "element-plus";
import {type NavigationGuard} from "vue-router";

export const loginGuard: NavigationGuard = (to) => {

  const user = useUserStore()
  if (!user.val) {
    ElMessage({type: 'error', message: '请先登录'})
    return {
      path: '/login',
      query: { redirect: to.fullPath },
    }
  }
}

export const adminGuard: NavigationGuard = () => {

  const user = useUserStore()
  if (!user.val?.admin) {
    ElMessage({type: 'error', message: '你不是管理员'})
    return false
  }
}
