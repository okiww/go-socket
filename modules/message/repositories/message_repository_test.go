package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"go-socket/modules/message/models"
	"testing"
	"time"
)

func TestMessageRepository_GetAll(t *testing.T) {
	t.Run("test get all message", func(t *testing.T) {
		message := "Test message"
		id := uuid.NewV1()
		createdAt := time.Now().UTC()
		storage := NewMessageStorage()

		storage.Messages = []models.Message{
			models.Message{
				ID:        id,
				Message:   message,
				CreatedAt: createdAt,
			},
		}
		repo := NewMessageRepositoryHandler(storage)
		modelRepo, err := repo.GetAll()
		assert.Nil(t, err)

		// assert model repo
		for _, item := range modelRepo {
			assert.Equal(t, item.ID, id)
			assert.Equal(t, item.Message, message)
			assert.Equal(t, item.CreatedAt, createdAt)
		}
	})
}

func TestMessageRepository_Save(t *testing.T) {
	t.Run("Test save message", func(t *testing.T) {
		message := "Test message"
		id := uuid.NewV1()
		createdAt := time.Now().UTC()
		storage := NewMessageStorage()

		model := models.Message{
			ID:        id,
			Message:   message,
			CreatedAt: createdAt,
		}

		repo := NewMessageRepositoryHandler(storage)
		err := repo.Save(model)

		assert.Nil(t, err)
	})

	t.Run("Test message empty string", func(t *testing.T) {
		message := ""
		id := uuid.NewV1()
		createdAt := time.Now().UTC()
		storage := NewMessageStorage()

		model := models.Message{
			ID:        id,
			Message:   message,
			CreatedAt: createdAt,
		}

		repo := NewMessageRepositoryHandler(storage)
		err := repo.Save(model)

		assert.NotNil(t, err)
	})
}
