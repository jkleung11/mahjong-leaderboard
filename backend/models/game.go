package models

import "time"

type Game struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Date          time.Time `gorm:"index" json:"date"`
	WinnerID      *int      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"winner_id"`
	WinningPoints *int      `json:"winning_points"`
	DateCreated   time.Time `gorm:"autoCreateTime" json:"date_created"`
	DateModified  time.Time `gorm:"autoUpdateTime" json:"date_modified"`
}
