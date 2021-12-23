package interact

import (
	"Raid-Client/constants"
	"Raid-Client/gateway"
	"Raid-Client/utils"
	"fmt"
	"time"
)

var TOKENS []string
var updating bool

func ChangeStatus() {
	defer handlePanic()
	if !updating {
		time.Sleep(100 * time.Millisecond)
		for _, token := range TOKENS {
			ws := gateway.SetupWebSocket(token)
			go gateway.RecieveIncomingPayloads(ws, token)
			for {
				if gateway.WSConnected {
					gateway.SetStatus(gateway.Status, ws)
					utils.Logger(fmt.Sprintf("%s has updated their status", token))
					fmt.Printf("%s %s\n", constants.White(token), constants.Green("| Successfully set the status"))
					ws.Close()
					break
				}
			}
		}
	}
}

// Loop presence message every 60s
func loopMessage() {
	for {
		time.Sleep(100 * time.Second)
		updating = true
		for _, tkn := range TOKENS {
			time.Sleep(100 * time.Millisecond)
			ws := gateway.SetupWebSocket(tkn)
			go gateway.RecieveIncomingPayloads(ws, tkn)

			for {
				if gateway.WSConnected {
					gateway.SetStatus(gateway.Status, ws)
					ws.Close()
					break
				}
			}
		}
		updating = false
	}
}
