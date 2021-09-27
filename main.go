package main

// Token we are using for testing: ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk
// Invite to server for testing: https://discord.gg/7XZNPEcHza
import (
	"Raid-Client/server"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/patrickmn/go-cache"
)

var white = color.FgWhite.Render
var red = color.FgRed.Render

func main() {
	for {
		text := Input("")
		spawner(text)
	}

}

func spawner(Tool string) {
	switch Tool {
	case "exit", ".exit", "EXIT", "close":
		os.Exit(0)
	case "help", "h", "HELP ME", "menu", "home", "HELP":
		fmt.Printf("%s %s\n", white("1. Join Server - Params:"), red("<Invite Code>"))
		fmt.Printf("%s %s\n", white("2. Leave Server - Params:"), red("<Server ID>"))
		fmt.Printf("%s %s\n", white("3. Spam Message - Params:"), red("<Server ID> <Channel ID> <Message To Spam>"))
	}
}

func Input(DisplayText string) string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	if DisplayText == "" {
		fmt.Printf("%s@DiscSpam > ", user.Name)
	} else {
		fmt.Printf("%s > ", DisplayText)
	}
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func init() {
	server.C = cache.New(60*time.Minute, 120*time.Minute)
}

// err := server.JoinServer("https://discord.gg/7XZNPEcHza", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk")
// if err != nil {
// 	log.Fatal(err)
// }
// time.Sleep(3 * time.Second)
// err = server.LeaveServer("889537520688332851", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk")
// if err != nil {
// 	log.Fatal(err)
// }
// err := interact.SendMessage("889537520688332851", "889537520688332855", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk", "Hello from other side")
// if err != nil {
// 	log.Fatal(err)
// }
// err := interact.AddReaction("889537520688332855", "889913197119828019", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk", "🇦")
// if err != nil {
// 	log.Fatal(err)
// }
// err := interact.ReactionMessage("889537520688332855", "889913197119828019", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk", "loi")
// if err != nil {
// 	log.Fatal(err)
// }
