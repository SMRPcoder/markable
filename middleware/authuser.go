package middleware

import (
	"strings"

	"github.com/SMRPcoder/markable/database"
	"github.com/SMRPcoder/markable/errorlog"
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
			return errorlog.Log_n_send(err.Error(), c, 400, "Jwt Error")
		}
		var user models.User
		result := database.DB.Where("id = ?", data.Id).First(&user)
		if result.Error != nil {
			return errorlog.Log_n_send(err.Error(), c, 401, "UnAuthorized Error")
		}
		c.Locals("user_id", user.ID)
		c.Next()
	} else {
		return c.Status(401).JSON(fiber.Map{"message": "Auth Token Not Provided", "status": false})
	}

	return nil
}
