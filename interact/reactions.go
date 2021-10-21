package interact

import (
	"Raid-Client/cloudflare"
	"Raid-Client/constants"
	"Raid-Client/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kyokomi/emoji/v2"
)

var cmt int

func AddReaction(ChannelID string, MessageID string, Token string, Emoji string) error {
	if cmt >= 2 {
		return errors.New("error working")
	} else {
		Emoji = strings.TrimSuffix(emoji.Sprint(Emoji), " ")
		request, err := http.NewRequest("PUT", fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s/reactions/%s/%s", ChannelID, MessageID, Emoji, "%40me"), nil)
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

		switch res.StatusCode {
		case 204:
			utils.Logger(fmt.Sprintf("%s has successfully added the reaction %s to the message %s", Token, Emoji, MessageID))
			fmt.Printf("%s %s %s\n", constants.Red(Token), constants.Green("Success:"), constants.White(fmt.Sprintf("Added reaction %s to %s", Emoji, MessageID)))
			cmt--
		default:
			utils.Logger(fmt.Sprintf("%s was unable to add the reaction %s to %s", Token, Emoji, MessageID))
			fmt.Printf("%s %s\n", constants.White(Token), constants.Red(fmt.Sprintf("Error: Unable to add reaction %s to %s", Emoji, MessageID)))
			cmt++
		}

		return nil
	}
}

func ReactionMessage(ChannelID string, MessageID string, Token string, Word string) error {
	letters := map[string]string{
		"A": "🇦",
		"B": "🇧",
		"C": "🇨",
		"D": "🇩",
		"E": "🇪",
		"F": "🇫",
		"G": "🇬",
		"H": "🇭",
		"I": "🇮",
		"J": "🇯",
		"K": "🇰",
		"L": "🇱",
		"M": "🇲",
		"N": "🇳",
		"O": "🇴",
		"P": "🇵",
		"Q": "🇶",
		"R": "🇷",
		"S": "🇸",
		"T": "🇹",
		"U": "🇺",
		"V": "🇻",
		"W": "🇼",
		"X": "🇽",
		"Y": "🇾",
		"Z": "🇿",
	}
	for _, letter := range Word {
		l := strings.ToUpper(string(letter))
		AddReaction(ChannelID, MessageID, Token, letters[l])
		time.Sleep(1 * time.Second)
	}
	return nil
}
