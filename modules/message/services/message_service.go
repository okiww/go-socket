package services

import (
	"github.com/jinzhu/copier"
	"go-socket/modules/message/dto"
	"go-socket/modules/message/models"
	"go-socket/modules/message/repositories"
)

type MessageService struct {
	messageRepo repositories.MessageInterface
}

func NewMessageService(storage *repositories.MessageStorage) *MessageService {
	return &MessageService{
		messageRepo: repositories.NewMessageRepositoryHandler(storage),
	}
}

// GetAllMessage service for get all message
func (svc MessageService) GetAllMessage() ([]dto.Message, error) {
	model, err := svc.messageRepo.GetAll()
	var result []dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return result, err
}

// Submit service for get all message
func (svc MessageService) Submit(message string) (*dto.Message, error) {
	model := models.NewMessage(message)
	err := svc.messageRepo.Save(*model)
	var result dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, err
}
