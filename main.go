package main

// Token we are using for testing: ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk
// Invite to server for testing: https://discord.gg/7XZNPEcHza
import (
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
	err := server.LeaveServer("889537520688332851", "ODg5NTUzMTY2MTU5NDUwMTY0.YUi7nA.6yXPQf3fWb-qFj6pncNP97Ie_pk")
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	server.C = cache.New(60*time.Minute, 120*time.Minute)
}
