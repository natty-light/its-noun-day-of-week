package main

import "os"

type Env struct {
	discordToken string
}

func getEnv() Env {
	var env Env = Env{}

	env.discordToken = os.Getenv("discordToken")

	return env
}
