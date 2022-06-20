package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/muratdemir0/faceit-task/internal/user"
	mocks "github.com/muratdemir0/faceit-task/mocks/user"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateUserHandler(t *testing.T) {
	Convey("Given create user request is invalid", t, func() {
		app := createTestApp()
		c := gomock.NewController(t)
		defer c.Finish()
		createUserReq := `{`
		Convey("When create request is required", func() {
			mockService := mocks.NewMockService(c)

			handler := user.NewHandler(mockService)
			handler.RegisterRoutes(app)

			req := NewHTTPRequestWithJSON(http.MethodPost, "/users", createUserReq)
			actualResponse, _ := app.Test(req)
			defer actualResponse.Body.Close()
			Convey("Then response status code should be 400", func() {
				So(actualResponse.StatusCode, ShouldEqual, http.StatusBadRequest)
			})
		})
	})
	Convey("Given create user request is valid", t, func() {
		app := createTestApp()
		c := gomock.NewController(t)
		defer c.Finish()
		createUserReq := user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "doe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When create is called with valid request", func() {
			mockService := mocks.NewMockService(c)
			mockService.EXPECT().Create(gomock.Any(), &createUserReq).Return(nil)

			handler := user.NewHandler(mockService)
			handler.RegisterRoutes(app)

			req := NewHTTPRequestWithJSON(http.MethodPost, "/users", createUserReq)
			actualResponse, err := app.Test(req)
			defer actualResponse.Body.Close()
			So(err, ShouldBeNil)
			Convey("Then response status code should be 201", func() {
				So(actualResponse.StatusCode, ShouldEqual, http.StatusCreated)
			})
		})
	})
	Convey("Given create user request is valid", t, func() {
		app := createTestApp()
		c := gomock.NewController(t)
		defer c.Finish()
		createUserReq := user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "doe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When create is called with valid request", func() {
			mockService := mocks.NewMockService(c)
			mockService.EXPECT().Create(gomock.Any(), &createUserReq).Return(errors.New("error"))

			handler := user.NewHandler(mockService)
			handler.RegisterRoutes(app)

			req := NewHTTPRequestWithJSON(http.MethodPost, "/users", createUserReq)
			actualResponse, _ := app.Test(req)
			defer actualResponse.Body.Close()
			Convey("Then response status code should be 500", func() {
				So(actualResponse.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
		})
	})
}

func TestHandler_UpdateUserHandler(t *testing.T) {
	userID := "1"
	path := "/users/" + userID
	Convey("Given update user request is invalid", t, func() {
		app := createTestApp()
		c := gomock.NewController(t)
		defer c.Finish()
		updateUserReq := `{`
		Convey("When update request is required", func() {
			mockService := mocks.NewMockService(c)

			handler := user.NewHandler(mockService)
			handler.RegisterRoutes(app)

			req := NewHTTPRequestWithJSON(http.MethodPut, path, updateUserReq)
			actualResponse, _ := app.Test(req)
			defer actualResponse.Body.Close()
			Convey("Then response status code should be 400", func() {
				So(actualResponse.StatusCode, ShouldEqual, http.StatusBadRequest)
			})
		})
	})
	Convey("Given update user request is valid", t, func() {
		app := createTestApp()
		c := gomock.NewController(t)
		defer c.Finish()
		updateUserReq := user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "doe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When update method is called with valid request", func() {
			mockService := mocks.NewMockService(c)
			mockService.EXPECT().Update(gomock.Any(), userID, &updateUserReq).Return(nil)

			handler := user.NewHandler(mockService)
			handler.RegisterRoutes(app)

			req := NewHTTPRequestWithJSON(http.MethodPut, path, updateUserReq)
			actualResponse, err := app.Test(req)
			defer actualResponse.Body.Close()
			So(err, ShouldBeNil)
			Convey("Then response status code should be 200", func() {
				So(actualResponse.StatusCode, ShouldEqual, http.StatusOK)
			})
		})
	})
	Convey("Given update user request is valid", t, func() {
		app := createTestApp()
		c := gomock.NewController(t)
		defer c.Finish()
		updateUserReq := user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "doe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When update method is called with valid request", func() {
			mockService := mocks.NewMockService(c)
			mockService.EXPECT().Update(gomock.Any(), userID, &updateUserReq).Return(errors.New("error"))

			handler := user.NewHandler(mockService)
			handler.RegisterRoutes(app)

			req := NewHTTPRequestWithJSON(http.MethodPut, path, updateUserReq)
			actualResponse, _ := app.Test(req)
			defer actualResponse.Body.Close()
			Convey("Then response status code should be 500", func() {
				So(actualResponse.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
		})
	})
}

func createTestApp() *fiber.App {
	return fiber.New()
}

func NewHTTPRequestWithJSON(method, url string, request interface{}) *http.Request {
	body, _ := json.Marshal(request)
	req := httptest.NewRequest(method, url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}
