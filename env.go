package main

import "os"

type Env struct {
	discordToken string
	channelId    string
}

func getEnv() Env {
	var env Env = Env{}

	env.discordToken = os.Getenv("discordToken")
	env.channelId = os.Getenv("channelId")

	return env
}
