import { defineStore } from 'pinia'
import axios from 'axios'

export const useChatStore = defineStore('chat', {
  state: () => ({
    conversations: [],
    messages: [],
    activeConversationId: null,
    input: '',
    loading: false,
    showSidebar: true,
    showArtifacts: false,
    currentArtifact: null
  }),
  actions: {
    async fetchConversations() {
      try {
        const res = await axios.get('/api/v1/conversations')
        this.conversations = res.data.conversations || []
      } catch (e) {
        console.error('Failed to fetch conversations:', e)
      }
    },

    async selectConversation(id) {
      this.loading = true
      this.activeConversationId = id
      try {
        const res = await axios.get(`/api/v1/conversations/${id}`)
        this.messages = res.data.messages || []
        // Hide artifacts when switching
        this.showArtifacts = false
      } catch (e) {
        console.error('Failed to fetch conversation:', e)
      } finally {
        this.loading = false
      }
    },

    async createNewChat() {
      try {
        const res = await axios.post('/api/v1/conversations', { title: 'New Chat' })
        await this.fetchConversations()
        await this.selectConversation(res.data.id)
      } catch (e) {
        console.error('Failed to create conversation:', e)
      }
    },

    async deleteConversation(id) {
      if (!confirm('Are you sure you want to delete this chat?')) return
      try {
        await axios.delete(`/api/v1/conversations/${id}`)
        await this.fetchConversations()
        if (this.activeConversationId === id) {
          this.messages = []
          this.activeConversationId = null
        }
      } catch (e) {
        console.error('Failed to delete conversation:', e)
      }
    },

    async sendMessage() {
      if (!this.input.trim() || this.loading) return

      if (!this.activeConversationId) {
        await this.createNewChat()
      }

      const userMsg = { id: Date.now(), role: 'user', content: this.input }
      this.messages.push(userMsg)
      const userInput = this.input
      const conversationId = this.activeConversationId
      this.input = ''
      this.loading = true

      try {
        const res = await axios.post('/api/v1/chat', {
          message: userInput,
          conversation_id: conversationId
        })
        const assistantMsg = {
          id: Date.now() + 1,
          role: 'assistant',
          content: res.data.content,
          timestamp: res.data.timestamp
        }
        this.messages.push(assistantMsg)

        // Detect artifacts (simple heuristic for demo)
        if (res.data.content.includes('```html') || res.data.content.includes('```javascript')) {
          this.currentArtifact = res.data.content
          this.showArtifacts = true
        }

      } catch (e) {
        const errorMsg = {
          id: Date.now() + 1,
          role: 'assistant',
          content: 'Error: Connection failed. Please check your LLM settings.'
        }
        this.messages.push(errorMsg)
      } finally {
        this.loading = false
      }
    },

    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },

    toggleArtifacts() {
      this.showArtifacts = !this.showArtifacts
    }
  }
})