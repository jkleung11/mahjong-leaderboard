package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Date          time.Time `json:"date"`
	WinnerID      *uint     `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:WinnerID;references:ID" json:"winner_id"`
	Winner        *Player   `gorm:"foreignKey:WinnerID;references:ID"`
	WinningPoints *uint     `json:"winning_points"`
}
