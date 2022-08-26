package gateway

import (
	"encoding/json"
	"math"
	"strconv"
	"sync"
	"time"

	"Raid-Client/utils"

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

var (
	scrapedMemb      []Member
	ScrapedGuilds    []Guild
	wg               sync.WaitGroup
	data             = make(chan string)
	finishedFetching = make(chan bool)
)

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
		switch {
		case decodedMessage.Opcode == 10:
			data := decodedMessage.EventData.(map[string]interface{})
			heartbeatInterval := data["heartbeat_interval"].(float64)

			go setupHeartbeat(heartbeatInterval, ws)
			identify(ws, token)
		case decodedMessage.EventName == "READY":
			utils.Logger("Received READY event")
			go getGuildsData(string(p))
		case decodedMessage.EventName == "GUILD_MEMBER_LIST_UPDATE":
			data <- string(p)
			utils.Logger(string(p))
		default:
			utils.Logger("received payload:", string(p))
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

		utils.Logger("sending payload (heartbeat): ", string(b))
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

	utils.Logger("sending payload:", string(b))
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
	ws.connection, _, dialErr = websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?encoding=json&v=9", nil)
	if dialErr != nil {
		return ws.connection
	}
	return ws.connection
}

/*
	https://arandomnewaccount.gitlab.io/discord-unofficial-docs/lazy_guilds.html
*/

func SearchGuildMembers(ws *websocket.Conn, result []Member, guildID, channelID, token, method string, delay int) []Member {
	scrapedMemb = []Member{}
	result = []Member{}
	// Get membercount from READY EVENT
	memberCount, _ := strconv.Atoi(getMemberCount(guildID))
	var multiplier int
	// Get multiplier

	switch method {
	case "overlap":
		multiplier = 100
	case "no-overlap":
		multiplier = 200
	default:
		multiplier = 100
	}

	// Listen for data
	go captureData(guildID)
	wg.Add(1)
	// First lazyGuild request that is needed
	go func() {
		/*
			{"op":14,"d":{"guild_id":"632655162875445258","typing":true,"threads":true,"activities":true,"members":[],"channels":{"762838406526402591":[[0,99]]},"thread_member_lists":[]}}
		*/
		// Ugly af but works, still looking for a way to make it look better
		defer wg.Done()
		first, err := json.Marshal(DiscordGatewayFetchMembers{
			Opcode: 14,
			EventData: struct {
				GuildID           string        "json:\"guild_id\""
				Typing            bool          "json:\"typing,omitempty\""
				Threads           bool          "json:\"threads,omitempty\""
				Activities        bool          "json:\"activities,omitempty\""
				Members           []interface{} "json:\"members,omitempty\""
				Channels          interface{}   "json:\"channels\""
				ThreadMemberLists []interface{} "json:\"thread_member_lists,omitempty\""
			}{
				GuildID:    guildID,
				Typing:     true,
				Threads:    false,
				Activities: true,
				Members:    []interface{}{nil},
				Channels: map[string]interface{}{
					channelID: [][]int{{0, 99}},
				},
				ThreadMemberLists: []interface{}{nil},
			},
		})
		if err != nil {
			utils.Logger("Error: ", err)
		}
		utils.Logger("sending payload (HELLO_REQUEST_GUILD_MEMBERS): ", string(first))
		err = ws.WriteMessage(websocket.TextMessage, first)
		if err != nil {
			utils.Logger("Error: ", err)
		}
	}()
	wg.Wait()

	/*
		https://arandomnewaccount.gitlab.io/discord-unofficial-docs/lazy_guilds.html#op-14-lazy-request-what-to-send
	*/
	var z int

	if memberCount < 100 {
		z = 1
	} else {
		z = int(math.Round(float64(memberCount / 100)))
	}
Fetching:

	for i := 0; i < z+3; i++ { // +3 because in certain small guilds you will need to send more requests
		// Gotta work on high members guilds but won't on really small one
		select {
		case <-finishedFetching:
			utils.Logger("Finished Fetching Guild")
			result = scrapedMemb
			break Fetching
		default:
			wg.Add(1)
			go func(index, mul int) {
				defer wg.Done()
				ranges := getRanges(index, 200, memberCount)
				// Ugly af but works, still looking for a way to make it look better
				second, err := json.Marshal(DiscordGatewayFetchMembers{
					Opcode: 14,
					EventData: struct {
						GuildID           string        "json:\"guild_id\""
						Typing            bool          "json:\"typing,omitempty\""
						Threads           bool          "json:\"threads,omitempty\""
						Activities        bool          "json:\"activities,omitempty\""
						Members           []interface{} "json:\"members,omitempty\""
						Channels          interface{}   "json:\"channels\""
						ThreadMemberLists []interface{} "json:\"thread_member_lists,omitempty\""
					}{
						GuildID:    guildID,
						Typing:     false,
						Threads:    false,
						Activities: false,
						Members:    []interface{}{},
						Channels: map[string]interface{}{
							channelID: ranges,
						},
						ThreadMemberLists: []interface{}{},
					},
				})
				if err != nil {
					utils.Logger("Error: ", err)
				}
				utils.Logger("sending payload (REQUEST_GUILD_MEMBERS): ", string(second))
				time.Sleep(time.Millisecond * time.Duration(delay))
				err = ws.WriteMessage(websocket.TextMessage, second)
				if err != nil {
					if err == websocket.ErrCloseSent {
						utils.Logger("Error writing message: ", err)
						time.Sleep(time.Second * 1)
						result = scrapedMemb
					}
					utils.Logger("Error writing message: ", err)

				}

			}(i, multiplier)
			wg.Wait() // Can't conccurently write to websocket
		}
	}
	scrapedMemb = []Member{}
	return result
}

func captureData(guildID string) {
	// loop until channel is closed, we gotta close it after finishing fetching
	for d := range data {
		// command to exit program
		if d == "q" {
			return
		}

		go parseMemberData(d, guildID)

	}
}
