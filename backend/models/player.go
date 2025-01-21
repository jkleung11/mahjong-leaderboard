package models

type Player struct {
	ID   int    `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}
