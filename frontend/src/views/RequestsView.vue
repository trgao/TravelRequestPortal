<template>
  <NavBar />
  <div class="flex min-h-full flex-col justify-center py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <h2 class="text-center text-2xl/9 font-bold text-gray-900">
        View all requests
      </h2>
    </div>
    <AdminTable v-if="authStore.isAdmin && authStore.isAuthenticated" />
    <RequestsTable v-if="!authStore.isAdmin && authStore.isAuthenticated" />
  </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'
import RequestsTable from '@/components/RequestsTable.vue'
import AdminTable from '@/components/AdminTable.vue'
import { useAuthStore } from '@/stores/AuthStore'

export default {
  name: 'RequestsView',
  components: {
    NavBar,
    RequestsTable,
    AdminTable,
  },
  computed: {
    authStore: () => useAuthStore(),
  },
}
</script>
