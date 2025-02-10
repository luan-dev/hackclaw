package main

import (
	"log"
	"os"

	bot "codecosta.com/hackclaw/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discordAppID := os.Getenv("DISCORD_APP_ID")
	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	// COSTA_DASH_TOKEN := os.Getenv("COSTA_DASH_TOKEN")

	bot.DiscordAppID = discordAppID
	bot.DiscordBotToken = discordBotToken
	bot.Run()
}
