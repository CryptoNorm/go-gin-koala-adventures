package model

import "time"

type GameEvent struct {
	Id        int16     `json:"id"`
	Player    string    `json:"player"`
	GameLevel int8      `json:"game_level"`
	Score     int16     `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}
