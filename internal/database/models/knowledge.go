package models

import (
	"time"
)

// KnowledgeBase represents a collection of documents
type KnowledgeBase struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Document represents a source document
type Document struct {
	ID             int       `json:"id" db:"id"`
	KnowledgeBaseID int       `json:"knowledge_base_id" db:"knowledge_base_id"`
	Title          string    `json:"title" db:"title"`
	Content        string    `json:"content" db:"content"`
	SourceType     string    `json:"source_type" db:"source_type"`
	SourceURL      string    `json:"source_url" db:"source_url"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// DocumentChunk represents a chunk of a document
type DocumentChunk struct {
	ID         int       `json:"id" db:"id"`
	DocumentID int       `json:"document_id" db:"document_id"`
	Content    string    `json:"content" db:"content"`
	ChunkOrder int       `json:"chunk_order" db:"chunk_order"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// ChunkEmbedding represents vector embeddings for document chunks
type ChunkEmbedding struct {
	ID        int       `json:"id" db:"id"`
	ChunkID   int       `json:"chunk_id" db:"chunk_id"`
	Embedding []byte    `json:"embedding" db:"embedding"` // Stored as binary
	ModelUsed string    `json:"model_used" db:"model_used"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}