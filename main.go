package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func setupRoutes(app *fiber.App) {

	app.Get("/api", welcome)

	app.Post("/api/contructs", routes.CreateContruct)
	app.Get("/api/contructs", routes.GetContructs)
	app.Get("/api/contructs/:id", routes.GetContruct)
	app.Put("/api/contructs/:id", routes.UpdateContruct)
	app.Delete("/api/contructs/:id", routes.DeleteContruct)
}

func main() {

	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
