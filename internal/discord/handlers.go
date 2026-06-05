package discord

import (
	"fmt"
	"cluster-mvp/internal/services"

	"github.com/bwmarrin/discordgo"
)

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	if data.Name != "verify" {
		return
	}

	code := data.Options[0].StringValue()

	result, err := services.VerifyCode(code, i.Member.User.ID)

	if err != nil {
		respond(s, i, "❌ Invalid or used code")
		return
	}

	roleID := getRoleID(result.House)

	if roleID != "" {
		_ = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, roleID)

		// Replace with your real Recruit role ID
		_ = s.GuildMemberRoleRemove(
			i.GuildID,
			i.Member.User.ID,
			"1512493065166655700",
		)
	}

	houseNames := map[string]string{
		"KER": "Kernel",
		"COM": "Compiler",
		"RNT": "Runtime",
		"ALG": "Algorithm",
	}

	houseName := houseNames[result.House]

	message := fmt.Sprintf(
		"✅ Verification Successful\n\nHouse: %s\nTable: %d\nSeat: %s\n\nWelcome to House %s.",
		houseName,
		result.TableID,
		result.Slot,
		houseName,
	)

	respond(s, i, message)
}

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}