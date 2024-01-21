package models

import "time"

type Job struct {
	JobId       uint `gorm:"primary_key"`
	Title       string
	CategoryIds string
	Description string
	Contact     string
	CreateBy    int
	CreateAt    time.Time
	UpdatedAt   time.Time
	Status      int
}
