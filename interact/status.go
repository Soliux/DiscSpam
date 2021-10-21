package interact

import (
	"Raid-Client/constants"
	"Raid-Client/utils"
	"fmt"
	"time"
)

var TOKENS []string
var updating bool

func ChangeStatus() {
	if !updating {
		time.Sleep(100 * time.Millisecond)
		for _, token := range TOKENS {
			ws := utils.SetupWebSocket(token)
			go utils.RecieveIncomingPayloads(ws, token)
			for {
				if utils.WSConnected {
					utils.SetStatus(utils.Status, ws)
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
