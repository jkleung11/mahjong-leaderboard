package models

import (
	"gorm.io/gorm"
)

type GamePlayers struct {
	gorm.Model
	GameID   uint   `gorm:"index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"game_id"`
	Game     Game   `gorm:"foreignKey:GameID;references:ID"`
	PlayerID uint   `gorm:"index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"player_id"`
	Player   Player `gorm:"foreignKey:PlayerID;references:ID"`
}
