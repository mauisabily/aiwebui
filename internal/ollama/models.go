package ollama

// GenerateRequest represents a request to generate text
type GenerateRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream,omitempty"`
	Context []int                  `json:"context,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// GenerateResponse represents a response from text generation
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

// ChatMessage represents a single message in a chat
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents a request for chat completion
type ChatRequest struct {
	Model    string                 `json:"model"`
	Messages []ChatMessage          `json:"messages"`
	Stream   bool                   `json:"stream,omitempty"`
	Options  map[string]interface{} `json:"options,omitempty"`
}

// ChatResponse represents a response from chat completion
type ChatResponse struct {
	Model     string      `json:"model"`
	CreatedAt string      `json:"created_at"`
	Message   ChatMessage `json:"message"`
	Done      bool        `json:"done"`
}

// EmbeddingsRequest represents a request for embeddings
type EmbeddingsRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// EmbeddingsResponse represents a response from embeddings generation
type EmbeddingsResponse struct {
	Embedding []float64 `json:"embedding"`
}

// Model represents information about a model
type Model struct {
	Name      string            `json:"name"`
	Digest    string            `json:"digest"`
	ModifiedAt string           `json:"modified_at"`
	Size      int64             `json:"size"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// ListModelsResponse represents a response from listing models
type ListModelsResponse struct {
	Models []Model `json:"models"`
}

// ShowModelRequest represents a request to show model information
type ShowModelRequest struct {
	Name string `json:"name"`
}

// ShowModelResponse represents a response from showing model information
type ShowModelResponse struct {
	License    string `json:"license,omitempty"`
	Modelfile  string `json:"modelfile,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	Template   string `json:"template,omitempty"`
}