package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Result string

const (
	Win  Result = "win"
	Loss Result = "loss"
	Draw Result = "draw"
)

func (r Result) Validate() error {
	switch r {
	case Win, Loss, Draw:
		return nil
	default:
		return errors.New("invalid result value")
	}
}

type GamePlayer struct {
	// uniquely defines a player's participation in a game
	gorm.Model
	GameID       uint      `gorm:"index" json:"game_id"`
	Date         time.Time `gorm:"index" json:"date"`
	PlayerID     uint      `gorm:"index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"player_id"`
	Player       Player    `gorm:"foreignKey:PlayerID;references:ID"`
	Result       string    `gorm:"index" json:"result"`
	PointsEarned uint
}
