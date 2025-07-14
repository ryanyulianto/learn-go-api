package main

import (
	"belajar-go-api/internal/config"
	"belajar-go-api/internal/connection"
	"belajar-go-api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Get()
	dbConnection := connection.GetDatabase(config.Database)
	_ = repository.NewCustomer(dbConnection)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test lagi")
	})
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
