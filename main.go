package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func main() {
	token := os.Getenv("TOKEN")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return
	}
	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection to Discord API.", err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			return
		}
	}(discord)

	<-make(chan struct{})
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	content := strings.Fields(message.Content)
	if content[1] == "ping" {
		_, err := session.ChannelMessageSend(message.ChannelID, "pong!")
		if err != nil {
			return
		}
	}
}
