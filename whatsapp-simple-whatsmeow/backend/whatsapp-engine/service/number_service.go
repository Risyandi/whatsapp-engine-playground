package service

import (
	"context"
	"engine/database"
	"engine/entity"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NumberService struct {
	Database *database.MongoDB
}

func (numberService *NumberService) GetNumberByVirtualMachineId(id string) []*entity.NumberEntity {
	var numbers []*entity.NumberEntity

	numberCollection := numberService.Database.Collection("numbers")

	objectId, _ := primitive.ObjectIDFromHex(id)
	cursor, err := numberCollection.Find(context.TODO(), bson.M{"virtualMachineId": objectId})
	if err != nil {
		fmt.Println("error in GetNumberByVirtualMachineId", err.Error())
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &numbers); err != nil {
		fmt.Println("error in GetNumberByVirtualMachineId", err.Error())
	}

	return numbers
}

func (numberService *NumberService) GetOneNumberById(id string) *entity.NumberEntity {
	var number *entity.NumberEntity

	numberCollection := numberService.Database.Collection("numbers")

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := numberCollection.FindOne(context.TODO(), bson.M{"_id": objectId})

	err := result.Decode(&number)
	if err != nil {
		return nil
	}

	return number
}

func (numberService *NumberService) UpdateStatusConnect(id, name, jid string) {
	numberCollection := numberService.Database.Collection("numbers")

	splitedColon := strings.Split(jid, ":")[0]
	splitedAt := strings.Split(splitedColon, "@")[0]

	objectId, _ := primitive.ObjectIDFromHex(id)
	numberCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectId},
		bson.M{
			"$set": bson.M{
				"isConnected": true,
				"name":        name,
				"phoneNumber": splitedAt,
				"jid":         jid,
				"qrcode":      "",
			},
		},
	)
}

func (numberService *NumberService) UpdateStatusDisconnect(id, qrcode string) {
	numberCollection := numberService.Database.Collection("numbers")

	objectId, _ := primitive.ObjectIDFromHex(id)
	numberCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectId},
		bson.M{
			"$set": bson.M{
				"isConnected": false,
				"qrcode":      qrcode,
			},
		},
	)
}
