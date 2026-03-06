# Frontend JavaScript Implementation

## Overview
This document describes the JavaScript implementation for the AI WebUI frontend, focusing on Vue.js components and their interactions with the backend API.

## Project Structure

```
web/
├── vue/
│   ├── components/
│   │   ├── chat/
│   │   │   ├── ChatView.vue
│   │   │   ├── MessageList.vue
│   │   │   ├── MessageInput.vue
│   │   │   └── MessageBubble.vue
│   │   ├── sidebar/
│   │   │   ├── Sidebar.vue
│   │   │   ├── ChatList.vue
│   │   │   └── KnowledgeBaseList.vue
│   │   ├── modals/
│   │   │   ├── SettingsModal.vue
│   │   │   ├── ModelManagerModal.vue
│   │   │   └── KnowledgeBaseModal.vue
│   │   ├── header/
│   │   │   ├── Header.vue
│   │   │   ├── ModelSelector.vue
│   │   │   └── UserProfile.vue
│   │   └── common/
│   │       ├── Button.vue
│   │       ├── Input.vue
│   │       └── Modal.vue
│   ├── composables/
│   │   ├── useChat.js
│   │   ├── useModels.js
│   │   ├── useKnowledgeBase.js
│   │   └── useSettings.js
│   ├── stores/
│   │   ├── chatStore.js
│   │   ├── modelStore.js
│   │   ├── knowledgeBaseStore.js
│   │   └── settingsStore.js
│   ├── utils/
│   │   ├── api.js
│   │   ├── formatter.js
│   │   └── helpers.js
│   ├── App.vue
│   └── main.js
├── static/
│   ├── css/
│   │   └── tailwind.css
│   └── js/
│       └── vendor/
└── templates/
    └── index.html
```

## Main Application Entry Point

### main.js
```javascript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import './static/css/tailwind.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.mount('#app')
```

### App.vue
```vue
<template>
  <div class="flex flex-col h-screen bg-gray-50 dark:bg-gray-900">
    <Header />
    <div class="flex flex-1 overflow-hidden">
      <Sidebar />
      <main class="flex-1 overflow-auto">
        <ChatView v-if="currentView === 'chat'" />
        <KnowledgeBaseView v-if="currentView === 'knowledge'" />
      </main>
    </div>
    <Footer />
    <Modals />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Header from './components/header/Header.vue'
import Sidebar from './components/sidebar/Sidebar.vue'
import ChatView from './components/chat/ChatView.vue'
import KnowledgeBaseView from './components/knowledge/KnowledgeBaseView.vue'
import Footer from './components/Footer.vue'
import Modals from './components/modals/Modals.vue'

const currentView = ref('chat')
</script>
```

## API Utility Layer

### api.js
```javascript
const API_BASE = '/api/v1'

class ApiClient {
  constructor() {
    this.baseURL = API_BASE
  }

  async request(endpoint, options = {}) {
    const url = `${this.baseURL}${endpoint}`
    const config = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    }

    try {
      const response = await fetch(url, config)
      
      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`)
      }
      
      return await response.json()
    } catch (error) {
      console.error('API request failed:', error)
      throw error
    }
  }

  // Chat endpoints
  async sendMessage(conversationId, message, model) {
    return this.request('/chat', {
      method: 'POST',
      body: JSON.stringify({ conversation_id: conversationId, message, model })
    })
  }

  async getConversation(id) {
    return this.request(`/conversations/${id}`)
  }

  async listConversations(limit = 20, offset = 0) {
    return this.request(`/conversations?limit=${limit}&offset=${offset}`)
  }

  async createConversation(title) {
    return this.request('/conversations', {
      method: 'POST',
      body: JSON.stringify({ title })
    })
  }

  async deleteConversation(id) {
    return this.request(`/conversations/${id}`, {
      method: 'DELETE'
    })
  }

  // Model endpoints
  async listModels() {
    return this.request('/models')
  }

  async getModelInfo(name) {
    return this.request(`/models/${name}`)
  }

  // Knowledge base endpoints
  async listKnowledgeBases() {
    return this.request('/knowledge-bases')
  }

  async getKnowledgeBase(id) {
    return this.request(`/knowledge-bases/${id}`)
  }

  async createKnowledgeBase(data) {
    return this.request('/knowledge-bases', {
      method: 'POST',
      body: JSON.stringify(data)
    })
  }

  async searchKnowledgeBase(id, query, limit = 10) {
    return this.request(`/knowledge-bases/${id}/search?query=${encodeURIComponent(query)}&limit=${limit}`)
  }

  async uploadDocument(knowledgeBaseId, formData) {
    return this.request(`/knowledge-bases/${knowledgeBaseId}/documents`, {
      method: 'POST',
      body: formData
    })
  }

  // Settings endpoints
  async getSettings() {
    return this.request('/settings')
  }

  async updateSettings(settings) {
    return this.request('/settings', {
      method: 'PUT',
      body: JSON.stringify(settings)
    })
  }
}

