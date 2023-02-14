package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/petersephrin/wdbstudio/controllers/appcontroller"
)

func SetWebRoutes(router *fiber.App) {
	router.Get("/", appcontroller.Home)
	router.Get("/about", appcontroller.About)
	router.Get("/contact", appcontroller.Contact)
	router.Get("/portfolio", appcontroller.Portfolio)
}
