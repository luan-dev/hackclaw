package commands

import (
	"fmt"

	"codecosta.com/hackclaw/app/models"
	"codecosta.com/hackclaw/app/utils"
	"github.com/bwmarrin/discordgo"
)

func SendSpawns(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	utils.LogCommand(interaction.Member.User.Username, "/spawns")

	options := interaction.ApplicationCommandData().Options
	selectedMap := options[0].Value.(string)

	switch selectedMap {
	case string(models.ZERO_DAM):
		break
	default:
		err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       fmt.Sprintf("Sorry %s is not supported", selectedMap),
						Description: "We currently only support: Zero Dam",
					},
				},
			},
		})
		if err != nil {
			utils.LogDiscordError("SendSpawns.InteractionRespond", err.Error())
		}

		return
	}

	embeds := []*discordgo.MessageEmbed{
		{
			Title: selectedMap + " Spawns",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.imgur.com/hN7W003.jpeg",
			},
		},
	}

	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Title:  selectedMap + " Spawns",
			Embeds: embeds,
		},
	})
	if err != nil {
		utils.LogDiscordError("SendSpawns.InteractionRespond", err.Error())
	}
}
