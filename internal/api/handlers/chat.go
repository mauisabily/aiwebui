package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"aiwebui/internal/database/models"
	"aiwebui/internal/ollama"
)

// ChatRequest represents a chat request
type ChatRequest struct {
	Model          string `json:"model" binding:"required"`
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
		// Create new conversation
		conversation := &models.Conversation{
			UserID: 1, // Default user for now
			Title:  "New Conversation",
		}
		// In a real implementation, we would save this to the database
		// For now, we'll just use a dummy ID
		conversationID = 1
	}

	// Add user message to conversation
	userMessage := &models.Message{
		ConversationID: conversationID,
		Role:           "user",
		Content:        req.Message,
	}
	// In a real implementation, we would save this to the database

	// Prepare chat request to Ollama
	chatReq := &ollama.ChatRequest{
		Model: req.Model,
		Messages: []ollama.ChatMessage{
			{
				Role:    "user",
				Content: req.Message,
			},
		},
	}

	// Send request to Ollama
	response, err := h.Ollama.Chat(chatReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to generate response",
			Code:  "GENERATION_ERROR",
		})
		return
	}

	// Create response
	chatResponse := ChatResponse{
		ID:             1, // Dummy ID
		ConversationID: conversationID,
		Role:           "assistant",
		Content:        response.Message.Content,
		Timestamp:      response.CreatedAt,
		Model:          response.Model,
	}

	// Add assistant message to conversation
	assistantMessage := &models.Message{
		ConversationID: conversationID,
		Role:           "assistant",
		Content:        response.Message.Content,
	}
	// In a real implementation, we would save this to the database

	c.JSON(http.StatusOK, chatResponse)
}

// GetConversation retrieves a conversation by ID
func (h *Handler) GetConversation(c *gin.Context) {
	id := c.Param("id")
	conversationID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid conversation ID",
			Code:  "INVALID_ID",
		})
		return
	}

	// In a real implementation, we would retrieve the conversation from the database
	// For now, we'll return a dummy response
	conversation := struct {
		ID        int             `json:"id"`
		Title     string          `json:"title"`
		CreatedAt string          `json:"created_at"`
		UpdatedAt string          `json:"updated_at"`
		Messages  []models.Message `json:"messages"`
	}{
		ID:        conversationID,
		Title:     "Sample Conversation",
		CreatedAt: "2026-03-06T12:00:00Z",
		UpdatedAt: "2026-03-06T12:34:56Z",
		Messages: []models.Message{
			{
				ID:             1,
				ConversationID: conversationID,
				Role:           "user",
				Content:        "Hello, how are you?",
				CreatedAt:      "2026-03-06T12:00:00Z",
			},
			{
				ID:             2,
				ConversationID: conversationID,
				Role:           "assistant",
				Content:        "I'm doing well, thank you for asking!",
				CreatedAt:      "2026-03-06T12:00:30Z",
			},
		},
	}

	c.JSON(http.StatusOK, conversation)
}

// ListConversations lists all conversations
func (h *Handler) ListConversations(c *gin.Context) {
	// Parse query parameters
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	// In a real implementation, we would retrieve conversations from the database
	// For now, we'll return a dummy response
	conversations := struct {
		Conversations []struct {
			ID           int    `json:"id"`
			Title        string `json:"title"`
			CreatedAt    string `json:"created_at"`
			UpdatedAt    string `json:"updated_at"`
			MessageCount int    `json:"message_count"`
		} `json:"conversations"`
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}{
		Conversations: []struct {
			ID           int    `json:"id"`
			Title        string `json:"title"`
			CreatedAt    string `json:"created_at"`
			UpdatedAt    string `json:"updated_at"`
			MessageCount int    `json:"message_count"`
		}{
			{
				ID:           123,
				Title:        "General Discussion",
				CreatedAt:    "2026-03-06T12:00:00Z",
				UpdatedAt:    "2026-03-06T12:34:56Z",
				MessageCount: 15,
			},
			{
				ID:           124,
				Title:        "Technical Questions",
				CreatedAt:    "2026-03-05T10:30:00Z",
				UpdatedAt:    "2026-03-05T11:45:22Z",
				MessageCount: 8,
			},
		},
		Total:  42,
		Limit:  20,
		Offset: 0,
	}

	lim, _ := strconv.Atoi(limit)
	off, _ := strconv.Atoi(offset)
	conversations.Limit = lim
	conversations.Offset = off

	c.JSON(http.StatusOK, conversations)
}

// CreateConversation creates a new conversation
func (h *Handler) CreateConversation(c *gin.Context) {
	var req struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
			Code:  "INVALID_REQUEST",
		})
		return
	}

	// In a real implementation, we would save the conversation to the database
	// For now, we'll return a dummy response
	conversation := struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		ID:        125,
		Title:     req.Title,
		CreatedAt: "2026-03-06T12:45:00Z",
		UpdatedAt: "2026-03-06T12:45:00Z",
	}

	c.JSON(http.StatusOK, conversation)
}

// DeleteConversation deletes a conversation by ID
func (h *Handler) DeleteConversation(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid conversation ID",
			Code:  "INVALID_ID",
		})
		return
	}

	// In a real implementation, we would delete the conversation from the database
	// For now, we'll return a success response
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: "Conversation deleted successfully",
	})
}