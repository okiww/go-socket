package services

import (
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"go-socket/modules/message/dto"
	"go-socket/modules/message/models"
	"go-socket/modules/message/repositories"
)

type MessageService struct {
	messageRepo repositories.MessageInterface
}

func NewMessageService() *MessageService {
	return &MessageService{
		messageRepo: repositories.NewMessageRepositoryHandler(),
	}
}

func (svc MessageService) GetMessageByID(id uuid.UUID) (*dto.Message, error) {
	model, err := svc.messageRepo.Get(id)
	var result dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, err
}

func (svc MessageService) GetAllMessage() ([]dto.Message, error) {
	model, err := svc.messageRepo.GetAll()
	var result []dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return result, err
}

func (svc MessageService) Submit(message string) (*dto.Message, error) {
	model := models.NewMessage(message)
	err := svc.messageRepo.Save(*model)
	var result dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, err
}
