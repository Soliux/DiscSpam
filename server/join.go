package server

import (
	"Raid-Client/cloudflare"
	"Raid-Client/constants"
	"Raid-Client/tools"
	"Raid-Client/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/patrickmn/go-cache"
)

var C *cache.Cache

func JoinServer(inviteCode string, token string) error {
	defer handlePanic()
	code := ""
	if strings.Contains(inviteCode, "https://discord") {
		j := strings.Split(inviteCode, "/")
		code = j[3]
	} else {
		code = inviteCode
	}
	payload := map[string]string{"": ""}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)
	request, err := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/v9/invites/%s", code), payloadBuf)
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

	body, err := ioutil.ReadAll(res.Body)
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) || os.IsTimeout(err) {
		fmt.Printf("%s %s\n", constants.Yellow(token), constants.Red("[!] Timed out"))
		return err
	}

	switch res.StatusCode {
	case 200:
		var responseJson map[string]interface{}
		if err := json.Unmarshal(body, &responseJson); err != nil {
			return err
		}
		ParseServerID := responseJson["guild"].(map[string]interface{})
		ServerID := ParseServerID["id"].(string)
		ServerName := ParseServerID["name"].(string)
		C.Set("JoinServerID", ServerID, cache.NoExpiration)
		utils.Logger(fmt.Sprintf("%s has successfully joined %s", token, ServerName))
		fmt.Printf("%s %s %s\n", constants.White(token), constants.Green("| Successfully Joined"), constants.White(ServerName))
	default:
		utils.Logger(fmt.Sprintf("%s was unable to join %s", token, code))
		fmt.Printf("%s %s %s\n", constants.White(token), constants.Red("| Unable To Join"), constants.White(code))
	}

	return nil
}

func handlePanic() {
	if err := recover(); err != nil {

	}
}
