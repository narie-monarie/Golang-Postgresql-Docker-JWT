package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/narie-monarie/config"
	"github.com/narie-monarie/models"
)

type Product = models.Product

func CreateProduct(c *fiber.Ctx) error {
	newProduct := Product{}
	if err := c.BodyParser(newProduct); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	result := config.DB.Create(&newProduct)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Product Created",
		"product": newProduct,
	})
}

func GetProducts(c *fiber.Ctx) error {
	products := []Product{}
	result := config.DB.Find(&products)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get products",
		})
	}
	return c.Status(http.StatusOK).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := Product{}
	result := config.DB.Preload("Products").First(&product, id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No product found",
		})
	}
	return c.Status(http.StatusOK).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := Product{}
	result := config.DB.Delete(&product, id)

	if result.Error != nil {

		if result.RowsAffected == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "No product found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "product deleted",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := Product{}
	result := config.DB.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No product found",
		})
	}
	c.BodyParser(&product)
	result = config.DB.Save(&product)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Product updated",
		"product": product,
	})
}
