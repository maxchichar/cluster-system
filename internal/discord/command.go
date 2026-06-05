package discord

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "verify",
		Description: "Verify yourself",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "code",
				Description: "Your invite code",
				Required:    true,
			},
		},
	},
}

func registerCommands(s *discordgo.Session) {
	appID := s.State.User.ID

	_, err := s.ApplicationCommandBulkOverwrite(appID, "", Commands)
	if err != nil {
		panic(err)
	}
}