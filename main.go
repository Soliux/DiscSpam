package main

// Token we are using for testing: ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk
// Invite to server for testing: https://discord.gg/7XZNPEcHza
import (
	"Raid-Client/interact"
	"Raid-Client/server"
	"Raid-Client/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	"github.com/patrickmn/go-cache"
)

var white = color.FgWhite.Render
var red = color.FgRed.Render
var Tokens []string
var wg sync.WaitGroup

func main() {
	Help()
	for {
		text := Input("")
		spawner(text)
	}

}

func spawner(Tool string) {
	switch Tool {
	case "exit", ".exit", "EXIT", "close":
		os.Exit(0)
	case "help", "h", "HELP ME", "menu", "home", "HELP", ".":
		Help()
	case "cls", "clear", "CLS", "CLEAR", "Clear", "Cls":
		utils.ClearScreen()
	// There most likely is a more elegant way of doing this but I am just going to do this because it is simple and easy to do
	case "1", "1.", "join", "join server", "JOIN", "JOIN SERVER":
		invite := Input("Enter Server Invite")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Server string) {
				server.JoinServer(Server, TOKEN)
				wg.Done()
			}(tkn, invite)
		}
		wg.Wait()
	case "2", "2.", "leave", "Leave Server", "leave server", "Leave":
		ServerID := Input("Enter Server ID")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Server string) {
				server.LeaveServer(Server, TOKEN)
				wg.Done()
			}(tkn, ServerID)
		}
		wg.Wait()
	case "3", "3.", "spam message", "send messages", "message spammer", "spam":
		ServerID := Input("Enter Server ID")
		ChannelID := Input("Enter Channel ID")
		MessageToSpam := Input("Enter Message To Spam")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Server string, Message string, Channel string) {
				interact.SendMessage(Server, Channel, TOKEN, Message)
				wg.Done()
			}(tkn, ServerID, MessageToSpam, ChannelID)
		}
		wg.Wait()
	case "4", "4.", "reaction message", "add reaction":
		ChannelID := Input("Enter Channel ID")
		MessageID := Input("Etner Message ID")
		Emoji := Input("Enter Emoji")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Emoji string, Message string, Channel string) {
				interact.AddReaction(Channel, MessageID, TOKEN, Emoji)
				wg.Done()
			}(tkn, Emoji, MessageID, ChannelID)
		}
		wg.Wait()
	case "5", "5.", "react message", "message reaction":
		ChannelID := Input("Enter Channel ID")
		MessageID := Input("Etner Message ID")
		Word := Input("Enter Word")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Word string, Message string, Channel string) {
				interact.ReactionMessage(Channel, MessageID, TOKEN, Word)
				wg.Done()
			}(tkn, Word, MessageID, ChannelID)
		}
		wg.Wait()
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

func Help() {
	fmt.Printf("%s %s\n", white("1. Join Server - Params:"), red("<Invite Code>"))
	fmt.Printf("%s %s\n", white("2. Leave Server - Params:"), red("<Server ID>"))
	fmt.Printf("%s %s\n", white("3. Spam Message - Params:"), red("<Server ID> <Channel ID> <Message To Spam>"))
	fmt.Printf("%s %s\n", white("4. Add Reaction - Params:"), red("<Channel ID> <Message ID> <Emoji>"))
	fmt.Printf("%s %s\n", white("5. Add Reaction Message - Params:"), red("<Channel ID> <Message ID> <Reaction Message>"))
}

func init() {
	server.C = cache.New(60*time.Minute, 120*time.Minute)
	tmp, err := utils.ReadTokens("./tokens.txt")
	if err != nil {
		fmt.Printf("%s\n", red("Error reading tokens in tokens.txt"))
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}
	Tokens = tmp
}
