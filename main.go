package main

import (
	"api/database"
	"api/database/migration"
	"api/route"

	"github.com/gofiber/fiber"
)

func main() {
	// database connetion
	database.DatabaseInit()
	// database migrations
	migration.RunMigration()
	// go fiber
	app := fiber.New()

	route.RouteInit(app)
	// app listen to port?
	app.Listen("3000")
}
