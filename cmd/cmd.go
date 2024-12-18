package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Execute() {
	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}
