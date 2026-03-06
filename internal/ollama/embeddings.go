package ollama

import (
	"math"
)

// CosineSimilarity calculates the cosine similarity between two vectors
func CosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0.0
	}

	// Calculate dot product
	dotProduct := 0.0
	for i := range a {
		dotProduct += a[i] * b[i]
	}

	// Calculate magnitudes
	magnitudeA := 0.0
	magnitudeB := 0.0
	for i := range a {
		magnitudeA += a[i] * a[i]
		magnitudeB += b[i] * b[i]
	}

	magnitudeA = math.Sqrt(magnitudeA)
	magnitudeB = math.Sqrt(magnitudeB)

	// Avoid division by zero
	if magnitudeA == 0 || magnitudeB == 0 {
		return 0.0
	}

	return dotProduct / (magnitudeA * magnitudeB)
}

// EmbeddingGenerator wraps the Ollama client for embedding generation
type EmbeddingGenerator struct {
	client *Client
	model  string
}

// NewEmbeddingGenerator creates a new embedding generator
func NewEmbeddingGenerator(client *Client, model string) *EmbeddingGenerator {
	return &EmbeddingGenerator{
		client: client,
		model:  model,
	}
}

// GenerateEmbedding generates an embedding for the given text
func (eg *EmbeddingGenerator) GenerateEmbedding(text string) ([]float64, error) {
	req := &EmbeddingsRequest{
		Model:  eg.model,
		Prompt: text,
	}

	resp, err := eg.client.CreateEmbeddings(req)
	if err != nil {
		return nil, err
	}

	return resp.Embedding, nil
}