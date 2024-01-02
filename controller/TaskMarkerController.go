package controller

import (
	"encoding/json"
	"time"

	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddTask(c *fiber.Ctx) error {
	validate := validator.New()
	var reqTask struct {
		TaskName  string `json:"task_name" validate:"required"`
		TotalDays int    `json:"total_days" validate:"required"`
	}
	if err := c.BodyParser(&reqTask); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	if err := validate.Struct(reqTask); err != nil {
		return c.Status(206).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	var newTask models.Taskmark
	newTask.TaskName = reqTask.TaskName
	newTask.TotalDays = reqTask.TotalDays
	taskdays := make(map[string]bool, reqTask.TotalDays)
	for i := 0; i <= reqTask.TotalDays; i++ {
		currentTime := time.Now().UTC()
		newTime := currentTime.Add(time.Duration(i) * 24 * time.Hour)
		currentDate := newTime.Format("2006-01-02")
		taskdays[currentDate] = false
	}
	jsontask, err := json.Marshal(taskdays)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	strtaskdays := string(jsontask)
	newTask.TaskDays = strtaskdays
	result := database.DB.Create(&newTask)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Daily Todo Added Successfully", "status": true})
	return nil
}

func MarkToday(c *fiber.Ctx) error {
	var reqTask struct {
		Id uuid.UUID `json:"id"`
	}
	if err := c.BodyParser(&reqTask); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	var taskmark models.Taskmark
	result := database.DB.Where("id = ?", reqTask.Id).First(&taskmark)
	if result.Error != nil {
		return c.Status(417).JSON(fiber.Map{"message": result.Error.Error(), "status": false})
	}
	jsonTaskDays := make(map[string]bool)
	json.Unmarshal([]byte(taskmark.TaskDays), &jsonTaskDays)
	currentTime := time.Now().UTC()
	currentDate := currentTime.Format("2006-01-02")
	jsonTaskDays[currentDate] = true
	jsonData, err := json.Marshal(jsonTaskDays)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	taskmark.TaskDays = string(jsonData)
	updateResult := database.DB.Model(&models.Taskmark{}).Updates(&taskmark)
	if updateResult.Error != nil {
		return c.Status(417).JSON(fiber.Map{"message": updateResult.Error.Error(), "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Successfully Task Completed Today", "status": true})
	return nil
}
