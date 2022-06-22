package user_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/muratdemir0/faceit-task/internal/user"
	mocks "github.com/muratdemir0/faceit-task/mocks/user"
	"github.com/muratdemir0/faceit-task/pkg/store"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_Create(t *testing.T) {
	Convey("Given create user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		cur := &user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's create is called", func() {
			cu := &store.User{
				FirstName: cur.FirstName,
				LastName:  cur.LastName,
				Nickname:  cur.Nickname,
				Password:  cur.Password,
				Email:     cur.Email,
				Country:   cur.Country,
			}
			mockUserStore.EXPECT().Create(gomock.Any(), cu).Return(nil)
			service := user.NewService(mockUserStore)
			err := service.Create(context.TODO(), cur)
			Convey("Then user should be created", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given create user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		cur := &user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's create is called", func() {
			cu := &store.User{
				FirstName: cur.FirstName,
				LastName:  cur.LastName,
				Nickname:  cur.Nickname,
				Password:  cur.Password,
				Email:     cur.Email,
				Country:   cur.Country,
			}
			mockUserStore.EXPECT().Create(gomock.Any(), cu).Return(errors.New("error"))
			service := user.NewService(mockUserStore)
			err := service.Create(context.TODO(), cur)
			Convey("Then repo should be return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestService_Update(t *testing.T) {
	userID := "123"
	Convey("Given update user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		updateUserReq := &user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's update method is called", func() {
			cu := &store.User{
				ID:        userID,
				FirstName: updateUserReq.FirstName,
				LastName:  updateUserReq.LastName,
				Nickname:  updateUserReq.Nickname,
				Password:  updateUserReq.Password,
				Email:     updateUserReq.Email,
				Country:   updateUserReq.Country,
			}
			mockUserStore.EXPECT().Update(gomock.Any(), cu).Return(nil)
			service := user.NewService(mockUserStore)
			err := service.Update(context.TODO(), userID, updateUserReq)
			Convey("Then user should be updated", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given update user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		updateUserReq := &user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's update method is called", func() {
			cu := &store.User{
				ID:        userID,
				FirstName: updateUserReq.FirstName,
				LastName:  updateUserReq.LastName,
				Nickname:  updateUserReq.Nickname,
				Password:  updateUserReq.Password,
				Email:     updateUserReq.Email,
				Country:   updateUserReq.Country,
			}
			mockUserStore.EXPECT().Update(gomock.Any(), cu).Return(errors.New("error"))
			service := user.NewService(mockUserStore)
			err := service.Update(context.TODO(), userID, updateUserReq)
			Convey("Then repo should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestService_Delete(t *testing.T) {
	userID := "123"
	Convey("Given user id", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		Convey("When repo's delete method is called", func() {
			mockUserStore.EXPECT().Delete(gomock.Any(), userID).Return(nil)
			service := user.NewService(mockUserStore)
			err := service.Delete(context.TODO(), userID)
			Convey("Then repo should be deleted", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given user id", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		Convey("When repo's delete method is called", func() {
			mockUserStore.EXPECT().Delete(gomock.Any(), userID).Return(errors.New("error"))
			service := user.NewService(mockUserStore)
			err := service.Delete(context.TODO(), userID)
			Convey("Then repo should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestService_FindBy(t *testing.T) {
	Convey("Given country code request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		req := user.FindUserRequest{
			Country: "UK",
		}
		findUserStore := store.FindBy{
			Country: req.Country,
		}
		users := []store.User{
			{
				ID:        "",
				FirstName: "",
				LastName:  "",
				Nickname:  "",
				Password:  "",
				Email:     "",
				Country:   "",
			},
		}
		expectedResponse := user.Response{
			Users: []user.User{
				{
					ID:        "",
					FirstName: "",
					LastName:  "",
					Nickname:  "",
					Password:  "",
					Email:     "",
					Country:   "",
				},
			},
		}
		Convey("When repo's findBy method is called", func() {
			mockUserStore.EXPECT().FindBy(gomock.Any(), findUserStore).Return(users, nil)
			service := user.NewService(mockUserStore)
			acutalResponse, _ := service.FindBy(context.TODO(), &req)
			Convey("Then repo should return users", func() {
				So(acutalResponse, ShouldResemble, &expectedResponse)
			})
		})

	})
	Convey("Given country code request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		req := user.FindUserRequest{
			Country: "UK",
		}
		findUserStore := store.FindBy{
			Country: req.Country,
		}
		Convey("When repo's findBy method is called", func() {
			mockUserStore.EXPECT().FindBy(gomock.Any(), findUserStore).Return(nil, errors.New("error"))
			service := user.NewService(mockUserStore)
			_, err := service.FindBy(context.TODO(), &req)
			Convey("Then repo should return users", func() {
				So(err, ShouldNotBeNil)
			})
		})

	})
}
