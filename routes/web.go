package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/petersephrin/sephrintech/controllers/appcontroller"
)

func SetWebRoutes(router *fiber.App) {
	router.Get("/", appcontroller.Home)
	router.Get("/about", appcontroller.About)
	router.Get("/contact", appcontroller.Contact)
	router.Post("/contact", appcontroller.SendEmailContact)
	router.Get("/portfolio", appcontroller.Portfolio)

}
