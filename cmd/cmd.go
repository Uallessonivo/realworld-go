package cmd

import (
	"log"

	"Github.com/Uallessonivo/RealWorld/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Execute() {
	database.InitSqlite()

	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}
