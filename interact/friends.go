package interact

import (
	"Raid-Client/cloudflare"
	"Raid-Client/utils"
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
	case 400, 403:
		fmt.Printf("%s %s\n", white(Token), red(fmt.Sprintf("Error: Unable to send friend request to %s", User)))
	case 204:
		fmt.Printf("%s %s %s\n", red(Token), green("Success:"), white(fmt.Sprintf("Sent friend request to %s", User)))
	default:
		fmt.Printf("%s %s\n", white(Token), red(fmt.Sprintf("Error: Unable to send friend request to %s", User)))
	}

	return nil
}

func RemoveFriend(Token string, UserID string) error {
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

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 400, 403:
		fmt.Printf("%s %s\n", white(Token), red(fmt.Sprintf("Error: Unable to unfriend user id: %s", UserID)))
	case 204:
		fmt.Printf("%s %s %s\n", red(Token), green("Success:"), white(fmt.Sprintf("Unfriended user id: %s", UserID)))
	default:
		fmt.Printf("%s %s\n", white(Token), red(fmt.Sprintf("Error: Unable to unfriend user id: %s", UserID)))
	}

	return nil
}
