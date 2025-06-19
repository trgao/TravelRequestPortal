<script setup>
defineProps(['name'])
</script>

<template>
  <div>
    <button
      type="submit"
      class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 hover:cursor-pointer"
    >
      {{ name }}
    </button>
  </div>
</template>
