package repositories

import (
	"errors"
	"go-socket/modules/message/models"
)

type messageRepository struct {
	storage *MessageStorage
}

func NewMessageRepositoryHandler(storage *MessageStorage) MessageInterface {
	return &messageRepository{
		storage: storage,
	}
}

func (repo messageRepository) GetAll() ([]models.Message, error) {
	messages := []models.Message{}
	for _, item := range repo.storage.Messages {
		messages = append(messages, item)
	}

	return messages, nil
}

func (repo messageRepository) Save(message models.Message) error {
	if message.Message == "" {
		return errors.New("message data not empty string")
	}
	repo.storage.Messages = append(repo.storage.Messages, message)
	return nil
}
