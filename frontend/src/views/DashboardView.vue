<template>
  <NavBar />
  <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">Dashboard</h2>
    </div>

    <div class="mt-5 sm:mx-auto flex gap-3 flex-wrap justify-center">
      <div
        class="block w-full md:w-72 p-6 bg-white border border-gray-200 rounded-lg shadow-sm"
      >
        <h5 class="mb-2 text-xl font-semibold text-gray-600">
          Pending requests
        </h5>
        <p class="text-xl font-normal text-gray-700">{{ pending }}</p>
      </div>
      <div
        class="block w-full md:w-72 p-6 bg-white border border-gray-200 rounded-lg shadow-sm"
      >
        <h5 class="mb-2 text-xl font-semibold text-green-600">
          Approved requests
        </h5>
        <p class="text-xl font-normal text-gray-700">{{ approved }}</p>
      </div>
      <div
        class="block w-full md:w-72 p-6 bg-white border border-gray-200 rounded-lg shadow-sm"
      >
        <h5 class="mb-2 text-xl font-semibold text-red-600">
          Rejected requests
        </h5>
        <p class="text-xl font-normal text-gray-700">{{ rejected }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'
import { useAuthStore } from '@/stores/AuthStore'
import client from '@/axios'

export default {
  name: 'DashboardView',
  components: {
    NavBar,
  },
  computed: {
    authStore: () => useAuthStore(),
  },
  data() {
    return {
      pending: 0,
      approved: 0,
      rejected: 0,
    }
  },
  methods: {
    getSummary() {
      client
        .get('http://localhost:3000/api/summary', { withCredentials: true })
        .then((res) => {
          console.log(res)
          this.pending = res.data.pending
          this.approved = res.data.approved
          this.rejected = res.data.rejected
        })
        .catch((err) => {
          console.error(err)
        })
    },
  },
  mounted() {
    this.getSummary()
  },
}
</script>
