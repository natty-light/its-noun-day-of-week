package main

import (
	"bytes"
	"fmt"
	"its-noun-day-of-week/utils"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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

	buffer := aws.NewWriteAtBuffer([]byte{})
	_, err = downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(env.S3Bucket),
		Key:    aws.String("THURSDAY.png"),
	})
	if err != nil {
		fmt.Println("S3 download error", err)
	}
	file = &discordgo.File{Reader: bytes.NewReader(buffer.Bytes()), Name: "THURSDAY.png", ContentType: "image/png"}
	messageData.Files = append(messageData.Files, file)
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
		break
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
