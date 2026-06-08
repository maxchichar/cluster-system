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
	var arrivalMessage string

	switch result.House {

	case "KER":
	arrivalMessage = fmt.Sprintf(
		"━━━━━━━━━━━━━━━━━━━━━━\n\n"+
			"🔷 HOUSE KERNEL\n\n"+
			"Welcome <@%s>.\n\n"+
			"You have been assigned to the Core of CLUSTER ASCENSION.\n\n"+
			"House Kernel values leadership, discipline, stability, and responsibility.\n\n"+
			"📜 Motto:\n"+
			"\"Strength through Structure.\"\n\n"+
			"⚔️ Table: %d\n"+
			"📍 Seat: %s\n\n"+
			"Welcome to the Core Ascendant.\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━",
		i.Member.User.ID,
		result.TableID,
		result.Slot,
	)

	case "COM":
	arrivalMessage = fmt.Sprintf(
		"━━━━━━━━━━━━━━━━━━━━━━\n\n"+
			"⚙️ HOUSE COMPILER\n\n"+
			"Welcome <@%s>.\n\n"+
			"You have been assigned to the Forge of CLUSTER ASCENSION.\n\n"+
			"House Compiler transforms ideas into reality through creativity and execution.\n\n"+
			"📜 Motto:\n"+
			"\"Ideas Become Reality Here.\"\n\n"+
			"⚔️ Table: %d\n"+
			"📍 Seat: %s\n\n"+
			"Welcome to the Forge Ascendant.\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━",
		i.Member.User.ID,
		result.TableID,
		result.Slot,
	)

	case "RNT":
	arrivalMessage = fmt.Sprintf(
		"━━━━━━━━━━━━━━━━━━━━━━\n\n"+
			"🚀 HOUSE RUNTIME\n\n"+
			"Welcome <@%s>.\n\n"+
			"You have been assigned to the Engine of CLUSTER ASCENSION.\n\n"+
			"House Runtime thrives on action, adaptability, and performance.\n\n"+
			"📜 Motto:\n"+
			"\"Performance Defines Potential.\"\n\n"+
			"⚔️ Table: %d\n"+
			"📍 Seat: %s\n\n"+
			"Welcome to the Engine Ascendant.\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━",
		i.Member.User.ID,
		result.TableID,
		result.Slot,
	)

	case "ALG":
	arrivalMessage = fmt.Sprintf(
		"━━━━━━━━━━━━━━━━━━━━━━\n\n"+
			"🧠 HOUSE ALGORITHM\n\n"+
			"Welcome <@%s>.\n\n"+
			"You have been assigned to the Mind of CLUSTER ASCENSION.\n\n"+
			"House Algorithm excels in strategy, logic, and problem solving.\n\n"+
			"📜 Motto:\n"+
			"\"Wisdom Creates Possibility.\"\n\n"+
			"⚔️ Table: %d\n"+
			"📍 Seat: %s\n\n"+
			"Welcome to the Mind Ascendant.\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━",
		i.Member.User.ID,
		result.TableID,
		result.Slot,
	)
	}

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
