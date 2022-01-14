package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	FirstName string `json:"fisrtname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	db.AutoMigrate(&User{})
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	db.Find(&users)
	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	db.Find(&user, id)
	return c.JSON(&user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Create(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	db.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}

	db.Delete(&user)
	return c.SendString("User is deleted!!!")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(User)
	db.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Save(&user)
	return c.JSON(&user)
}
