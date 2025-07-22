package main

import (
	"log"
	"os"
	"url_shortener/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	routes.RegisterRoute(app)
	app.Listen(":"+os.Getenv("PORT"))
	log.Println("Server up on 3000")
}