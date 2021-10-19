package main

import (
	"Raid-Client/interact"
	"Raid-Client/server"
	"Raid-Client/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
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
	case "help", "h", "HELP ME", "menu", "home", "HELP", ".", "ls", "LS":
		utils.ClearScreen()
		Help()
	case "cls", "clear", "CLS", "CLEAR", "Clear", "Cls":
		utils.ClearScreen()
		Help()
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
		MessageID := Input("Enter Message ID")
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
		MessageID := Input("Enter Message ID")
		Word := Input("Enter Word")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Word string, Message string, Channel string) {
				interact.ReactionMessage(Channel, MessageID, TOKEN, Word)
				wg.Done()
			}(tkn, Word, MessageID, ChannelID)
		}
		wg.Wait()
	case "6", "6.", "change nickname", "nick":
		ServerID := Input("Enter Server ID")
		Nickname := Input("Enter Nickname")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, Server string, Nick string) {
				server.ChangeNickname(ServerID, TOKEN, Nickname)
				wg.Done()
			}(tkn, ServerID, Nickname)
		}
		wg.Wait()
	case "7", "7.", "status", "change status":
		Type := 0
		fmt.Println("[NOTE] The status will only remain active for 2 minutes")
		Content := Input("Enter Status Content (e.g. hello world)")
		utils.Status = Input("Enter Status (e.g. online, idle, dnd)")
		Activity := strings.ToLower(Input("Enter Type (e.g. playing, watching, listening)"))

		switch Activity {
		case "playing":
			Type = utils.ActivityGame
		case "watching":
			Type = utils.ActivityWatching
		case "listening":
			Type = utils.ActivityListening
		}

		utils.Presence = []utils.Activity{{Name: Content, Type: utils.ActivityType(Type)}}

		wg.Add(1)
		go func() {
			interact.TOKENS = Tokens
			interact.ChangeStatus()
			wg.Done()
		}()
		wg.Wait()

	case "8", "8.", "friend", "add friends":
		Username := Input("Enter Username")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, User string) {
				interact.AddFriend(TOKEN, Username)
				wg.Done()
			}(tkn, Username)
		}
		wg.Wait()
	case "9", "9.", "unfriend", "remove friends":
		UserID := Input("Enter User ID")
		for _, tkn := range Tokens {
			wg.Add(1)
			go func(TOKEN string, User string) {
				interact.RemoveFriend(TOKEN, UserID)
				wg.Done()
			}(tkn, UserID)
		}
		wg.Wait()
	}
}

func Input(DisplayText string) string {
	var rtnInput string
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	if user.Name == "" {
		user.Name = "Raider"
	}
	if DisplayText == "" {
		fmt.Printf("%s@DiscSpam > ", user.Name)
	} else {
		fmt.Printf("%s > ", DisplayText)
	}
	text, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		rtnInput = strings.Replace(text, "\r\n", "", -1)
	} else {
		rtnInput = strings.Replace(text, "\n", "", -1)
	}
	return rtnInput
}

func Help() {
	fmt.Printf("%s %s\n", white("1. Join Server - Params:"), red("<Invite Code>"))
	fmt.Printf("%s %s\n", white("2. Leave Server - Params:"), red("<Server ID>"))
	fmt.Printf("%s %s\n", white("3. Spam Message - Params:"), red("<Server ID> <Channel ID> <Message To Spam>"))
	fmt.Printf("%s %s\n", white("4. Add Reaction - Params:"), red("<Channel ID> <Message ID> <Emoji>"))
	fmt.Printf("%s %s\n", white("5. Add Reaction Message - Params:"), red("<Channel ID> <Message ID> <Reaction Message>"))
	fmt.Printf("%s %s\n", white("6. Change Nickname - Params:"), red("<Server ID> <Nickname>"))
	fmt.Printf("%s %s\n", white("7. Change Status - Params:"), red("<Content> <Status> <Type>"))
	fmt.Printf("%s %s\n", white("8. Add Friend - Params:"), red("<Username> i.e Wumpus#0000"))
	fmt.Printf("%s %s\n", white("9. Remove Friend - Params:"), red("<User ID>"))
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
