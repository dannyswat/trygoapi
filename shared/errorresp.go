package shared

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

func BadRequest(c *fiber.Ctx, err error) error {
	c.Status(400)
	return c.JSON(&ErrorResponse{
		Message:   err.Error(),
		ErrorCode: "invalid_request",
	})
}

func Unauthorized(c *fiber.Ctx, errMsg string) error {
	c.Status(401)
	return c.JSON(&ErrorResponse{
		Message:   errMsg,
		ErrorCode: "unauthorized",
	})
}
