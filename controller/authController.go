package controller

import (
	"errors"

	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/functions"
	"github.com/SMRPcoder/markable/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	validate := validator.New()
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}
	var thisuser models.User
	userresult := database.DB.Where("username= ?", user.Username).Take(&thisuser)
	if errors.Is(userresult.Error, gorm.ErrRecordNotFound) {
		if err := validate.Struct(user); err != nil {
			return c.Status(206).JSON(fiber.Map{"message": err.Error(), "status": false})
		}
		result := database.DB.Create(&user)
		if result.Error != nil {
			return c.Status(206).JSON(fiber.Map{"message": "Error While Creating A User", "status": false})
		}
		c.Status(200).JSON(fiber.Map{"message": "data saved", "data": user, "status": true})

	} else if userresult.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Error At User Querying", "status": false})
	}

	// log.Println(result)

	return nil
}

func Login(c *fiber.Ctx) error {
	validate := validator.New()
	var requser struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var user models.User
	if err := c.BodyParser(&requser); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while Body parsing", "status": false})
	}
	if err := validate.Struct(requser); err != nil {
		return c.Status(206).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	result := database.DB.Where("username= ?", requser.Username).First(&user)
	if result.Error != nil {
		return c.Status(417).JSON(fiber.Map{"message": result.Error.Error(), "status": false})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requser.Password)); err != nil {
		return c.Status(200).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	token, err := functions.EncodeJwt(functions.JWTuser{Username: requser.Username, Name: user.Name, Id: user.ID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error(), "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "Loggedin Successfully", "status": true, "token": "Bearer " + token})
	return nil

}
