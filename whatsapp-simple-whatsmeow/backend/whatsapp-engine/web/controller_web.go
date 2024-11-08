package web

import (
	"engine/library"
	"engine/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	NumberService   *service.NumberService
	WhatsAppLibrary *library.WhatsAppLibrary
}

type BodySendMessage struct {
	Receiver string `validate:"required" json:"receiver"`
	Text     string `validate:"required" json:"text"`
}

type ResponseSendMessage struct {
	MessageID string `json:"messageId"`
}

func (controller *Controller) authorize(context *gin.Context) {
	var id = context.Param("id")

	number := controller.NumberService.GetOneNumberById(id)
	if number == nil {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errorMessage": "404__number not found"})
		return
	}

	go controller.WhatsAppLibrary.Authorize(number.ID, number.JID)

	context.IndentedJSON(http.StatusOK, gin.H(nil))
}

func (controller *Controller) removeAuthorize(context *gin.Context) {
	var id = context.Param("id")

	number := controller.NumberService.GetOneNumberById(id)
	if number == nil {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errorMessage": "404__number not found"})
		return
	}

	controller.WhatsAppLibrary.RemoveAuthorize(number.ID, number.JID)

	context.IndentedJSON(http.StatusOK, gin.H(nil))
}

func (controller *Controller) sendMessage(context *gin.Context) {
	var id = context.Param("id")

	var requestBody BodySendMessage
	if err := context.BindJSON(&requestBody); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"errorMessage": "error bind json from request body"})
		return
	}

	number := controller.NumberService.GetOneNumberById(id)
	if number == nil {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errorMessage": "number not found"})
		return
	}

	messageId, err := controller.WhatsAppLibrary.SendMessage(number, requestBody.Receiver, requestBody.Text)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, ResponseSendMessage{MessageID: messageId})
}
