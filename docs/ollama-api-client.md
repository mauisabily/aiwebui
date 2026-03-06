# Ollama API Client Design

## Overview
This document describes the design of the Ollama API client for the AI WebUI application. The client will communicate with the Ollama service running at http://192.168.1.50:11434.

## API Endpoints to Implement

Based on the Ollama API documentation, we need to implement the following endpoints:

1. **List Models** - GET `/api/tags`
2. **Generate Text** - POST `/api/generate`
3. **Chat Completion** - POST `/api/chat`
4. **Create Embeddings** - POST `/api/embeddings`
5. **Show Model Info** - POST `/api/show`

## Client Structure

```go
type OllamaClient struct {
    baseURL    string
    httpClient *http.Client
}

func NewOllamaClient(baseURL string) *OllamaClient {
    return &OllamaClient{
        baseURL:    baseURL,
        httpClient: &http.Client{Timeout: 30 * time.Second},
    }
}
```

## Data Structures

### Generate Request
```go
type GenerateRequest struct {
    Model   string `json:"model"`
    Prompt  string `json:"prompt"`
    Stream  bool   `json:"stream,omitempty"`
    Context []int  `json:"context,omitempty"`
    Options map[string]interface{} `json:"options,omitempty"`
}
```

### Generate Response
```go
type GenerateResponse struct {
    Model              string   `json:"model"`
    CreatedAt          string   `json:"created_at"`
    Response           string   `json:"response"`
    Done               bool     `json:"done"`
    Context            []int    `json:"context,omitempty"`
    TotalDuration      int64    `json:"total_duration,omitempty"`
    LoadDuration       int64    `json:"load_duration,omitempty"`
    PromptEvalCount    int      `json:"prompt_eval_count,omitempty"`
    PromptEvalDuration int64    `json:"prompt_eval_duration,omitempty"`
    EvalCount          int      `json:"eval_count,omitempty"`
    EvalDuration       int64    `json:"eval_duration,omitempty"`
}
```

### Chat Request
```go
type ChatMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type ChatRequest struct {
    Model   string        `json:"model"`
    Messages []ChatMessage `json:"messages"`
    Stream  bool          `json:"stream,omitempty"`
    Options map[string]interface{} `json:"options,omitempty"`
}
```

### Chat Response
```go
type ChatResponse struct {
    Model     string      `json:"model"`
    CreatedAt string      `json:"created_at"`
    Message   ChatMessage `json:"message"`
    Done      bool        `json:"done"`
}
```

### Embeddings Request
```go
type EmbeddingsRequest struct {
    Model  string `json:"model"`
    Prompt string `json:"prompt"`
}
```

### Embeddings Response
```go
type EmbeddingsResponse struct {
    Embedding []float64 `json:"embedding"`
}
```

## Methods Implementation

### 1. List Models
```go
func (c *OllamaClient) ListModels() ([]Model, error) {
    // GET /api/tags
}
```

### 2. Generate Text
```go
func (c *OllamaClient) Generate(req *GenerateRequest) (*GenerateResponse, error) {
    // POST /api/generate
}
```

### 3. Chat Completion
```go
func (c *OllamaClient) Chat(req *ChatRequest) (*ChatResponse, error) {
    // POST /api/chat
}
```

### 4. Create Embeddings
```go
func (c *OllamaClient) CreateEmbeddings(req *EmbeddingsRequest) (*EmbeddingsResponse, error) {
    // POST /api/embeddings
}
```

### 5. Show Model Info
```go
func (c *OllamaClient) ShowModel(modelName string) (map[string]interface{}, error) {
    // POST /api/show
}
```

## Error Handling

The client should handle various types of errors:

1. Network errors (timeouts, connection refused, etc.)
2. HTTP errors (4xx, 5xx status codes)
3. JSON parsing errors
4. Ollama-specific errors returned in the response body

## Streaming Responses

For streaming responses, we'll implement a separate method that accepts a callback function:

```go
func (c *OllamaClient) GenerateStream(req *GenerateRequest, callback func(response *GenerateResponse) error) error {
    // Handle streaming response
}
```

## Configuration

The client should be configurable with:

1. Base URL for the Ollama service
2. Timeout settings
3. Default model selection
4. Optional authentication (if needed in the future)

## Usage Examples

### Simple Text Generation
```go
client := NewOllamaClient("http://192.168.1.50:11434")
req := &GenerateRequest{
    Model:  "llama3",
    Prompt: "Why is the sky blue?",
}
response, err := client.Generate(req)
if err != nil {
    log.Fatal(err)
}
fmt.Println(response.Response)
```

### Chat Completion
```go
client := NewOllamaClient("http://192.168.1.50:11434")
req := &ChatRequest{
    Model: "llama3",
    Messages: []ChatMessage{
        {Role: "user", Content: "Hello, how are you?"},
    },
}
response, err := client.Chat(req)
if err != nil {
    log.Fatal(err)
}
fmt.Println(response.Message.Content)
```

### Creating Embeddings
```go
client := NewOllamaClient("http://192.168.1.50:11434")
req := &EmbeddingsRequest{
    Model:  "llama3",
    Prompt: "The sky is blue because of Rayleigh scattering.",
}
response, err := client.CreateEmbeddings(req)
if err != nil {
    log.Fatal(err)
}
fmt.Println(response.Embedding)
```

## Integration with RAG

The embeddings functionality will be crucial for the RAG implementation:

1. Generate embeddings for knowledge base documents
2. Generate embeddings for user queries
3. Calculate similarity between query and document embeddings
4. Retrieve most relevant documents based on similarity scores

## Security Considerations

1. Validate all inputs to prevent injection attacks
2. Implement proper timeout handling to prevent resource exhaustion
3. Log API interactions for monitoring and debugging
4. Handle sensitive data appropriately (no credentials in logs)

This design provides a robust foundation for interacting with the Ollama API while maintaining flexibility for future enhancements.