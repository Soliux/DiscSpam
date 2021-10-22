package constants

import (
	"os"
	"sync"

	"github.com/gookit/color"
)

var (
	// Define some of our colours so that we can print nicely to the console
	White  = color.FgWhite.Render
	Red    = color.FgRed.Render
	Green  = color.FgGreen.Render
	Yellow = color.FgYellow.Render

	// Define some other variables that will exist for the runtime of the program
	Tokens  []string
	Proxies []string
	Wg      sync.WaitGroup
	Logging bool
	Proxy   bool
	LogFile *os.File
)
