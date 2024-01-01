package controller

import (
	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Body parsing", "status": false})
	}
	if userID, ok := c.Locals("user_id").(uuid.UUID); ok {
		todo.UserID = userID
	}
	result := database.DB.Create(&todo)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while creating Todo", "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Todo Added To List", "status": true})
	return nil
}

func ChangeStatus(c *fiber.Ctx) error {
	var reqtodo struct {
		Finished bool      `json:"finished"`
		Id       uuid.UUID `json:"id"`
	}
	if err := c.BodyParser(&reqtodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Body parsing", "status": false})
	}
	var todo models.Todo
	result := database.DB.Where("id = ?", reqtodo.Id).First(&todo)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"message": "No Todo Found!!!", "status": false})
	}
	todo.Finished = reqtodo.Finished
	database.DB.Save(&todo)
	c.Status(200).JSON(fiber.Map{"message": "Todo Status Changed Successfully", "status": true})
	return nil
}

func DeleteTodo(c *fiber.Ctx) error {
	var reqtodo struct {
		Id uuid.UUID `json:"id"`
	}
	if err := c.BodyParser(&reqtodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Body parsing", "status": false})
	}

	result := database.DB.Delete(&models.Todo{}, reqtodo.Id)
	if result.Error != nil || result.RowsAffected < 1 {
		return c.Status(400).JSON(fiber.Map{"message": "Todo Not Found", "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Todo Deleted Successfully", "status": true})
	return nil
}

func DeleteAllCompleted(c *fiber.Ctx) error {
	result := database.DB.Where("finished = ?", true).Delete(&models.User{})
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"message": result.Error.Error(), "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Completed Todo is Deleted Successfully", "status": true})
	return nil
}

func ViewAllTodos(c *fiber.Ctx) error {
	var todo []models.Todo
	result := database.DB.Where("user_id = ?", c.Locals("user_id").(uuid.UUID)).Find(&todo)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"message": result.Error.Error(), "status": false})
	}
	c.Status(200).JSON(fiber.Map{"data": todo, "status": true})
	return nil
}
