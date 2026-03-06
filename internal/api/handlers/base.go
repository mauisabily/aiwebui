package handlers

import (
	"aiwebui/internal/database"
	"aiwebui/internal/ollama"
)

// Handler represents the API handler structure
type Handler struct {
	DB     *database.DB
	Ollama *ollama.Client
}

// NewHandler creates a new API handler
func NewHandler(db *database.DB, ollama *ollama.Client) *Handler {
	return &Handler{
		DB:     db,
		Ollama: ollama,
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