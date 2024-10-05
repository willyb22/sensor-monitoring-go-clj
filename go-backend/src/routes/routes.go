package routes

import (
	"go-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/", controllers.GetHome)
	app.Post("/sensor", controllers.PostSensorHandler)
	// ...
}