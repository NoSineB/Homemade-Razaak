package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"homemade.com/database"
	"homemade.com/models"
)

func Checkout(c *fiber.Ctx) error {
	db := database.ConnectDatabase()

	type RequestBody struct {
		UserID uint `json:"user_id"`
	}

	var requestBody RequestBody
	var user models.User

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	db.Find(&user, requestBody.UserID)
	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	db.Delete(&models.CartItem{}, "user_id = ?", user.ID)

	return c.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}

func GetCartLength(c *fiber.Ctx) error {
	db := database.ConnectDatabase()

	type RequestBody struct {
		UserID uint `json:"user_id"`
	}

	var requestBody RequestBody
	var user models.User

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	db.Find(&user, requestBody.UserID)

	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var cartItems []models.CartItem
	var length int

	db.Model(&models.CartItem{}).Where("user_id = ?", user.ID).Find(&cartItems)
	length = len(cartItems)

	return c.JSON(fiber.Map{
		"length": length,
	})
}

func GetCartTotal(c *fiber.Ctx) error {
	db := database.ConnectDatabase()

	type RequestBody struct {
		UserID uint `json:"user_id"`
	}

	var requestBody RequestBody
	var user models.User

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	db.Find(&user, requestBody.UserID)

	fmt.Println(user)

	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var cartItems []models.CartItem
	var total int

	db.Model(&models.CartItem{}).Where("user_id = ?", user.ID).Find(&cartItems)

	for _, cartItem := range cartItems {
		var item models.Item
		db.Find(&item, cartItem.ItemID)
		total += item.Price
	}

	return c.JSON(fiber.Map{
		"total": total,
	})
}

func AddToCart(c *fiber.Ctx) error {
	db := database.ConnectDatabase()

	type RequestBody struct {
		UserID uint `json:"user_id"`
		ItemID uint `json:"item_id"`
	}

	var requestBody RequestBody
	var user models.User
	var item models.Item

	cartItem := new(models.CartItem)

	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	db.Find(&user, requestBody.UserID)

	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := db.Where("id = ?", requestBody.ItemID).First(&item).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	cartItem.UserID = user.ID
	cartItem.ItemID = item.ID

	if err := db.Create(cartItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to cart"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Item added to cart successfully",
	})
}
