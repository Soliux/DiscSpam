package server

import (
	"Raid-Client/cloudflare"
	"Raid-Client/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var responseJson map[string]interface{}
	if err := json.Unmarshal(body, &responseJson); err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		fmt.Printf("%s %s %s\n", red(Token), green("Success:"), white(fmt.Sprintf("Changed nickname to %s", Nickname)))
	default:
		fmt.Printf("%s %s\n", white(Token), red(fmt.Sprintf("Error: Unable to change nickname to %s", Nickname)))
	}

	return nil
}
