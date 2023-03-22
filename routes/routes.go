package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/controllers"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Auth routes
	v1.Post("/register", controllers.Register)
	v1.Post("/login", controllers.Login)
	v1.Get("/me", controllers.Me)

	// Article routes
	v1.Post("/article", controllers.CreateArticle)
	v1.Get("/article", controllers.GetAllArticle)
}
