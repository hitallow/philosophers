package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes : setup main rountes
func SetupRoutes(app *fiber.App){

	
	app.Get("/",func(c *fiber.Ctx) error {
	
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"message":"Welcome to my project 😎",
			},
		)
	})
	api := app.Group("/api")
	SetupPhilosopherRoutes(api)

}