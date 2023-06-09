package main

//github.com/200103395/MidGo

import (
	"MiddleProject/config"
	"MiddleProject/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	config.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
