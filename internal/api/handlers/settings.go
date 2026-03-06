package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSettings represents user settings
type UserSettings struct {
	DefaultModel          string `json:"default_model"`
	Theme                 string `json:"theme"`
	Language              string `json:"language"`
	AutoSaveConversations bool   `json:"auto_save_conversations"`
}

// GetSettings retrieves user settings
func (h *Handler) GetSettings(c *gin.Context) {
	// In a real implementation, we would retrieve settings from the database
	// For now, we'll return dummy settings
	settings := UserSettings{
		DefaultModel:          "llama3",
		Theme:                 "dark",
		Language:              "en",
		AutoSaveConversations: true,
	}

	c.JSON(http.StatusOK, settings)
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

	// In a real implementation, we would save settings to the database
	// For now, we'll return the updated settings
	settings := UserSettings{
		DefaultModel:          req.DefaultModel,
		Theme:                 req.Theme,
		Language:              "en", // Default value
		AutoSaveConversations: true, // Default value
	}

	// Merge with request data
	if req.Language != "" {
		settings.Language = req.Language
	}
	if req.AutoSaveConversations {
		settings.AutoSaveConversations = req.AutoSaveConversations
	}

	c.JSON(http.StatusOK, settings)
}