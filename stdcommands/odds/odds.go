package odds

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common/config"
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
		query := strings.ToLower(data.Args[0].Str())
		confAPIKey := config.RegisterOption("yagpdb.oddsapikey", "API key for querying sports odds.", "")
		apiKey := confAPIKey.GetString()
		apiEndpoint := "https://api.the-odds-api.com/v4/sports/upcoming/odds/?regions=us&markets=h2h,spreads&oddsFormat=american&bookmakers=draftkings&sport=" + url.QueryEscape(query) + "&apiKey=" + url.QueryEscape(apiKey)
		// var to build out request from cmd args and api endpoint

		output, err := getOdds(apiEndpoint)
		if err != nil {
			return nil, err
		}
		// gets odds

		broadcast := "not available"
		if !(len(output.Team.NextEvent[0].Competitions[0].Broadcasts) == 0) {
			broadcast = output.Team.NextEvent[0].Competitions[0].Broadcasts[0].Media.ShortName
		}
		// this builds an embed for an upcoming game with no "live event", it just shows the next game with a start time.

		embed := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Game: %s", output.Team.NextEvent[0].Name),
			Description: fmt.Sprintf("TV: %s\n %s\n %s %s - %s %s ", broadcast, output.Team.NextEvent[0].Competitions[0].Status.Type.ShortDetail, score.Competitions[0].Competitors[0].Team.Name, score.Competitions[0].Competitors[0].Score, score.Competitions[0].Competitors[1].Team.Name, score.Competitions[0].Competitors[1].Score),
			Color:       int(5493559),
		}
		// this builds the embed for a game that is currently live. It includes the score.

		return embed, nil
	},
}

func getOdds(apiEndpoint string) (*Output, error) { // this is gets the odds of all current and upcoming games
	resp, err := http.Get(apiEndpoint)
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
