package dtos

import "time"

type PlayerResults struct {
	Name         string `json:"name"`
	Result       string `json:"result"`
	PointsEarned uint   `json:"points_earned"`
}

type GameDetails struct {
	GameID  uint      `json:"game_id"`
	Date    time.Time `json:"date"`
	Winner  *string   `json:"winner"`
	Results []PlayerResults
}
