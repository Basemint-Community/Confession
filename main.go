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
	applicationID := os.Getenv("ApplicationID")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(interactionCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	err = registerSlashCommand(dg, applicationID)
	if err != nil {
		fmt.Println("Error registering slash command: ", err)
		dg.Close()
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func registerSlashCommand(s *discordgo.Session, applicationID string) error {
    command := &discordgo.ApplicationCommand{
        Name:        "confessions",
        Description: "Request confessions from the bot with an optional message.",
        Options: []*discordgo.ApplicationCommandOption{
            {
                Type:        discordgo.ApplicationCommandOptionString,
                Name:        "message",
                Description: "Your confession message",
                Required:    false,
            },
        },
    }
    
    _, err := s.ApplicationCommandCreate(applicationID, "", command)
    if err != nil {
        return fmt.Errorf("error registering slash command: %v", err)
    }
    
    fmt.Println("Slash command '/confessions' registered successfully.")
    return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "!hello") {
		s.ChannelMessageSend(m.ChannelID, "meow meow")
	}
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
    if i.Type == discordgo.InteractionApplicationCommand {
        switch i.ApplicationCommandData().Name {
        case "confessions":
            confessionMessage := ""
            for _, opt := range i.ApplicationCommandData().Options {
                if opt.Name == "message" && opt.Type == discordgo.ApplicationCommandOptionString {
                    confessionMessage = opt.StringValue()
                    break
                }
            }
            
            channelID := os.Getenv("ChannelID")
            embed := &discordgo.MessageEmbed{
                Type:        discordgo.EmbedTypeRich,
                Title:       "Confession",
                Description: confessionMessage,
                Color:       0xff0000, 
            }
            
            if confessionMessage != "" {
                _, err := s.ChannelMessageSendEmbed(channelID, embed)
                if err != nil {
                    fmt.Println("Error sending embed message: ", err)
                }
            }
        }
    }
}



 