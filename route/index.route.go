package route

import (
	"api/handler"

	"github.com/gofiber/fiber"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserIndex)
}
