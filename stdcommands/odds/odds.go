package odds

import (
	"encoding/json"
	"html"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common/config"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
)

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
		confAPIKey := config.RegisterOption("yagpdb.oddsapikey", "API key for querying sports odds.", "")
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
		err = json.Unmarhsall(body, &res)
		if err != nil || len(res[0].SportTitle) == 0 {
			logrus.WithError(err).Error("Failed getting response from Odds API")
			return "No Odds found.", err
		}

		var odd = &res[0]
		if len(odd.ID) == 1 || data.Context().Value(paginatedMessages.CtxKeyNoPagination) != nil {
			return createOddsEmbed(odd, &Odd.ID[0]), nil
		}

		_, err = paginatedMessages.CreatePaginatedMessage(data.GuildData.GS.ID, data.ChannelID, 1, len(odd.ID), func(p *paginatedMessages.PaginatedMessage, page int) (*discordgo.MessageEmbed, error) {
			if page > len(odd.ID) {
				return nil, paginatedMessages.ErrNoResults
			}

			return createOddsEmbed(odd, &odd.ID[page-1]), nil
		})

		return nil, err
	},
}

func createOddsEmbed(res *OddsResponse, outcomes *Outcomes) *discordgo.MessageEmbed {
	title := res.HomeTeam + " vs " + res.AwayTeam

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: "testing",
		Color:       0x53d337,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	return embed
}

var policy = bluemonday.StrictPolicy()

func normalizeOutput(s string) string {
	// The API occasionally returns HTML tags and escapes as part of output, remove them.
	decoded := html.UnescapeString(policy.Sanitize(s))
	// It also sometimes returns non-printable characters, strip them out too.
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, decoded)
}

type Bookmakers struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	LastUpdate string    `json:"last_update"`
	Markets    []Markets `json:"markets"`
}

type Markets struct {
	Key        string     `json:"key"`
	LastUpdate string     `json:"last_update"`
	Outcomes   []Outcomes `json:"outcomes"`
}

type Outcomes struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Point float64 `json:"point"`
}

type OddsResponse struct {
	ID           string       `json:"id"`
	SportKey     string       `json:"sport_key"`
	SportTitle   string       `json:"sport_title"`
	CommenceTime string       `json:"commence_time"`
	HomeTeam     string       `json:"home_team"`
	AwayTeam     string       `json:"away_team"`
	Bookmakers   []Bookmakers `json:"bookmakers"`
}
