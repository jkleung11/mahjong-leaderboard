package models

type GamePlayer struct {
	GameID   int  `gorm:"primaryKey;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"game_id"`
	PlayerID int  `gorm:"primaryKey;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"player_id"`
	IsWinner bool `json:"is_winner"`
}
