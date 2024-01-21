package models

type Role struct {
	RoleId uint `gorm:"primary_key"`
	Name   string
}
