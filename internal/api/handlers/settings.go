package handlers

import (
	"aiwebui/internal/ollama"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSettings represents user settings
type UserSettings struct {
	DefaultModel          string `json:"default_model"`
	Theme                 string `json:"theme"`
	Language              string `json:"language"`
	AutoSaveConversations bool   `json:"auto_save_conversations"`
	LLMMode               string `json:"llm_mode"` // "ollama" or "airllm"
	OllamaURL             string `json:"ollama_url"`
	AirLLMURL             string `json:"airllm_url"`
}

// GetSettings retrieves user settings
func (h *Handler) GetSettings(c *gin.Context) {
	// Ensure table exists
	_, _ = h.DB.Exec("CREATE TABLE IF NOT EXISTS settings (" +
		"`key` VARCHAR(255) PRIMARY KEY," +
		"`value` TEXT NOT NULL" +
		")")

	settings := UserSettings{
		DefaultModel: "llama3",
		Theme:        "dark",
		Language:     "en",
		LLMMode:      "ollama",
		OllamaURL:    "http://192.168.1.50:11434",
		AirLLMURL:    "http://localhost:8000",
	}

	rows, err := h.DB.Query("SELECT `key`, `value` FROM settings")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var k, v string
			if err := rows.Scan(&k, &v); err == nil {
				switch k {
				case "default_model":
					settings.DefaultModel = v
				case "theme":
					settings.Theme = v
				case "language":
					settings.Language = v
				case "llm_mode":
					settings.LLMMode = v
				case "ollama_url":
					settings.OllamaURL = v
				case "airllm_url":
					settings.AirLLMURL = v
				}
			}
		}
	}

	c.JSON(http.StatusOK, settings)
}

// TestConnectionRequest represents a request to test a connection
type TestConnectionRequest struct {
	LLMMode   string `json:"llm_mode"`
	OllamaURL string `json:"ollama_url"`
	AirLLMURL string `json:"airllm_url"`
}

// TestConnectionResponse represents a response from testing a connection
type TestConnectionResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Models  []string `json:"models"`
}

// TestConnection tests the connection to an LLM provider
func (h *Handler) TestConnection(c *gin.Context) {
	var req TestConnectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
			Code:  "INVALID_REQUEST",
		})
		return
	}

	url := req.OllamaURL
	if req.LLMMode == "airllm" {
		url = req.AirLLMURL
	}

	// Create a temporary client to test the connection
	testClient := ollama.NewClient(url)
	modelsResp, err := testClient.ListModels()
	if err != nil {
		c.JSON(http.StatusOK, TestConnectionResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var models []string
	for _, m := range modelsResp.Models {
		models = append(models, m.Name)
	}

	c.JSON(http.StatusOK, TestConnectionResponse{
		Success: true,
		Message: "Connection successful",
		Models:  models,
	})
}

// UpdateSettings updates user settings
func (h *Handler) UpdateSettings(c *gin.Context) {
	var req UserSettings

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
			Code:  "INVALID_REQUEST",
		})
		return
	}

	upsertQuery := "INSERT INTO settings (`key`, `value`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `value` = VALUES(`value`)"
	
	settingsMap := map[string]string{
		"default_model": req.DefaultModel,
		"theme":         req.Theme,
		"language":      req.Language,
		"llm_mode":      req.LLMMode,
		"ollama_url":    req.OllamaURL,
		"airllm_url":    req.AirLLMURL,
	}

	for k, v := range settingsMap {
		if v != "" {
			_, _ = h.DB.Exec(upsertQuery, k, v)
		}
	}

	c.JSON(http.StatusOK, req)
}

// getLLMSettings retrieves LLM settings from the database
func (h *Handler) getLLMSettings() (*UserSettings, error) {
	settings := &UserSettings{
		DefaultModel: "llama3",
		LLMMode:      "ollama",
		OllamaURL:    "http://192.168.1.50:11434",
		AirLLMURL:    "http://localhost:8000",
	}

	rows, err := h.DB.Query("SELECT `key`, `value` FROM settings")
	if err != nil {
		return settings, err
	}
	defer rows.Close()

	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err == nil {
			switch k {
			case "default_model":
				settings.DefaultModel = v
			case "llm_mode":
				settings.LLMMode = v
			case "ollama_url":
				settings.OllamaURL = v
			case "airllm_url":
				settings.AirLLMURL = v
			}
		}
	}
	return settings, nil
}