package imdb

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	//"sort"

	"github.com/botlabs-gg/yagpdb/v2/bot/paginatedmessages"
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
)

var Command = &commands.YAGCommand{
	CmdCategory:         commands.CategoryFun,
	Name:                "IMDB",
	RequiredArgs:        1,
	Description:         "Title search of IMDB.com.",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	Arguments: []*dcmd.ArgDef{
		{Name: "Title", Type: dcmd.String},
	},
	ArgSwitches: []*dcmd.ArgDef{
		{Name: "c", Help: "Compact output"},
		{Name: "p", Help: "Paginated output"},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		var compactView, paginatedView bool
		imdbtitle := data.Args[0].Str()
		if data.Switches["c"].Value != nil && data.Switches["c"].Value.(bool) {
			compactView = true
		}
		if data.Switches["p"].Value != nil && data.Switches["p"].Value.(bool) {
			compactView = false
			paginatedView = true
		}
		titles, err := getTitleData(imdbtitle)
		if err != nil {
			return nil, err
		}
		if len(titles.Results) == 0 {
			return "No results", nil
		}
		if len(titles.ErrorMessage) > 0 {
			return titles.ErrorMessage, nil
		}
		if compactView {
			compactData := fmt.Sprintf("%s: %s", titles.Results[0].Title,
				"https://imdb.com/title/"+titles.Results[0].ID,
			)
			return compactData, nil
		}
		info, err := getInfo(titles.Results[0].ID)
		if err != nil {
			return nil, err
		}
		imdbEmbed := embedCreator(*info)
		if paginatedView {
			_, err := paginatedmessages.CreatePaginatedMessage(
				data.GuildData.GS.ID, data.ChannelID, 1, len(titles.Results), func(p *paginatedmessages.PaginatedMessage, page int) (*discordgo.MessageEmbed, error) {
					i := page - 1
					info, err := getInfo(titles.Results[i].ID)
					if err != nil {
						return nil, err
					}
					paginatedEmbed := embedCreator(*info)
					return paginatedEmbed, nil
				})
			if err != nil {
				return "Something went wrong", nil
			}
		} else {
			return imdbEmbed, nil
		}
		return nil, nil
	},
}

func embedCreator(results Info) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: results.FullTitle,
			URL:  "https://imdb.com/title/" + results.ID,
		},
		Color: int(rand.Int63n(16777215)),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: results.Image,
		},
		Description: fmt.Sprintf("Release Date : %s\nLanguages : %s\nIMDB Ratings : %s\nIMDB Ratings Votes: %s\nPlot: %s", results.ReleaseDate, results.Languages, results.ImDbRating, results.ImDbRatingVotes, results.Plot),
	}
	return embed
}
func getTitleData(imdbtitle string) (*Response, error) {
	url := "https://imdb-api.com/en/API/SearchTitle/k_8r16tz55/" + imdbtitle
	body, err := handleAPI(url)
	if err != nil {
		return nil, err
	}
	result := &Response{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func getInfo(id string) (*Info, error) {
	url := "https://imdb-api.com/en/API/Title/k_8r16tz55/" + id
	body, err := handleAPI(url)
	if err != nil {
		return nil, err
	}
	result := &Info{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}
func handleAPI(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, commands.NewPublicError("HTTP err: ", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type Response struct {
	SearchType   string    `json:"searchType"`
	Expression   string    `json:"expression"`
	Results      []Results `json:"results"`
	ErrorMessage string    `json:"errorMessage"`
}
type Results struct {
	ID          string `json:"id"`
	ResultType  string `json:"resultType"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Info struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	FullTitle       string `json:"fullTitle"`
	Image           string `json:"image"`
	ReleaseDate     string `json:"releaseDate"`
	Plot            string `json:"plot"`
	Languages       string `json:"languages"`
	ImDbRating      string `json:"imDbRating"`
	ImDbRatingVotes string `json:"imDbRatingVotes"`
}
