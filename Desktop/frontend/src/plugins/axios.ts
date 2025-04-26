import axios from 'axios'

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8088/',
  timeout: 1000,
})

export default axiosInstance