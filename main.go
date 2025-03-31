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

	app.Post("/post", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString("Error parsing request body")
		}

		if user.Username == "" && user.Password == "" {
			return c.Status(400).SendString("Username and password are required")
		}

		return c.JSON(fiber.Map{
			"username": user.Username,
			"password": user.Password,
			"message":  "User created successfully",
		})
	})

	app.Listen(":3000")
}
