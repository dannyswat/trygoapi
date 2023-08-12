package main

import (
	"log"

	"time"

	"dannys/hk/trygoapi/users"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Status  int
	Message string
}

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		elapsed := time.Since(start)
		log.Printf("%s: %s", c.OriginalURL(), elapsed)
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&ApiResponse{Status: 0, Message: "Hello World!"})
	})
	users.RegisterHandlers(app)
	log.Fatal(app.Listen(":8080"))
}
