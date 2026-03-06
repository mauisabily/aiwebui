package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"aiwebui/internal/config"
	"aiwebui/internal/database"
	"aiwebui/internal/ollama"
	"aiwebui/internal/api/handlers"
	"aiwebui/internal/rag"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Ollama client
	ollamaClient := ollama.NewClient(cfg.GetOllamaURL())

	// Initialize RAG engine
	ragEngine := rag.NewEngine(cfg, db, ollamaClient)

	// Create API handler
	handler := handlers.NewHandler(db, ollamaClient)

	// Set up Gin router
	router := gin.Default()

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "AI WebUI backend is running",
			})
		})

		// Chat endpoints
		api.POST("/chat", handler.SendMessage)
		api.POST("/chat/rag", handler.SendRAGMessage)
		api.GET("/conversations/:id", handler.GetConversation)
		api.GET("/conversations", handler.ListConversations)
		api.POST("/conversations", handler.CreateConversation)
		api.DELETE("/conversations/:id", handler.DeleteConversation)
		
		// Model endpoints
		api.GET("/models", handler.ListModels)
		api.GET("/models/:name", handler.GetModelInfo)
		
		// Knowledge base endpoints
		api.GET("/knowledge-bases", handler.ListKnowledgeBases)
		api.POST("/knowledge-bases", handler.CreateKnowledgeBase)
		api.GET("/knowledge-bases/:id", handler.GetKnowledgeBase)
		api.POST("/knowledge-bases/:id/documents", handler.UploadDocument)
		api.GET("/knowledge-bases/:id/search", handler.SearchKnowledgeBase)
		api.DELETE("/knowledge-bases/:id", handler.DeleteKnowledgeBase)
		
		// Settings endpoints
		api.GET("/settings", handler.GetSettings)
		api.PUT("/settings", handler.UpdateSettings)
		
		// RAG endpoints
		api.PUT("/conversations/:id/rag", handler.EnableRAG)
		api.GET("/conversations/:id/rag", handler.GetRAGStatus)
	}

	// Start server
	log.Printf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := router.Run(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}