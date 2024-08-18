import axios from 'axios'
import cookie from 'js-cookie'
import { userStore } from '../stores/counter'

const http = axios.create({
    baseURL: 'https://flood.axtl.cn:520/api',
    timeout: 50000,
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer yourToken'
    }
})

http.interceptors.request.use(config => {
    const token = cookie.get('floodtoken')
    if (token) {
        config.headers['Authorization'] = `Bearer ${token}` // 将 Token 添加到请求头
    }
    return config
}, error => {
    return Promise.reject(error)
})

http.interceptors.response.use(response => {
    return response
}, error => {
    if (error.response) {
        // Vue.prototype.$Message.error(error.response.data.msg)
    }
  return Promise.reject(error)
})

export default http