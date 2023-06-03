package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	env := getEnv()

	authStr := "Bot " + env.discordToken
	fmt.Println(authStr)
	s, err := discordgo.New(authStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	today := time.Now().Weekday()
	dayString := today.String()

	messageData := &discordgo.MessageSend{}
	messageData.Files = make([]*discordgo.File, 0)
	var file *discordgo.File

	// TODO: Videos
	// file = &discordgo.File{Reader: img, Name: "THURSDAY.mp4", ContentType: "video/mp4"}
	switch dayString {
	case "Sunday":
		break
	case "Monday":
		break
	case "Tuesday":
		break
	case "Wednesday":
		break
	case "Thursday":
		img, err := os.Open("./images/THURSDAY.png")
		if err != nil {
			return
		}
		defer img.Close()
		file = &discordgo.File{Reader: img, Name: "THURSDAY.png", ContentType: "image/png"}
		messageData.Files = append(messageData.Files, file)
	case "Friday":
		break
	case "Saturday":
		break
	}

	err = s.Open()
	if err != nil {
		fmt.Println("s.Open error", err)
	}
	defer s.Close()
	res, err := s.ChannelMessageSendComplex(env.channelId, messageData)

	if err != nil {
		fmt.Println("s.ChannelMessageSendComplex error", err)
	} else {
		fmt.Println("Send successful", res)
	}
}
