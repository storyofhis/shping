package models

import "time"

type Categories struct {
	Id                uint `gorm:"primaryKey;autoIncrement"`
	Type              string
	SoldProductAmount uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
