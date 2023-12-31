package main

import (
	"github.com/SMRPcoder/markable/controller"
	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connetion()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200).JSON(fiber.Map{"message": "Hii Hello World"})
		return nil
	})

	authRoutes := app.Group("/auth")
	authRoutes.Post("/register", controller.Register)
	authRoutes.Post("/login", controller.Login)

	todoRoutes := app.Group("/todo", middleware.Authenticate)
	todoRoutes.Post("/add", controller.AddTodo)
	todoRoutes.Post("/changeStatus", controller.ChangeStatus)
	todoRoutes.Post("/delete", controller.DeleteTodo)
	todoRoutes.Post("/deleteCompleted", controller.DeleteAllCompleted)
	todoRoutes.Post("/viewAll", controller.ViewAllTodos)

	app.Listen(":5000")
}
