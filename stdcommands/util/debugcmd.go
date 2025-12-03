package util

import (
	"math/rand"
	"time"

	"github.com/botlabs-gg/yagpdb/v2/bot"
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

func isExecedByCC(data *dcmd.Data) bool {
	if v := data.Context().Value(commands.CtxKeyExecutedByCC); v != nil {
		if cast, _ := v.(bool); cast {
			return true
		}
	}

	return false
}

func RequireBotAdmin(inner dcmd.RunFunc) dcmd.RunFunc {
	rand.Seed(time.Now().UnixNano())

	adminresponses := []string{
		"No",
		"How dare you.",
		"Absolutely not.",
		"Who the fuck are you?",
		"This command only works for non-ducks",
		"You are not a bot admin. This incident will be reported.",
	}

	return func(data *dcmd.Data) (interface{}, error) {
		if isExecedByCC(data) {
			return "", nil
		}

		if admin, err := bot.IsBotAdmin(data.Author.ID); admin && err == nil {
			return inner(data)
		}

		adminrandomIndex := rand.Intn(len(adminresponses))
		return adminresponses[adminrandomIndex], nil
	}
}

func RequireOwner(inner dcmd.RunFunc) dcmd.RunFunc {
	rand.Seed(time.Now().UnixNano())

	ownerresponses := []string{
		"Nah",
		"You aren't my owner.",
		"You don't smell like my owner.",
		"No, fuck off duck.",
		"This command only works for non-ducks",
		"You are not the bot owner. I'm tellin'.",
	}
	return func(data *dcmd.Data) (interface{}, error) {
		if isExecedByCC(data) {
			return " ", nil
		}

		if common.IsOwner(data.Author.ID) {
			return inner(data)
		}

		ownerrandomIndex := rand.Intn(len(ownerresponses))
		return ownerresponses[ownerrandomIndex], nil
	}
}
