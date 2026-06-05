package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"cluster-mvp/internal/config"
)

var Bot *discordgo.Session

func Start(cfg config.Config) *discordgo.Session {

	dg, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMembers

	// IMPORTANT: slash command handler
	dg.AddHandler(interactionHandler)

	// register commands when bot is ready
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		registerCommands(s)
	})

	// auto assign UNVERIFIED role on join
	dg.AddHandler(memberJoinHandler)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	Bot = dg

	log.Println("Discord bot running")

	return dg
}