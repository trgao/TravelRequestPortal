<template>
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">
        Login to your account
      </h2>
    </div>

    <FormError v-model="error" />

    <div class="mt-5 sm:mx-auto sm:w-full sm:max-w-sm">
      <form class="space-y-6" @submit.prevent="login" @input="resetError">
        <TextField
          v-model:input="user.email"
          type="email"
          id="email"
          name="Email"
          required
        />
        <TextField
          v-model:input="user.password"
          type="password"
          id="password"
          name="Password"
          required
        />
        <SubmitButton name="Login" />
      </form>

      <p class="mt-10 text-center text-sm/6 text-gray-500">
        Don't have an account?
        <RouterLink to="/register" class="text-yellow-600 hover:underline">
          Register
        </RouterLink>
      </p>
    </div>
  </div>
</template>

<script>
import FormError from '@/components/FormError.vue'
import TextField from '@/components/TextField.vue'
import SubmitButton from '@/components/SubmitButton.vue'
import { RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/AuthStore'
import router from '@/router/index'
import client from '@/axios'

export default {
  name: 'LoginView',
  components: {
    FormError,
    TextField,
    SubmitButton,
    RouterLink,
  },
  computed: {
    authStore: () => useAuthStore(),
  },
  data() {
    return {
      error: '',
      user: {
        email: '',
        password: '',
      },
    }
  },
  methods: {
    resetError() {
      this.error = ''
    },
    login() {
      client
        .post(
          'http://localhost:3000/api/auth/login',
          { user: this.user },
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.authStore.isAuthenticated = true
          this.authStore.isAdmin = res.data.is_admin
          this.authStore.userId = res.data.user_id
          if (this.authStore.isAdmin) {
            router.push('/dashboard')
          } else {
            router.push('/')
          }
        })
        .catch((err) => {
          console.error(err)
          this.error = 'Incorrect email or password'
        })
    },
  },
}
</script>
