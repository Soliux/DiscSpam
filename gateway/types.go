package gateway

import (
	"time"

	"github.com/tidwall/gjson"
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
type DiscordGatewayFetchMembers struct {
	Opcode    int `json:"op"`
	EventData struct {
		GuildID           string        `json:"guild_id"`
		Typing            bool          `json:"typing,omitempty"`
		Threads           bool          `json:"threads,omitempty"`
		Activities        bool          `json:"activities,omitempty"`
		Members           []interface{} `json:"members,omitempty"`
		Channels          interface{}   `json:"channels"`
		ThreadMemberLists []interface{} `json:"thread_member_lists,omitempty"`
	} `json:"d"`
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

type MemberData struct {
	OnlineCount  string
	MemberCount  string
	ID           string
	GuildID      string
	HoistedRoles gjson.Result
	Types        []string
	Locations    gjson.Result
	Updates      gjson.Result
}

type Guild struct {
	ID          string
	MemberCount string
}

type Member struct {
	User struct {
		Username      string `json:"username"`
		PublicFlags   int    `json:"public_flags"`
		ID            string `json:"id"`
		Discriminator string `json:"discriminator"`
		Avatar        string `json:"avatar"`
	} `json:"user"`
	Roles    []interface{} `json:"roles"`
	Presence struct {
		User struct {
			ID string `json:"id"`
		} `json:"user"`
		Status       string `json:"status"`
		ClientStatus struct {
			Desktop string `json:"desktop"`
			Web     string `json:"web"`
		} `json:"client_status"`
		Activities []interface{} `json:"activities"`
	} `json:"presence"`
	Mute        bool        `json:"mute"`
	JoinedAt    time.Time   `json:"joined_at"`
	HoistedRole interface{} `json:"hoisted_role"`
	Deaf        bool        `json:"deaf"`
}