export const apiClient = new ApiClient()
```

## Composables for Logic Reuse

### useChat.js
```javascript
import { ref, reactive } from 'vue'
import { apiClient } from '../utils/api'

export function useChat() {
  const conversations = ref([])
  const currentConversation = ref(null)
  const messages = ref([])
  const isLoading = ref(false)
  const isStreaming = ref(false)

  const loadConversations = async () => {
    try {
      const data = await apiClient.listConversations()
      conversations.value = data.conversations
    } catch (error) {
      console.error('Failed to load conversations:', error)
    }
  }

  const createConversation = async (title) => {
    try {
      const data = await apiClient.createConversation(title)
      conversations.value.unshift(data)
      return data
    } catch (error) {
      console.error('Failed to create conversation:', error)
      throw error
    }
  }

  const loadConversation = async (id) => {
    try {
      const data = await apiClient.getConversation(id)
      currentConversation.value = data
      messages.value = data.messages
    } catch (error) {
      console.error('Failed to load conversation:', error)
      throw error
    }
  }

  const sendMessage = async (message, model) => {
    if (!currentConversation.value) {
      const newConv = await createConversation('New Conversation')
      currentConversation.value = newConv
    }

    // Add user message immediately
    const userMessage = {
      id: Date.now(),
      role: 'user',
      content: message,
      timestamp: new Date().toISOString()
    }
    messages.value.push(userMessage)

    // Add placeholder for AI response
    const aiMessage = {
      id: Date.now() + 1,
      role: 'assistant',
      content: '',
      timestamp: new Date().toISOString()
    }
    messages.value.push(aiMessage)

    isLoading.value = true
    isStreaming.value = true

    try {
      const response = await apiClient.sendMessage(
        currentConversation.value.id,
        message,
        model
      )
      
      // Update AI message with actual response
      aiMessage.content = response.content
      isLoading.value = false
      isStreaming.value = false
      
      return response
    } catch (error) {
      // Update AI message with error
      aiMessage.content = `Error: ${error.message}`
      isLoading.value = false
      isStreaming.value = false
      throw error
    }
  }

  const deleteConversation = async (id) => {
    try {
      await apiClient.deleteConversation(id)
      conversations.value = conversations.value.filter(conv => conv.id !== id)
      if (currentConversation.value?.id === id) {
        currentConversation.value = null
        messages.value = []
      }
    } catch (error) {
      console.error('Failed to delete conversation:', error)
      throw error
    }
  }

  return {
    conversations,
    currentConversation,
    messages,
    isLoading,
    isStreaming,
    loadConversations,
    createConversation,
    loadConversation,
    sendMessage,
    deleteConversation
  }
}
```

### useModels.js
```javascript
import { ref } from 'vue'
import { apiClient } from '../utils/api'

export function useModels() {
  const models = ref([])
  const selectedModel = ref('')
  const isLoading = ref(false)

  const loadModels = async () => {
    isLoading.value = true
    try {
      const data = await apiClient.listModels()
      models.value = data.models
      if (data.models.length > 0 && !selectedModel.value) {
        selectedModel.value = data.models[0].name
      }
    } catch (error) {
      console.error('Failed to load models:', error)
    } finally {
      isLoading.value = false
    }
  }

  const getModelInfo = async (name) => {
    try {
      return await apiClient.getModelInfo(name)
    } catch (error) {
      console.error(`Failed to get info for model ${name}:`, error)
      throw error
    }
  }

  return {
    models,
    selectedModel,
    isLoading,
    loadModels,
    getModelInfo
  }
}
```

### useKnowledgeBase.js
```javascript
import { ref } from 'vue'
import { apiClient } from '../utils/api'

