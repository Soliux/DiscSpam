package utils

import (
	"Raid-Client/constants"
	"fmt"
	"log"
	"os"
	"strings"
)

func SetupLogger() {
	if constants.Logging {
		constants.LogFile, _ = os.OpenFile("DiscSpam.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(constants.LogFile)
	}
}

func Logger(contents ...interface{}) {
	if constants.Logging {
		content := formatMessage(contents...)
		log.Println(content)

	}
}
func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}
