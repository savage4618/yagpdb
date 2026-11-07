package nflscore

import (
	"fmt"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common/espn"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"github.com/botlabs-gg/yagpdb/v2/stdcommands/sportsteams"
)

var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryFun,
	Name:        "nflscore",
	Aliases:     []string{"pfbscore", "yobitchwhatsthescoreoftheprofessionalfootballgame"},
	Description: "NFL Football Scores",
	Arguments: []*dcmd.ArgDef{
		{Name: "Team", Type: dcmd.String, Help: "Team name", AutocompleteFunc: sportsteams.NFLAutocomplete},
	},
	SlashCommandEnabled: true,
	DefaultEnabled:      true,
	RequiredArgs:        1,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		game, err := espn.Fetch("football/nfl", data.Args[0].Str())
		if err != nil {
			return nil, commands.NewPublicError(err.Error())
		}
		embed := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Game: %s", game.Name),
			Description: fmt.Sprintf("TV: %s\n %s\n", game.Broadcast, game.Status),
			Color:       game.Color,
			Fields: []*discordgo.MessageEmbedField{
				{Name: game.Teams[0], Value: game.Scores[0], Inline: true},
				{Name: game.Teams[1], Value: game.Scores[1], Inline: true},
			},
		}
		if game.Logo != "" {
			embed.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: game.Logo}
		}
		return embed, nil
	},
}
