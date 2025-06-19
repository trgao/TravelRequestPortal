import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => {
    return {
      isAuthenticated: false,
      isAdmin: false,
      userId: 0,
    }
  },
  persist: true,
})
