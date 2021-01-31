package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hitallow/philosophers/src/controllers"
)

// SetupPhilosopherRoutes : set a philosopher routes
func SetupPhilosopherRoutes(
	router fiber.Router,
){

	router.Get(
		"/philosophers",
		func(c *fiber.Ctx) error {
			return controllers.ListPhilosophers(*c)
		},
	)

	router.Get(
		"/philosophers/:name/random",
		func(c *fiber.Ctx) error {
			
			return	controllers.GetPhrase(
				c.Params("name"),
				*c,
			)
		},
	)

	router.Get(
		"/philosophers/:name/phrases",
		func(c *fiber.Ctx) error {
			return controllers.ListPhilosopherPhrases(
				c.Params("name"),
				*c,
			)
		},
	)
}
