package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"aiwebui/internal/rag"
	"aiwebui/internal/ollama"
)

// RAGStatus represents the RAG status for a conversation
type RAGStatus struct {
	Enabled           bool    `json:"enabled"`
	KnowledgeBaseIDs  []int   `json:"knowledge_base_ids"`
	LastUpdated       string  `json:"last_updated"`
}

// EnableRAG enables or disables RAG for a conversation
func (h *Handler) EnableRAG(c *gin.Context) {
	conversationID := c.Param("id")
	_, err := strconv.Atoi(conversationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid conversation ID",
			Code:  "INVALID_ID",
		})
		return
	}

	var req struct {
		Enabled          bool  `json:"enabled"`
		KnowledgeBaseIDs []int `json:"knowledge_base_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
			Code:  "INVALID_REQUEST",
		})
		return
	}

	// In a real implementation, we would save the RAG status to the database
	// For now, we'll return a success response
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: "RAG enabled for conversation",
	})
}

// GetRAGStatus retrieves the RAG status for a conversation
func (h *Handler) GetRAGStatus(c *gin.Context) {
	conversationID := c.Param("id")
	_, err := strconv.Atoi(conversationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid conversation ID",
			Code:  "INVALID_ID",
		})
		return
	}

	// In a real implementation, we would retrieve the RAG status from the database
	// For now, we'll return a dummy response
	status := RAGStatus{
		Enabled:          true,
		KnowledgeBaseIDs: []int{1, 2},
		LastUpdated:      "2026-03-06T12:30:00Z",
	}

	c.JSON(http.StatusOK, status)
}

// RAGChatRequest represents a chat request with RAG enhancement
type RAGChatRequest struct {
	Model             string `json:"model" binding:"required"`
	Message           string `json:"message" binding:"required"`
	ConversationID    int    `json:"conversation_id"`
	KnowledgeBaseIDs  []int  `json:"knowledge_base_ids"`
	EnableRAG         bool   `json:"enable_rag"`
}

// RAGChatResponse represents a chat response with RAG sources
type RAGChatResponse struct {
	ID             int    `json:"id"`
	ConversationID int    `json:"conversation_id"`
	Role           string `json:"role"`
	Content        string `json:"content"`
	Sources        []struct {
		DocumentID     int     `json:"document_id"`
		DocumentTitle  string  `json:"document_title"`
		RelevanceScore float64 `json:"relevance_score"`
	} `json:"sources"`
	Timestamp string `json:"timestamp"`
}

// SendRAGMessage handles sending a chat message with RAG enhancement
func (h *Handler) SendRAGMessage(c *gin.Context) {
	var req RAGChatRequest
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
		conversationID = 1 // Dummy ID
	}

	// Add user message to conversation
	// In a real implementation, we would save this to the database

	// Get current LLM settings
	settings, _ := h.getLLMSettings()
	
	// If RAG is enabled, retrieve relevant context and enhance prompt
	var contextResults []rag.SearchResult
	enhancedPrompt := req.Message
	if req.EnableRAG {
		var err error
		enhancedPrompt, contextResults, err = h.RAG.ProcessQuery(req.Message, 1) // Using dummy knowledgeBaseID 1
		if err != nil {
			c.JSON(http.StatusOK, RAGChatResponse{
				Role:    "assistant",
				Content: "ERROR: Failed to process RAG query. Details: " + err.Error(),
			})
			return
		}
	}

	// Prepare chat request
	chatReq := &ollama.ChatRequest{
		Model: req.Model,
		Messages: []ollama.ChatMessage{
			{Role: "user", Content: enhancedPrompt},
		},
	}
	if req.Model == "" {
		chatReq.Model = settings.DefaultModel
	}

	url := settings.OllamaURL
	if settings.LLMMode == "airllm" {
		url = settings.AirLLMURL
	}

	// Send request to configured provider
	client := ollama.NewClient(url)
	ollamaRes, err := client.Chat(chatReq)
	if err != nil {
		c.JSON(http.StatusOK, RAGChatResponse{
			Role:    "assistant",
			Content: "ERROR: Failed to connect to AI provider at " + url + ". Please check your settings. Details: " + err.Error(),
		})
		return
	}

	resContent := ollamaRes.Message.Content
	resTimestamp := ollamaRes.CreatedAt

	response := RAGChatResponse{
		ID:             12345,
		ConversationID: conversationID,
		Role:           "assistant",
		Content:        resContent,
		Timestamp:      resTimestamp,
	}

	// Map RAG results to response sources
	for _, res := range contextResults {
		response.Sources = append(response.Sources, struct {
			DocumentID     int     `json:"document_id"`
			DocumentTitle  string  `json:"document_title"`
			RelevanceScore float64 `json:"relevance_score"`
		}{
			DocumentID:     res.DocumentID,
			DocumentTitle:  res.DocumentTitle,
			RelevanceScore: res.RelevanceScore,
		})
	}

	// Add assistant message to conversation
	// In a real implementation, we would save this to the database

	c.JSON(http.StatusOK, response)
}