package routes

import (
	"MiddleProject/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
}
