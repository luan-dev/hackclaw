package commands

import (
	"codecosta.com/hackclaw/app/utils"
	"github.com/bwmarrin/discordgo"
)

var CommandList = []*discordgo.ApplicationCommand{
	{
		Name:        "spawns",
		Description: "Show spawns",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "map",
				Description: "Map Name",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "test",
		Description: "Test command",
	},
}

var CommandHandlers = map[string]func(discord *discordgo.Session, interaction *discordgo.InteractionCreate){
	"test":   test,
	"spawns": SendSpawns,
}

func test(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "test")

	content := "test response"

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
