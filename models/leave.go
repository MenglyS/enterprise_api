package models

import "time"

type Leave struct {
	LeaveId       uint `gorm:"primary_key"`
	Title         string
	LeaveFile     string
	Description   string
	LeaveDateFrom time.Time
	LeaveDateTo   time.Time
	CreatedBy     int
	ApprovedBy    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        int
}

func (Leave) TableName() string {
	return "leave_requests"
}
