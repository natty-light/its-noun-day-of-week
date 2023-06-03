package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	env := getEnv()

	discordClient, err := discordgo.New("Bot " + env.discordToken)
	if err != nil {
		return
	}

	today := time.Now().Weekday()
	dayString := today.String()

	messageData := &discordgo.MessageSend{}
	file := &discordgo.File{}

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
		break
	case "Friday":
		break
	case "Saturday":
		break
	}

	discordClient.ChannelMessageSendComplex()
}
