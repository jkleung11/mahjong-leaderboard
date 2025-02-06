package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model        // use built in struct
	Name       string `gorm:"unique;not null" json:"name"`
}
