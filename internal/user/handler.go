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
	Update(ctx context.Context, userID string, req *UpdateUserRequest) error
	Delete(ctx context.Context, userID string) error
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.CreateUserHandler)
	app.Put("/users/:userID", h.UpdateUserHandler)
	app.Delete("/users/:userID", h.DeleteUserHandler)
}

func (h Handler) CreateUserHandler(ctx *fiber.Ctx) error {
	createUserReq := CreateUserRequest{}
	if err := ctx.BodyParser(&createUserReq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	err := h.service.Create(ctx.Context(), &createUserReq)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	ctx.Status(http.StatusCreated)
	return nil
}

func (h Handler) UpdateUserHandler(ctx *fiber.Ctx) error {
	updateUserReq := UpdateUserRequest{}
	userID := ctx.Params("userID")
	if err := ctx.BodyParser(&updateUserReq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	err := h.service.Update(ctx.Context(), userID, &updateUserReq)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	ctx.Status(http.StatusOK)
	return nil
}

func (h Handler) DeleteUserHandler(ctx *fiber.Ctx) error {
	userID := ctx.Params("userID")
	err := h.service.Delete(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	ctx.Status(http.StatusOK)
	return nil
}
