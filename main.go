package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/narie-monarie/config"
	"github.com/narie-monarie/routes"
)

func main() {
	app := fiber.New()
	routes.UserRoutes(app)
	config.InitDB()
	app.Listen(":3000")
}
