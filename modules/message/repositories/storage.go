package repositories

import "go-socket/modules/message/models"

type MessageStorage struct {
	Messages []models.Message
}

func NewMessageStorage() *MessageStorage {
	return &MessageStorage{
		Messages: []models.Message{},
	}
}