package controller

import (
	"time"

	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddRemainder(c *fiber.Ctx) error {
	validate := validator.New()
	var reqReminder struct {
		Remind string    `json:"remind" validate:"required"`
		Time   time.Time `json:"time" validate:"required"`
	}
	if err := c.BodyParser(&reqReminder); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Body parsing", "status": false})
	}
	if err := validate.Struct(reqReminder); err != nil {
		return c.Status(206).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	var newReminder models.Reminder
	newReminder.Remind = reqReminder.Remind
	newReminder.Time = reqReminder.Time
	if userID, ok := c.Locals("user_id").(uuid.UUID); ok {
		newReminder.UserID = userID
	}
	result := database.DB.Create(newReminder)
	if result.Error != nil {
		return c.Status(417).JSON(fiber.Map{"message": "Error while Creating Remainder", "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Reminder Created Successfully", "status": true})
	return nil
}

func DeleteRemainder(c *fiber.Ctx) error {
	validate := validator.New()
	var reminder struct {
		Id uuid.UUID `json:"id" validate:"required"`
	}
	if err := c.BodyParser(&reminder); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Body parsing", "status": false})
	}
	if err := validate.Struct(reminder); err != nil {
		return c.Status(206).JSON(fiber.Map{"message": err.Error(), "status": false})
	}

	result := database.DB.Delete(&models.Reminder{}, reminder.Id)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Delete", "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Deleted Successfuly", "status": true})
	return nil
}

// func EditRemainder(c *fiber.Ctx) error{
// 	valiadate :=validator.New()
// 	var reqRemainder struct{
// 		ID uuid.UUID `json:"id" validate:"required"`
// 	}

// 	return nil
// }
