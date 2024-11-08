package web

import (
	"engine/database"
	"engine/library"
	"engine/service"

	"github.com/gin-gonic/gin"
)

func InitRoute(mongodb *database.MongoDB, whatsappLibrary *library.WhatsAppLibrary) {
	router := gin.Default()

	controller := Controller{
		NumberService:   &service.NumberService{Database: mongodb},
		WhatsAppLibrary: whatsappLibrary,
	}

	router.POST("/authorize/:id", controller.authorize)
	router.DELETE("/authorize/:id", controller.removeAuthorize)
	router.POST("/message/:id", controller.sendMessage)

	router.Run("0.0.0.0:5002")
}
