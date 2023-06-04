package main

import (
	"fmt"
	"its-noun-day-of-week/utils"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	env := utils.GetEnv()

	s, err := discordgo.New("Bot " + env.DiscordToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	dayString := strings.ToLower(time.Now().Weekday().String())

	fmt.Println(dayString)
	client, err := utils.CreateS3Client(env)
	if err != nil {
		fmt.Println("S3 Client config error", err)
		return
	}
	downloader := utils.CreateS3Downloader(client)
	if err != nil {
		fmt.Println("S3 Downloader error", err)
		return
	}
	d := utils.S3DataSource{Client: client, Downloader: downloader}

	messageData, err := prepareDailyMessage(env, d, "thursday")
	// messageData, err := prepareDailyMessage(env, d, dayString)
	if err != nil {
		fmt.Println("Prepare message error", err)
		return
	}

	err = s.Open()
	if err != nil {
		fmt.Println("s.Open error", err)
		return
	}
	defer s.Close()
	res, err := s.ChannelMessageSendComplex(env.ChannelId, messageData)
	if err != nil {
		fmt.Println("s.ChannelMessageSendComplex error", err)
	} else {
		fmt.Println("Send successful", res)
	}
}

func prepareDailyMessage(env utils.Env, d utils.S3DataSource, dayOfWeek string) (*discordgo.MessageSend, error) {
	messageData := &discordgo.MessageSend{}
	messageData.Files = make([]*discordgo.File, 0)
	keys, _ := d.ListAllFilesInFolder(env, dayOfWeek)
	if len(keys) == 0 {
		return nil, fmt.Errorf("no images for %s", dayOfWeek)
	}
	for _, item := range keys[1:] {
		key := *item.Key
		file, err := d.DownloadAndParseFile(env, key)
		if err != nil {
			continue
		}
		messageData.Files = append(messageData.Files, file)
	}
	return messageData, nil
}
