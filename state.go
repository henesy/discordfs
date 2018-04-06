package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type DiscordState struct {
	Session		*discordgo.Session
	Guilds		[]*discordgo.UserGuild
	User		*discordgo.User
	Channel		*discordgo.Channel
	Messages	[]*discordgo.Message
	MaxMessages	int
}

// Create a new Discord Session
func (s *DiscordState) Init() error {
	s.Session = new(discordgo.Session)
	
	fmt.Printf("\nConnecting…")

	dg, err := discordgo.New(Config.Username, Config.Password)
	if err != nil {
		return err
	}

	// Open the websocket and begin listening.
	dg.Open()

	//Retrieve GuildID's from current User
	//need index of Guilds[] rather than UserGuilds[] (maybe)
	Guilds, err := dg.UserGuilds(0, "", "")
	if err != nil {
		return err
	}

	s.Guilds = Guilds

	s.Session = dg

	s.User, _ = s.Session.User("@me")

	fmt.Printf(" PASSED!\n")

	return nil
}

// Add Message to State
func (State *DiscordState) AddMessage(Message *discordgo.Message) {
	//Do not add if Amount <= 0
	if State.MaxMessages <= 0 {
		return
	}

	//Remove First Message if next message is going to increase length past MessageAmount
	if len(State.Messages) == State.MaxMessages {
		State.Messages = append(State.Messages[:0], State.Messages[1:]...)
	}

	State.Messages = append(State.Messages, Message)
}
