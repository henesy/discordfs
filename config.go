package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Configuration is a struct that contains all configuration fields
type Configuration struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	MessageDefault bool   `json:"messagedefault"`
	Messages       int    `json:"messages"`
}

// Config is the global configuration of discord-cli
var Config Configuration

// Get Username/Password and load into Config
func (c *Configuration) Init() {
	var EmptyStruct Configuration
	//Set Default values
	fmt.Print("Input your email: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	EmptyStruct.Username = scan.Text()
	fmt.Print("Input your password: ")
	//password, err := terminal.ReadPassword(0)
	
	plan9 := true
	/* Plan 9 raw mode for rio */
	consctl, err := Rawon()
	if err != nil {
		fmt.Println("Failed to set rawon")
		plan9 = false
	}
	
	password := "";
	
	if plan9 {
		password = GetCons()
	
		err = RawOff(consctl)
		if err != nil {
			fmt.Println("\nFailed to set rawoff")
			fmt.Print(err, "\n")
		}
	} else {
		log.Fatal("Discordfs only support 9front and being run on rio atm, sorry.")
	}
	
	EmptyStruct.Password = password
	EmptyStruct.Messages = 10
	EmptyStruct.MessageDefault = true
	
	Config = EmptyStruct
}
