package modules

import (
	"belajar-golang/server/modules/user"

	"github.com/gofiber/fiber/v2"
)

func Modules(app *fiber.App) {

	user.Router(app)

}
