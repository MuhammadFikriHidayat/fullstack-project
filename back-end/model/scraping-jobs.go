package model

import "gorm.io/gorm"

type Scrap struct {
	gorm.Model
	UserID uint `json:"user_id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Status string `gorm:"type:varchar(20);default:waiting" json:"status"`
	User User `gorm:"foreignKey:UserID"`
}