package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	url := "https://statsapi.web.nhl.com/api/v1/schedule?expand=schedule.linescore"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch NHL data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var parsed struct {
		Dates []struct {
			Games []struct {
				GameID   int    `json:"gamePk"`
				GameDate string `json:"gameDate"`
				Status   struct {
					DetailedState string `json:"detailedState"`
				} `json:"status"`
				Teams struct {
					Home struct {
						Team struct {
							Name string `json:"name"`
						} `json:"team"`
						Score int `json:"score"`
					} `json:"home"`
					Away struct {
						Team struct {
							Name string `json:"name"`
						} `json:"team"`
						Score int `json:"score"`
					} `json:"away"`
				} `json:"teams"`
				Linescore struct {
					CurrentPeriod             int    `json:"currentPeriod"`
					CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining"`
				} `json:"linescore"`
			} `json:"games"`
		} `json:"dates"`
	}

	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse NHL API response: %w", err)
	}

	var games []Game
	for _, date := range parsed.Dates {
		for _, g := range date.Games {
			game := Game{
				GameID:    g.GameID,
				GameDate:  g.GameDate,
				Status:    g.Status.DetailedState,
				HomeTeam:  g.Teams.Home.Team.Name,
				AwayTeam:  g.Teams.Away.Team.Name,
				HomeScore: g.Teams.Home.Score,
				AwayScore: g.Teams.Away.Score,
				Period:    g.Linescore.CurrentPeriod,
				TimeLeft:  g.Linescore.CurrentPeriodTimeRemaining,
			}
			games = append(games, game)
		}
	}

	return games, nil
}