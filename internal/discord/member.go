package discord

import "github.com/bwmarrin/discordgo"

func memberJoinHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	// Replace with your real role ID
	_ = s.GuildMemberRoleAdd(
		m.GuildID,
		m.User.ID,
		"UNVERIFIED_ROLE_ID",
	)
}