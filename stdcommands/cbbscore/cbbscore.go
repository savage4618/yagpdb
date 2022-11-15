package cbbscore

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

type NextEvent struct {
	Name         string         `json:"name"`
	Competitions []Competitions `json:"competitions"`
}

type Competitions struct {
	Notes      []Notes      `json:"notes"`
	Broadcasts []Broadcasts `json:"broadcasts"`
	Status     Status       `json:"status"`
}

type Notes struct {
	Headline string `json:"headline"`
}

type Broadcasts struct {
	Media Media `json:"media"`
}

type Media struct {
	ShortName string `json:"shortName"`
}

type Status struct {
	Clock  float64 `json:"clock"`
	Period int     `json:"period"`
	Type   Type    `json:"type"`
}

type Type struct {
	Description string `json:"description"`
	ShortDetail string `json:"shortDetail"`
}

var Command = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryFun,
	Name:        "cbbscore",
	Description: "Grabs College Basketball score from ESPN API",
	Arguments: []*dcmd.ArgDef{
		{Name: "Team", Type: dcmd.String},
	},
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RequiredArgs:        1,

	RunFunc: func(data *dcmd.Data) (interface{}, error) {

		addr := "https://site.api.espn.com/apis/site/v2/sports/basketball/mens-college-basketball/teams/" + url.QueryEscape(data.Args[0].Str())

		resp, err := http.Get(addr)
		if err != nil {
			return nil, err
		}

		msg := fmt.Printf("Game: %s TV: %s\n Scheduled for: %s",
			resp.Body(NextEvent).Name,
			resp.Body(Broadcasts).Media,
			resp.Body(Type).ShortDetail,
		)
	},
}
