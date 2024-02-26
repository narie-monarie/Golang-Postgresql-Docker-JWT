package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string    `gorm:"unique"`
	Products  []Product `gorm:"foreignKey:UserID"`
}

type Product struct {
	ID          uint `gorm:"primaryKey"`
	ProductName string
	Price       float64
	Image       string
	Description string
	Location    string
	Available   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint `gorm:"foreignKey:User"`
}
