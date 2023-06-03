package main

import (
	"fmt"
	"its-noun-day-of-week/utils"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	env := utils.GetEnv()

	authStr := "Bot " + env.DiscordToken
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

	downloader, err := utils.CreateS3Downloader(env)
	if err != nil {
		return
	}

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
		file, err = utils.DownloadAndParseFile(downloader, env, "THURSDAY.mp4", "video/mp4")
		if err != nil {
			return
		}
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
	res, err := s.ChannelMessageSendComplex(env.ChannelId, messageData)

	if err != nil {
		fmt.Println("s.ChannelMessageSendComplex error", err)
	} else {
		fmt.Println("Send successful", res)
	}
}
