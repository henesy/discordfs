package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type DiscordState struct {
	Session	*discordgo.Session
	Guilds	[]*discordgo.UserGuild
	User	*discordgo.User
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
