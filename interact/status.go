package interact

import (
	"Raid-Client/utils"
	"time"
)

var tokens []string
var updating bool

/*
Instead of having multiple websocket connections,
we setup one and send our presence message and close
the connection. Our status stays for around 2 mins, meaning
we can keep connecting and sending a message every 100 secs
*/

// Instantly update status instead of waiting for loop
func ChangeStatus(Tokens []string) {
	tokens = Tokens

	if !updating { 
		for _, tkn := range tokens {
			time.Sleep(100 * time.Millisecond)
			ws := utils.SetupWebSocket(tkn)
			go utils.RecieveIncomingPayloads(ws, tkn)

			for {
				if utils.WSConnected {
					utils.SetStatus(utils.Status, ws)
					ws.Close()
					break
				}
			}
		}
	}

	if !utils.Looping {
		utils.Looping = true
		go loopMessage()
	}
}

// Loop presence message every 60s
func loopMessage() {
	for {
		time.Sleep(100 * time.Second)
		updating = true
		for _, tkn := range tokens {
			time.Sleep(100 * time.Millisecond)
			ws := utils.SetupWebSocket(tkn)
			go utils.RecieveIncomingPayloads(ws, tkn)

			for {
				if utils.WSConnected {
					utils.SetStatus(utils.Status, ws)
					ws.Close()
					break
				}
			}
		}
		updating = false
	}
}
