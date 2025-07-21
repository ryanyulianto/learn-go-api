package main

import (
	"belajar-go-api/internal/api"
	"belajar-go-api/internal/config"
	"belajar-go-api/internal/connection"
	"belajar-go-api/internal/repository"
	"belajar-go-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.Get()
	dbConnection := connection.GetDatabase(config.Database)
	app := fiber.New()
	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService)
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
