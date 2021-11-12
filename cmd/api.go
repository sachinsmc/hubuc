package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/hubuc-task/routes"
)

// Run function will create API server
func Run() {
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
