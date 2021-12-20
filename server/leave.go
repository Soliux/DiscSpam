package server

import (
	"Raid-Client/cloudflare"
	"Raid-Client/constants"
	"Raid-Client/tools"
	"Raid-Client/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func LeaveServer(serverID string, token string) error {
	defer handlePanic()
	payload := map[string]string{"lurking": "false"}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("https://discord.com/api/v9/users/@me/guilds/%s", serverID), payloadBuf)
	if err != nil {
		return err
	}
	cf := cloudflare.Cookie()
	xprop := utils.XSuperProperties()
	request.Header = http.Header{
		"Accept":             []string{"*/*"},
		"Accept-language":    []string{"en-GB"},
		"Authorization":      []string{token},
		"Content-length":     []string{"2"},
		"Content-type":       []string{"application/json"},
		"Cookie":             []string{cf},
		"Origin":             []string{"https://discord.com"},
		"Sec-fetch-dest":     []string{"empty"},
		"Sec-fetch-mode":     []string{"cors"},
		"Sec-fetch-site":     []string{"same-origin"},
		"User-agent":         []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.15 Chrome/83.0.4103.122 Electron/9.3.5 Safari/537.36"},
		"X-debug-options":    []string{"bugReporterEnabled"},
		"X-super-properties": []string{xprop},
	}

	client := tools.CreateHttpClient()
	defer client.CloseIdleConnections()

	res, err := client.Do(request)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		utils.Logger(fmt.Sprintf("%s has successfully left %s", token, serverID))
		fmt.Printf("%s %s %s\n", constants.White(token), constants.Green("| Successfully Left"), constants.White(serverID))
	case 200:
		utils.Logger(fmt.Sprintf("%s has successfully left %s", token, serverID))
		fmt.Printf("%s %s %s\n", constants.White(token), constants.Green("| Successfully Left"), constants.White(serverID))
	default:
		utils.Logger(fmt.Sprintf("%s was unable to leave %s", token, serverID))
		fmt.Printf("%s %s %s\n", constants.White(token), constants.Red("| Cannot Leave"), constants.White(serverID))
	}

	return nil
}
