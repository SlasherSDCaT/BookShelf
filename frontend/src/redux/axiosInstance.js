// axiosInstance.js
import axios from 'axios';
import Cookies from 'js-cookie';

const axiosInstance = axios.create({
  // baseURL: 'https://bookshelf-cq3i.onrender.com/api',
  baseURL: 'https://bookshelf-cq3i.onrender.com/api',
});

axiosInstance.interceptors.request.use((config) => {
  const token = Cookies.get('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default axiosInstance;