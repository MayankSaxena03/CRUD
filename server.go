package main

import (
	"github.com/MayankSaxena03/CRUD/models"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/user/:id", models.GetUser)
	app.Put("/user/:id", models.UpdateUser)
	app.Post("/user", models.CreateUser)
	app.Delete("/user/:id", models.DeleteUser)
}

func main() {
	models.InitDB()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	Router(app)
	app.Listen(":8080")
}
