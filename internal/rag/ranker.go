package rag

import (
	"sort"
)

// ResultRanker handles ranking of search results
type ResultRanker struct {
	keywordWeight   float64
	semanticWeight  float64
	minRelevanceScore float64
}

// NewResultRanker creates a new result ranker
func NewResultRanker(keywordWeight, semanticWeight, minRelevanceScore float64) *ResultRanker {
	return &ResultRanker{
		keywordWeight:     keywordWeight,
		semanticWeight:    semanticWeight,
		minRelevanceScore: minRelevanceScore,
	}
}

// RankResults ranks and filters search results
func (rr *ResultRanker) RankResults(keywordResults, vectorResults []SearchResult) []SearchResult {
	// Combine results and calculate weighted scores
	resultMap := make(map[int]*SearchResult)
	
	// Process keyword results
	for _, result := range keywordResults {
		if result.RelevanceScore >= rr.minRelevanceScore {
			if existing, exists := resultMap[result.DocumentID]; exists {
				// Combine scores if document already exists
				existing.RelevanceScore = (existing.RelevanceScore + result.RelevanceScore*rr.keywordWeight) / 2
			} else {
				// Add new result with weighted score
				result.RelevanceScore = result.RelevanceScore * rr.keywordWeight
				resultMap[result.DocumentID] = &result
			}
		}
	}
	
	// Process vector results
	for _, result := range vectorResults {
		if result.RelevanceScore >= rr.minRelevanceScore {
			if existing, exists := resultMap[result.DocumentID]; exists {
				// Combine scores if document already exists
				existing.RelevanceScore = (existing.RelevanceScore + result.RelevanceScore*rr.semanticWeight) / 2
			} else {
				// Add new result with weighted score
				result.RelevanceScore = result.RelevanceScore * rr.semanticWeight
				resultMap[result.DocumentID] = &result
			}
		}
	}
	
	// Convert map to slice
	var results []SearchResult
	for _, result := range resultMap {
		results = append(results, *result)
	}
	
	// Sort by relevance score (descending)
	sort.Slice(results, func(i, j int) bool {
		return results[i].RelevanceScore > results[j].RelevanceScore
	})
	
	return results
}