package router

import (
	"main/reciver"
	"main/sender"

	"github.com/gofiber/fiber/v2"
)

func Rout(app fiber.Router) {
	//for rendering the templates
	app.Get("/recivemq", reciver.ReciverTemplates)

	app.Post("/send", sender.Send)

	app.Get("/recive", reciver.Recive)

}
