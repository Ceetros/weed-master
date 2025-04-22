package Discord

import "github.com/bwmarrin/discordgo"

var Bot Discord

type Discord struct {
	MainGuild           *discordgo.Guild
	Bot                 *discordgo.Session
	NotificationChannel string
}

func GetDiscord() Discord {
	return Bot
}

func SetBot(bot Discord) {
	Bot = bot
}
