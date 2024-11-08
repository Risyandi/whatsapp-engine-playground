package entity

import "time"

type NumberEntity struct {
	ID               string    `json:"_id" bson:"_id"`
	IsConnected      bool      `json:"isConnected" bson:"isConnected"`
	Name             string    `json:"name" bson:"name"`
	PhoneNumber      string    `json:"phoneNumber" bson:"phoneNumber"`
	JID              string    `json:"jid" bson:"jid"`
	QRCode           string    `json:"qrcode" bson:"qrcode"`
	VirtualMachineID string    `json:"virtualMachineId" bson:"virtualMachineId"`
	CreatedAt        time.Time `json:"createdAt" bson:"createdAt"`
}
