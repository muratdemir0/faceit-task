package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	service Service
}

type Service interface {
	Create(ctx context.Context, req *CreateUserRequest) error
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.CreateUserHandler)
}

func (h Handler) CreateUserHandler(ctx *fiber.Ctx) error {
	cur := CreateUserRequest{}
	if err := ctx.BodyParser(&cur); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	err := h.service.Create(ctx.Context(), &cur)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	ctx.Status(http.StatusCreated)
	return nil
}
