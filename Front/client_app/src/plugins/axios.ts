import axios from 'axios'

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8088/',
  timeout: 1000,
})

// axiosInstance.interceptors.request.use(
//   config => {
//     // TODO uncoment when do the do with auth
//     const authStore = useAuthStore()
//     const token = authStore.getToken()
//     if (token) {
//       config.headers.Authorization = `Bearer ${token}`
//     }
//       return config
//   },
//   error => {
//     return Promise.reject(error)
//   }
// )

export default axiosInstance