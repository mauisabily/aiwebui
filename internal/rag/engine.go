package rag

import (
	"aiwebui/internal/database"
	"aiwebui/internal/ollama"
	"aiwebui/internal/config"
)

// Engine represents the RAG engine
type Engine struct {
	config     *config.RAGConfig
	chunker    *DocumentChunker
	searcher   *Searcher
	ranker     *ResultRanker
	injector   *ContextInjector
}

// NewEngine creates a new RAG engine
func NewEngine(cfg *config.Config, db *database.DB, ollamaClient *ollama.Client) *Engine {
	// Create embedding generator
	embeddingGenerator := ollama.NewEmbeddingGenerator(ollamaClient, cfg.Ollama.DefaultModel)
	
	// Create components
	chunker := NewDocumentChunker(cfg.RAG.ChunkSize, cfg.RAG.ChunkOverlap)
	searcher := NewSearcher(db, embeddingGenerator)
	ranker := NewResultRanker(0.6, 0.4, 0.3) // Default weights
	injector := NewContextInjector(2000)     // Max 2000 chars for context
	
	return &Engine{
		config:   &cfg.RAG,
		chunker:  chunker,
		searcher: searcher,
		ranker:   ranker,
		injector: injector,
	}
}

// Retrieve retrieves relevant documents for a query
func (e *Engine) Retrieve(query string, knowledgeBaseID int) ([]SearchResult, error) {
	// Perform search
	results, err := e.searcher.Search(query, knowledgeBaseID, e.config.MaxResults)
	if err != nil {
		return nil, err
	}
	
	return results, nil
}

// EnhancePrompt enhances a prompt with retrieved context
func (e *Engine) EnhancePrompt(prompt string, results []SearchResult) string {
	return e.injector.InjectContext(prompt, results)
}

// ProcessQuery processes a query through the full RAG pipeline
func (e *Engine) ProcessQuery(query string, knowledgeBaseID int) (string, []SearchResult, error) {
	// Retrieve relevant documents
	results, err := e.Retrieve(query, knowledgeBaseID)
	if err != nil {
		return "", nil, err
	}
	
	// Enhance prompt with context
	enhancedPrompt := e.EnhancePrompt(query, results)
	
	return enhancedPrompt, results, nil
}