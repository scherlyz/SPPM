package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Route test
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "SPPM API running...",
        })
    })

    log.Fatal(app.Listen(":8081"))
}
