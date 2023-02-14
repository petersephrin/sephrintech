package appcontroller

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"title": "Sephrin Tech",
	}, "layouts/master")
}

func About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"title": "About Sephrin Tech",
	}, "layouts/master")
}

func Portfolio(c *fiber.Ctx) error {
	return c.Render("portfolio", fiber.Map{
		"title": "Talk to Us-Sephrin Tech",
	}, "layouts/master")
}

func Contact(c *fiber.Ctx) error {
	return c.Render("contact", fiber.Map{
		"title": "Talk to Us-Sephrin Tech",
	}, "layouts/master")
}
