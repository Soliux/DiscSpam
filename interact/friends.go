package interact

import (
	"Raid-Client/cloudflare"
	"Raid-Client/constants"
	"Raid-Client/tools"
	"Raid-Client/utils"
	"context"
	"errors"
	"os"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func AddFriend(Token string, User string) error {
	username := strings.Split(User, "#")
	payload := map[string]string{
		"username":      username[0],
		"discriminator": username[1],
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)

	request, err := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/v9/users/%s/relationships", "%40me"), payloadBuf)
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
	case 204:
		utils.Logger(fmt.Sprintf("%s has successfully added %s", Token, User))
		fmt.Printf("%s %s %s\n", constants.Red(Token), constants.Green("Success:"), constants.White(fmt.Sprintf("Sent friend request to %s", User)))
	default:
		utils.Logger(fmt.Sprintf("%s was unable to add %s", Token, User))
		fmt.Printf("%s %s\n", constants.White(Token), constants.Red(fmt.Sprintf("Error: Unable to send friend request to %s", User)))
	}

	return nil
}

func RemoveFriend(Token string, UserID string) error {
	defer handlePanic()
	request, err := http.NewRequest("DELETE", fmt.Sprintf("https://discord.com/api/v9/users/%s/relationships/%s", "%40me", UserID), nil)
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
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) || os.IsTimeout(err) {
		fmt.Printf("%s %s\n", constants.Yellow(Token), constants.Red("[!] Timed out"))
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 204:
		utils.Logger(fmt.Sprintf("%s has removed %s", Token, UserID))
		fmt.Printf("%s %s %s\n", constants.Red(Token), constants.Green("Success:"), constants.White(fmt.Sprintf("Unfriended user id: %s", UserID)))
	default:
		utils.Logger(fmt.Sprintf("%s was unable to remove %s", Token, UserID))
		fmt.Printf("%s %s\n", constants.White(Token), constants.Red(fmt.Sprintf("Error: Unable to unfriend user id: %s", UserID)))
	}

	return nil
}

func handlePanic() {
	if err := recover(); err != nil {
	}
}
