package server

import (
	"belajar-golang/server/modules"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Application() {
	app := fiber.New()

	modules.Modules(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Request yang kamu kirim salah",
		})
	})

	log.Fatal(app.Listen(":3000"))

}
