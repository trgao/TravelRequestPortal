<script setup>
defineProps(['id', 'name', 'required'])
const input = defineModel()
</script>

<template>
  <div>
    <div class="flex justify-between flex-wrap">
      <label :for="id" class="block text-sm/6 font-medium text-gray-900">
        {{ name }}
      </label>
    </div>
    <div>
      <input
        v-model="input"
        type="datetime-local"
        :name="id"
        :id="id"
        :required="required"
        class="block w-full px-3 py-1.5 rounded-md bg-white text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
      />
    </div>
  </div>
</template>
