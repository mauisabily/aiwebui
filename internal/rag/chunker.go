package rag

import (
	"strings"
)

// DocumentChunker handles splitting documents into chunks
type DocumentChunker struct {
	chunkSize    int
	chunkOverlap int
}

// NewDocumentChunker creates a new document chunker
func NewDocumentChunker(chunkSize, chunkOverlap int) *DocumentChunker {
	return &DocumentChunker{
		chunkSize:    chunkSize,
		chunkOverlap: chunkOverlap,
	}
}

// ChunkDocument splits a document into overlapping chunks
func (dc *DocumentChunker) ChunkDocument(content string) []string {
	if len(content) <= dc.chunkSize {
		return []string{content}
	}

	var chunks []string
	runes := []rune(content)
	
	for i := 0; i < len(runes); i += (dc.chunkSize - dc.chunkOverlap) {
		end := i + dc.chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		
		chunk := string(runes[i:end])
		chunks = append(chunks, strings.TrimSpace(chunk))
		
		// If we've reached the end, break
		if end == len(runes) {
			break
		}
	}
	
	return chunks
}