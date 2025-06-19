<script setup>
defineProps(['id', 'name', 'type', 'required', 'placeholder'])
const input = defineModel('input')
const error = defineModel('error')
</script>

<template>
  <div>
    <div class="flex justify-between flex-wrap">
      <label :for="id" class="block text-sm/6 font-medium text-gray-900">
        {{ name }}
      </label>
      <p v-if="error" class="text-sm/6 text-red-500">{{ error }}</p>
    </div>
    <div>
      <input
        v-model="input"
        :type="type"
        :name="id"
        :id="id"
        :required="required"
        :placeholder="placeholder"
        class="block w-full px-3 py-1.5 rounded-md bg-white text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 focus:outline-2 focus:-outline-offset-2 sm:text-sm/6"
        :class="error ? 'focus:outline-red-500' : 'focus:outline-indigo-600'"
      />
    </div>
  </div>
</template>
