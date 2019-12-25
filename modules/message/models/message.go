package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Model message
type Message struct {
	ID        uuid.UUID
	Message   string
	CreatedAt time.Time
}

// Init object of Message
func NewMessage(message string) *Message {
	return &Message{
		ID:        uuid.NewV1(),
		Message:   message,
		CreatedAt: time.Now().UTC(),
	}
}
