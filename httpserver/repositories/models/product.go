package models

import "time"

type Products struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	Title      string
	Price      uint
	Stock      int
	CategoryId uint
	Category   Categories `gorm:"foreignKey:CategoryId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
