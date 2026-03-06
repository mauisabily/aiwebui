<template>
  <div class="flex h-screen bg-white dark:bg-gray-950 overflow-hidden font-sans text-gray-900 dark:text-gray-100">
    
    <!-- Left Sidebar: Chat History -->
    <Sidebar />

    <!-- Main Content Area -->
    <main class="flex-1 flex flex-col relative min-w-0 bg-white dark:bg-gray-950 transition-all duration-300">
      
      <!-- Top Navigation / Toolbar -->
      <header class="h-14 border-b border-gray-100 dark:border-gray-800 flex items-center justify-between px-4 bg-white/80 dark:bg-gray-950/80 backdrop-blur-md sticky top-0 z-10">
        <div class="flex items-center space-x-3">
          <button 
            @click="chatStore.toggleSidebar" 
            class="p-2 hover:bg-gray-100 dark:hover:bg-gray-800 rounded-lg transition"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
          <div class="flex flex-col">
            <span class="text-sm font-bold tracking-tight">OpenWebUI Style</span>
            <span class="text-[10px] text-gray-400 uppercase tracking-widest">{{ chatStore.activeConversationId ? 'Session Active' : 'New Session' }}</span>
          </div>
        </div>
        
        <div class="flex items-center space-x-2">
            <button 
                @click="chatStore.toggleArtifacts"
                :class="chatStore.showArtifacts ? 'text-indigo-600 bg-indigo-50 dark:bg-indigo-900/20' : 'text-gray-500 hover:bg-gray-100 dark:hover:bg-gray-800'"
                class="p-2 rounded-lg transition flex items-center space-x-1"
                title="Toggle Artifacts"
            >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 21h6l-.75-4M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
            </button>
        </div>
      </header>

      <!-- Messages Area -->
      <div 
        ref="messageContainer"
        class="flex-1 overflow-y-auto scroll-smooth custom-scrollbar"
      >
        <div class="max-w-4xl mx-auto px-4 py-8 space-y-8">
            
          <!-- Welcome Message -->
          <div v-if="chatStore.messages.length === 0" class="flex flex-col items-center justify-center py-20 space-y-6 animate-fade-in">
              <div class="w-16 h-16 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-2xl flex items-center justify-center shadow-lg shadow-indigo-200 dark:shadow-none rotate-3">
                  <span class="text-3xl">🤖</span>
              </div>
              <h2 class="text-2xl font-black tracking-tight text-center">Apa yang boleh saya bantu hari ini?</h2>
              <p class="text-gray-500 dark:text-gray-400 text-center max-w-md">Tanya apa sahaja, atau muat naik dokumen untuk mula menganalisis.</p>
          </div>

          <!-- Message Bubbles -->
          <div 
            v-for="(msg, index) in chatStore.messages" 
            :key="msg.id" 
            class="flex space-x-4 group animate-message-in"
          >
            <div class="flex-shrink-0 mt-1">
                <div 
                    :class="msg.role === 'assistant' ? 'bg-indigo-600 text-white' : 'bg-gray-200 dark:bg-gray-800 text-gray-500'"
                    class="w-8 h-8 rounded-full flex items-center justify-center shadow-sm"
                >
                    <span v-if="msg.role === 'assistant'" class="text-xs font-bold">AI</span>
                    <span v-else class="text-xs">👤</span>
                </div>
            </div>
            <div class="flex-1 min-w-0">
                <div class="flex items-center space-x-2 mb-1">
                    <span class="font-bold text-xs uppercase tracking-wider text-gray-400">
                        {{ msg.role === 'assistant' ? 'Assistant' : 'You' }}
                    </span>
                    <span class="text-[10px] text-gray-300">{{ msg.timestamp }}</span>
                </div>
                <div 
                    :class="msg.role === 'assistant' ? 'prose dark:prose-invert max-w-none' : 'text-gray-700 dark:text-gray-300'"
                    class="leading-relaxed"
                >
                    <p v-if="!msg.content.startsWith('ERROR:')" class="whitespace-pre-wrap selection:bg-indigo-100 dark:selection:bg-indigo-900/40">{{ msg.content }}</p>
                    <p v-else class="text-red-500 dark:text-red-400 font-mono text-sm bg-red-50 dark:bg-red-900/20 p-4 rounded-xl border border-red-100 dark:border-red-900/30">
                        {{ msg.content }}
                    </p>
                </div>
            </div>
          </div>
          
          <!-- Thinking Indicator -->
          <div v-if="chatStore.loading" class="flex space-x-4 animate-pulse">
            <div class="flex-shrink-0">
                <div class="w-8 h-8 rounded-full bg-indigo-100 dark:bg-indigo-900/30 flex items-center justify-center">
                    <span class="text-sm">🧠</span>
                </div>
            </div>
            <div class="flex-1 space-y-2 py-1">
                <div class="h-2 bg-gray-200 dark:bg-gray-800 rounded-full w-3/4"></div>
                <div class="h-2 bg-gray-200 dark:bg-gray-800 rounded-full w-1/2"></div>
            </div>
          </div>

          <!-- Bottom Spacer -->
          <div class="h-32"></div>
        </div>
      </div>

      <!-- Input Area -->
      <div class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-white dark:from-gray-950 via-white/90 dark:via-gray-950/90 to-transparent pointer-events-none">
        <div class="max-w-3xl mx-auto flex flex-col space-y-2 pointer-events-auto">
            <form 
                @submit.prevent="chatStore.sendMessage" 
                class="relative bg-white dark:bg-gray-900 rounded-2xl border-2 border-gray-100 dark:border-gray-800 focus-within:border-indigo-500/50 transition-all shadow-xl shadow-gray-200/50 dark:shadow-none pr-2 py-2 flex items-center overflow-hidden"
            >
                <textarea 
                    v-model="chatStore.input" 
                    placeholder="Tanya apa-apa sahaja..." 
                    class="flex-1 bg-transparent border-none focus:ring-0 px-4 py-2 resize-none max-h-32 min-h-[44px] text-gray-700 dark:text-gray-200 placeholder:text-gray-400"
                    rows="1"
                    @keydown.enter.exact.prevent="chatStore.sendMessage"
                ></textarea>
                
                <div class="flex items-center space-x-1 px-2">
                    <button 
                        type="submit" 
                        :disabled="!chatStore.input.trim() || chatStore.loading"
                        class="p-2 bg-indigo-600 hover:bg-indigo-700 disabled:bg-gray-200 dark:disabled:bg-gray-800 text-white rounded-xl transition-all shadow-md shadow-indigo-100 dark:shadow-none active:scale-95"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
                        </svg>
                    </button>
                </div>
            </form>
            <p class="text-[10px] text-center text-gray-400">LLM can make mistakes. Check important info.</p>
        </div>
      </div>
    </main>

    <!-- Right Sidebar: Artifacts/Browser -->
    <ArtifactView />

  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { useChatStore } from '@/stores/chat'
import Sidebar from './Sidebar.vue'
import ArtifactView from './ArtifactView.vue'

const chatStore = useChatStore()
const route = useRoute()
const messageContainer = ref(null)

const scrollToBottom = async () => {
    await nextTick()
    if (messageContainer.value) {
        messageContainer.value.scrollTop = messageContainer.value.scrollHeight
    }
}

const loadChat = () => {
  const id = route.params.id
  if (id) {
    chatStore.selectConversation(id)
  } else {
    chatStore.messages = []
    chatStore.activeConversationId = null
  }
}

watch(() => route.params.id, loadChat)
watch(() => chatStore.messages.length, scrollToBottom)

onMounted(() => {
    loadChat()
    scrollToBottom()
})
</script>

<style scoped>
.animate-fade-in {
    animation: fadeIn 0.5s ease-out forwards;
}

.animate-message-in {
    animation: messageIn 0.3s ease-out forwards;
}

@keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
}

@keyframes messageIn {
    from { opacity: 0; transform: translateY(4px); }
    to { opacity: 1; transform: translateY(0); }
}

.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
    @apply bg-gray-200 dark:bg-gray-800 rounded-full;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    @apply bg-gray-300 dark:bg-gray-700;
}
</style>