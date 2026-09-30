package user

import (
	user "belajar-golang/server/modules/user/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/user")

	api.Get("v1/users", user.GetUsers)

}
