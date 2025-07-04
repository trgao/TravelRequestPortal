<template>
  <NavBar />
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">
        Onboard a new user
      </h2>
    </div>

    <FormError v-model="error.general" />
    <FormSuccess v-model="success" />

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
        <div>
          <div class="flex justify-between flex-wrap">
            <label
              for="is-admin"
              class="block text-sm/6 font-medium text-gray-900"
            >
              Account Type
            </label>
          </div>
          <div class="flex gap-2">
            <input type="radio" v-model="user.is_admin" :value="true" />
            <label for="admin">Admin</label>
            <input type="radio" v-model="user.is_admin" :value="false" />
            <label for="employee">Employee</label>
          </div>
        </div>
        <SubmitButton name="Create Account" />
      </form>
    </div>
  </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'
import FormError from '@/components/FormError.vue'
import FormSuccess from '@/components/FormSuccess.vue'
import TextField from '@/components/TextField.vue'
import SubmitButton from '@/components/SubmitButton.vue'
import { useAuthStore } from '@/stores/AuthStore'
import client from '@/axios'

export default {
  name: 'OnboardView',
  components: {
    NavBar,
    FormError,
    FormSuccess,
    TextField,
    SubmitButton,
  },
  computed: {
    authStore: () => useAuthStore(),
  },
  data() {
    return {
      success: '',
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
        is_admin: false,
      },
    }
  },
  methods: {
    checkPassword() {
      this.error.general = ''
      this.success = ''
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
        is_admin: this.user.is_admin,
      }

      client
        .post(
          'http://localhost:3000/api/auth/onboard',
          { user: newUser },
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.error.general = ''
          this.success = 'Employee account has been successfully created'
          this.user = {
            email: '',
            firstName: '',
            lastName: '',
            password: '',
            confirmpassword: '',
            is_admin: false,
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
