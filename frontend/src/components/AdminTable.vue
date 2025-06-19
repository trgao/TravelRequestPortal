<script setup>
import { format } from 'date-fns'
</script>

<template>
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
            <th class="text-left p-3 px-5">Employee</th>
            <th class="text-left p-3 px-5">Start Location</th>
            <th class="text-left p-3 px-5">Destination</th>
            <th class="text-left p-3 px-5">Date and Time</th>
            <th class="text-left p-3 px-5">Purpose of Travel</th>
            <th class="text-left p-3 px-5">Status</th>
            <th class="text-left p-3 px-5">Admin remarks</th>
            <th class="text-left p-3 px-5">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr
            v-for="(request, index) in requests"
            :key="request.id"
            class="hover:bg-gray-100"
          >
            <td class="p-3 px-5">{{ request.id }}</td>
            <td class="p-3 px-5">
              {{ request.first_name }} {{ request.last_name }}
            </td>
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
              <div class="flex gap-1 justify-between items-center">
                <textarea
                  rows="3"
                  v-if="isEditable[index]"
                  v-model="adminRemarks[index]"
                  class="min-h-11 w-full border border-indigo-600 rounded-lg p-2"
                ></textarea>
                <div v-else class="text-left">
                  <span
                    v-if="request.admin_remarks == ''"
                    class="italic opacity-50"
                    >No remarks</span
                  >
                  <span v-else>{{ request.admin_remarks }}</span>
                </div>
                <div
                  v-if="isEditable[index]"
                  class="flex flex-col gap-2 items-center"
                >
                  <button
                    title="Update admin remarks"
                    v-on:click="updateRemarks(request.id, index)"
                    class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer"
                  >
                    &#10003;
                  </button>
                  <button
                    title="Cancel"
                    v-on:click="handleCancel(index)"
                    class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer"
                  >
                    &#10005;
                  </button>
                </div>
                <button
                  v-on:click="handleEdit(index)"
                  v-else
                  title="Edit admin remarks"
                  class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-3 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer"
                >
                  <svg
                    data-v-9dd4e508=""
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 18 18"
                  >
                    <g
                      data-v-9dd4e508=""
                      fill="currentColor"
                      data-darkreader-inline-fill=""
                      style="--darkreader-inline-fill: currentColor"
                    >
                      <g data-v-9dd4e508="">
                        <path
                          d="M16.001 11a1 1 0 0 0-1 1v2.5a.5.5 0 0 1-.5.5H3.499a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5h2.5a1 1 0 1 0 0-2h-4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14.002a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1z"
                        ></path>
                        <path
                          d="M13.169 3.403l1.429 1.429L8.65 11H7V9.35l6.169-5.947zM13.182 1a.994.994 0 0 0-.706.292L5 8.5V13h4.501l7.209-7.475a.997.997 0 0 0 0-1.411l-2.822-2.822A.992.992 0 0 0 13.183 1z"
                        ></path>
                      </g>
                    </g>
                  </svg>
                </button>
              </div>
            </td>
            <td class="p-3 px-5">
              <div class="flex gap-1">
                <button
                  v-on:click="updateStatus(request.id, index, 'approved')"
                  title="Approve travel request"
                  class="bg-transparent hover:bg-green-500 text-green-500 font-semibold hover:text-white py-2 px-4 border border-green-500 hover:border-transparent rounded-lg cursor-pointer"
                >
                  &#10003;
                </button>
                <button
                  v-on:click="updateStatus(request.id, index, 'rejected')"
                  title="Reject travel request"
                  class="bg-transparent hover:bg-red-500 text-red-500 font-semibold hover:text-white py-2 px-4 border border-red-500 hover:border-transparent rounded-lg cursor-pointer"
                >
                  &#10005;
                </button>
              </div>
            </td>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <td colspan="9">
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
        v-for="(request, index) in requests"
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
            <span class="font-semibold">Employee:</span>
            {{ request.first_name }} {{ request.last_name }}
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
            <span class="font-semibold">Admin remarks: </span>
            <div class="flex gap-1 justify-between items-center">
              <textarea
                rows="3"
                v-if="isEditable[index]"
                v-model="adminRemarks[index]"
                class="min-h-11 w-full border border-indigo-600 rounded-lg p-2"
              ></textarea>
              <div v-else class="text-left">
                <span
                  v-if="request.admin_remarks == ''"
                  class="italic opacity-50"
                  >No remarks</span
                >
                <span v-else>{{ request.admin_remarks }}</span>
              </div>
              <div
                v-if="isEditable[index]"
                class="flex flex-col gap-2 items-center"
              >
                <button
                  title="Update admin remarks"
                  v-on:click="updateRemarks(request.id, index)"
                  class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer"
                >
                  &#10003;
                </button>
                <button
                  title="Cancel"
                  v-on:click="handleCancel(index)"
                  class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-4 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer"
                >
                  &#10005;
                </button>
              </div>
              <button
                v-on:click="handleEdit(index)"
                v-else
                title="Edit admin remarks"
                class="bg-transparent hover:bg-indigo-600 text-indigo-600 font-semibold hover:text-white py-2 px-3 border border-indigo-600 hover:border-transparent rounded-lg cursor-pointer"
              >
                <svg
                  data-v-9dd4e508=""
                  xmlns="http://www.w3.org/2000/svg"
                  width="18"
                  height="18"
                  viewBox="0 0 18 18"
                >
                  <g
                    data-v-9dd4e508=""
                    fill="currentColor"
                    data-darkreader-inline-fill=""
                    style="--darkreader-inline-fill: currentColor"
                  >
                    <g data-v-9dd4e508="">
                      <path
                        d="M16.001 11a1 1 0 0 0-1 1v2.5a.5.5 0 0 1-.5.5H3.499a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5h2.5a1 1 0 1 0 0-2h-4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14.002a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1z"
                      ></path>
                      <path
                        d="M13.169 3.403l1.429 1.429L8.65 11H7V9.35l6.169-5.947zM13.182 1a.994.994 0 0 0-.706.292L5 8.5V13h4.501l7.209-7.475a.997.997 0 0 0 0-1.411l-2.822-2.822A.992.992 0 0 0 13.183 1z"
                      ></path>
                    </g>
                  </g>
                </svg>
              </button>
            </div>
          </div>
          <div class="flex gap-2 justify-between">
            <button
              v-on:click="updateStatus(request.id, index, 'approved')"
              title="Approve travel request"
              class="w-full bg-transparent hover:bg-green-500 text-green-500 font-semibold hover:text-white py-2 px-4 border border-green-500 hover:border-transparent rounded-lg cursor-pointer"
            >
              &#10003; Approve
            </button>
            <button
              v-on:click="updateStatus(request.id, index, 'rejected')"
              title="Reject travel request"
              class="w-full bg-transparent hover:bg-red-500 text-red-500 font-semibold hover:text-white py-2 px-4 border border-red-500 hover:border-transparent rounded-lg cursor-pointer"
            >
              &#10005; Reject
            </button>
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
import TextField from './TextField.vue'
import client from '@/axios'

