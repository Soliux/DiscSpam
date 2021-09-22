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

	"github.com/patrickmn/go-cache"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

}

func init() {
	server.C = cache.New(60*time.Minute, 120*time.Minute)
}

func Input(DisplayText string) string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s@%s: ", user.Name, DisplayText)
	text, _ := reader.ReadString('\n')
	return text
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
