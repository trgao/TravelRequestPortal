<template>
  <nav class="flex gap-2 justify-end bg-indigo-600 p-1">
    <div class="hidden md:block">
      <button
        v-on:click="goToDashboard"
        v-if="authStore.isAdmin"
        class="hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Dashboard
      </button>
      <button
        v-on:click="goToRequests"
        class="hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        View Requests
      </button>
      <button
        v-on:click="goToOnboard"
        v-if="authStore.isAdmin"
        class="hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Onboard a User
      </button>
      <button
        v-on:click="goToForm"
        v-else
        class="hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Submit a Request
      </button>
      <button
        v-on:click="logout"
        class="hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Logout
      </button>
    </div>
    <button
      class="rounded-lg md:hidden focus:outline-none focus:shadow-outline p-2"
      @click="isOpen = !isOpen"
    >
      <svg fill="white" viewBox="0 0 20 20" class="w-6 h-6">
        <path
          v-show="!isOpen"
          fill-rule="evenodd"
          d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM9 15a1 1 0 011-1h6a1 1 0 110 2h-6a1 1 0 01-1-1z"
          clip-rule="evenodd"
        ></path>
        <path
          v-show="isOpen"
          fill-rule="evenodd"
          d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
          clip-rule="evenodd"
        ></path>
      </svg>
    </button>
  </nav>
  <nav class="md:hidden sm:absolute right-0 w-full origin-top-right z-30">
    <div
      class="absolute z-30 flex flex-col items-center justify-center w-full px-2 py-4 space-y-2 text-center bg-indigo-600"
      :class="{ block: isOpen, hidden: !isOpen }"
    >
      <button
        v-on:click="goToDashboard"
        v-if="authStore.isAdmin"
        class="w-full hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Dashboard
      </button>
      <button
        v-on:click="goToRequests"
        class="w-full hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        View Requests
      </button>
      <button
        v-on:click="goToOnboard"
        v-if="authStore.isAdmin"
        class="w-full hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Onboard a User
      </button>
      <button
        v-on:click="goToForm"
        v-else
        class="w-full hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Submit a Request
      </button>
      <button
        v-on:click="logout"
        class="w-full hover:cursor-pointer text-white hover:bg-indigo-800 rounded-lg p-2"
      >
        Logout
      </button>
    </div>
  </nav>
</template>

<script>
import client from '@/axios'
import router from '@/router'
import { useAuthStore } from '@/stores/AuthStore'

export default {
  name: 'NavBar',
  computed: {
    authStore: () => useAuthStore(),
  },
  data() {
    return {
      isOpen: false,
    }
  },
  methods: {
    goToDashboard() {
      router.push('/dashboard')
    },
    goToRequests() {
      router.push('/')
    },
    goToOnboard() {
      router.push('/onboard')
    },
    goToForm() {
      router.push('/form')
    },
    logout() {
      client
        .post(
          'http://localhost:3000/api/auth/logout',
          {},
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.authStore.isAuthenticated = false
          this.authStore.isAdmin = false
          router.push('/login')
          console.log('Logged out')
        })
        .catch((err) => console.error(err))
    },
  },
}
</script>
