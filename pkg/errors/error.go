package errors

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/muratdemir0/faceit-task/pkg/store"
	errs "github.com/pkg/errors"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func makeResponse(err error) ErrorResponse {
	var errorResponse ErrorResponse
	if errs.As(err, &errorResponse) {
		return errorResponse
	}
	if errors.Is(err, sql.ErrNoRows) {
		return NotFound("")
	}
	if errors.Is(err, store.NotFoundError) {
		return NotFound("")
	}
	if errors.Is(err, fiber.ErrMethodNotAllowed) {
		return MethodNotAllowedError("")
	}

	return InternalServerError("")
}

func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "Well, this is unexpected."
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

func MethodNotAllowedError(msg string) ErrorResponse {
	if msg == "" {
		msg = "Method not allowed."
	}
	return ErrorResponse{
		Status:  http.StatusMethodNotAllowed,
		Message: msg,
	}
}

func NotFound(msg string) ErrorResponse {
	if msg == "" {
		msg = "Request was not found."
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}
func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Request was not valid."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}
