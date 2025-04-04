package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"Name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Item struct {
	gorm.Model
	ItemName string `json:"recipe"`
	Price    int    `json:"price"`
}

type CartItem struct {
	gorm.Model
	UserID uint `json:"user_id"`
	ItemID uint `json:"item_id"`
}
