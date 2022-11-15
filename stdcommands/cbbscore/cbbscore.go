// package cbbscore

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"

// 	"github.com/botlabs-gg/yagpdb/v2/commands"
// 	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
// )

// // type NextEvent struct {
// // 	Name         string         `json:"name"`
// // 	Competitions []Competitions `json:"competitions"`
// // }

// // type Competitions struct {
// // 	Notes      []Notes      `json:"notes"`
// // 	Broadcasts []Broadcasts `json:"broadcasts"`
// // 	Status     Status       `json:"status"`
// // }

// // type Notes struct {
// // 	Headline string `json:"headline"`
// // }

// // type Broadcasts struct {
// // 	Media Media `json:"media"`
// // }

// // type Media struct {
// // 	ShortName string `json:"shortName"`
// // }

// // type Status struct {
// // 	Clock  float64 `json:"clock"`
// // 	Period int     `json:"period"`
// // 	Type   Type    `json:"type"`
// // }

// // type Type struct {
// // 	Description string `json:"description"`
// // 	ShortDetail string `json:"shortDetail"`
// // }

// type Output struct {
// 	Team struct {
// 		NextEvent []struct {
// 			Name         string `json:"name"`
// 			Competitions []struct {
// 				Broadcasts []struct {
// 					Media struct {
// 						ShortName string `json:"shortName"`
// 					} `json:"media"`
// 				} `json:"broadcasts"`
// 				Status struct {
// 					Type struct {
// 						ShortDetail string `json:"shortDetail"`
// 					} `json:"type"`
// 				} `json:"status"`
// 			} `json:"competitions"`
// 		} `json:"nextEvent"`
// 	} `json:"team"`
// }

// var Command = &commands.YAGCommand{
// 	Cooldown:    5,
// 	CmdCategory: commands.CategoryFun,
// 	Name:        "cbbscore",
// 	Description: "Grabs College Basketball score from ESPN API",
// 	Arguments: []*dcmd.ArgDef{
// 		{Name: "Team", Type: dcmd.String},
// 	},
// 	DefaultEnabled:      true,
// 	SlashCommandEnabled: true,
// 	RequiredArgs:        1,

// 	RunFunc: func(data *dcmd.Data) (interface{}, error) {

// 		addr := "https://site.api.espn.com/apis/site/v2/sports/basketball/mens-college-basketball/teams/" + url.QueryEscape(data.Args[0].Str())

// 		resp, err := http.Get(addr)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer resp.Body.Close()
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}
// 		output := &Output{}
// 		err = json.Unmarshall([]byte(body), &output)
// 		if err != nil {
// 			return
// 		}
// 		msg := fmt.Printf("%+v\n", output)
// 		fmt.Printf("%s\n", output.Team.NextEvent[0].Name)
// 		fmt.Printf("%s\n", output.Team.NextEvent[0].Competitions[0].Broadcasts[0].Media.ShortName)
// 		fmt.Printf("%s\n", output.Team.NextEvent[0].Competitions[0].Status.Type.ShortDetail)
// 	},
// }

package cbbscore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
)

var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryFun,
	Name:        "cbbscore",
	Description: "College Basketball Scores",
	Arguments: []*dcmd.ArgDef{
		{Name: "Team", Type: dcmd.String},
	},
	SlashCommandEnabled: true,
	DefaultEnabled:      true,
	RequiredArgs:        1,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		addr := "https://site.api.espn.com/apis/site/v2/sports/basketball/mens-college-basketball/teams/" + data.Args[0].Str()

		output, err := apiSearch(addr)
		if err != nil {
			return nil, err
		}
		color, err := strconv.ParseInt(output.Team.Color, 16, 16)
		if err != nil {
			return nil, err
		}

		embed := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Game: %s", output.Team.NextEvent[0].Name),
			Description: fmt.Sprintf("Game: %s\n TV: %s\n Scheduled for: %s", output.Team.NextEvent[0].Name, output.Team.NextEvent[0].Competitions[0].Broadcasts[0].Media.ShortName, output.Team.NextEvent[0].Competitions[0].Status.Type.ShortDetail),
			Color:       int(color),
		}
		return embed, nil
	},
}

func apiSearch(addr string) (*Output, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	output := &Output{}
	err = json.Unmarshal([]byte(body), &output)
	if err != nil {
		return nil, err
	}
	return output, nil

}

type Output struct {
	Team struct {
		Color     string `json:"color"`
		NextEvent []struct {
			Name         string `json:"name"`
			Competitions []struct {
				Broadcasts []struct {
					Media struct {
						ShortName string `json:"shortName"`
					} `json:"media"`
				} `json:"broadcasts"`
				Status struct {
					Type struct {
						ShortDetail string `json:"shortDetail"`
					} `json:"type"`
				} `json:"status"`
			} `json:"competitions"`
		} `json:"nextEvent"`
	} `json:"team"`
}
