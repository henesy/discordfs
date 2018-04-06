package main

import (
	"os"
	"log"
	"bufio"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Rawon() (*os.File, error){
	consctl, err := os.OpenFile("/dev/consctl", os.O_WRONLY, 0200)
	if err != nil {
		/* not on Plan 9 */
		log.Println("\nNot running on Plan 9")
		return consctl, err
	}
	
	rawon := []byte("rawon")
	_, err = consctl.Write(rawon)
	if err != nil {
		consctl.Close()
		return consctl, err
	}
	
	return consctl, nil
}

func RawOff(consctl *os.File) error {
	//consctl, err := os.OpenFile("/dev/consctl", os.O_WRONLY, 0200)
	//if err != nil {
	//	/* not on Plan 9 */
	//	return err
	//}
	
	rawoff := []byte("rawoff")
	_, err := consctl.Write(rawoff)
	if err != nil {
		consctl.Close()
		return err
	}
	
	consctl.Close()
	return nil
}

func GetCons() string {
	cons, err := os.OpenFile("/dev/cons", os.O_RDWR, 0600)
	if err != nil {
		log.Println("Failed to open /dev/cons")
	}
	consScan := bufio.NewScanner(cons)
	consScan.Scan()
	return consScan.Text()
}

//ReceivingMessageParser parses receiving message for mentions, images and MultiLine and returns string array
func ReceivingMessageParser(m *discordgo.Message) []string {
	Message := m.ContentWithMentionsReplaced()

	//Parse images
	for _, Attachment := range m.Attachments {
		Message = Message + " " + Attachment.URL
	}

	// MultiLine comment parsing
	Messages := strings.Split(Message, "\n")

	return Messages
}
