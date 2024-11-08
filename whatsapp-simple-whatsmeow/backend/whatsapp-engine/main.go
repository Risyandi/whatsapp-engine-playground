package main

import (
	"context"
	"engine/database"
	"engine/library"
	"engine/service"
	"engine/web"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("whatsapp engine running at port 5002")

	mongodb := initMongoDB()
	whatsappLibrary := initWhatsAppLibrary(mongodb)

	defer mongodb.Client.Disconnect(context.TODO())
	initNumber(whatsappLibrary)

	web.InitRoute(mongodb, whatsappLibrary)
}

func initMongoDB() *database.MongoDB {
	MONGODB_DATABASE := os.Getenv("MONGODB_DATABASE")
	MONGODB_URL := os.Getenv("MONGODB_URL")

	var mongodb = database.MongoDB{
		URL:          MONGODB_URL,
		DatabaseName: MONGODB_DATABASE,
	}

	mongodb.Connect()
	return &mongodb
}

func initWhatsAppLibrary(mongodb *database.MongoDB) *library.WhatsAppLibrary {
	numberService := &service.NumberService{Database: mongodb}

	return &library.WhatsAppLibrary{
		NumberService: numberService,
	}
}

func initNumber(whatsappLibrary *library.WhatsAppLibrary) {
	ID_VM := os.Getenv("ID_VM")

	numbers := whatsappLibrary.NumberService.GetNumberByVirtualMachineId(ID_VM)
	fmt.Println("found numbers in", ID_VM, "total", len(numbers))

	for _, number := range numbers {
		go whatsappLibrary.Authorize(number.ID, number.JID)
	}
}
