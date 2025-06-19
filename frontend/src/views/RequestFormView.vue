<template>
  <NavBar />
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">
        Submit a new travel request
      </h2>
    </div>

    <FormError v-model="error" />
    <FormSuccess v-model="success" />

    <div class="mt-5 sm:mx-auto sm:w-full sm:max-w-sm">
      <form
        class="space-y-6"
        @submit.prevent="submitForm"
        @input="resetMessages"
      >
        <TextField
          v-model:input="request.startLocation"
          id="start-location"
          name="Start Location"
          required
        />
        <TextField
          v-model:input="request.destination"
          id="destination"
          name="Destination"
          required
        />
        <DatePicker
          v-model="request.dateTime"
          id="date-time"
          name="Date and Time"
          required
        />
        <TextField
          v-model:input="request.purpose"
          id="purpose"
          name="Purpose"
          required
        />
        <SubmitButton name="Submit Request" />
      </form>
    </div>
  </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'
import FormError from '@/components/FormError.vue'
import FormSuccess from '@/components/FormSuccess.vue'
import TextField from '@/components/TextField.vue'
import DatePicker from '@/components/DatePicker.vue'
import SubmitButton from '@/components/SubmitButton.vue'
import { useAuthStore } from '@/stores/AuthStore'
import router from '@/router'
import client from '@/axios'

export default {
  name: 'RequestFormView',
  components: {
    NavBar,
    FormError,
    FormSuccess,
    TextField,
    DatePicker,
    SubmitButton,
  },
  computed: {
    authStore: () => useAuthStore(),
  },
  data() {
    return {
      success: '',
      error: '',
      request: {
        startLocation: '',
        destination: '',
        dateTime: '',
        purpose: '',
      },
    }
  },
  methods: {
    resetMessages() {
      this.error = ''
      this.success = ''
    },
    submitForm() {
      const newRequest = {
        start_location: this.request.startLocation,
        destination: this.request.destination,
        date_time: this.request.dateTime,
        purpose: this.request.purpose,
      }

      client
        .post(
          'http://localhost:3000/api/request',
          { request: newRequest },
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.error = ''
          this.success = 'Travel request has been successfully submitted'
          this.request = {
            startLocation: '',
            destination: '',
            dateTime: '',
            purpose: '',
          }
          router.push('/')
        })
        .catch((err) => {
          console.error(err)
          this.error = err.response.data.error
        })
    },
  },
}
</script>
