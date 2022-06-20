package user_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/muratdemir0/faceit-task/internal/user"
	mocks "github.com/muratdemir0/faceit-task/mocks/user"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_Create(t *testing.T) {
	Convey("Given create user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserRepo := mocks.NewMockRepository(c)
		cur := &user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's create is called", func() {
			cu := &user.CreateUser{
				FirstName: cur.FirstName,
				LastName:  cur.LastName,
				Nickname:  cur.Nickname,
				Password:  cur.Password,
				Email:     cur.Email,
				Country:   cur.Country,
			}
			mockUserRepo.EXPECT().Create(gomock.Any(), cu).Return(nil)
			service := user.NewService(mockUserRepo)
			err := service.Create(context.TODO(), cur)
			Convey("Then user should be created", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given create user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserRepo := mocks.NewMockRepository(c)
		cur := &user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's create is called", func() {
			cu := &user.CreateUser{
				FirstName: cur.FirstName,
				LastName:  cur.LastName,
				Nickname:  cur.Nickname,
				Password:  cur.Password,
				Email:     cur.Email,
				Country:   cur.Country,
			}
			mockUserRepo.EXPECT().Create(gomock.Any(), cu).Return(errors.New("error"))
			service := user.NewService(mockUserRepo)
			err := service.Create(context.TODO(), cur)
			Convey("Then repo should be return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
