package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Date          time.Time
	WinnerID      *uint `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:WinnerID;references:ID" json:"winner_id"`
	WinningPoints *uint `json:"winning_points"`
}
