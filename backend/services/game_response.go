package services

type PlayersResponse struct {
	Name         string `json:"name"`
	Result       string `json:"result"`
	PointsEarned uint   `json:"points_earned"`
}

type GameResponse struct {
	Message string            `json:"message"`
	GameID  uint              `json:"game_id"`
	Players []PlayersResponse `json:"players"`
}
