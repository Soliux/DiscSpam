package server

import (
	"Raid-Client/cloudflare"
	"Raid-Client/constants"
	"Raid-Client/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ChangeNickname(ServerID string, Token string, Nickname string) error {
	payload := map[string]string{
		"nick": Nickname,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)

	request, err := http.NewRequest("PATCH", fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", ServerID, "%40me"), payloadBuf)
	if err != nil {
		return err
	}

	cf := cloudflare.Cookie()
	xprop := utils.XSuperProperties()

	request.Header = http.Header{
		"Accept":             []string{"*/*"},
		"Accept-language":    []string{"en-GB"},
		"Authorization":      []string{Token},
		"Content-length":     []string{"2"},
		"Content-type":       []string{"application/json"},
		"Cookie":             []string{cf},
		"Origin":             []string{"https://discord.com"},
		"Referrer":           []string{fmt.Sprintf("https://discord.com/channels/%s", ServerID)},
		"Sec-fetch-dest":     []string{"empty"},
		"Sec-fetch-mode":     []string{"cors"},
		"Sec-fetch-site":     []string{"same-origin"},
		"User-agent":         []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.15 Chrome/83.0.4103.122 Electron/9.3.5 Safari/537.36"},
		"X-debug-options":    []string{"bugReporterEnabled"},
		"X-super-properties": []string{xprop},
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		utils.Logger(fmt.Sprintf("%s has successfully updated their nickname to %s", Token, Nickname))
		fmt.Printf("%s %s %s\n", constants.Red(Token), constants.Green("Success:"), constants.White(fmt.Sprintf("Changed nickname to %s", Nickname)))
	default:
		utils.Logger(fmt.Sprintf("%s was unable to update their nickname to %s", Token, Nickname))
		fmt.Printf("%s %s\n", constants.White(Token), constants.Red(fmt.Sprintf("Error: Unable to change nickname to %s", Nickname)))
	}

	return nil
}
