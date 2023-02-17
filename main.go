package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
	"github.com/joho/godotenv"
	"github.com/petersephrin/sephrintech/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("SephrinTech Website v0.1")

	engine := jet.New("./public/views", ".jet")

	config := fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
		Views:     engine,
	}
	app := fiber.New(config)

	app.Static("/images", "./public/images")
	app.Static("/nodemodules", "./node_modules")
	app.Static("/css", "./public/css")
	app.Static("/js", "./public/js")

	//just keeping my shit together so it's together
	//app.Get("/del/delete", artistcontroller.Logout)
	routes.SetWebRoutes(app)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Env not loaded")
	}
	//os.Setenv("PORT", "3000")
	port := os.Getenv("PORT")
	if err := app.Listen(":" + port); err != nil {
		fmt.Println("Failed to listen at port 9030")
		fmt.Println(err)
		log.Fatal(err)
	}
}
