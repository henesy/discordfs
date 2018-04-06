package main

import (
//	"github.com/bwmarrin/discordgo"
	"log"
	"flag"
	"os"
)

// Global Message Types
const (
	ErrorMsg  = "Error"
	InfoMsg   = "Info"
	HeaderMsg = "Head"
	TextMsg   = "Text"
)

// Global state
var State DiscordState

// Global message type
type MsgType string


/* Connects as a given account to Discord and serves an fs */
func main() {
	var status string
	flag.StringVar(&status, "s", "⑨", "Set the default status message")

	// TODO: Read from factotum
	Config.Init()
	
	err := State.Init()
	if err != nil {
		log.Fatalln("Unable to start Discord Session: ", err)
	}
	
	// Connect event handlers (see: discordgo/events.go)
	State.Session.AddHandler(newMessage)
	State.Session.AddHandler(newReaction)
	
	log.SetOutput(os.Stderr)
	State.Session.UpdateStatus(0, status)
}
