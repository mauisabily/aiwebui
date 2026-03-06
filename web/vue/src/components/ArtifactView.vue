<template>
  <aside 
    v-if="chatStore.showArtifacts"
    class="w-1/3 flex-shrink-0 bg-white dark:bg-gray-800 border-l border-gray-200 dark:border-gray-700 flex flex-col transition-all duration-300 shadow-xl z-20"
  >
    <!-- Header -->
    <div class="p-4 flex items-center justify-between border-b border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-900/50">
      <div class="flex items-center space-x-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-indigo-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 21h6l-.75-4M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
        </svg>
        <span class="font-bold text-gray-700 dark:text-gray-200">Artifact / Preview</span>
      </div>
      <button 
        @click="chatStore.toggleArtifacts"
        class="p-1 hover:bg-gray-200 dark:hover:bg-gray-700 rounded-md transition"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-hidden flex flex-col">
        <!-- Tabs -->
        <div class="flex border-b border-gray-200 dark:border-gray-700">
            <button 
                @click="activeTab = 'code'" 
                :class="activeTab === 'code' ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500'"
                class="flex-1 py-2 text-sm font-medium border-b-2 transition"
            >
                Code
            </button>
            <button 
                @click="activeTab = 'preview'" 
                :class="activeTab === 'preview' ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500'"
                class="flex-1 py-2 text-sm font-medium border-b-2 transition"
            >
                Browser
            </button>
        </div>

        <div class="flex-1 overflow-auto p-4 bg-gray-50 dark:bg-gray-900">
            <!-- Code Tab -->
            <div v-if="activeTab === 'code'" class="h-full">
                <pre class="text-xs p-4 bg-gray-900 text-gray-300 rounded-lg overflow-auto h-full shadow-inner">{{ chatStore.currentArtifact }}</pre>
            </div>

            <!-- Preview Tab -->
            <div v-if="activeTab === 'preview'" class="h-full rounded-lg bg-white overflow-hidden shadow-sm border border-gray-200 dark:border-gray-700">
                <iframe 
                    v-if="htmlContent"
                    :srcdoc="htmlContent" 
                    class="w-full h-full border-none bg-white"
                ></iframe>
                <div v-else class="flex flex-col items-center justify-center h-full space-y-4 text-gray-400">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 opacity-20" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                    </svg>
                    <p class="text-sm">No HTML content detected to preview.</p>
                </div>
            </div>
        </div>
    </div>
  </aside>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useChatStore } from '@/stores/chat'

const chatStore = useChatStore()
const activeTab = ref('preview')

const htmlContent = computed(() => {
  if (!chatStore.currentArtifact) return null
  const match = chatStore.currentArtifact.match(/```html\n([\s\S]*?)\n```/)
  return match ? match[1] : null
})
</script>