export function useKnowledgeBase() {
  const knowledgeBases = ref([])
  const currentKnowledgeBase = ref(null)
  const documents = ref([])
  const isLoading = ref(false)
  const searchResults = ref([])

  const loadKnowledgeBases = async () => {
    isLoading.value = true
    try {
      const data = await apiClient.listKnowledgeBases()
      knowledgeBases.value = data.knowledge_bases
    } catch (error) {
      console.error('Failed to load knowledge bases:', error)
    } finally {
      isLoading.value = false
    }
  }

  const getKnowledgeBase = async (id) => {
    isLoading.value = true
    try {
      const data = await apiClient.getKnowledgeBase(id)
      currentKnowledgeBase.value = data
      documents.value = data.documents
    } catch (error) {
      console.error('Failed to load knowledge base:', error)
    } finally {
      isLoading.value = false
    }
  }

  const createKnowledgeBase = async (data) => {
    try {
      const kb = await apiClient.createKnowledgeBase(data)
      knowledgeBases.value.push(kb)
      return kb
    } catch (error) {
      console.error('Failed to create knowledge base:', error)
      throw error
    }
  }

  const searchKnowledgeBase = async (id, query) => {
    try {
      const data = await apiClient.searchKnowledgeBase(id, query)
      searchResults.value = data.results
      return data.results
    } catch (error) {
      console.error('Failed to search knowledge base:', error)
      throw error
    }
  }

  const uploadDocument = async (knowledgeBaseId, file, title) => {
    const formData = new FormData()
    formData.append('file', file)
    if (title) {
      formData.append('title', title)
    }

    try {
      const doc = await apiClient.uploadDocument(knowledgeBaseId, formData)
      // Refresh documents list
      await getKnowledgeBase(knowledgeBaseId)
      return doc
    } catch (error) {
      console.error('Failed to upload document:', error)
      throw error
    }
  }

  return {
    knowledgeBases,
    currentKnowledgeBase,
    documents,
    isLoading,
    searchResults,
    loadKnowledgeBases,
    getKnowledgeBase,
    createKnowledgeBase,
    searchKnowledgeBase,
    uploadDocument
  }
}
```

## Pinia Store Implementation

### chatStore.js
```javascript
import { defineStore } from 'pinia'
import { useChat } from '../composables/useChat'

export const useChatStore = defineStore('chat', () => {
  const {
    conversations,
    currentConversation,
    messages,
    isLoading,
    isStreaming,
    loadConversations,
    createConversation,
    loadConversation,
    sendMessage,
    deleteConversation
  } = useChat()

  return {
    // State
    conversations,
    currentConversation,
    messages,
    isLoading,
    isStreaming,

    // Actions
    loadConversations,
    createConversation,
    loadConversation,
    sendMessage,
    deleteConversation
  }
})
```

## Key Vue Components

### MessageInput.vue
```vue
<template>
  <div class="border-t border-gray-200 dark:border-gray-700 p-4">
    <div class="flex items-end gap-2">
      <textarea
        v-model="inputText"
        @keydown="handleKeydown"
        :disabled="isLoading"
        placeholder="Type your message..."
        class="flex-1 border border-gray-300 dark:border-gray-600 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-indigo-500 resize-none max-h-32"
        rows="1"
      ></textarea>
      <button
        @click="sendMessage"
        :disabled="!inputText.trim() || isLoading"
        class="btn-primary p-3"
      >
        <PaperAirplaneIcon class="h-5 w-5" />
      </button>
    </div>
    <div class="mt-2 text-sm text-gray-500 dark:text-gray-400">
      Press Enter to send, Shift+Enter for new line
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { PaperAirplaneIcon } from '@heroicons/vue/24/solid'
import { useChatStore } from '../../stores/chatStore'

const chatStore = useChatStore()
const inputText = ref('')

const emit = defineEmits(['send'])

const isLoading = computed(() => chatStore.isLoading)

const handleKeydown = (event) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendMessage()
  }
}

const sendMessage = () => {
  if (!inputText.value.trim()) return
  
  emit('send', inputText.value)
  inputText.value = ''
}

// Auto-resize textarea
watch(inputText, () => {
  const textarea = event.target
  if (textarea) {
    textarea.style.height = 'auto'
    textarea.style.height = textarea.scrollHeight + 'px'
  }
})
</script>
```

### ChatView.vue
```vue
<template>
  <div class="flex flex-col h-full">
    <div 
      ref="messagesContainer"
      class="flex-1 overflow-y-auto p-4 space-y-6"
    >
      <div v-if="messages.length === 0" class="flex flex-col items-center justify-center h-full text-gray-500">
        <ChatBubbleLeftRightIcon class="h-12 w-12 mb-4" />
        <p>Start a conversation with your AI assistant</p>
      </div>
      <MessageBubble
        v-for="message in messages"
        :key="message.id"
        :message="message"
      />
    </div>
    <MessageInput 
      @send="handleSendMessage"
      :is-loading="chatStore.isLoading"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { ChatBubbleLeftRightIcon } from '@heroicons/vue/24/outline'
