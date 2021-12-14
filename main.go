package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yormangomez/test_cars/db"
	"github.com/yormangomez/test_cars/router"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cors.New())

	router.UserRouter(app)

	// config db
	db.ConnectDB()

	log.Fatal(app.Listen(":3000"))

}
