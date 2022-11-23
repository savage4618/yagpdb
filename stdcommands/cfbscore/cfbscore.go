package cfbscore

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
	Name:        "cfbscore",
	Aliases:     []string{"fbscore", "yobitchwhatsthescoreofthefootballgame"},
	Description: "College Football Scores",
	Arguments: []*dcmd.ArgDef{
		{Name: "Team", Type: dcmd.String},
	},
	SlashCommandEnabled: true,
	DefaultEnabled:      true,
	RequiredArgs:        1,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		addrTeam := "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/" + data.Args[0].Str()

		output, err := apiTeamSearch(addrTeam)
		if err != nil {
			return nil, err
		}
		eventID := output.Team.NextEvent[0].Id
		addrScore := "http://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard/" + eventID

		score, err := apiScoreSearch(addrScore)
		if err != nil {
			return nil, err
		}

		color, err := strconv.ParseInt(output.Team.Color, 16, 64)
		if err != nil {
			return nil, err
		}

		broadcast := "not available"
		if !(len(output.Team.NextEvent[0].Competitions[0].Broadcasts) == 0) {
			broadcast = output.Team.NextEvent[0].Competitions[0].Broadcasts[0].Media.ShortName
		}

		embed := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Game: %s", output.Team.NextEvent[0].Name),
			Description: fmt.Sprintf("TV: %s\n %s\n %s %s - %s %s ", broadcast, output.Team.NextEvent[0].Competitions[0].Status.Type.ShortDetail, score.Competitions[0].Competitors[0].Team.Name, score.Competitions[0].Competitors[0].Score, score.Competitions[0].Competitors[1].Team.Name, score.Competitions[0].Competitors[1].Score),
			Color:       int(color),
		}

		return embed, nil
	},
}

func apiTeamSearch(addrTeam string) (*Output, error) {
	resp, err := http.Get(addrTeam)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 502 {
		return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode, " You probably entered a team with a space in the name. Check this link for a Team ID: <https://docs.google.com/spreadsheets/d/1QXwnLD-OPOoDkoqNxx9iSI_Nc1WgbTVYQueyx6UPPWY/edit?usp=sharing>")
	} else if resp.StatusCode == 400 {
		return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode, " You have entered an invalid team. Please refer to this list to grab your Team ID: <https://docs.google.com/spreadsheets/d/1QXwnLD-OPOoDkoqNxx9iSI_Nc1WgbTVYQueyx6UPPWY/edit?usp=sharing>")
	} else if resp.StatusCode != 200 {
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

func apiScoreSearch(addrScore string) (*Score, error) {
	resp, err := http.Get(addrScore)
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

	score := &Score{}
	err = json.Unmarshal([]byte(body), &score)
	if err != nil {
		return nil, err
	}
	return score, nil

}

type Output struct {
	Team struct {
		Color string `json:"color"`
		Logos []struct {
			Href string `json:"href"`
		} `json:"Logos"`
		NextEvent []struct {
			Id           string `json:"id"`
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
						Description string `json:"description"`
					} `json:"type"`
				} `json:"status"`
			} `json:"competitions"`
		} `json:"nextEvent"`
	} `json:"team"`
}

type Score struct {
	Id           string `json:"id"`
	ShortName    string `json:"shortname"`
	Competitions []struct {
		Status struct {
			Displayclock string `json:"displayClock"`
			Period       int    `json:"period"`
			Type         struct {
				Completed bool `json:"completed"`
			} `json:"Type"`
		} `json:"Status"`
		Competitors []struct {
			Team struct {
				Name string `json:"shortdisplayname"`
			} `json:"team"`
			Score string `json:"score"`
		} `json:"competitors"`
	} `json:"competitions"`
}
