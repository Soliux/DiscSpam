package tools

import (
	"Raid-Client/constants"
	"Raid-Client/utils"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func Create_http_client() *http.Client {
	var http_client *http.Client
	if constants.Proxy {
		proxy := constants.Proxies[rand.Intn(len(constants.Proxies))]
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			fmt.Println(err)
		}
		// Create our transport so we can setup our proxy
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		// Create our client that is told to run traffic through the proxy
		http_client = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 5,
		}
	} else {
		http_client = &http.Client{
			Timeout: time.Second * 5,
		}
	}
	return http_client
}

func Populate_proxies() {
	var err error
	constants.Proxies, err = utils.ReadTokens("./proxies.txt")
	if err != nil {
		utils.Logger("Unable to load in proxies from proxies.txt... scraping the web for proxies.")
		fmt.Println("Unable to load in proxies from proxies.txt... scraping the web for proxies.")
		// TODO: Create the function to scrape the proxies and then write them to our proxies.txt so that we do not have to do that again.
	}
}
