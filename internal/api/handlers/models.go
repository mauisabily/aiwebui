package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"aiwebui/internal/ollama"
)

// ListModels lists all available models
func (h *Handler) ListModels(c *gin.Context) {
	response, err := h.Ollama.ListModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to list models",
			Code:  "LIST_MODELS_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetModelInfo retrieves information about a specific model
func (h *Handler) GetModelInfo(c *gin.Context) {
	modelName := c.Param("name")

	response, err := h.Ollama.ShowModel(modelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to get model info",
			Code:  "MODEL_INFO_ERROR",
		})
		return
	}

	// Create a combined response with model details
	modelInfo := struct {
		Name    string                 `json:"name"`
		Digest  string                 `json:"digest"`
		ModifiedAt string            `json:"modified_at"`
		Size    int64                  `json:"size"`
		Details map[string]interface{} `json:"details"`
	}{
		Name:       modelName,
		Digest:     "", // Not provided by ShowModel
		ModifiedAt: "", // Not provided by ShowModel
		Size:       0,  // Not provided by ShowModel
		Details: map[string]interface{}{
			"license":    response.License,
			"modelfile":  response.Modelfile,
			"parameters": response.Parameters,
			"template":   response.Template,
		},
	}

	c.JSON(http.StatusOK, modelInfo)
}