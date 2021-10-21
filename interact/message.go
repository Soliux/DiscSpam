package interact

import (
	"Raid-Client/cloudflare"
	"Raid-Client/constants"
	"Raid-Client/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var BadCount int

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
			utils.Logger(fmt.Sprintf("%s was not able ot send a message in %s", Token, ServerID))
			fmt.Printf("%s %s\n", constants.White(Token), constants.Red("Unable to send message make sure it is the server"))
			BadCount++
		case "Missing Permissions":
			utils.Logger(fmt.Sprintf("%s does not have the correct permissions to send a message in %s", Token, ServerID))
			fmt.Printf("%s %s %s\n", constants.Red(Token), constants.Yellow("is missing permissions I.E needs a role to message in"), constants.White(ServerID))
			BadCount++
		case nil:
			utils.Logger(fmt.Sprintf("%s has sent the message %s in %s", Token, Content, ServerID))
			fmt.Printf("%s %s %s\n", constants.Red(Token), constants.Green("Success:"), constants.White("Message has been sent to ", ServerID))
			BadCount--
		}
		return nil
	}
}
