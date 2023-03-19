package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/controllers"
)

func Setup(app *fiber.App) {

	var apiVersion = "/api/v1/%s"

	// Auth routes
	app.Post(fmt.Sprintf(apiVersion, "register"), controllers.Register)
	app.Post(fmt.Sprintf(apiVersion, "login"), controllers.Login)
	app.Get(fmt.Sprintf(apiVersion, "me"), controllers.Me)

	// Article routes
	app.Post(fmt.Sprintf(apiVersion, "article"), controllers.CreateArticle)
	app.Get(fmt.Sprintf(apiVersion, "article"), controllers.GetAllArticle)
}
