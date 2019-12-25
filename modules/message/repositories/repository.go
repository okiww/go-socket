package repositories

import (
	"go-socket/modules/message/models"
)

type MessageInterface interface {
	GetAll() ([]models.Message, error)
	Save(message models.Message) error
}
