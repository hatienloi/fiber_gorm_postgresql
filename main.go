package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/database"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/handlers"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/initializers"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/routes"
)

func main() {

	// Load .env
	initializers.LoadConfig()

	// Connect to database
	database.Connect()

	// Add init fiber and add fiber config
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.DefaultErrorHandler,
	})

	// Add middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	// Add logger
	app.Use(logger.New(logger.Config{
		Format:     "${pid} [${time}] ${locals:requestid} ${status} - ${method} ${path}​ ${error}\n",
		TimeFormat: "15:04:05 02-Jan-2006",
	}))

	// Setup router
	routes.Setup(app)

	// Start listen
	err := app.Listen(fmt.Sprintf(":%s", initializers.Config.ServerPort))
	if err != nil {
		return
	}
}
