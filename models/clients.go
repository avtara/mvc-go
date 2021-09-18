package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model   `json:"-"`
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ClientKey    string `gorm:"type:varchar(255);not null" json:"client_key"`
	ClientSecret string `gorm:"type:varchar(255);not null" json:"client_secret"`
	Status       bool   `gorm:"type:bool" json:"status"`
}
