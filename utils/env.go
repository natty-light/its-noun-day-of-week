package utils

import "os"

type Env struct {
	DiscordToken  string
	ChannelId     string
	AwsRegion     string
	S3Bucket      string
	TimestampFile string
}

func GetEnv() Env {
	var env Env = Env{}

	env.DiscordToken = os.Getenv("discordToken")
	env.ChannelId = os.Getenv("channelId")
	env.AwsRegion = os.Getenv("BUCKET_REGION")
	env.S3Bucket = os.Getenv("AWS_S3_BUCKET")
	env.TimestampFile = os.Getenv("timestampFile")

	return env
}
