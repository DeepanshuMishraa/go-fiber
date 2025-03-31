package main

import "github.com/gofiber/fiber/v2"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}
