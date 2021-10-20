package utils

import (
	"Raid-Client/cloudflare"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gookit/color"
)

var (
	wg         sync.WaitGroup
	mutex      sync.Mutex
	goodTokens []string
	good       int
	bad        int
	locked     int

	green = color.FgGreen.Render
	white = color.FgWhite.Render
	red   = color.FgRed.Render
)

func CheckTokens(tokens []string) []string {
	good = 0
	bad = 0
	locked = 0
	fmt.Printf("Checking %d tokens\n", len(tokens))
	for _, t := range tokens {
		wg.Add(1)
		go func(t string) {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()
			request, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/guild-events", nil)
			if err != nil {
				fmt.Println(err)
			}
			cf := cloudflare.Cookie()
			xprop := XSuperProperties()
			request.Header = http.Header{
				"Accept":             []string{"*/*"},
				"Accept-language":    []string{"en-GB"},
				"Authorization":      []string{t},
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
				fmt.Println(err)
			}
			defer res.Body.Close()

			switch res.StatusCode {
			case 200:
				good++
				fmt.Printf("%s %s \n", white(t), green("| is Valid"))
				goodTokens = append(goodTokens, t)
			case 401:
				bad++
				fmt.Printf("%s %s \n", white(t), red("| is Invalid"))
			case 403:
				locked++
				fmt.Printf("%s %s \n", white(t), red("| is Phone locked"))
			default:
				bad++
				fmt.Printf("%s %s \n", white(t), red("| is Invalid"))
			}
		}(t)
	}
	wg.Wait()
	fmt.Printf("%s\n%s%s\n%s%s\n%s%s\n", green("Finished Checking: "), green("Good tokens: "), green(good), red("Bad tokens: "), red(bad), red("Locked tokens: "), red(locked))
	return goodTokens
}