import MessageBubble from './MessageBubble.vue'
import MessageInput from './MessageInput.vue'
import { useChatStore } from '../../stores/chatStore'
import { useModelsStore } from '../../stores/modelsStore'

const chatStore = useChatStore()
const modelsStore = useModelsStore()
const messagesContainer = ref(null)

const messages = computed(() => chatStore.messages)

onMounted(async () => {
  await chatStore.loadConversations()
  await modelsStore.loadModels()
})

const handleSendMessage = async (message) => {
  try {
    await chatStore.sendMessage(message, modelsStore.selectedModel)
    // Scroll to bottom after new message
    await nextTick()
    scrollToBottom()
  } catch (error) {
    console.error('Failed to send message:', error)
  }
}

const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}
</script>
```

## WebSocket Integration for Streaming

### websocket.js
```javascript
class WebSocketClient {
  constructor(url) {
    this.url = url
    this.ws = null
    this.listeners = {}
    this.reconnectInterval = 5000
    this.maxReconnectAttempts = 5
    this.reconnectAttempts = 0
  }

  connect() {
    this.ws = new WebSocket(this.url)

    this.ws.onopen = () => {
      console.log('WebSocket connected')
      this.reconnectAttempts = 0
      this.emit('open')
    }

    this.ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        this.emit(data.type, data.data)
      } catch (error) {
        console.error('Failed to parse WebSocket message:', error)
      }
    }

    this.ws.onclose = () => {
      console.log('WebSocket disconnected')
      this.emit('close')
      this.handleReconnect()
    }

    this.ws.onerror = (error) => {
      console.error('WebSocket error:', error)
      this.emit('error', error)
    }
  }

  handleReconnect() {
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      this.reconnectAttempts++
      setTimeout(() => {
        console.log(`Attempting to reconnect (${this.reconnectAttempts}/${this.maxReconnectAttempts})`)
        this.connect()
      }, this.reconnectInterval)
    }
  }

  send(type, data) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify({ type, data }))
    }
  }

  emit(event, data) {
    if (this.listeners[event]) {
      this.listeners[event].forEach(callback => callback(data))
    }
  }

  on(event, callback) {
    if (!this.listeners[event]) {
      this.listeners[event] = []
    }
    this.listeners[event].push(callback)
  }

  off(event, callback) {
    if (this.listeners[event]) {
      this.listeners[event] = this.listeners[event].filter(cb => cb !== callback)
    }
  }

  disconnect() {
    if (this.ws) {
      this.ws.close()
    }
  }
}

export const wsClient = new WebSocketClient('ws://localhost:8080/api/v1/ws/chat')
```

## Event Handling and User Interactions

### Event Flow Examples

1. **Sending a Message:**
   - User types message and presses Enter
   - MessageInput component emits 'send' event
   - ChatView handles the event and calls chatStore.sendMessage()
   - API request is made to backend
   - Response is received and displayed in MessageBubble

2. **Creating a New Chat:**
   - User clicks "New Chat" button
   - Sidebar component calls chatStore.createConversation()
   - API request creates new conversation
   - Conversation is added to conversations list
   - UI updates to show new empty chat

3. **Uploading a Document:**
   - User selects file in KnowledgeBaseView
   - useKnowledgeBase().uploadDocument() is called
   - FormData is sent to API endpoint
   - Backend processes file and stores in database
   - UI refreshes to show new document

## Error Handling and User Feedback

### Error Handling Strategy

1. **Network Errors:**
   - Display toast notification
   - Provide retry option
   - Log to console for debugging

2. **Validation Errors:**
   - Show inline error messages
   - Highlight problematic fields
   - Prevent form submission

3. **Server Errors:**
   - Parse error response
   - Display user-friendly message
   - Suggest corrective actions

### Loading States

1. **Global Loading:**
   - Show spinner in header during API requests
   - Disable interactive elements

2. **Component Loading:**
   - Skeleton loaders for lists
   - Progress indicators for uploads
   - Disabled buttons during operations

## Performance Optimization

### Techniques Used

1. **Virtual Scrolling:**
   - For long conversation histories
   - Using vue-virtual-scroller library

2. **Lazy Loading:**
   - Components loaded on demand
   - Images loaded when in viewport

3. **Caching:**
   - Store recent API responses
   - Use localStorage for persistent data

4. **Debouncing:**
   - Search input throttling
   - Window resize handlers

This JavaScript implementation provides a solid foundation for the frontend interactions of the AI WebUI, leveraging Vue.js reactivity and composition APIs to create a responsive and maintainable codebase.