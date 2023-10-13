package odds

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
	Name:        "odds",
	// Aliases:     []string{"fbscore", "yobitchwhatsthescoreofthefootballgame"},
	Description: "Betting odds",
	Arguments: []*dcmd.ArgDef{
		{Name: "Sport", Type: dcmd.String},
	},
	SlashCommandEnabled: true,
	DefaultEnabled:      true,
	RequiredArgs:        1,
	//Options: ,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		addrTeam := "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/" + data.Args[0].Str()
		// var to build out request from cmd args and api endpoint

		output, err := apiTeamSearch(addrTeam)
		if err != nil {
			return nil, err
		}
		// grabs list of teams in json

		eventID := output.Team.NextEvent[0].Id // var ID of searched team's next event

		addrScore := "http://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard/" + eventID
		// var to build out eventid and api endpoint

		score, err := apiScoreSearch(addrScore)
		if err != nil {
			return nil, err
		}
		// gets information of the "next event" or "current event"

		color, err := strconv.ParseInt(output.Team.Color, 16, 64)
		if err != nil {
			return nil, err
		}
		// gets searched team's color for the embed

		broadcast := "not available"
		if !(len(output.Team.NextEvent[0].Competitions[0].Broadcasts) == 0) {
			broadcast = output.Team.NextEvent[0].Competitions[0].Broadcasts[0].Media.ShortName
		}
		// this builds an embed for an upcoming game with no "live event", it just shows the next game with a start time.

		embed := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Game: %s", output.Team.NextEvent[0].Name),
			Description: fmt.Sprintf("TV: %s\n %s\n %s %s - %s %s ", broadcast, output.Team.NextEvent[0].Competitions[0].Status.Type.ShortDetail, score.Competitions[0].Competitors[0].Team.Name, score.Competitions[0].Competitors[0].Score, score.Competitions[0].Competitors[1].Team.Name, score.Competitions[0].Competitors[1].Score),
			Color:       int(color),
		}
		// this builds the embed for a game that is currently live. It includes the score.

		return embed, nil
	},
}

func apiTeamSearch(addrTeam string) (*Output, error) { // this is the team search function
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

func apiScoreSearch(addrScore string) (*Score, error) { // this is the score search function
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
	ID           string `json:"id"`
	SportKey     string `json:"sport_key"`
	SportTitle   string `json:"sport_title"`
	CommenceTime string `json:"commence_time"`
	HomeTeam     string `json:"home_team"`
	AwayTeam     string `json:"away_team"`
	Bookmakers   []struct {
		Key        string `json:"key"`
		Title      string `json:"title"`
		LastUpdate string `json:"last_update"`
		Markets    []struct {
			Key        string `json:"key"`
			LastUpdate string `json:"last_update"`
			Outcomes   []struct {
				Name  string  `json:"name"`
				Price float64 `json:"price"`
				Point float64 `json:"point"`
			} `json:"outcomes"`
		} `json:"markets"`
	} `json:"bookmakers"`
}
