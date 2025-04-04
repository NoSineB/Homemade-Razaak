package handlers

import (
	"github.com/gofiber/fiber/v2"
	"homemade.com/database"
	"homemade.com/models"
)

func Register(c *fiber.Ctx) error {
	db := database.ConnectDatabase()
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := db.Create(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user_id": user.ID,
		"name":    user.Name,
	})
}

func Login(c *fiber.Ctx) error {
	db := database.ConnectDatabase()
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var foundUser models.User
	if err := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&foundUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{
		"user_id": foundUser.ID,
		"name":    foundUser.Name,
	})
}
