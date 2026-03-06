<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Welcome to AI WebUI</h1>
    <p class="mb-6">This is a simple home page. Navigate to the chat page to start a conversation.</p>
    <div class="flex space-x-4">
      <router-link to="/chat" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition">Go to Chat</router-link>
      <router-link to="/settings" class="px-4 py-2 bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-white rounded hover:bg-gray-300 dark:hover:bg-gray-600 transition">Settings</router-link>
    </div>
  </div>
</template>

<script setup>
// Home component
</script>

<style scoped>
</style>
