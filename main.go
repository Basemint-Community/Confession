package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Basemint-Community/Confession/Confessions"
	"github.com/bwmarrin/discordgo"
)

func init() {
	Confessions.LoadEnv()
}

func main() {
	token := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "!hello") {
		s.ChannelMessageSend(m.ChannelID, "Hello, I'm a Discord bot!")
	}
}
