package routes

import (
	"os"
	"time"

	"github.com/DeepanshuMishraa/go-fiber/database"
	"github.com/DeepanshuMishraa/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type User struct {
	ID uint `json:"id"`
}

func CreateResponseUser(user models.User) User {
	return User{
		ID: user.ID,
	}
}

func RegisterUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, email, and password are required",
		})
	}

	// Check if user already exists

	var existingUser models.User

	if err := database.Database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	database.Database.DB.Create(&user)

	response := CreateResponseUser(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"id":      response.ID,
	})
}

func CreateToken(email string) (string, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return "", err
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	var secretKey = []byte(JWT_SECRET)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "go-fiber-topshelf",
		"exp": time.Now().Add(time.Hour).Unix(), // 1 hour expiration
		"lat": time.Now().Unix(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func SignInUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	var existingUser models.User

	if err := database.Database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	if existingUser.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := CreateToken(existingUser.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	response := CreateResponseUser(existingUser)

	return c.JSON(fiber.Map{
		"message": "User signed in successfully",
		"id":      response.ID,
		"token":   token,
	})
}