export default {
  name: 'AdminTable',
  data() {
    return {
      requests: [],
      isEditable: [],
      adminRemarks: [],
      page: 1,
      limit: 10,
      totalPages: 1,
      employee: '',
    }
  },
  computed: {
    url() {
      const authStore = useAuthStore()
      return `http://localhost:3000/api/admin/${authStore.userId}`
    },
  },
  watch: {
    employee() {
      this.page = 1
      this.getRequests()
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
    handleEdit(index) {
      this.isEditable[index] = true
      this.adminRemarks[index] = this.requests[index].admin_remarks
    },
    handleCancel(index) {
      this.isEditable[index] = false
    },
    updateStatus(id, index, status) {
      client
        .patch(
          `http://localhost:3000/api/request/${id}/status`,
          { status: status },
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.requests[index].status = status
        })
        .catch((err) => {
          console.error(err)
          alert('Unable to update status')
        })
    },
    updateRemarks(id, index) {
      client
        .patch(
          `http://localhost:3000/api/request/${id}/remarks`,
          { remarks: this.adminRemarks[index] },
          { withCredentials: true },
        )
        .then((res) => {
          console.log(res)
          this.requests[index].admin_remarks = this.adminRemarks[index]
          this.isEditable[index] = false
        })
        .catch((err) => {
          console.error(err)
          alert('Unable to update remarks')
        })
    },
    getRequests() {
      client
        .get(
          `${this.url}?page=${this.page}&limit=${this.limit}&employee=${this.employee}`,
          {
            withCredentials: true,
          },
        )
        .then((res) => {
          console.log(res)
          this.requests = res.data.requests
          this.isEditable = Array(this.requests.length).fill(false)
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
