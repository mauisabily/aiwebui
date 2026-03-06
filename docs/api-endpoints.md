# REST API Endpoints Design

## Overview
This document describes the REST API endpoints for the AI WebUI application, covering chat functionality, knowledge base management, and user settings.

## API Standards

- All endpoints will be prefixed with `/api/v1`
- JSON request/response format
- Standard HTTP status codes
- Consistent error response format
- Token-based authentication (if implemented)

## Error Response Format

```json
{
  "error": "Description of the error",
  "code": "ERROR_CODE"
}
```

## Chat Endpoints

### 1. Send Chat Message
```
POST /api/v1/chat
```

**Request Body:**
```json
{
  "model": "llama3",
  "message": "Hello, how are you?",
  "conversation_id": 123,
  "context_ids": [456, 789]
}
```

**Response:**
```json
{
  "id": 12345,
  "conversation_id": 123,
  "role": "assistant",
  "content": "I'm doing well, thank you for asking!",
  "timestamp": "2026-03-06T12:34:56Z",
  "model": "llama3"
}
```

### 2. Get Conversation History
```
GET /api/v1/conversations/{id}
```

**Response:**
```json
{
  "id": 123,
  "title": "General Discussion",
  "created_at": "2026-03-06T12:00:00Z",
  "updated_at": "2026-03-06T12:34:56Z",
  "messages": [
    {
      "id": 101,
      "role": "user",
      "content": "Hello, how are you?",
      "timestamp": "2026-03-06T12:00:00Z"
    },
    {
      "id": 102,
      "role": "assistant",
      "content": "I'm doing well, thank you for asking!",
      "timestamp": "2026-03-06T12:00:30Z"
    }
  ]
}
```

### 3. List Conversations
```
GET /api/v1/conversations
```

**Query Parameters:**
- `limit` (optional): Number of conversations to return (default: 20)
- `offset` (optional): Offset for pagination (default: 0)

**Response:**
```json
{
  "conversations": [
    {
      "id": 123,
      "title": "General Discussion",
      "created_at": "2026-03-06T12:00:00Z",
      "updated_at": "2026-03-06T12:34:56Z",
      "message_count": 15
    },
    {
      "id": 124,
      "title": "Technical Questions",
      "created_at": "2026-03-05T10:30:00Z",
      "updated_at": "2026-03-05T11:45:22Z",
      "message_count": 8
    }
  ],
  "total": 42,
  "limit": 20,
  "offset": 0
}
```

### 4. Create New Conversation
```
POST /api/v1/conversations
```

**Request Body:**
```json
{
  "title": "New Discussion Topic"
}
```

**Response:**
```json
{
  "id": 125,
  "title": "New Discussion Topic",
  "created_at": "2026-03-06T12:45:00Z",
  "updated_at": "2026-03-06T12:45:00Z"
}
```

### 5. Delete Conversation
```
DELETE /api/v1/conversations/{id}
```

**Response:**
```json
{
  "success": true,
  "message": "Conversation deleted successfully"
}
```

## Model Management Endpoints

### 1. List Available Models
```
GET /api/v1/models
```

**Response:**
```json
{
  "models": [
    {
      "name": "llama3",
      "digest": "sha256:...",
      "modified_at": "2026-03-01T10:00:00Z",
      "size": 4650000000
    },
    {
      "name": "mistral",
      "digest": "sha256:...",
      "modified_at": "2026-02-28T15:30:00Z",
      "size": 3200000000
    }
  ]
}
```

### 2. Get Model Information
```
GET /api/v1/models/{name}
```

**Response:**
```json
{
  "name": "llama3",
  "digest": "sha256:...",
  "modified_at": "2026-03-01T10:00:00Z",
  "size": 4650000000,
  "details": {
    "format": "gguf",
    "family": "llama",
    "families": ["llama"],
    "parameter_size": "8B",
    "quantization_level": "Q4_K_M"
  }
}
```

## Knowledge Base Endpoints

### 1. List Knowledge Bases
```
GET /api/v1/knowledge-bases
```

**Response:**
```json
{
  "knowledge_bases": [
    {
      "id": 1,
      "name": "Default Knowledge Base",
      "description": "Main knowledge base for general information",
      "document_count": 42,
      "created_at": "2026-03-01T09:00:00Z"
    }
  ]
}
```

### 2. Create Knowledge Base
```
POST /api/v1/knowledge-bases
```

**Request Body:**
```json
{
  "name": "Technical Documentation",
  "description": "Knowledge base for technical documentation"
}
```

