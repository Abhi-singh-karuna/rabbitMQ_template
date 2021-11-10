package main

import (
	"main/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Static("/", "./templates")

	router.Rout(app)

	app.Listen(":9000")
}
