package handler

import "github.com/gofiber/fiber"

func UserIndex(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"success":   true,
		"messsage":  "This api is ready!",
		"version":   "1.14.6",
		"framework": "GO Fiber",
	})
}
