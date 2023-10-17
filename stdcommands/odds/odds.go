package odds

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/botlabs-gg/yagpdb/v2/bot/paginatedmessages"
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common/config"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"github.com/sirupsen/logrus"
)

var confAPIKey = config.RegisterOption("yagpdb.oddsapikey", "API key for querying sports odds.", "")
var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryFun,
	Name:        "odds",
	Aliases:     []string{"betting"},
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
		apiKey := confAPIKey.GetString()
		apiEndpoint := "https://api.the-odds-api.com/v4/sports/upcoming/odds/?regions=us&markets=h2h,spreads&oddsFormat=american&bookmakers=draftkings&sport=" + url.QueryEscape(query) + "&apiKey=" + url.QueryEscape(apiKey)
		// var to build out request from cmd args and api endpoint
		req, err := http.NewRequest("GET", apiEndpoint, nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 {
			return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode, " Too many requests in a short time. Slow down.")
		} else if resp.StatusCode == 404 {
			return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode, " You probably entered an invalid sport.")
		} else if resp.StatusCode != 200 {
			return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var res []OddsResponse
		err = json.Unmarshal(body, &res)
		if err != nil || len(res[0].SportTitle) == 0 {
			logrus.WithError(err).Error("Failed getting response from Odds API")
			return "No Odds found.", err
		}

		var odd = &res[0]
		if len(odd.Bookmakers) == 1 || data.Context().Value(paginatedmessages.CtxKeyNoPagination) != nil {
			return createOddsEmbed(odd, &odd.Bookmakers[0]), nil
		}

		_, err = paginatedmessages.CreatePaginatedMessage(data.GuildData.GS.ID, data.ChannelID, 1, len(odd.ID), func(p *paginatedmessages.PaginatedMessage, page int) (*discordgo.MessageEmbed, error) {
			if page > len(odd.ID) {
				return nil, paginatedmessages.ErrNoResults
			}

			return createOddsEmbed(odd, &odd.Bookmakers[page-1]), nil
		})

		return nil, err
	},
}

func createOddsEmbed(data *dcmd.Data, res []OddsResponse, bm []Bookmaker, mkt []Market, otcm []Outcome, i int) *discordgo.MessageEmbed {
	title := "Odds for upcoming games in " + data.Args[0].Str()

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: res[i].HomeTeam + " vs " + res[i].AwayTeam,
		Color:       0x53d337,
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer:      &discordgo.MessageEmbedFooter{Text: "Odds updated :" + bm[i].LastUpdate},
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: "https://randomdad.xyz/SB_Green_Icon.png"},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "DraftKings Sportsbook Odds",
			URL:     "https://sportsbook.draftkings.com/",
			IconURL: "https://randomdad.xyz/SB_Green_Icon.png"},
	}

	return embed
}

type Bookmaker struct {
	Key        string   `json:"key"`
	Title      string   `json:"title"`
	LastUpdate string   `json:"last_update"`
	Markets    []Market `json:"markets"`
}

type Market struct {
	Key        string    `json:"key"`
	LastUpdate string    `json:"last_update"`
	Outcomes   []Outcome `json:"outcomes"`
}

type Outcome struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Point float64 `json:"point"`
}

type OddsResponse struct {
	ID           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime string      `json:"commence_time"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Bookmakers   []Bookmaker `json:"bookmakers"`
}
