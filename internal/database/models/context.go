package models

import (
	"time"
)

// ConversationContext represents additional context for conversations
type ConversationContext struct {
	ID             int       `json:"id" db:"id"`
	ConversationID int       `json:"conversation_id" db:"conversation_id"`
	ContextType    string    `json:"context_type" db:"context_type"`
	ContextData    string    `json:"context_data" db:"context_data"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}