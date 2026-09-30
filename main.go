package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})

	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Produk : " + id)
	})

	app.Get("/GetAllUser", func(c *fiber.Ctx) error {
		user := []User{
			{
				ID:    1,
				Name:  "Andi",
				Email: "[EMAIL_ADDRESS]",
			},
			{
				ID:    2,
				Name:  "Budi",
				Email: "[EMAIL_ADDRESS]",
			},
		}

		return c.JSON(user)
	})

	app.Post("/daftar", func(c *fiber.Ctx) error {
		// var body User
		body := User{}

		if err := c.BodyParser(&body); err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"id":    body.ID,
			"name":  body.Name,
			"email": body.Email,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
