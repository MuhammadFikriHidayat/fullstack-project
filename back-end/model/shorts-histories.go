package model

import "gorm.io/gorm"

type History struct {
	gorm.Model
	JobID uint `json:"job_id"`
	Title string `gorm:"type:text;not null" json:"title"`
	URL string `gorm:"type:text;not null" json:"url"`
	Job Scrap `gorm:"foreignKey:JobID"`
}