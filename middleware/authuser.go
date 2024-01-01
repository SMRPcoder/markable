package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/functions"
	"github.com/SMRPcoder/markable/models"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization != "" && strings.HasPrefix(authorization, "Bearer") {
		token := strings.Split(authorization, " ")[1]
		data, err := functions.DecodeJwt(token)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return c.Status(401).JSON(fiber.Map{"message": "Wrong Token Provided", "status": false})
		}
		if reflect.DeepEqual(data, functions.JWTuser{}) {
			return c.Status(401).JSON(fiber.Map{"message": "Invalid JWT claims", "status": false})
		}
		var user models.User
		result := database.DB.Where("id = ?", data.Id).First(&user)
		if result.Error != nil {
			return c.Status(401).JSON(fiber.Map{"message": "UnAuthorized Error", "status": false})
		}
		c.Locals("user_id", user.ID)
		c.Next()
	} else {
		return c.Status(401).JSON(fiber.Map{"message": "Auth Token Not Provided", "status": false})
	}

	return nil
}
