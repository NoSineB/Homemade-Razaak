package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"homemade.com/models"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Item{}, &models.CartItem{})

	// Migrate()

	return db
}

func Migrate() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Item{}, &models.CartItem{})

	user := models.User{
		Name:     "benison",
		Password: "password",
		Email:    "benison@example.com",
	}
	db.Create(&user)

	items := []models.Item{
		{
			ItemName: "Meals",
			Price:    150,
		},
		{
			ItemName: "Chicken Briyani",
			Price:    250,
		},
		{
			ItemName: "Fish Fry",
			Price:    75,
		},
		{
			ItemName: "Kanji",
			Price:    50,
		},
		{
			ItemName: "Fish Curry",
			Price:    350,
		},
		{
			ItemName: "Achar",
			Price:    110,
		},
		{
			ItemName: "Chappathi",
			Price:    20,
		},
		{
			ItemName: "Beaf Curry",
			Price:    110,
		},
	}

	db.Create(&items)
}
