import axios from "axios";
import {ElMessage} from "element-plus";

const apiAxios = axios.create({
  baseURL: 'http://axtl.cn:8700',
  timeout: 5000,
})

apiAxios.interceptors.request.use(config => {

  const token = localStorage.getItem('token')
  if (token) {
    config.headers.set('Authorization', `Bearer ${token}`)
  }

  return config
})

apiAxios.interceptors.response.use(res => {

  if (res.data.message) {
    switch (res.status) {
      case 200:
        ElMessage({
          message: res.data.message,
          type: 'success',
        })
        break
      case 401:
      case 403:
      case 400:
      case 500:
        let message: string
        if (res.data.error) {
          message = `${res.data.message}: ${res.data.error}`
        } else {
          message = res.data.message
        }
        ElMessage({
          message: message,
          type: 'error',
        })
        break
    }
  }

  return res.data
}, err => {
  return Promise.reject(err)
})

export default apiAxios
