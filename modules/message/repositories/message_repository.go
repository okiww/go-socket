package repositories

import (
	"go-socket/modules/message/models"
)

type messageRepository struct {
	storage *MessageStorage
}

func NewMessageRepositoryHandler() MessageInterface {
	return &messageRepository{
		storage: NewMessageStorage(),
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
	repo.storage.Messages = append(repo.storage.Messages, message)
	return nil
}
