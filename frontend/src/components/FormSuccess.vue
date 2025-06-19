<script setup>
const success = defineModel()
</script>

<template>
  <div
    class="sm:mx-auto sm:w-full sm:max-w-sm bg-green-500 p-3 mt-5 rounded-md"
    v-if="success"
  >
    <p class="text-center text-white">{{ success }}</p>
  </div>
</template>
