<script setup>
const error = defineModel()
</script>

<template>
  <div
    class="sm:mx-auto sm:w-full sm:max-w-sm bg-red-500 p-3 mt-5 rounded-md"
    v-if="error"
  >
    <p class="text-center text-white">{{ error }}</p>
  </div>
</template>
