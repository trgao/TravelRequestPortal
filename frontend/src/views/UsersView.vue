<template>
  <NavBar />
  <div class="flex min-h-full flex-col justify-center py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">
        View all users in company
      </h2>
    </div>
    <div>
      <div
        class="flex align-middle justify-center items-center w-full overflow-visible"
      >
        <TextField
          v-model:input="employee"
          id="employee"
          placeholder="Search employee"
          class="w-full md:w-96 px-6"
          required
        />
      </div>
      <div class="p-4 hidden md:flex flex-col gap-2 overflow-x-scroll">
        <table
          class="p-3 w-full bg-white border border-gray-100 shadow-md rounded-lg mb-4 divide-y-2 divide-gray-200"
        >
          <thead>
            <tr class="*:font-semibold *:text-gray-900">
              <th class="text-left p-3 px-5">ID</th>
              <th class="text-left p-3 px-5">Email</th>
              <th class="text-left p-3 px-5">First Name</th>
              <th class="text-left p-3 px-5">Last Name</th>
              <th class="text-left p-3 px-5">Account type</th>
              <th class="text-left p-3 px-5">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="user in users" :key="user.id" class="hover:bg-gray-100">
              <td class="p-3 px-5">{{ user.id }}</td>
              <td class="p-3 px-5">
                {{ user.email }}
              </td>
              <td class="p-3 px-5">
                {{ user.first_name }}
              </td>
              <td class="p-3 px-5">
                {{ user.last_name }}
              </td>
              <td class="p-3 px-5">
                {{ user.is_admin ? 'Admin' : 'Employee' }}
              </td>
              <td class="p-3 px-5">
                <button
                  v-on:click="handleDelete(user.id)"
                  title="Remove user account"
                  class="bg-transparent hover:bg-red-600 text-red-600 font-semibold hover:text-white py-2 px-3 border border-red-600 hover:border-transparent rounded-lg cursor-pointer"
                >
                  <svg
                    width="18"
                    height="18"
                    fill="currentColor"
                    aria-hidden="true"
                    viewBox="0 0 20 20"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot>
            <tr>
              <td colspan="5">
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
                  <div class="flex gap-2 items-center">
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
          v-for="user in users"
          :key="user.id"
          class="rounded-lg shadow-md border border-gray-300 px-8 pb-8 pt-4 flex flex-col gap-1"
        >
          <div class="font-semibold text-xl">User #{{ user.id }}</div>
          <div>
            <span class="font-semibold">Email:</span>
            {{ user.email }}
          </div>
          <div>
            <span class="font-semibold">First name:</span>
            {{ user.first_name }}
          </div>
          <div>
            <span class="font-semibold">Last name:</span>
            {{ user.last_name }}
          </div>
          <div>
            <span class="font-semibold">Account type:</span>
            {{ user.is_admin ? 'Admin' : 'Employee' }}
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
  </div>
</template>

<script>
import client from '@/axios'
import NavBar from '@/components/NavBar.vue'
import TextField from '@/components/TextField.vue'
import { useAuthStore } from '@/stores/AuthStore'

export default {
  name: 'UsersView',
  components: {
    NavBar,
    TextField,
  },
  data() {
    return {
      users: [],
      page: 1,
      limit: 10,
      totalPages: 1,
      employee: '',
    }
  },
  computed: {
    authStore: () => useAuthStore(),
  },
  watch: {
    employee() {
      this.page = 1
      this.getUsers()
    },
  },
  methods: {
    handleDelete(id) {
      client
        .delete(`http://localhost:3000/api/auth/users/${id}`, {
          withCredentials: true,
        })
        .then((res) => {
          console.log(res)
        })
        .catch((err) => {
          console.error(err)
        })
    },
    handlePrevious() {
      this.page--
      this.getUsers()
    },
    handleNext() {
      this.page++
      this.getUsers()
    },
    handleChangeLimit() {
      this.page = 1
      this.getUsers()
    },
    getUsers() {
      client
        .get(
          `http://localhost:3000/api/auth/users?page=${this.page}&limit=${this.limit}&employee=${this.employee}`,
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.users = res.data.users
          this.totalPages = res.data.total_pages
        })
        .catch((err) => {
          console.error(err)
        })
    },
  },
  mounted() {
    this.getUsers()
  },
}
</script>
