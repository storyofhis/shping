package models

import "time"

type TransactionHistory struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	ProductId  uint
	Product    Products `gorm:"foreignKey:ProductId"`
	UserId     uint
	User       Users `gorm:"foreignKey:UserId"`
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
