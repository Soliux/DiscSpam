package utils

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

var WSConnected bool
var Looping bool
var Presence []Activity
var Status string

type Connection struct {
	connection *websocket.Conn
}

type ActivityType int

const (
	ActivityGame      = 0
	ActivityListening = 2
	ActivityWatching  = 3
)

type Activity struct {
	Name string       `json:"name"`
	Type ActivityType `json:"type"`
}

type DiscordGatewayPayload struct {
	Opcode    int         `json:"op"`
	EventData interface{} `json:"d"`
	EventName string      `json:"t,omitempty"`
}

type DiscordGatewayEventDataIdentify struct {
	Token      string                 `json:"token"`
	Intents    int                    `json:"intents"`
	Properties map[string]interface{} `json:"properties"`
}

type DiscordGatewayEventDataUpdateStatus struct {
	TimeSinceIdle *int       `json:"since"`
	Activities    []Activity `json:"activities,omitempty"`
	Status        string     `json:"status"`
	IsAfk         bool       `json:"afk"`
}

func RecieveIncomingPayloads(ws *websocket.Conn, token string) error {
	for {
		_, p, readErr := ws.ReadMessage()
		if readErr != nil {
			return readErr
		}

		var decodedMessage DiscordGatewayPayload
		decodeErr := json.Unmarshal(p, &decodedMessage)
		if decodeErr != nil {
			return decodeErr
		}
		//log.Println("recieved payload: ", decodedMessage)

		if decodedMessage.Opcode == 10 {
			data := decodedMessage.EventData.(map[string]interface{})
			heartbeatInterval := data["heartbeat_interval"].(float64)

			go setupHeartbeat(heartbeatInterval, ws)
			identify(ws, token)
		}
	}
}

func setupHeartbeat(interval float64, ws *websocket.Conn) error {
	c := time.Tick(time.Duration(interval) * time.Millisecond)
	for range c {
		b, marshalErr := json.Marshal(DiscordGatewayPayload{1, nil, ""})
		if marshalErr != nil {
			return marshalErr
		}

		//log.Println("sending payload (heartbeat): ", string(b))
		err := ws.WriteMessage(websocket.TextMessage, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func identify(ws *websocket.Conn, token string) error {
	b, marshalErr := json.Marshal(DiscordGatewayPayload{2,
		DiscordGatewayEventDataIdentify{token, 0, map[string]interface{}{
			"$os":      "Windows",
			"$browser": "Chrome",
			"$device":  "",
		}}, ""})
	if marshalErr != nil {
		return marshalErr
	}

	err := ws.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	WSConnected = true
	return nil
}

func SetStatus(status string, ws *websocket.Conn) error {
	b, marshalErr := json.Marshal(DiscordGatewayPayload{3,
		DiscordGatewayEventDataUpdateStatus{
			nil,
			Presence,
			status,
			false,
		}, ""})
	if marshalErr != nil {
		return marshalErr
	}

	//log.Println("sending payload:", string(b))
	err := ws.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}
	return nil
}

func SetupWebSocket(Token string) *websocket.Conn {
	WSConnected = false
	ws := Connection{
		connection: &websocket.Conn{},
	}

	var dialErr error
	ws.connection, _, dialErr = websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=9&encoding=json", nil)
	if dialErr != nil {
		return ws.connection
	}
	return ws.connection
}
