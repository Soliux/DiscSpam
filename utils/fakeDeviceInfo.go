package utils

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type DeviceInformation struct {
	OS                       string `json:"os"`
	Browser                  string `json:"browser"`
	UA                       string `json:"browser_user_agent"`
	BrowserVersion           string `json:"browser_version"`
	OSVersion                string `json:"os_version"`
	Referrer                 string `json:"referrer"`
	ReferrerDomain           string `json:"referring_domain"`
	ReferrerCurrent          string `json:"referrer_current"`
	ReferreringCurrentDomain string `json:"referring_domain_current"`
	ReleaseChannel           string `json:"release_channel"`
	ClientBuild              string `json:"client_build_number"`
	ClientEventSource        string `json:"client_event_source"`
}

func FakeDevice() DeviceInformation {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(8-1) + 1 {
	case 1:
		return DeviceInformation{OS: "Windows", Browser: "Chrome", UA: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", BrowserVersion: "69.0.3497.100", OSVersion: "10", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	case 2:
		return DeviceInformation{OS: "Windows", Browser: "Chrome", UA: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763", BrowserVersion: "18.17763", OSVersion: "10", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	case 3:
		return DeviceInformation{OS: "Windows", Browser: "Edge", UA: "Mozilla/5.0 (Windows NT 5.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36", BrowserVersion: "60.0.3112.90", OSVersion: "10", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	case 4:
		return DeviceInformation{OS: "Windows", Browser: "Chrome", UA: "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36", BrowserVersion: "60.0.3112.113", OSVersion: "8.1", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	case 5:
		return DeviceInformation{OS: "Windows", Browser: "Internet Explorer", UA: "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729; rv:11.0) like Gecko", BrowserVersion: "11.0", OSVersion: "7", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	case 6:
		return DeviceInformation{OS: "Windows", Browser: "FireFox", UA: "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0", BrowserVersion: "54.0", OSVersion: "7", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	case 7:
		return DeviceInformation{OS: "Windows", Browser: "FireFox", UA: "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/66.0", BrowserVersion: "66.0", OSVersion: "7", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	default:
		return DeviceInformation{OS: "Windows", Browser: "FireFox", UA: "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/66.0", BrowserVersion: "66.0", OSVersion: "7", ReleaseChannel: "stable", ClientBuild: "36127", ClientEventSource: "None"}
	}
}

func XSuperProperties() string {
	fDevice := FakeDevice()
	byteArray, err := json.Marshal(fDevice)
	if err != nil {
		log.Fatal(err)
	}
	prop := base64.StdEncoding.EncodeToString([]byte(byteArray))
	return string(prop)
}
