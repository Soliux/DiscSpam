package server

import (
	"Raid-Client/cloudflare"
	"Raid-Client/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LeaveServer(serverID string, token string) error {
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

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(request)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		fmt.Printf("%s %s %s", white(token), green("| Successfully Left"), white(serverID))
	case 404:
		fmt.Printf("%s %s %s", white(token), red("| Cannot Leave"), white(serverID))
	default:
		fmt.Printf("%s %s %s", white(token), red("| Cannot Leave"), white(serverID))
	}

	return nil
}
