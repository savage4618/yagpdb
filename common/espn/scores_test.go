package espn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func testEvent(id, date, state string, completed bool) Event {
	var event Event
	data := fmt.Sprintf(`{"id":%q,"date":%q,"name":"Duke at UNC","competitions":[{"status":{"type":{"state":%q,"completed":%t,"shortDetail":"Final"}},"competitors":[{"team":{"id":"150","shortDisplayName":"Duke"},"score":"82"},{"team":{"id":"153","shortDisplayName":"UNC"},"score":"80"}],"broadcasts":[{"names":["ESPN"]}]}]}`, id, date, state, completed)
	if err := json.Unmarshal([]byte(data), &event); err != nil {
		panic(err)
	}
	return event
}

func jsonString(value interface{}) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestSelectEvent(t *testing.T) {
	live := testEvent("live", "2026-11-06T20:00Z", "in", false)
	final := testEvent("final", "2026-11-06T18:00Z", "post", true)
	older := testEvent("older", "2026-11-06T15:00Z", "post", true)
	upcoming := testEvent("upcoming", "2026-11-06T22:00Z", "pre", false)
	later := testEvent("later", "2026-11-07T22:00Z", "pre", false)
	for _, tc := range []struct {
		name   string
		events []Event
		want   string
	}{
		{"live before final and upcoming", []Event{upcoming, final, live}, "live"},
		{"live order independent", []Event{live, upcoming, final}, "live"},
		{"completed before upcoming", []Event{upcoming, final}, "final"},
		{"latest completed", []Event{final, older}, "final"},
		{"earliest upcoming", []Event{later, upcoming}, "upcoming"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := selectEvent(tc.events, "150")
			if got == nil || got.ID != tc.want {
				t.Fatalf("selected %+v, want %s", got, tc.want)
			}
		})
	}
	if got := selectEvent([]Event{live, {}}, "999"); got != nil {
		t.Fatalf("selected another team's game: %+v", got)
	}
	// ESPN can put status at event level instead of competition level.
	live.Status = live.Competitions[0].Status
	live.Competitions[0].Status = Status{}
	if got := selectEvent([]Event{final, live}, "150"); got.ID != "live" {
		t.Fatal("ignored event-level live status")
	}
}

func TestFetch(t *testing.T) {
	for _, league := range []string{"basketball/mens-college-basketball", "football/college-football", "basketball/nba", "football/nfl"} {
		for _, current := range []bool{true, false} {
			t.Run(fmt.Sprintf("%s/current=%t", league, current), func(t *testing.T) {
				fallbackCalls := 0
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					switch r.URL.Path {
					case "/" + league + "/teams/duke":
						// Team responses may use score objects, unlike the scoreboard's
						// strings. Only the fallback event ID is needed from this data.
						fmt.Fprint(w, `{"team":{"id":"150","color":"003087","nextEvent":[{"id":"next","competitions":[{"competitors":[{"score":{"value":33,"displayValue":"33"}}]}]}]}}`)
					case "/" + league + "/scoreboard":
						group := ""
						if strings.Contains(league, "mens-college") {
							group = "50"
						} else if strings.Contains(league, "college-football") {
							group = "90"
						}
						if r.URL.Query().Get("limit") != "300" || r.URL.Query().Get("groups") != group {
							t.Error("missing full-slate query parameters")
						}
						if current {
							fmt.Fprint(w, `{"events":[`+jsonString(testEvent("current", "2026-11-06T20:00Z", "in", false))+`]}`)
						} else {
							fmt.Fprint(w, `{"events":[]}`)
						}
					case "/" + league + "/scoreboard/next":
						fallbackCalls++
						fmt.Fprint(w, jsonString(testEvent("next", "2026-11-07T20:00Z", "pre", false)))
					default:
						t.Errorf("unexpected request: %s", r.URL)
						http.NotFound(w, r)
					}
				}))
				defer server.Close()
				game, err := fetch(server.Client(), server.URL+"/"+league, "DUKE")
				if err != nil {
					t.Fatal(err)
				}
				if game.Teams != [2]string{"Duke", "UNC"} || game.Scores != [2]string{"82", "80"} || game.Broadcast != "ESPN" || game.Color != 0x003087 || game.Status != "Final" {
					t.Fatalf("incorrect embed data: %+v", game)
				}
				if (current && fallbackCalls != 0) || (!current && fallbackCalls != 1) {
					t.Fatalf("unexpected nextEvent lookups: %d", fallbackCalls)
				}
			})
		}
	}
}

func TestMissingData(t *testing.T) {
	if _, err := makeGame(Team{}, Event{}); err == nil {
		t.Fatal("expected an error for missing competitions")
	}
	event := testEvent("next", "", "pre", false)
	event.Competitions[0].Broadcasts = nil
	event.Competitions[0].Competitors[0].Score = ""
	event.Competitions[0].Competitors[0].Team.ShortDisplayName = ""
	event.Competitions[0].Competitors[0].Team.DisplayName = "Duke Blue Devils"
	game, err := makeGame(Team{Color: "invalid"}, event)
	if err != nil || game.Logo != "" || game.Color != 0 || game.Broadcast != "not available" || game.Scores[0] != "—" || game.Teams[0] != "Duke Blue Devils" {
		t.Fatalf("missing optional data not handled: %+v, %v", game, err)
	}
	if err := json.Unmarshal([]byte(`[{"media":{"shortName":"ABC"}}]`), &event.Competitions[0].Broadcasts); err != nil {
		t.Fatal(err)
	}
	game, err = makeGame(Team{}, event)
	if err != nil || game.Broadcast != "ABC" {
		t.Fatalf("media broadcast not handled: %+v, %v", game, err)
	}
}

func TestFetchErrors(t *testing.T) {
	for _, body := range []string{`{"team":{"id":"150"}}`, `{}`, `not json`} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, body)
		}))
		_, err := fetch(server.Client(), server.URL, "150")
		server.Close()
		if err == nil {
			t.Fatalf("expected an error for %q", body)
		}
	}
	for _, code := range []int{400, 404, 500, 502} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(code) }))
		_, err := fetch(server.Client(), server.URL, "150")
		server.Close()
		if err == nil {
			t.Fatalf("expected an error for HTTP %d", code)
		}
	}
}

// Opt in to checking the actual endpoints without making routine tests depend on ESPN.
func TestLiveESPN(t *testing.T) {
	if os.Getenv("ESPN_LIVE_TESTS") != "1" {
		t.Skip("set ESPN_LIVE_TESTS=1 to check ESPN")
	}
	for _, tc := range []struct{ league, team string }{
		{"basketball/mens-college-basketball", "150"},
		{"basketball/mens-college-basketball", "96"},
		{"basketball/mens-college-basketball", "41"},
		{"football/college-football", "150"},
		{"basketball/nba", "12"},
		{"basketball/nba", "2"},
		{"football/nfl", "12"},
		{"football/nfl", "33"},
	} {
		t.Run(tc.league+"/"+tc.team, func(t *testing.T) {
			game, err := Fetch(tc.league, tc.team)
			if err != nil {
				t.Fatal(err)
			}
			if game.Name == "" || game.Status == "" || game.Teams[0] == "" || game.Teams[1] == "" {
				t.Fatalf("incomplete game: %+v", game)
			}
			t.Logf("%s: %s; %s %s / %s %s; TV: %s", game.Name, game.Status, game.Teams[0], game.Scores[0], game.Teams[1], game.Scores[1], game.Broadcast)
		})
	}
}
