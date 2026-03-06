package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"aiwebui/internal/ollama"
)

// ChatRequest represents a chat request
type ChatRequest struct {
    Model          string `json:"model"`
    Message        string `json:"message" binding:"required"`
    ConversationID int    `json:"conversation_id"`
    ContextIDs     []int  `json:"context_ids"`
}

// ChatResponse represents a chat response
type ChatResponse struct {
    ID             int    `json:"id"`
    ConversationID int    `json:"conversation_id"`
    Role           string `json:"role"`
    Content        string `json:"content"`
    Timestamp      string `json:"timestamp"`
    Model          string `json:"model"`
}

// SendMessage handles sending a chat message
func (h *Handler) SendMessage(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
			Code:  "INVALID_REQUEST",
		})
		return
	}

	// Get or create conversation
	conversationID := req.ConversationID
	if conversationID == 0 {
		res, err := h.DB.Exec("INSERT INTO conversations (title) VALUES (?)", "New Chat")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create conversation"})
			return
		}
		newID, _ := res.LastInsertId()
		conversationID = int(newID)
	}

	// Get current LLM settings
	settings, _ := h.getLLMSettings()

	// Get message history for context (messages already saved in DB - do NOT include current message yet)
	var history []ollama.ChatMessage
	rows, err := h.DB.Query("SELECT role, content FROM messages WHERE conversation_id = ? ORDER BY timestamp ASC", conversationID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var hMsg ollama.ChatMessage
			if err := rows.Scan(&hMsg.Role, &hMsg.Content); err == nil {
				history = append(history, hMsg)
			}
		}
	}

	// Append current user message to context (NOT saved to DB yet)
	history = append(history, ollama.ChatMessage{
		Role:    "user",
		Content: req.Message,
	})

	// Determine model
	model := req.Model
	if model == "" {
		model = settings.DefaultModel
	}

	// Determine URL
	url := settings.OllamaURL
	if settings.LLMMode == "airllm" {
		url = settings.AirLLMURL
	}

	// Prepare chat request to Ollama
	chatReq := &ollama.ChatRequest{
		Model:    model,
		Messages: history,
		Stream:   false,
	}

	fmt.Printf("[Chat] Sending to Ollama (%s) conv=%d model=%s messages=%d\n", url, conversationID, model, len(history))

	// === CALL AI FIRST ===
	client := ollama.NewClient(url)
	response, err := client.Chat(chatReq)
	if err != nil {
		errMsg := fmt.Sprintf("ERROR: Gagal sambung ke AI provider (%s). Sila semak Settings. Butiran: %v", url, err)
		fmt.Printf("[Chat] %s\n", errMsg)
		c.JSON(http.StatusOK, ChatResponse{
			ConversationID: conversationID,
			Role:           "assistant",
			Content:        errMsg,
		})
		return
	}

	assistantContent := response.Message.Content
	fmt.Printf("[Chat] Got response from Ollama: model=%s done=%v content_len=%d\n", response.Model, response.Done, len(assistantContent))

	if assistantContent == "" {
		errMsg := "ERROR: AI membalas dengan mesej kosong. Sila cuba lagi atau tukar model."
		fmt.Printf("[Chat] Empty response from model %s\n", model)
		c.JSON(http.StatusOK, ChatResponse{
			ConversationID: conversationID,
			Role:           "assistant",
			Content:        errMsg,
		})
		return
	}

	// === SAVE BOTH MESSAGES ONLY AFTER SUCCESSFUL AI RESPONSE ===
	h.DB.Exec("INSERT INTO messages (conversation_id, role, content) VALUES (?, 'user', ?)", conversationID, req.Message)
	h.DB.Exec("INSERT INTO messages (conversation_id, role, content) VALUES (?, 'assistant', ?)", conversationID, assistantContent)

	// Update conversation title if still "New Chat"
	if req.Message != "" {
		title := req.Message
		if len(title) > 50 {
			title = title[:50]
		}
		h.DB.Exec("UPDATE conversations SET title = ? WHERE id = ? AND title = 'New Chat'", title, conversationID)
	}

	c.JSON(http.StatusOK, ChatResponse{
		ConversationID: conversationID,
		Role:           "assistant",
		Content:        assistantContent,
		Timestamp:      response.CreatedAt,
		Model:          response.Model,
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetConversation retrieves a conversation and its messages
func (h *Handler) GetConversation(c *gin.Context) {
	id := c.Param("id")

	// Get messages
	rows, err := h.DB.Query("SELECT role, content, timestamp FROM messages WHERE conversation_id = ? ORDER BY timestamp ASC", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}
	defer rows.Close()

	type MessageItem struct {
		Role      string `json:"role"`
		Content   string `json:"content"`
		CreatedAt string `json:"created_at"`
	}

	var messages []MessageItem
	for rows.Next() {
		var m MessageItem
		var createdAt time.Time
		if err := rows.Scan(&m.Role, &m.Content, &createdAt); err == nil {
			m.CreatedAt = createdAt.Format(time.RFC3339) // Format time for JSON
			messages = append(messages, m)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"messages": messages,
	})
}

// ListConversations lists all conversations
func (h *Handler) ListConversations(c *gin.Context) {
	rows, err := h.DB.Query(`
		SELECT c.id, c.title, c.created_at, c.updated_at, COUNT(m.id) as message_count
		FROM conversations c
		LEFT JOIN messages m ON c.id = m.conversation_id
		GROUP BY c.id
		ORDER BY c.updated_at DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list conversations"})
		return
	}
	defer rows.Close()

	type ConversationItem struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		MessageCount int    `json:"message_count"`
	}

	var list []ConversationItem
	for rows.Next() {
		var item ConversationItem
		var createdAt, updatedAt time.Time
		var messageCount sql.NullInt64 // Use NullInt64 for COUNT which can be 0 (NULL if no messages)
		if err := rows.Scan(&item.ID, &item.Title, &createdAt, &updatedAt, &messageCount); err == nil {
			item.CreatedAt = createdAt.Format(time.RFC3339)
			item.UpdatedAt = updatedAt.Format(time.RFC3339)
			item.MessageCount = int(messageCount.Int64)
			list = append(list, item)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"conversations": list,
		"total":         len(list),
	})
}

// CreateConversation creates a new conversation
func (h *Handler) CreateConversation(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// If title is empty, default to "New Chat"
	if req.Title == "" {
		req.Title = "New Chat"
	}

	res, err := h.DB.Exec("INSERT INTO conversations (title) VALUES (?)", req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create conversation"})
		return
	}

	id, _ := res.LastInsertId()
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": req.Title,
	})
}

// DeleteConversation deletes a conversation by ID
func (h *Handler) DeleteConversation(c *gin.Context) {
	id := c.Param("id")
	_, err := h.DB.Exec("DELETE FROM conversations WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete conversation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
