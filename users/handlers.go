package users

import (
	"dannys/hk/trygoapi/shared"

	"github.com/gofiber/fiber/v2"
)

type LoginStatus string

const (
	Success      LoginStatus = "success"
	InvalidCred  LoginStatus = "invalid"
	AcctLocked   LoginStatus = "account_locked"
	AcctDisabled LoginStatus = "account_disabled"
	Need2FA      LoginStatus = "proceed_2fa"
)

type LoginModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	AccessToken string      `json:"token"`
	Status      LoginStatus `json:"status"`
}

func login(c *fiber.Ctx) error {

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
	app.Post("/login", login)
}
