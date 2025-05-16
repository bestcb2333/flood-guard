import axios from 'axios'
import { ElMessage } from 'element-plus'
import usePersistedStore from './stores/persisted'
import qs from 'qs'

export const apiAxios = axios.create({
  timeout: 5000,
  paramsSerializer: params => qs.stringify(params, {arrayFormat: 'repeat'})
})

apiAxios.interceptors.request.use((config) => {
  const persisted = usePersistedStore()
  const token = persisted.token
  config.baseURL = persisted.setting.apiAddr
  if (token) {
    config.headers.set('Authorization', `Bearer ${token}`)
  }
  return config
})

apiAxios.interceptors.response.use(
  (res) => {
    if (res.data.message) {
      ElMessage({type: 'success', message: res.data.message})
    }
    return res.data.data
  },
  (err) => {
    if (err.response) {
      const data = err.response.data
      if (err.response.data) {
        ElMessage({type: 'error', message: data.message})
        return Promise.reject(err)
      } else {
        return Promise.reject(err)
      }
    } else {
      return Promise.reject(err)
    }
  },
)
