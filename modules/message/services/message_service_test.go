package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	mockRepository "go-socket/gen/mocks"
	"go-socket/modules/message/models"
	"testing"
	"time"
)

func TestMessageService_GetAllMessage(t *testing.T) {
	t.Run("Test get all message", func(t *testing.T) {
		//init gomock
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mockRepository.NewMockMessageInterface(mockCtrl)
		var expected []models.Message

		message := models.Message{
			ID:        uuid.NewV1(),
			Message:   "test",
			CreatedAt: time.Now().UTC(),
		}

		expected = append(expected, message)
		errMock := errors.New("Mock Error")
		mock.EXPECT().GetAll().Return(expected, errMock).Times(1)

		svc := MessageService{
			messageRepo: mock,
		}

		res, _ := svc.GetAllMessage()
		assert.Equal(t, expected, res)
		assert.NotNil(t, expected)
	})
}

func TestMessageService_Submit(t *testing.T) {
	t.Run("Test submit message", func(t *testing.T) {
		//init gomock
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mockRepository.NewMockMessageInterface(mockCtrl)
		message := "Testing"
		model := models.NewMessage(message)
		mock.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		svc := MessageService{
			messageRepo: mock,
		}
		res, err := svc.Submit(message)
		assert.Equal(t, model.Message, res.Message)
		assert.Nil(t, err)
	})
}
