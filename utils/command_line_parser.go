package utils

import (
	"flag"
)

func Get_commandline_values() (*bool, *bool) {
	log := flag.Bool("log", false, "Log to an external file")
	proxy := flag.Bool("proxy", false, "Use proxies when interacting with discord.com")
	flag.Parse()
	return log, proxy
}
