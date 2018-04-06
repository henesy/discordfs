package main

import (
//	"github.com/bwmarrin/discordgo"
	"log"
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
	// TODO: Read from factotum
	Config.Init()
	
	err := State.Init()
	if err != nil {
		log.Fatalln("Unable to start Discord Session: ", err)
	}
	
	
}
