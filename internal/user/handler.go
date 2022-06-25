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
	List(ctx context.Context, criteria *ListUserRequest) (*Response, error)
	Delete(ctx context.Context, userID string) error
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.Create)
	app.Put("/users/:userID", h.Update)
	app.Delete("/users/:userID", h.Delete)
	app.Get("/users", h.List)
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	createUserReq := CreateUserRequest{}
	if err := ctx.BodyParser(&createUserReq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	err := h.service.Create(ctx.Context(), &createUserReq)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusCreated).JSON(DefaultResponse{})
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	updateUserReq := UpdateUserRequest{}
	userID := ctx.Params("userID")
	if err := ctx.BodyParser(&updateUserReq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	err := h.service.Update(ctx.Context(), userID, &updateUserReq)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(DefaultResponse{})
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Params("userID")
	err := h.service.Delete(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(DefaultResponse{})
}

func (h Handler) List(ctx *fiber.Ctx) error {
	params := &ListUserRequest{}
	if err := ctx.QueryParser(params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	users, err := h.service.List(ctx.Context(), params)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(DefaultResponse{Data: users})
}
