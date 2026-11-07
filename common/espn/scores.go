// Package espn selects a team's current ESPN game independently of nextEvent.
package espn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 15 * time.Second}

type Team struct {
	ID               string `json:"id"`
	ShortDisplayName string `json:"shortDisplayName"`
	DisplayName      string `json:"displayName"`
	Color            string `json:"color"`
	Logos            []struct {
		Href string `json:"href"`
	} `json:"logos"`
	NextEvent []struct {
		ID string `json:"id"`
	} `json:"nextEvent"`
}

type Status struct {
	Type struct {
		State       string `json:"state"`
		Completed   bool   `json:"completed"`
		ShortDetail string `json:"shortDetail"`
	} `json:"type"`
}

type Event struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Date         string        `json:"date"`
	Status       Status        `json:"status"`
	Competitions []Competition `json:"competitions"`
}

type Competition struct {
	Status      Status `json:"status"`
	Competitors []struct {
		Team  Team   `json:"team"`
		Score string `json:"score"`
	} `json:"competitors"`
	Broadcasts []struct {
		Names []string `json:"names"`
		Media struct {
			ShortName string `json:"shortName"`
		} `json:"media"`
	} `json:"broadcasts"`
}

// Game contains only the fields needed to render a score command's embed.
type Game struct {
	Name, Logo, Broadcast, Status string
	Color                         int
	Teams, Scores                 [2]string
}

func Fetch(league, query string) (*Game, error) {
	return fetch(client, "https://site.api.espn.com/apis/site/v2/sports/"+league, query)
}

func fetch(client *http.Client, baseURL, query string) (*Game, error) {
	var output struct {
		Team Team `json:"team"`
	}
	if err := getJSON(client, baseURL+"/teams/"+url.PathEscape(strings.ToLower(strings.TrimSpace(query))), &output); err != nil {
		return nil, err
	}
	if output.Team.ID == "" {
		return nil, fmt.Errorf("Could not find that team. Use the team autocomplete or an ESPN team ID.")
	}

	// Use ESPN's current slate. Include all Division I teams (FBS and FCS
	// for football), rather than only ranked teams. ESPN accepts 300 here;
	// larger limits can silently revert to the default of 25 football games.
	addr := baseURL + "/scoreboard?limit=300"
	if strings.HasSuffix(baseURL, "/mens-college-basketball") {
		addr += "&groups=50"
	} else if strings.HasSuffix(baseURL, "/college-football") {
		addr += "&groups=90"
	}
	var scoreboard struct {
		Events []Event `json:"events"`
	}
	if err := getJSON(client, addr, &scoreboard); err != nil {
		return nil, err
	}
	event := selectEvent(scoreboard.Events, output.Team.ID)
	if event == nil {
		// nextEvent remains useful when the team has no game on the current
		// slate, but must never take precedence over a live or completed game.
		if len(output.Team.NextEvent) == 0 || output.Team.NextEvent[0].ID == "" {
			return nil, fmt.Errorf("No current or upcoming games found for that team. The season may be over or not yet scheduled.")
		}
		event = &Event{}
		if err := getJSON(client, baseURL+"/scoreboard/"+url.PathEscape(output.Team.NextEvent[0].ID), event); err != nil {
			return nil, err
		}
	}
	return makeGame(output.Team, *event)
}

func getJSON(client *http.Client, addr string, output interface{}) error {
	resp, err := client.Get(addr)
	if err != nil {
		return fmt.Errorf("Could not reach ESPN: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("ESPN could not find that team or game. Use the team autocomplete or find your ESPN team ID here: <https://docs.google.com/spreadsheets/d/1QXwnLD-OPOoDkoqNxx9iSI_Nc1WgbTVYQueyx6UPPWY/edit?usp=sharing>. There are separate tabs for college, NBA, and NFL teams.")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ESPN returned HTTP %d. Please try again later.", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return fmt.Errorf("Could not read ESPN's response: %w", err)
	}
	return nil
}

func eventStatus(event Event) Status {
	if len(event.Competitions) > 0 && event.Competitions[0].Status.Type.State != "" {
		return event.Competitions[0].Status
	}
	return event.Status
}

func priority(event Event) int {
	status := eventStatus(event)
	if status.Type.State == "in" {
		return 3
	}
	if status.Type.Completed {
		return 2
	}
	return 1
}

func selectEvent(events []Event, teamID string) *Event {
	var selected *Event
	for i := range events {
		event := &events[i]
		if len(event.Competitions) == 0 || len(event.Competitions[0].Competitors) < 2 {
			continue
		}
		matches := false
		for _, competitor := range event.Competitions[0].Competitors {
			if competitor.Team.ID == teamID {
				matches = true
				break
			}
		}
		if !matches {
			continue
		}
		if selected == nil || priority(*event) > priority(*selected) {
			selected = event
		} else if priority(*event) == priority(*selected) {
			// ISO dates sort chronologically: latest result, earliest tipoff.
			if (priority(*event) == 1 && event.Date < selected.Date) || (priority(*event) > 1 && event.Date > selected.Date) {
				selected = event
			}
		}
	}
	return selected
}

func makeGame(team Team, event Event) (*Game, error) {
	if len(event.Competitions) == 0 || len(event.Competitions[0].Competitors) < 2 {
		return nil, fmt.Errorf("ESPN has not provided matchup details for that game yet.")
	}
	competition := event.Competitions[0]
	game := &Game{Name: event.Name, Broadcast: "not available", Status: eventStatus(event).Type.ShortDetail}
	color, _ := strconv.ParseInt(team.Color, 16, 32)
	game.Color = int(color)
	if len(team.Logos) > 0 {
		game.Logo = team.Logos[0].Href
	}
	var broadcasts []string
	for _, broadcast := range competition.Broadcasts {
		if len(broadcast.Names) > 0 {
			broadcasts = append(broadcasts, broadcast.Names...)
		} else if broadcast.Media.ShortName != "" {
			broadcasts = append(broadcasts, broadcast.Media.ShortName)
		}
	}
	if len(broadcasts) > 0 {
		game.Broadcast = strings.Join(broadcasts, ", ")
	}
	for i := range game.Teams {
		competitor := competition.Competitors[i]
		game.Teams[i] = competitor.Team.ShortDisplayName
		if game.Teams[i] == "" {
			game.Teams[i] = competitor.Team.DisplayName
		}
		game.Scores[i] = competitor.Score
		if game.Scores[i] == "" {
			game.Scores[i] = "—"
		}
	}
	return game, nil
}
