package models

import (
	"time"
)

type App struct {
	ID    uint      `gorm:"primarykey"`
	Ctime time.Time `gorm:"ctime" json:"ctime"`
	Utime time.Time `gorm:"utime" json:"utime"`
}

func (App) TableName() string {
	return "app"
}
