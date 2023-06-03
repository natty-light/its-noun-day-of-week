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

	client, err := utils.CreateS3Client(env)
	if err != nil {
		fmt.Println("S3 Client config error", err)
		return
	}
	downloader := utils.CreateS3Downloader(client)
	if err != nil {
		return
	}
	d := utils.S3DataSource{Client: client, Downloader: downloader}

	x, _ := d.ListAllFilesInFolder(env, "thursday")
	for _, item := range x.Contents[1:] {
		key := *item.Key
		mimetype, err := d.GetObjectMetadata(env, key)
		if err != nil {
			fmt.Println(err)
		}
		file, err := d.DownloadAndParseFile(env, *item.Key, mimetype)
		if err != nil {
			continue
		}
		messageData.Files = append(messageData.Files, file)
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
