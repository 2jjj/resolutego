package handler

import (
	"github.com/bwmarrin/discordgo"
)

func PingHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}
}
