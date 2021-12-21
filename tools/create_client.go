package tools

import (
	"Raid-Client/constants"
	"Raid-Client/utils"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func CreateHttpClient() *http.Client {
	var httpClient *http.Client
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
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 5,
		}
	} else {
		httpClient = &http.Client{
			Timeout: time.Second * 5,
		}
	}
	return httpClient
}

func PopulateProxies() {
	fmt.Println("Populating proxies...")
	var err error
	constants.Proxies, err = utils.ReadTokens("./proxies.txt")
	if err != nil || len(constants.Proxies) == 0 {
		utils.Logger("Unable to load in proxies from proxies.txt... scraping the web for proxies.")
		fmt.Println("Unable to load in proxies from proxies.txt... scraping the web for proxies.")
		ScrapeProxies()
	}
}

func ScrapeProxies() {
	fmt.Println("Scraping proxies...")
	request, err := http.Get("https://api.proxyscrape.com/?request=displayproxies&proxytype=http&timeout=10000&country=all&anonymity=all&ssl=no")
	if err != nil {
		log.Fatal(err)
	}
	// Read the data from the web page I.E Proxies
	fmt.Println("Reading proxies...")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Writing proxies to file...")
	proxies := strings.TrimSuffix(string(data), "\n")
	// Append to our proxies file, if it does not exist we are simple going to create it and then append to it.
	file, err := os.OpenFile("./proxies.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Proxies File Not Found!...")
		os.Exit(0)
	}
	defer file.Close()
	// Write to our file

	_, err = fmt.Fprint(file, proxies)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	fmt.Println("Proxies Scraped and Saved!")
	utils.Logger("Scraped the web for proxies and added to proxies.txt")
	fmt.Println("Proxies saved to proxies.txt")
	// save the proxy list to constants
	constants.Proxies, _ = utils.ReadTokens("./proxies.txt")

}
