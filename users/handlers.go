package users

import (
	"dannys/hk/trygoapi/shared"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	model := &LoginModel{}
	if parseErr := c.BodyParser(model); parseErr != nil {
		return shared.BadRequest(c, parseErr)
	}

	if model.UserName != "dannys" || model.Password != "123456" {
		return shared.Unauthorized(c, "Invalid username or password")
	}

	return c.JSON(&LoginResult{
		AccessToken: "I love HK",
		Status:      Success,
	})
}

func RegisterHandlers(app *fiber.App) {
	app.Post("/login", Login)
}
