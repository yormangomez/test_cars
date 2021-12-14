package router

import "github.com/gofiber/fiber/v2"

func UserRouter(app *fiber.App) {
	app.Get("/api/user")
	app.Post("/api/register")
	app.Get("/api/user/:id")
	app.Patch("/api/user/:id")
	app.Delete("/api/user/:id")
}
