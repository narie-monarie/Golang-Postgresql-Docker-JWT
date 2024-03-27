package main

import (
	"narie-monarie/config"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	if err := config.InitDB(); err != nil {
		panic(err)
	}
	app.Listen(":42069")
}
