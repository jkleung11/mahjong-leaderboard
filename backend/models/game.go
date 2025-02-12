package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Date          time.Time
	WinnerID      *uint `gorm:"index" json:"winner_id"`
	WinningPoints *uint `json:"winning_points"`
}
