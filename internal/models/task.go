package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID           int `gorm:"autoIncrement"`
	StartTime    time.Time
	EndTime      time.Time
	UserPassport int
}
