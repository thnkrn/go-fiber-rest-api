package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thnkrn/go-fiber-rest-api/middleware"
	"github.com/thnkrn/go-fiber-rest-api/user"
)

func hello(c *fiber.Ctx) error {
	fmt.Println("test")
	return c.SendString("Hello World")
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "lark" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Lark",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Routers(app *fiber.App) {
	// Auth
	app.Post("/login", login)

	// User
	users := app.Group("/", middleware.Protected())
	users.Get("/users", user.GetUsers)
	users.Get("/user/:id", user.GetUser)
	users.Post("/user", user.SaveUser)
	users.Delete("/user/:id", user.DeleteUser)
	users.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":3000")
}
