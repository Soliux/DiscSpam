package interact

import (
	"Raid-Client/cloudflare"
	"Raid-Client/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gookit/color"
)

var BadCount int
var green = color.FgGreen.Render
var white = color.FgWhite.Render
var red = color.FgRed.Render
var yellow = color.FgYellow.Render

func SendMessage(ServerID string, ChannelID string, Token string, Content string) error {
	if BadCount >= 15 {
		return errors.New("auto anti token lock feature triggered")
	} else {
		payload := map[string]string{
			"content": Content,
		}

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(payload)

		request, err := http.NewRequest("POST", fmt.Sprintf("https://discordapp.com/api/v9/channels/%s/messages", ChannelID), payloadBuf)
		if err != nil {
			return err
		}

		cf := cloudflare.Cookie()
		xprop := utils.XSuperProperties()

		request.Header = http.Header{
			"Accept":             []string{"*/*"},
			"Accept-language":    []string{"en-GB"},
			"Authorization":      []string{Token},
			"Content-type":       []string{"application/json"},
			"Cookie":             []string{cf},
			"Origin":             []string{"https://discord.com"},
			"Referrer":           []string{fmt.Sprintf("https://discord.com/channels/%s/%s", ServerID, ChannelID)},
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

		// Do some parsing of the response to check for issues.
		switch responseJson["message"] {
		case "Missing Access":
			fmt.Printf("%s %s\n", white(Token), red("Unable to send message make sure it is the server"))
			BadCount++
		case "Missing Permissions":
			fmt.Printf("%s %s %s\n", red(Token), yellow("is missing permissions I.E needs a role to message in"), white(ServerID))
			BadCount++
		case nil:
			fmt.Printf("%s %s %s\n", red(Token), green("Success:"), white("Message has been sent to ", ServerID))
			BadCount--
		}
		return nil
	}
}
