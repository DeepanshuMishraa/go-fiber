package routes

import (
	"errors"

	"github.com/DeepanshuMishraa/go-fiber/database"
	"github.com/DeepanshuMishraa/go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request",
		})
	}

	database.Database.DB.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"id":      responseUser.ID,
	})

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.DB.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": responseUsers,
	})

}

func findUser(id int, user models.User) error {
	database.Database.DB.Find(&user, "id=?", id)

	if user.ID == 0 {
		return errors.New("User not found")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Make sure to pass a vaild id",
		})
	}

	if err := findUser(id, user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": responseUser,
	})

}
