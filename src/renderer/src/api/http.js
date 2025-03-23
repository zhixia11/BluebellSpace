import axios from 'axios'
import { Notification } from '@arco-design/web-vue'

const http = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  }
})

http.interceptors.response.use(response => {
  if (response.data.code !== 0) {
    Notification.error({
      content: response.data.message,
      duration: 5000
    })
    return Promise.reject(response.data.message)
  } else {
    return response.data.data
  }
}, error => {
  Notification.error({
    content: `网络跑偏啦, 请稍后再试, ${error}`
  })
  return Promise.reject(error)
})

export default http