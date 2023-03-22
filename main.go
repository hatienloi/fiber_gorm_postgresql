package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/database"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/handlers"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/routes"
)

func main() {

	database.Connect()

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.DefaultErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} [${time}] ${locals:requestid} ${status} - ${method} ${path}​ ${error}\n",
		TimeFormat: "15:04:05 02-Jan-2006",
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
