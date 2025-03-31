package main

import (
	"github.com/DeepanshuMishraa/go-fiber/database"
	"github.com/DeepanshuMishraa/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,World!")
	})

	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/get-users", routes.GetUsers)
	app.Get("/api/get-users/:id", routes.GetUser)
}

func main() {

	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
