package rag

import (
	"fmt"
	"strings"
)

// ContextInjector handles injecting retrieved context into prompts
type ContextInjector struct {
	maxContextLength int
}

// NewContextInjector creates a new context injector
func NewContextInjector(maxContextLength int) *ContextInjector {
	return &ContextInjector{
		maxContextLength: maxContextLength,
	}
}

// InjectContext injects retrieved context into a prompt
func (ci *ContextInjector) InjectContext(prompt string, results []SearchResult) string {
	if len(results) == 0 {
		return prompt
	}

	// Build context from search results
	var contextBuilder strings.Builder
	contextBuilder.WriteString("Context information:\n\n")
	
	totalLength := 0
	for _, result := range results {
		docContext := fmt.Sprintf("Document: %s\nContent: %s\n\n", result.DocumentTitle, result.Content)
		
		// Check if adding this document would exceed max length
		if totalLength+len(docContext) > ci.maxContextLength {
			// Trim the document content to fit
			availableSpace := ci.maxContextLength - totalLength - len("...\n\n")
			if availableSpace > 0 && len(docContext) > availableSpace {
				docContext = docContext[:availableSpace] + "...\n\n"
			} else {
				break // Skip remaining documents
			}
		}
		
		contextBuilder.WriteString(docContext)
		totalLength += len(docContext)
		
		// Stop if we've reached the maximum context length
		if totalLength >= ci.maxContextLength {
			break
		}
	}
	
	// If we have context, inject it into the prompt
	if contextBuilder.Len() > 0 {
		return fmt.Sprintf("%s\n%sPlease use the context information above to answer the following question:\n%s", 
			contextBuilder.String(), 
			strings.Repeat("-", 50), 
			prompt)
	}
	
	return prompt
}