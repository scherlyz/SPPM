package main

import (
    "sppm/config"
    "sppm/database"
	"log"
	

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
    database.ConnectPostgres()
    database.ConnectMongo()

	app := fiber.New()

	// Route test
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "SPPM API running...",
		})
	})

	log.Fatal(app.Listen(":8081"))
}
