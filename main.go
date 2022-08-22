package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	Id       string
	Name     string
	Lastname string
	Admin    bool
}

type Event struct {
	Title            string
	ShortDescription string
	LargeDescription string
	Organizer        string
	Place            string
	State            bool
}

type Admin struct {
	IdUser string
}

// func handleUsers(c *fiber.Ctx) error {
// 	user := User{
// 		Name:     "Nacho",
// 		Lastname: "Nazar",
// 		Year:     20,
// 	}
// 	return c.Status(fiber.StatusOK).JSON(user)
// }

// func handleCreateUser(c *fiber.Ctx) error {
// 	user := User{}
// 	if err := c.BodyParser(&user); err != nil {
// 		return err
// 	}

// 	user.Id = uuid.NewString()
// 	return c.Status(fiber.StatusOK).JSON(user)
// }

func main() {
	app := fiber.New()

	//Middlewares
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	// userGroup := app.Group("/users")

	// userGroup.Get("", handleUsers)
	// userGroup.Post("", handleCreateUser)
	// app.Get("/user", handleUsers)
	// app.Post("/user", handleCreateUser)
	app.Listen(":3000")
}
