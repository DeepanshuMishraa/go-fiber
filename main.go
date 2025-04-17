package main

import (
	"github.com/DeepanshuMishraa/go-fiber/database"
	"github.com/DeepanshuMishraa/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	app.Post("/api/user/register", routes.RegisterUser)
	app.Post("/api/user/login", routes.SignInUser)
}

func main() {

	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
