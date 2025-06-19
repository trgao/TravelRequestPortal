<template>
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">
        Register a new account
      </h2>
    </div>

    <FormError v-model="error.general" />

    <div class="mt-5 sm:mx-auto sm:w-full sm:max-w-sm">
      <form class="space-y-6" @submit.prevent="register" @input="checkPassword">
        <TextField
          v-model:input="user.email"
          type="email"
          id="email"
          name="Email"
          required
        />
        <TextField
          v-model:input="user.firstName"
          id="first-name"
          name="First Name"
          required
        />
        <TextField
          v-model:input="user.lastName"
          id="last-name"
          name="Last Name"
          required
        />
        <TextField
          v-model:input="user.company"
          id="company"
          name="Company Name"
          required
        />
        <TextField
          v-model:input="user.password"
          v-model:error="error.password"
          type="password"
          id="password"
          name="Password"
          required
        />
        <TextField
          v-model:input="user.confirmpassword"
          v-model:error="error.password"
          type="password"
          id="confirmpassword"
          name="Confirm Password"
          required
        />
        <SubmitButton name="Register" />
      </form>

      <p class="mt-10 text-center text-sm/6 text-gray-500">
        Already have an account?
        <RouterLink to="/login" class="text-yellow-600 hover:underline">
          Login
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
  name: 'RegisterView',
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
      error: {
        password: '',
        general: '',
      },
      user: {
        email: '',
        firstName: '',
        lastName: '',
        password: '',
        confirmpassword: '',
        company: '',
      },
    }
  },
  methods: {
    checkPassword() {
      this.error.general = ''
      if (this.user.password != this.user.confirmpassword) {
        this.error.password = 'Passwords do not match'
      } else {
        this.error.password = ''
      }
    },
    register() {
      if (this.user.password != this.user.confirmpassword) {
        console.error('Passwords do not match')
        return
      }

      const newUser = {
        email: this.user.email,
        first_name: this.user.firstName,
        last_name: this.user.lastName,
        password: this.user.password,
        company: this.user.company,
      }

      client
        .post(
          'http://localhost:3000/api/auth/register',
          { user: newUser },
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
          this.error.general = err.response.data.error
        })
    },
  },
}
</script>
