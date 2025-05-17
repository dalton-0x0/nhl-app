package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Game struct {
	GameID    int    `json:"gameID"`
	GameDate  string `json:"gameDate"`
	Status    string `json:"status"`
	HomeTeam  string `json:"homeTeam"`
	AwayTeam  string `json:"awayTeam"`
	HomeScore int    `json:"homeScore"`
	AwayScore int    `json:"awayScore"`
	Period    int    `json:"period"`
	TimeLeft  string `json:"timeLeft"`
}

func FetchLiveGames() ([]Game, error) {
	url := "https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/scoreboard"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ESPN NHL data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var parsed struct {
		Events []struct {
			Status struct {
				Type struct {
					Description string `json:"description"`
				} `json:"type"`
			} `json:"status"`
			Competitions []struct {
				Competitors []struct {
					Team struct {
						DisplayName string `json:"displayName"`
					} `json:"team"`
					Score string `json:"score"`
				} `json:"competitors"`
			} `json:"competitions"`
		} `json:"events"`
	}

	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse ESPN API response: %w", err)
	}

	var games []Game
	for _, e := range parsed.Events {
		if len(e.Competitions) == 0 || len(e.Competitions[0].Competitors) < 2 {
			continue
		}
		comp := e.Competitions[0].Competitors
		game := Game{
			GameID:    0, // ESPN doesn’t expose a game ID — set 0 or generate later
			GameDate:  "",
			Status:    e.Status.Type.Description,
			HomeTeam:  comp[0].Team.DisplayName,
			AwayTeam:  comp[1].Team.DisplayName,
			HomeScore: atoiSafe(comp[0].Score),
			AwayScore: atoiSafe(comp[1].Score),
			Period:    0, // Can be parsed from status if needed
			TimeLeft:  "",
		}
		games = append(games, game)
	}

	return games, nil
}

func atoiSafe(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
