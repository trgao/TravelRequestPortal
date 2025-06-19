<script setup>
import { format } from 'date-fns'
</script>

<template>
  <div>
    <div class="hidden md:flex px-6 overflow-x-scroll">
      <table
        class="p-3 w-full bg-white border border-gray-100 shadow-md rounded-lg mb-4 divide-y-2 divide-gray-200"
      >
        <thead>
          <tr class="*:font-semibold *:text-gray-900">
            <th class="text-left p-3 px-5">ID</th>
            <th class="text-left p-3 px-5">Start Location</th>
            <th class="text-left p-3 px-5">Destination</th>
            <th class="text-left p-3 px-5">Date and Time</th>
            <th class="text-left p-3 px-5">Purpose of Travel</th>
            <th class="text-left p-3 px-5">Status</th>
            <th class="text-left p-3 px-5">Admin remarks</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr
            v-for="request in requests"
            :key="request.id"
            class="hover:bg-gray-100"
          >
            <td class="p-3 px-5">{{ request.id }}</td>
            <td class="p-3 px-5">{{ request.start_location }}</td>
            <td class="p-3 px-5">{{ request.destination }}</td>
            <td class="p-3 px-5">
              {{ format(new Date(request.date_time), 'dd MMMM yyyy, H:mm') }}
            </td>
            <td class="p-3 px-5">{{ request.purpose }}</td>
            <td class="p-3 px-5">
              <span
                class="capitalize rounded-lg py-2 px-3 text-white"
                :class="
                  request.status == 'approved'
                    ? 'bg-green-500'
                    : request.status == 'rejected'
                      ? 'bg-red-500'
                      : 'bg-gray-500'
                "
              >
                {{ request.status }}
              </span>
            </td>
            <td class="p-3 px-5 min-w-60">
              <span v-if="request.admin_remarks == ''" class="italic opacity-50"
                >No remarks</span
              >
              <span v-else>{{ request.admin_remarks }}</span>
            </td>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <td colspan="7">
              <div class="flex justify-between p-3">
                <div class="flex items-center">
                  <p class="px-2">Rows per page:</p>
                  <select v-model="limit" @change="handleChangeLimit">
                    <option :value="10">10</option>
                    <option :value="25">25</option>
                    <option :value="50">50</option>
                    <option :value="100">100</option>
                  </select>
                </div>
                <div class="flex gap-2">
                  <button
                    v-on:click="handlePrevious"
                    :disabled="page <= 1"
                    class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer disabled:cursor-not-allowed disabled:opacity-50"
                  >
                    Previous
                  </button>
                  <button
                    v-on:click="handleNext"
                    :disabled="page >= totalPages"
                    class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer disabled:cursor-not-allowed disabled:opacity-50"
                  >
                    Next
                  </button>
                </div>
              </div>
            </td>
          </tr>
        </tfoot>
      </table>
    </div>
    <div class="px-6 pt-4 flex flex-col gap-4 md:hidden">
      <div
        v-for="request in requests"
        :key="request.id"
        class="rounded-lg shadow-md border"
        :class="
          request.status == 'approved'
            ? 'border-green-500'
            : request.status == 'rejected'
              ? 'border-red-500'
              : 'border-gray-500'
        "
      >
        <div class="flex flex-col gap-3">
          <span
            class="capitalize rounded-t-md px-3 py-1 text-white text-center"
            :class="
              request.status == 'approved'
                ? 'bg-green-500'
                : request.status == 'rejected'
                  ? 'bg-red-500'
                  : 'bg-gray-500'
            "
          >
            {{ request.status }}
          </span>
        </div>
        <div class="px-8 pb-8 pt-4 flex flex-col gap-1">
          <div class="font-semibold text-xl">
            Travel Request #{{ request.id }}
          </div>
          <div>
            <span class="font-semibold">Start location:</span>
            {{ request.start_location }}
          </div>
          <div>
            <span class="font-semibold">Destination:</span>
            {{ request.destination }}
          </div>
          <div>
            <span class="font-semibold">Date and time:</span>
            {{ format(new Date(request.date_time), 'dd MMMM yyyy, H:mm') }}
          </div>
          <div>
            <span class="font-semibold">Purpose of travel:</span>
            {{ request.purpose }}
          </div>
          <div>
            <span class="font-semibold">Admin remarks: <br /></span>
            <span v-if="request.admin_remarks == ''" class="italic opacity-50"
              >No remarks</span
            >
            <span v-else>{{ request.admin_remarks }}</span>
          </div>
        </div>
      </div>
      <div class="flex justify-between p-3 rounded-lg shadow-md">
        <div class="flex items-center">
          <p class="px-2">Items per page:</p>
          <select v-model="limit" @change="handleChangeLimit">
            <option :value="10">10</option>
            <option :value="25">25</option>
            <option :value="50">50</option>
            <option :value="100">100</option>
          </select>
        </div>
        <div class="flex gap-2">
          <button
            v-on:click="handlePrevious"
            :disabled="page <= 1"
            class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer disabled:cursor-not-allowed disabled:opacity-50"
          >
            Previous
          </button>
          <button
            v-on:click="handleNext"
            :disabled="page >= totalPages"
            class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer disabled:cursor-not-allowed disabled:opacity-50"
          >
            Next
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useAuthStore } from '@/stores/AuthStore'
import client from '@/axios'

export default {
  name: 'RequestsTable',
  data() {
    return {
      requests: [],
      page: 1,
      limit: 10,
      totalPages: 1,
    }
  },
  computed: {
    url() {
      const authStore = useAuthStore()
      return `http://localhost:3000/api/employee/${authStore.userId}`
    },
  },
  methods: {
    handlePrevious() {
      this.page--
      this.getRequests()
    },
    handleNext() {
      this.page++
      this.getRequests()
    },
    handleChangeLimit() {
      this.page = 1
      this.getRequests()
    },
    getRequests() {
      client
        .get(`${this.url}?page=${this.page}&limit=${this.limit}`, {
          withCredentials: true,
        })
        .then((res) => {
          console.log(res)
          this.requests = res.data.requests
          this.totalPages = res.data.total_pages
        })
        .catch((err) => {
          console.error(err)
        })
    },
  },
  mounted() {
    this.getRequests()
  },
}
</script>
