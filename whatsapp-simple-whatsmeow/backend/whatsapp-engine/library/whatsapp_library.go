package library

import (
	"context"
	"encoding/base64"
	"engine/entity"
	"engine/service"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"

	_ "github.com/mattn/go-sqlite3"
	goQRCode "github.com/skip2/go-qrcode"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type Session struct {
	ID               string
	JID              string
	Client           *whatsmeow.Client
	IsGenerateQRCode bool
	IsConnected      bool
}

var sessions []Session
var container, _ = sqlstore.New("sqlite3", "file:sessions.db?_foreign_keys=on", waLog.Stdout("Database", "DEBUG", true))

type WhatsAppLibrary struct {
	NumberService *service.NumberService
}

func (whatsappLibrary *WhatsAppLibrary) Authorize(id, jid string) {
	store.DeviceProps.Os = proto.String("wa-message-example")

	_, session := findSession(id)
	if session != nil && session.ID != "" {
		if session.IsGenerateQRCode {
			return
		}

		if session.IsConnected {
			return
		}

		fmt.Println("authorize number", id, time.Now())
	}

	deviceStore := getDeviceStore(id, jid)
	client := whatsmeow.NewClient(deviceStore, waLog.Stdout("Client", "ERROR", true))

	sessions = append(sessions, Session{
		ID:               id,
		JID:              jid,
		Client:           client,
		IsGenerateQRCode: true,
		IsConnected:      false,
	})

	client.AddEventHandler(func(evt interface{}) {
		switch v := evt.(type) {

		case *events.Message:
			fmt.Println("incoming new message", v.Info.ID)

		case *events.Connected:
			fmt.Println("number connected", id)

			currentJID := client.Store.ID.String()
			name := client.Store.PushName

			if name == "" {
				name = client.Store.BusinessName
			}

			whatsappLibrary.NumberService.UpdateStatusConnect(id, name, currentJID)

			index, session := findSession(id)
			if session != nil && session.ID != "" {
				sessions[index].JID = currentJID
				sessions[index].IsGenerateQRCode = false
				sessions[index].IsConnected = true
			}

		case *events.Disconnected:
			fmt.Println("number disconnected", id)
			client.Disconnect()
			whatsappLibrary.NumberService.UpdateStatusDisconnect(id, "")
			whatsappLibrary.Authorize(id, jid)

		case *events.LoggedOut:
			fmt.Println("number logged out", id)
			client.Disconnect()
			whatsappLibrary.NumberService.UpdateStatusDisconnect(id, "")
			whatsappLibrary.Authorize(id, jid)
		}
	})

	if client.Store.ID == nil {
		qrcode, _ := client.GetQRChannel(context.Background())
		client.Connect()

		for event := range qrcode {
			if event.Event == "code" {
				var filepath = "media/session_" + id + ".jpg"
				err := goQRCode.WriteFile(event.Code, goQRCode.Medium, 256, filepath)
				if err != nil {
					fmt.Println("ERROR WA_LIBRARY goQRCode.WriteFile", err.Error())
				}

				bytes, err := ioutil.ReadFile(filepath)
				if err != nil {
					fmt.Println("ERROR WA_LIBRARY ioutil.ReadFile", err.Error())
				}

				fmt.Println("number generating qrcode", id)
				whatsappLibrary.NumberService.UpdateStatusDisconnect(id, base64.StdEncoding.EncodeToString(bytes))
			} else if event.Event == "success" {
				fmt.Println("number connected qrcode", id)

				currentJID := client.Store.ID.String()
				name := client.Store.PushName

				if name == "" {
					name = client.Store.BusinessName
				}

				whatsappLibrary.NumberService.UpdateStatusConnect(id, name, currentJID)

				index, session := findSession(id)
				if session != nil && session.ID != "" {
					sessions[index].JID = currentJID
					sessions[index].IsGenerateQRCode = false
					sessions[index].IsConnected = true
				}
			} else {
				fmt.Println("number timeout generating qrcode", id)
				client.Disconnect()
				whatsappLibrary.NumberService.UpdateStatusDisconnect(id, "")

				index, session := findSession(id)
				if session != nil && session.ID != "" {
					sessions[index].IsGenerateQRCode = false
					sessions[index].IsConnected = false
				}
			}
		}
	} else {
		client.Connect()

		index, session := findSession(id)
		if session != nil && session.ID != "" {
			sessions[index].IsGenerateQRCode = false
		}
	}
}

func (whatsappLibrary *WhatsAppLibrary) RemoveAuthorize(id, jid string) {
	index, session := findSession(id)
	if session != nil && session.ID != "" {
		sessions = append(sessions[:index], sessions[index+1:]...)
		session.Client.Disconnect()

		parsedJID, err := types.ParseJID(jid)
		if err == nil {
			deviceStore, err := container.GetDevice(parsedJID)
			if err == nil && deviceStore != nil {
				container.DeleteDevice(deviceStore)
			}
		}
	}
}

func (whatsappLibrary *WhatsAppLibrary) SendMessage(number *entity.NumberEntity, receiver, text string) (string, error) {
	_, session := findSession(number.ID)
	if session == nil {
		return "", errors.New("session not found")
	}

	client := session.Client

	receiverJID, err := types.ParseJID(receiver + "@s.whatsapp.net")
	if err != nil {
		return "", errors.New("invalid receiver with phone number: " + receiver)
	}

	var protoMessage = &waE2E.Message{
		Conversation: proto.String(text),
	}

	response, err := client.SendMessage(context.Background(), receiverJID, protoMessage)
	if err != nil {
		return "", err
	}

	return response.ID, nil
}

func findSession(id string) (int, *Session) {
	var currentSession *Session
	var currentIndex int

	for index, session := range sessions {
		if session.ID == id {
			currentSession = &sessions[index]
			currentIndex = index
		}
	}

	return currentIndex, currentSession
}

func getDeviceStore(id, jid string) *store.Device {
	currentJID := jid
	_, session := findSession(id)

	if session != nil && session.JID != "" {
		currentJID = session.JID
	}

	parsedJID, err := types.ParseJID(currentJID)
	if err != nil {
		return container.NewDevice()
	}

	deviceStore, err := container.GetDevice(parsedJID)
	if err != nil || deviceStore == nil {
		return container.NewDevice()
	}

	return deviceStore
}
