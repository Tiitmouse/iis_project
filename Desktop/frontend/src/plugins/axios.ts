import axios from 'axios'

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8088/',
  timeout: 5000,
})

axiosInstance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('accessToken');
    if (token && config.url && config.url.startsWith('/api/') && !config.url.endsWith('/login') && !config.url.endsWith('/refresh')) {
      config.headers.Authorization = `Bearer ${token}`;
      console.log(`Interceptor: Added Bearer token to request for ${config.url}`);
    }
    return config;
  },
  (error) => {
    console.error("Interceptor Request Error:", error);
    return Promise.reject(error);
  }
);

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      console.log('Interceptor: Access token expired or invalid (401). Attempting refresh...');
      const refreshToken = localStorage.getItem('refreshToken');
      if (refreshToken) {
        try {
          const { data } = await axiosInstance.post('/api/refresh', { refresh_token: refreshToken });
          const newAccessToken = data.access_token;
          localStorage.setItem('accessToken', newAccessToken);
          console.log('Interceptor: Token refresh successful. Retrying original request.');

          originalRequest.headers['Authorization'] = `Bearer ${newAccessToken}`;
          return axiosInstance(originalRequest);
        } catch (refreshError) {
          console.error('Interceptor: Token refresh failed:', refreshError);

          localStorage.removeItem('accessToken');
          localStorage.removeItem('refreshToken');

          return Promise.reject(refreshError);
        }
      } else {
        console.error('Interceptor: No refresh token available for refresh attempt.');
        // Handle logout logic
        localStorage.removeItem('accessToken');

      }
    }

    return Promise.reject(error);
  }
);

export default axiosInstance