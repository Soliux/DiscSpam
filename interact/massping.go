package interact

import (
	"Raid-Client/utils"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func MassPing(guildID, channelID string, scrapedID, tokens []string, amount int) {

	var toSend []string
	var message string
	for _, ID := range scrapedID {
		if len(message) < 1950 {
			message += fmt.Sprintf("<@!%s>ඞ", fmt.Sprint(ID))
		} else {
			toSend = append(toSend, message)
			message = ""
		}
	}
	toSend = append(toSend, message)

	/*
		Amount = 2
		Tokens = 50
		Messages = 20

		totalMessages = Amount*(Tokens*Messages)
	*/
	fmt.Println("Sending a total of:", amount*(len(toSend)*len(tokens)), "messages")

	for _, token := range tokens {
		wg.Add(1)
		go func(t, c, s string, a int) {
			defer wg.Done()
			for i := 0; i < amount; i++ {
				for _, m := range toSend {
					time.Sleep(time.Millisecond * 100)
					utils.Logger("Sending:", m[:200], t, "Amount:", i)
					err := SendMessage(s, c, t, m)
					if err != nil {
						break
					}
				}
			}
		}(token, channelID, guildID, amount)

	}
	wg.Wait()
}
