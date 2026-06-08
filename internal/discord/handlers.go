package discord

import (
	"fmt"
	"cluster-mvp/internal/services"

	"github.com/bwmarrin/discordgo"
)

const HouseArrivalChannelID = "1513488280631644220"

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

		err = s.GuildMemberRoleAdd(
			i.GuildID,
			i.Member.User.ID,
			roleID,
		)

		if err != nil {
			respond(
				s,
				i,
				"❌ Failed to assign house role: "+err.Error(),
			)
			return
		}

		err = s.GuildMemberRoleRemove(
			i.GuildID,
			i.Member.User.ID,
			"1512493065166655700", // Recruit role
		)

		if err != nil {
			respond(
				s,
				i,
				"⚠️ House assigned but failed to remove Recruit role: "+err.Error(),
			)
			return
		}
	}

	houseNames := map[string]string{
		"KER": "Kernel",
		"COM": "Compiler",
		"RNT": "Runtime",
		"ALG": "Algorithm",
	}

	houseName := houseNames[result.House]

	// House Arrival Message
	arrivalMessage := fmt.Sprintf(
		"🎉 Welcome <@%s> to House %s!\n\nPlease give them a warm welcome.",
		i.Member.User.ID,
		houseName,
	)

	_, _ = s.ChannelMessageSend(
		HouseArrivalChannelID,
		arrivalMessage,
	)

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