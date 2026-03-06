<template>
  <div class="max-w-4xl mx-auto py-8 px-4">
    <div class="flex items-center justify-between mb-8">
      <div class="flex items-center space-x-3">
        <div class="w-10 h-10 bg-blue-600 rounded-lg flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </div>
        <h1 class="text-2xl font-bold dark:text-white">Settings</h1>
      </div>
      <router-link to="/" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </router-link>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <!-- Sidebar -->
      <div class="md:col-span-1 space-y-1">
        <button class="w-full text-left px-4 py-2 rounded-lg bg-gray-100 dark:bg-gray-700 text-blue-600 font-medium"> General </button>
        <button class="w-full text-left px-4 py-2 rounded-lg text-gray-600 dark:text-gray-400 hover:bg-gray-50 dark:hover:bg-gray-700"> Models </button>
        <button class="w-full text-left px-4 py-2 rounded-lg text-gray-600 dark:text-gray-400 hover:bg-gray-50 dark:hover:bg-gray-700"> Interface </button>
      </div>

      <!-- Main Content -->
      <div class="md:col-span-3 space-y-6">
        <div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6">
          <h2 class="text-lg font-semibold mb-6 dark:text-white border-b dark:border-gray-700 pb-2">LLM Provider Configuration</h2>
          
          <div class="space-y-6">
            <!-- Mode Toggle -->
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">Active Provider</label>
              <div class="flex p-1 bg-gray-100 dark:bg-gray-900 rounded-lg w-fit">
                <button 
                  @click="settings.llm_mode = 'ollama'"
                  :class="settings.llm_mode === 'ollama' ? 'bg-white dark:bg-gray-700 shadow-sm text-blue-600' : 'text-gray-500'"
                  class="px-4 py-1.5 rounded-md text-sm font-medium transition"
                >
                  Ollama
                </button>
                <button 
                  @click="settings.llm_mode = 'airllm'"
                  :class="settings.llm_mode === 'airllm' ? 'bg-white dark:bg-gray-700 shadow-sm text-blue-600' : 'text-gray-500'"
                  class="px-4 py-1.5 rounded-md text-sm font-medium transition"
                >
                  AirLLM
                </button>
              </div>
            </div>

            <!-- URL Input and Status -->
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                  {{ settings.llm_mode === 'ollama' ? 'Ollama' : 'AirLLM' }} Base URL
                </label>
                <div class="flex items-center space-x-2">
                  <span :class="status === 'success' ? 'bg-green-500' : status === 'error' ? 'bg-red-500' : 'bg-gray-400'" class="w-2.5 h-2.5 rounded-full shadow-sm animate-pulse"></span>
                  <span class="text-xs font-medium text-gray-500 uppercase tracking-wider">{{ status === 'success' ? 'Online' : status === 'error' ? 'Offline' : 'Disconnected' }}</span>
                </div>
              </div>
              
              <div class="flex space-x-2">
                <input 
                  v-model="activeUrl" 
                  type="text" 
                  class="flex-1 px-3 py-2 bg-gray-50 dark:bg-gray-900 border border-gray-300 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none dark:text-white transition"
                  placeholder="http://localhost:11434"
                />
                <button 
                  @click="testConnection"
                  :disabled="testing"
                  class="px-4 py-2 bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-600 font-medium transition disabled:opacity-50"
                >
                  <span v-if="!testing">Verify</span>
                  <svg v-else class="animate-spin h-5 w-5" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                </button>
              </div>
            </div>

            <!-- Model Selection -->
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Default Model</label>
              <div class="relative">
                <select 
                  v-model="settings.default_model"
                  class="w-full appearance-none px-3 py-2 bg-gray-50 dark:bg-gray-900 border border-gray-300 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none dark:text-white transition"
                  :disabled="models.length === 0"
                >
                  <option v-if="models.length === 0" value="">No models found</option>
                  <option v-for="model in models" :key="model" :value="model">{{ model }}</option>
                </select>
                <div class="absolute inset-y-0 right-0 flex items-center px-2 pointer-events-none text-gray-500">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </div>
              </div>
              <p class="mt-2 text-xs text-gray-500">Select the primary model for chat completions.</p>
            </div>
          </div>

          <div class="mt-8 pt-6 border-t dark:border-gray-700 flex justify-end">
            <button 
              @click="saveSettings" 
              :disabled="saving"
              class="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-semibold shadow-md transition disabled:opacity-50"
            >
              {{ saving ? 'Updating...' : 'Save Settings' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import axios from 'axios'

const settings = ref({
  llm_mode: 'ollama',
  ollama_url: '',
  airllm_url: '',
  default_model: '',
  language: 'en'
})

const activeUrl = computed({
  get: () => settings.value.llm_mode === 'ollama' ? settings.value.ollama_url : settings.value.airllm_url,
  set: (val) => {
    if (settings.value.llm_mode === 'ollama') settings.value.ollama_url = val
    else settings.value.airllm_url = val
  }
})

const models = ref([])
const status = ref('unknown') // success, error, unknown
const testing = ref(false)
const saving = ref(false)

const fetchSettings = async () => {
  try {
    const token = localStorage.getItem('admin_token')
    const response = await axios.get('/api/v1/settings', {
      headers: { 'Authorization': token }
    })
    settings.value = response.data
    // Auto test on load
    testConnection()
  } catch (err) {
    status.value = 'error'
  }
}

const testConnection = async () => {
  testing.value = true
  status.value = 'unknown'
  try {
    const token = localStorage.getItem('admin_token')
    const response = await axios.post('/api/v1/settings/test-connection', {
      llm_mode: settings.value.llm_mode,
      ollama_url: settings.value.ollama_url,
      airllm_url: settings.value.airllm_url
    }, {
      headers: { 'Authorization': token }
    })
    
    if (response.data.success) {
      status.value = 'success'
      models.value = response.data.models
    } else {
      status.value = 'error'
      models.value = []
    }
  } catch (err) {
    status.value = 'error'
    models.value = []
  } finally {
    testing.value = false
  }
}

const saveSettings = async () => {
  saving.value = true
  try {
    const token = localStorage.getItem('admin_token')
    await axios.put('/api/v1/settings', settings.value, {
      headers: { 'Authorization': token }
    })
    // Success notification could go here
  } catch (err) {
    // Error notification could go here
  } finally {
    saving.value = false
  }
}

watch(() => settings.value.llm_mode, () => {
  status.value = 'unknown'
  models.value = []
})

onMounted(fetchSettings)
</script>

<style scoped>
.form-radio {
  @apply rounded-full border-gray-300 text-blue-600 focus:ring-blue-500;
}
</style>
