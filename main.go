package main

import (
	"belajar-go-api/internal/api"
	"belajar-go-api/internal/config"
	"belajar-go-api/internal/connection"
	"belajar-go-api/internal/repository"
	"belajar-go-api/internal/service"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func setupCustomerModule(app *fiber.Group, db *sql.DB) {
	repo := repository.NewCustomer(db)
	service := service.NewCustomer(repo)
	api.NewCustomer(app, service)
}

func main() {
	config := config.Get()
	dbConnection := connection.GetDatabase(config.Database)
	app := fiber.New()
	api_v1 := app.Group("/api/v1")

	setupCustomerModule(api_v1.(*fiber.Group), dbConnection)

	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
