package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/narie-monarie/config"
	"github.com/narie-monarie/models"
)

type User = models.User
type Products = models.Product

func CreateUser(c *fiber.Ctx) error {
	newUser := User{}
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	result := config.DB.Create(&newUser)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create newUser",
		})
	}
	return c.Status(201).JSON(newUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []User{}
	result := config.DB.Preload("Products").Find(&users)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get users",
		})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{}
	result := config.DB.Preload("Products").First(&user, id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No user found",
		})
	}
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{}
	result := config.DB.Preload("Products").Delete(&user, id)

	if result.Error != nil {

		if result.RowsAffected == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "No user found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "User deleted",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{}
	result := config.DB.Preload("Products").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No user found",
		})
	}
	c.BodyParser(&user)
	result = config.DB.Save(&user)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "User updated",
		"user":    user,
	})
}
