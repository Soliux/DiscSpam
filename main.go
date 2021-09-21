package main

// Token we are using for testing: ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk
// Invite to server for testing: https://discord.gg/7XZNPEcHza
import (
	"Raid-Client/interact"
	"Raid-Client/server"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
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
	err := interact.ReactionMessage("889537520688332855", "889913197119828019", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk", "gay")
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	server.C = cache.New(60*time.Minute, 120*time.Minute)
}
