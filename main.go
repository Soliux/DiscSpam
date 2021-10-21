package main

import (
	"Raid-Client/constants"
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
	"time"

	"github.com/patrickmn/go-cache"
)

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
		utils.Logger("Exiting...")
		os.Exit(0)
	case "help", "h", "HELP ME", "menu", "home", "HELP", ".", "ls", "LS":
		utils.ClearScreen()
		Help()
		utils.Logger("Printing help menu")
	case "cls", "clear", "CLS", "CLEAR", "Clear", "Cls":
		utils.ClearScreen()
		utils.Logger("Clearing screen")
	case "1", "1.", "join", "join server", "JOIN", "JOIN SERVER":
		invite := Input("Enter Server Invite")
		utils.Logger("Join server module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, Server string) {
				server.JoinServer(Server, TOKEN)
				constants.Wg.Done()
			}(tkn, invite)
		}
		constants.Wg.Wait()
	case "2", "2.", "leave", "Leave Server", "leave server", "Leave":
		ServerID := Input("Enter Server ID")
		utils.Logger("Leave server module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, Server string) {
				server.LeaveServer(Server, TOKEN)
				constants.Wg.Done()
			}(tkn, ServerID)
		}
		constants.Wg.Wait()
	case "3", "3.", "spam message", "send messages", "message spammer", "spam":
		ServerID := Input("Enter Server ID")
		ChannelID := Input("Enter Channel ID")
		MessageToSpam := Input("Enter Message To Spam")
		utils.Logger("Message spammer module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, Server string, Message string, Channel string) {
				interact.SendMessage(Server, Channel, TOKEN, Message)
				constants.Wg.Done()
			}(tkn, ServerID, MessageToSpam, ChannelID)
		}
		constants.Wg.Wait()
	case "4", "4.", "reaction message", "add reaction":
		ChannelID := Input("Enter Channel ID")
		MessageID := Input("Enter Message ID")
		Emoji := Input("Enter Emoji")
		utils.Logger("Add reaction module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, Emoji string, Message string, Channel string) {
				interact.AddReaction(Channel, MessageID, TOKEN, Emoji)
				constants.Wg.Done()
			}(tkn, Emoji, MessageID, ChannelID)
		}
		constants.Wg.Wait()
	case "5", "5.", "react message", "message reaction":
		ChannelID := Input("Enter Channel ID")
		MessageID := Input("Enter Message ID")
		Word := Input("Enter Word")
		utils.Logger("Message reaction module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, Word string, Message string, Channel string) {
				interact.ReactionMessage(Channel, MessageID, TOKEN, Word)
				constants.Wg.Done()
			}(tkn, Word, MessageID, ChannelID)
		}
		constants.Wg.Wait()
	case "6", "6.", "change nickname", "nick":
		ServerID := Input("Enter Server ID")
		Nickname := Input("Enter Nickname")
		utils.Logger("Nickname changer module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, Server string, Nick string) {
				server.ChangeNickname(ServerID, TOKEN, Nickname)
				constants.Wg.Done()
			}(tkn, ServerID, Nickname)
		}
		constants.Wg.Wait()
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

		constants.Wg.Add(1)
		utils.Logger("Status change module starting...")
		go func() {
			interact.TOKENS = constants.Tokens
			interact.ChangeStatus()
			constants.Wg.Done()
		}()
		constants.Wg.Wait()

	case "8", "8.", "friend", "add friends":
		Username := Input("Enter Username")
		utils.Logger("Add friend module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, User string) {
				interact.AddFriend(TOKEN, Username)
				constants.Wg.Done()
			}(tkn, Username)
		}
		constants.Wg.Wait()
	case "9", "9.", "unfriend", "remove friends":
		UserID := Input("Enter User ID")
		utils.Logger("Unfriend module starting...")
		for _, tkn := range constants.Tokens {
			constants.Wg.Add(1)
			go func(TOKEN string, User string) {
				interact.RemoveFriend(TOKEN, UserID)
				constants.Wg.Done()
			}(tkn, UserID)
		}
		constants.Wg.Wait()

	case "10", "10.", "check", "check token", "token check":
		utils.Logger("Token Checker module starting...")
		constants.Wg.Add(1)
		go func() {
			defer constants.Wg.Done()
			t := utils.CheckTokens(constants.Tokens)
			utils.Logger("Writing old tokens to old_tokens.txt")
			utils.WriteLines(constants.Tokens, "./old_tokens.txt")
			constants.Tokens = nil
			constants.Tokens = t
			utils.Logger("Writing working tokens to tokens.txt")
			utils.WriteLines(constants.Tokens, "./tokens.txt")
		}()
		constants.Wg.Wait()
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
	fmt.Printf("%s %s\n", constants.White("1. Join Server - Params:"), constants.Red("<Invite Code>"))
	fmt.Printf("%s %s\n", constants.White("2. Leave Server - Params:"), constants.Red("<Server ID>"))
	fmt.Printf("%s %s\n", constants.White("3. Spam Message - Params:"), constants.Red("<Server ID> <Channel ID> <Message To Spam>"))
	fmt.Printf("%s %s\n", constants.White("4. Add Reaction - Params:"), constants.Red("<Channel ID> <Message ID> <Emoji>"))
	fmt.Printf("%s %s\n", constants.White("5. Add Reaction Message - Params:"), constants.Red("<Channel ID> <Message ID> <Reaction Message>"))
	fmt.Printf("%s %s\n", constants.White("6. Change Nickname - Params:"), constants.Red("<Server ID> <Nickname>"))
	fmt.Printf("%s %s\n", constants.White("7. Change Status - Params:"), constants.Red("<Content> <Status> <Type>"))
	fmt.Printf("%s %s\n", constants.White("8. Add Friend - Params:"), constants.Red("<Username> i.e Wumpus#0000"))
	fmt.Printf("%s %s\n", constants.White("9. Remove Friend - Params:"), constants.Red("<User ID>"))
	fmt.Printf("%s %s\n", constants.White("10. Check Tokens - Params"), constants.Red("<None>"))
}

func init() {
	l, p := utils.Get_commandline_values()
	constants.Logging = *l
	constants.Proxy = *p

	// Call our logger function and set the file output if needed
	utils.SetupLogger()

	server.C = cache.New(60*time.Minute, 120*time.Minute)
	tmp, err := utils.ReadTokens("./tokens.txt")
	if err != nil {
		utils.Logger("Error reading discord tokens from tokens.txt")
		fmt.Printf("%s\n", constants.Red("Error reading discord tokens from tokens.txt"))
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}
	constants.Tokens = tmp
}
