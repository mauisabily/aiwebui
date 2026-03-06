<template>
  <aside 
    v-if="chatStore.showSidebar"
    class="w-64 flex-shrink-0 bg-gray-50 dark:bg-gray-900 border-r border-gray-200 dark:border-gray-800 flex flex-col transition-all duration-300"
  >
    <!-- Header -->
    <div class="p-4 flex items-center justify-between border-b border-gray-200 dark:border-gray-800">
      <button 
        @click="chatStore.createNewChat"
        class="flex-1 flex items-center justify-center space-x-2 bg-white dark:bg-gray-800 hover:bg-gray-100 dark:hover:bg-gray-700 py-2 rounded-lg border border-gray-200 dark:border-gray-700 transition shadow-sm"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        <span class="text-sm font-medium dark:text-gray-200">New Chat</span>
      </button>
    </div>

    <!-- History List -->
    <div class="flex-1 overflow-y-auto p-2 space-y-1">
      <div v-if="chatStore.conversations.length === 0" class="text-center py-8 text-gray-400 text-xs">
        No recent chats
      </div>
      <router-link 
        v-for="conv in chatStore.conversations" 
        :key="conv.id"
        :to="'/c/' + conv.id"
        :class="chatStore.activeConversationId === conv.id ? 'bg-blue-50 dark:bg-blue-900/20 text-blue-600 border-blue-200 dark:border-blue-800' : 'hover:bg-gray-100 dark:hover:bg-gray-800 text-gray-600 dark:text-gray-400 border-transparent'"
        class="group p-3 rounded-lg cursor-pointer border transition flex items-center justify-between"
      >
        <div class="flex items-center space-x-3 overflow-hidden text-left">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
          </svg>
          <span class="text-sm truncate pr-2">{{ conv.title || 'Untitled Chat' }}</span>
        </div>
        <button 
          @click.stop.prevent="chatStore.deleteConversation(conv.id)"
          class="opacity-0 group-hover:opacity-100 p-1 hover:text-red-500 transition"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
        </button>
      </router-link>
    </div>

    <!-- Footer -->
    <div class="p-4 border-t border-gray-200 dark:border-gray-800">
      <router-link 
        to="/settings"
        class="flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 text-gray-600 dark:text-gray-400 transition"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
        <span class="text-sm font-medium">Settings</span>
      </router-link>
    </div>
  </aside>
</template>

<script setup>
import { onMounted } from 'vue'
import { useChatStore } from '@/stores/chat'

const chatStore = useChatStore()

onMounted(() => {
  chatStore.fetchConversations()
})
</script>
