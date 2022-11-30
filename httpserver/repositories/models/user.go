package models

import "time"

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

// qerying role
// CREATE TYPE role AS ENUM ('admin', 'customer');
// UPDATE users SET role = 'admin' WHERE id = 1;

type Users struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	FullName  string
	Email     string
	Password  string
	Role      Role `gorm:"type:role;default:'customer'"`
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
