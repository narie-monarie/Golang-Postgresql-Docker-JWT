package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/narie-monarie/controllers"
)

func UserRoutes(app *fiber.App) {
	userRouter := app.Group("/users")
	userRouter.Get("/", controllers.GetUsers)
	userRouter.Get("/:id", controllers.GetUser)
	userRouter.Post("/", controllers.CreateUser)
	userRouter.Put("/:id", controllers.UpdateUser)
	userRouter.Delete("/:id", controllers.DeleteUser)

	productRouter := app.Group("/products")
	productRouter.Get("/", controllers.GetProducts)
	productRouter.Get("/:id", controllers.GetProduct)
	productRouter.Post("/", controllers.CreateProduct)
	productRouter.Put("/:id", controllers.UpdateProduct)
	productRouter.Delete("/:id", controllers.DeleteProduct)
}
