package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents a login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// Login handles admin login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
			Code:  "INVALID_REQUEST",
		})
		return
	}

	// Simple hardcoded check for admin
	if req.Username == "admin" && req.Password == "admin123" {
		c.JSON(http.StatusOK, LoginResponse{
			Success: true,
			Message: "Login successful",
			Token:   "dummy-admin-token", // In a real app, use JWT
		})
		return
	}

	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error: "Invalid credentials",
		Code:  "UNAUTHORIZED",
	})
}

// AuthMiddleware is a simple middleware to check for admin token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "dummy-admin-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
