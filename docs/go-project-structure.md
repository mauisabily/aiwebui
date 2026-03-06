# Go Project Structure

## Overview
This document outlines the Go project structure for the AI WebUI application.

## Project Structure

```
cmd/
├── server/
│   └── main.go              # Application entry point

internal/
├── api/                     # HTTP handlers and routes
│   ├── handlers/
│   │   ├── chat.go          # Chat-related handlers
│   │   ├── knowledge.go     # Knowledge base handlers
│   │   ├── models.go        # Model management handlers
│   │   └── users.go         # User management handlers
│   └── middleware/
│       └── auth.go          # Authentication middleware

├── config/                  # Configuration management
│   └── config.go            # Configuration loading and validation

├── database/                # Database connection and models
│   ├── connection.go        # Database connection setup
│   ├── models/
│   │   ├── user.go          # User model
│   │   ├── conversation.go   # Conversation model
│   │   ├── message.go       # Message model
│   │   ├── knowledge.go     # Knowledge base models
│   │   └── settings.go      # User settings model
│   └── migrations/          # Database migration utilities

├── ollama/                  # Ollama API client
│   ├── client.go            # Ollama client implementation
│   ├── models.go            # Ollama API data structures
│   └── embeddings.go        # Embedding generation utilities

├── rag/                     # Retrieval-Augmented Generation engine
│   ├── retriever.go         # Document retrieval logic
│   ├── searcher.go          # Search implementation
│   ├── ranker.go            # Result ranking utilities
│   └── context.go           # Context injection for prompts

├── utils/                   # Utility functions
│   ├── crypto.go            # Cryptographic utilities
│   ├── validation.go        # Input validation functions
│   └── helpers.go           # General helper functions

web/                         # Web assets
├── static/                  # Static assets (CSS, JS, images)
│   ├── css/
│   ├── js/
│   └── images/
├── templates/               # HTML templates
└── vue/                     # Vue.js components

configs/                     # Configuration files
    └── config.yaml          # Main configuration file

go.mod                       # Go module file
go.sum                       # Go checksum file
README.md                    # Project documentation
```

## Package Descriptions

### cmd/server
Contains the main application entry point. This package initializes all components and starts the HTTP server.

### internal/api
Handles HTTP routing and request/response processing:
- `handlers/`: Contains handler functions for different API endpoints
- `middleware/`: Contains middleware functions for authentication, logging, etc.

### internal/config
Manages application configuration:
- Loading configuration from files or environment variables
- Validating configuration values
- Providing centralized access to configuration throughout the application

### internal/database
Manages database connectivity and data models:
- `connection.go`: Establishes and manages database connections
- `models/`: Defines data structures and database operations for each entity
- `migrations/`: Handles database schema migrations

### internal/ollama
Provides a client for interacting with the Ollama API:
- `client.go`: Main client implementation for sending requests to Ollama
- `models.go`: Data structures that match Ollama API responses
- `embeddings.go`: Utilities for generating embeddings from text

### internal/rag
Implements the Retrieval-Augmented Generation functionality:
- `retriever.go`: Retrieves relevant documents from the knowledge base
- `searcher.go`: Implements search algorithms (keyword and vector-based)
- `ranker.go`: Ranks retrieved documents by relevance
- `context.go`: Injects retrieved context into prompts for the LLM

### internal/utils
Contains utility functions used across the application:
- Helper functions for common operations
- Validation functions for input data
- Cryptographic utilities for security

## Key Dependencies

The application will use the following Go modules:

1. `github.com/gin-gonic/gin` - HTTP web framework
2. `github.com/go-sql-driver/mysql` - MySQL driver for Go
3. `gopkg.in/yaml.v3` - YAML configuration parsing
4. `github.com/golang-jwt/jwt/v5` - JWT token handling (if authentication is implemented)
5. `github.com/sirupsen/logrus` - Structured logging

## Entry Point (main.go)

The main entry point will:
1. Load configuration
2. Initialize database connection
3. Set up HTTP routes and middleware
4. Start the HTTP server

## Configuration

The application will be configured using a YAML file (`configs/config.yaml`) with the following structure:

```yaml
server:
  port: 8080
  host: "localhost"

mysql:
  host: "localhost"
  port: 3306
  username: "ai_kpst"
  password: "c61762a01f19d8"
  database: "ai_kpst"
  charset: "utf8mb4"

ollama:
  base_url: "http://192.168.1.50:11434"
  default_model: "llama3"

rag:
  chunk_size: 1000
  chunk_overlap: 200
  max_results: 5
```

This structure provides a clean separation of concerns and follows Go best practices for project organization.