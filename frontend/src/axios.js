import axios from 'axios'
import router from './router'
import { useAuthStore } from './stores/AuthStore'

const client = axios.create()

client.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    if (error.response.status == 401) {
      const authStore = useAuthStore()
      authStore.isAuthenticated = false
      authStore.isAdmin = false
      router.push('/login')
    }
    return Promise.reject(error)
  },
)

export default client
