package utils

import (
	"Raid-Client/constants"
	"log"
	"os"
)

func SetupLogger() {
	if constants.Logging {
		constants.LogFile, _ = os.OpenFile("DiscSpam.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(constants.LogFile)
	}
}

func Logger(content string) {
	if constants.Logging {
		log.Print(content)
	}
}
