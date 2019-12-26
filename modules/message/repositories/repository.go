package repositories

import (
	"go-socket/modules/message/models"
)

// MessageInterface contract for repository
type MessageInterface interface {
	GetAll() ([]models.Message, error)
	Save(message models.Message) error
}
