package dto

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Message data object
type Message struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestMessage struct {
	Message string `json:"message"`
} 
