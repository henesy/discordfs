package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"log"
)

func removeReaction(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	
}

func newReaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Global Mentions
	Mention := "@" + State.User.Username
	if strings.Contains(m.ContentWithMentionsReplaced(), Mention) {
		//go Notify(m.Message)
	}

	//State Messages
	if m.ChannelID == State.Channel.ID {
		State.AddMessage(m.Message)

		Messages := ReceivingMessageParser(m.Message)

		for _, Msg := range Messages {
			//MessagePrint(string(m.Timestamp), m.Author.Username, Msg)
			log.Printf("> %s > %s\n", m.Author.Username, Msg)
		}
	}
}
