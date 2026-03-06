package handlers

import (
	"aiwebui/internal/database"
	"aiwebui/internal/ollama"
	"aiwebui/internal/rag"
)

// Handler represents the API handler structure
type Handler struct {
	DB     *database.DB
	Ollama *ollama.Client
	RAG    *rag.Engine
}

// NewHandler creates a new API handler
func NewHandler(db *database.DB, ollama *ollama.Client, ragEngine *rag.Engine) *Handler {
	return &Handler{
		DB:     db,
		Ollama: ollama,
		RAG:    ragEngine,
	}
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

// SuccessResponse represents a standard success response
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}