package main

import (
	"log"
	"pos-login/config"
	"pos-login/database"
	"pos-login/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env")
	}
	config.LoadEnv()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.Port))
}
