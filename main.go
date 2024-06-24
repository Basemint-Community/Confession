package main

import (
	"flag"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Token = os.Getenv("TOKEN")
	ConfessionsToken = flag.String("Bot", "", Token)
	TestGuild = flag.String("guild", "", "1233196678530728008")
	TestChannel = flag.String("channel", "", "1233196679008616490")
)

func init() {
	flag.Parse()
}

func main() {
	session, _ := discordgo.New(Token)
	session.Identify.Intents = discordgo.IntentMessageContent
}