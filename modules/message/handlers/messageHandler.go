package handlers

import (
	"github.com/gin-gonic/gin"
	"go-socket/modules/message/dto"
	"go-socket/modules/message/repositories"
	"go-socket/modules/message/services"
	"net/http"
)

type MessageHandler struct {
	messageService *services.MessageService
	storage *repositories.MessageStorage
}

func NewMessageHandler(storage *repositories.MessageStorage) *MessageHandler {
	return &MessageHandler{
		messageService: services.NewMessageService(),
		storage:        storage,
	}
}

func (ctrl MessageHandler) GetAllMessages(ctx *gin.Context)  {
	result, err := ctrl.messageService.GetAllMessage()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (ctrl MessageHandler) SubmitMessage(ctx *gin.Context)  {
	var request dto.RequestMessage
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
	}

	result, err := ctrl.messageService.Submit(request.Message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}