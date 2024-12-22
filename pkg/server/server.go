package server

import (
	"log"

	"github.com/gofiber/fiber"
)

func InitServerApp() {
	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}
