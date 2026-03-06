package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// If RAG is enabled, retrieve relevant context
	var context string
	if req.EnableRAG {
		// In a real implementation, we would search knowledge bases and retrieve relevant documents
		// For now, we'll use a dummy context
		context = "Based on the documentation, quantum computing is a type of computation that harnesses the principles of quantum mechanics."
	}

	// Prepare enhanced prompt with context
	enhancedPrompt := req.Message
	if context != "" {
		enhancedPrompt = "Context: " + context + "\n\nQuestion: " + req.Message
	}

	// Prepare chat request to Ollama
	// In a real implementation, we would use the ollama client to send the request
	// For now, we'll return a dummy response

	response := RAGChatResponse{
		ID:             12345,
		ConversationID: conversationID,
		Role:           "assistant",
		Content:        "Based on the information I have about quantum computing...",
		Sources: []struct {
			DocumentID     int     `json:"document_id"`
			DocumentTitle  string  `json:"document_title"`
			RelevanceScore float64 `json:"relevance_score"`
		}{
			{
				DocumentID:     456,
				DocumentTitle:  "Introduction to Quantum Mechanics",
				RelevanceScore: 0.92,
			},
		},
		Timestamp: "2026-03-06T12:34:56Z",
	}

	// Add assistant message to conversation
	// In a real implementation, we would save this to the database

	c.JSON(http.StatusOK, response)
}