**Response:**
```json
{
  "id": 2,
  "name": "Technical Documentation",
  "description": "Knowledge base for technical documentation",
  "document_count": 0,
  "created_at": "2026-03-06T12:50:00Z"
}
```

### 3. Get Knowledge Base Details
```
GET /api/v1/knowledge-bases/{id}
```

**Response:**
```json
{
  "id": 1,
  "name": "Default Knowledge Base",
  "description": "Main knowledge base for general information",
  "documents": [
    {
      "id": 101,
      "title": "Getting Started Guide",
      "source_type": "manual",
      "created_at": "2026-03-01T10:00:00Z"
    }
  ],
  "created_at": "2026-03-01T09:00:00Z"
}
```

### 4. Upload Document to Knowledge Base
```
POST /api/v1/knowledge-bases/{id}/documents
```

**Request Body (multipart/form-data):**
- `file`: The document file to upload
- `title` (optional): Title for the document

**Response:**
```json
{
  "id": 102,
  "title": "API Reference Manual",
  "source_type": "upload",
  "chunk_count": 15,
  "created_at": "2026-03-06T12:55:00Z"
}
```

### 5. Search Knowledge Base
```
GET /api/v1/knowledge-bases/{id}/search
```

**Query Parameters:**
- `query`: Search query text
- `limit` (optional): Maximum number of results (default: 10)

**Response:**
```json
{
  "query": "how to install",
  "results": [
    {
      "document_id": 101,
      "document_title": "Getting Started Guide",
      "content": "To install the application, follow these steps...",
      "relevance_score": 0.95
    }
  ]
}
```

### 6. Delete Knowledge Base
```
DELETE /api/v1/knowledge-bases/{id}
```

**Response:**
```json
{
  "success": true,
  "message": "Knowledge base deleted successfully"
}
```

## User Settings Endpoints

### 1. Get User Settings
```
GET /api/v1/settings
```

**Response:**
```json
{
  "default_model": "llama3",
  "theme": "dark",
  "language": "en",
  "auto_save_conversations": true
}
```

### 2. Update User Settings
```
PUT /api/v1/settings
```

**Request Body:**
```json
{
  "default_model": "mistral",
  "theme": "light"
}
```

**Response:**
```json
{
  "default_model": "mistral",
  "theme": "light",
  "language": "en",
  "auto_save_conversations": true
}
```

## RAG Integration Endpoints

### 1. Enable/Disable RAG for Conversation
```
PUT /api/v1/conversations/{id}/rag
```

**Request Body:**
```json
{
  "enabled": true,
  "knowledge_base_ids": [1, 2]
}
```

**Response:**
```json
{
  "success": true,
  "message": "RAG enabled for conversation"
}
```

### 2. Get RAG Status for Conversation
```
GET /api/v1/conversations/{id}/rag
```

**Response:**
```json
{
  "enabled": true,
  "knowledge_base_ids": [1, 2],
  "last_updated": "2026-03-06T12:30:00Z"
}
```

## WebSocket Endpoint (for streaming responses)

### 1. Chat Stream
```
WebSocket /api/v1/ws/chat
```

**Messages from client:**
```json
{
  "type": "chat_message",
  "data": {
    "model": "llama3",
    "message": "Explain quantum computing",
    "conversation_id": 123
  }
}
```

**Messages from server:**
```json
{
  "type": "chat_response_chunk",
  "data": {
    "content": "Quantum computing is..."
  }
}
```

```json
{
  "type": "chat_response_complete",
  "data": {
    "content": "Quantum computing is a type of computation that harnesses the principles of quantum mechanics.",
    "model": "llama3"
  }
}
```

## Rate Limiting

To prevent abuse, the API will implement rate limiting:

- 60 requests per minute per IP for most endpoints
- 10 requests per minute for model generation endpoints
- Exponential backoff for repeated violations

## CORS Policy

The API will allow CORS requests from the web frontend domain with the following headers:
- `Access-Control-Allow-Origin`: Web frontend domain
- `Access-Control-Allow-Methods`: GET, POST, PUT, DELETE, OPTIONS
- `Access-Control-Allow-Headers`: Content-Type, Authorization

## Authentication

If authentication is implemented, endpoints will require a JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

Unauthenticated requests will receive a 401 Unauthorized response.

This API design provides a comprehensive set of endpoints to support all the required functionality for the AI WebUI application while maintaining consistency and ease of use.