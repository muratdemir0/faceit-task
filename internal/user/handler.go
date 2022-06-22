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
	FindBy(ctx context.Context, criteria *FindUserRequest) (*Response, error)
	Delete(ctx context.Context, userID string) error
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.CreateHandler)
	app.Put("/users/:userID", h.UpdateHandler)
	app.Delete("/users/:userID", h.DeleteHandler)
	app.Get("/users", h.FindHandler)
}

func (h Handler) CreateHandler(ctx *fiber.Ctx) error {
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

func (h Handler) UpdateHandler(ctx *fiber.Ctx) error {
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

func (h Handler) DeleteHandler(ctx *fiber.Ctx) error {
	userID := ctx.Params("userID")
	err := h.service.Delete(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	ctx.Status(http.StatusOK)
	return nil
}

func (h Handler) FindHandler(ctx *fiber.Ctx) error {
	params := &FindUserRequest{}
	if err := ctx.QueryParser(params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	users, err := h.service.FindBy(ctx.Context(), params)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(users)
}
