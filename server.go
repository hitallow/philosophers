package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hitallow/philosophers/src/routes"
)

func main() {
  app := fiber.New()

  routes.SetupRoutes(app)
  
  app.Listen(":8000")
}