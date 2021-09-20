package cloudflare

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Cookie() string {
	url := "https://discord.com"
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responseCookies := resp.Cookies()
	dfcCookie := responseCookies[0].Value
	sdcCookie := responseCookies[1].Value
	return fmt.Sprintf("__dcfduid=%s; __sdcfduid=%s; locale=en-GB", dfcCookie, sdcCookie)
}
