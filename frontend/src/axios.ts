import axios from 'axios'
import { ElMessage } from 'element-plus'
import usePersistedStore from './stores/persisted'

export const apiAxios = axios.create({
  baseURL: 'http://axtl.cn:8700',
  timeout: 5000,
})

apiAxios.interceptors.request.use((config) => {
  const {token} = usePersistedStore()
  if (token) {
    config.headers.set('Authorization', `Bearer ${token}`)
  }
  return config
})

apiAxios.interceptors.response.use(
  (res) => {
    return res.data.data
  },
  (err) => {
    if (err.response) {
      const message = 'Error'
      ElMessage({type: 'error', message: message})
      return Promise.reject(new Error(message))
    } else {
      return Promise.reject(err)
    }
  },
)
