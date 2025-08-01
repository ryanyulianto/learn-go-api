package api

import (
	"belajar-go-api/domain"
	"belajar-go-api/dto"
	"belajar-go-api/internal/util"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type customerApi struct {
	customerService domain.CustomerService
}

func timeOut(ctx *fiber.Ctx, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx.Context(), timeout)
}
func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := customerApi{
		customerService: customerService,
	}

	app.Get("/customer", ca.Index)
	app.Post("/customer", ca.Create)
	app.Post("/customer/update", ca.Update)

}
func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := timeOut(ctx, 10*time.Second)
	defer cancel()
	res, err := ca.customerService.Index(c, ctx.Queries())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	newResponse := map[string]interface{}{
		"data":       res,
		"total_data": len(res),
	}

	return ctx.JSON(dto.CreateResponseSuccess(newResponse))
}
func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := timeOut(ctx, 10*time.Second)
	defer cancel()
	req := dto.CreateCustomerRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseErrorData("Validation error", fails))
	}
	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(req))
}
func (ca customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := timeOut(ctx, 10*time.Second)
	defer cancel()
	req := dto.UpdateCustomerRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseErrorData("Validation error", fails))
	}
	err := ca.customerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(req))
}
