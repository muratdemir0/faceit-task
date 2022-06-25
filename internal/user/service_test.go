package user_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/muratdemir0/faceit-task/internal/user"
	mocks2 "github.com/muratdemir0/faceit-task/mocks"
	producerMock "github.com/muratdemir0/faceit-task/mocks/event"
	mocks "github.com/muratdemir0/faceit-task/mocks/user"
	"github.com/muratdemir0/faceit-task/pkg/store"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Create(t *testing.T) {
	Convey("Given create user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		cur := &user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}

		ids := []string{
			"24f945a5-d4a3-455b-9555-05c50d06a77e",
		}
		mockUUID := mocks2.NewMockGenerateIDFunc(ids...)
		Convey("When repo's create is called", func() {
			cu := &store.User{
				ID:        ids[0],
				FirstName: cur.FirstName,
				LastName:  cur.LastName,
				Nickname:  cur.Nickname,
				Password:  cur.Password,
				Email:     cur.Email,
				Country:   cur.Country,
			}
			mockUserStore.EXPECT().Create(gomock.Any(), cu).Return(nil)
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Create(context.TODO(), cur)
			Convey("Then user should be created", func() {
				So(err, ShouldBeNil)
				Convey("Then user created event should be produced", func() {
					actualEvents := mockProducer.Events[user.KafkaUserCreatedTopic]
					expectedCreatedEvent := []interface{}{user.Event{
						UserID:    cu.ID,
						FirstName: cu.FirstName,
						LastName:  cu.LastName,
						Email:     cu.Email,
					}}
					So(actualEvents, ShouldResemble, expectedCreatedEvent)
				})
			})

		})
	})

	Convey("Given create user request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		cur := &user.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		ids := []string{
			"24f945a5-d4a3-455b-9555-05c50d06a77e",
		}
		mockUUID := mocks2.NewMockGenerateIDFunc(ids...)
		Convey("When repo's create is called", func() {
			cu := &store.User{
				ID:        ids[0],
				FirstName: cur.FirstName,
				LastName:  cur.LastName,
				Nickname:  cur.Nickname,
				Password:  cur.Password,
				Email:     cur.Email,
				Country:   cur.Country,
			}
			mockUserStore.EXPECT().Create(gomock.Any(), cu).Return(errors.New("error"))
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
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
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		updateUserReq := &user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's update method is called", func() {
			expectedUser := store.User{
				ID:        userID,
				FirstName: updateUserReq.FirstName,
				LastName:  updateUserReq.LastName,
				Nickname:  updateUserReq.Nickname,
				Password:  updateUserReq.Password,
				Email:     updateUserReq.Email,
				Country:   updateUserReq.Country,
			}
			mockUserStore.EXPECT().Get(gomock.Any(), userID).Return(expectedUser, nil)
			mockUserStore.EXPECT().Update(gomock.Any(), &expectedUser).Return(nil)
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Update(context.TODO(), userID, updateUserReq)
			Convey("Then user should be updated", func() {
				So(err, ShouldBeNil)
				Convey("Then user updated event should be produced", func() {
					actualEvents := mockProducer.Events[user.KafkaUserUpdatedTopic]
					expectedCreatedEvent := []interface{}{user.Event{
						UserID:    expectedUser.ID,
						FirstName: expectedUser.FirstName,
						LastName:  expectedUser.LastName,
						Email:     expectedUser.Email,
					}}
					So(actualEvents, ShouldResemble, expectedCreatedEvent)
				})
			})
		})
	})

	Convey("Given update user request is valid and user id is not exist", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		updateUserReq := &user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's update method is called", func() {
			mockUserStore.EXPECT().Get(gomock.Any(), userID).Return(store.User{}, errors.Wrap(store.NotFoundError, "failed to get user"))
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Update(context.TODO(), userID, updateUserReq)
			expectedError := errors.Wrap(errors.Wrap(store.NotFoundError, "failed to get user"), "failed to get user")
			Convey("Then repo should return an error which is not found", func() {
				assert.EqualError(t, err, expectedError.Error())
			})
		})
	})

	Convey("Given update user request is valid and user id is exist", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		updateUserReq := &user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's update method is called", func() {
			expectedUser := store.User{
				FirstName: updateUserReq.FirstName,
				LastName:  updateUserReq.LastName,
				Nickname:  updateUserReq.Nickname,
				Password:  updateUserReq.Password,
				Email:     updateUserReq.Email,
				Country:   updateUserReq.Country,
			}
			mockUserStore.EXPECT().Get(gomock.Any(), userID).Return(expectedUser, nil)
			mockUserStore.EXPECT().Update(gomock.Any(), &expectedUser).Return(errors.New("error"))
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Update(context.TODO(), userID, updateUserReq)
			Convey("Then repo should return an error which is not found", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given update user request and user id is exist", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		updateUserReq := &user.UpdateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "jdoe",
			Password:  "123456",
			Email:     "john@doe.com",
			Country:   "UK",
		}
		Convey("When repo's update method is called", func() {
			expectedUser := store.User{
				FirstName: updateUserReq.FirstName,
				LastName:  updateUserReq.LastName,
				Nickname:  updateUserReq.Nickname,
				Password:  updateUserReq.Password,
				Email:     updateUserReq.Email,
				Country:   updateUserReq.Country,
			}
			mockUserStore.EXPECT().Get(gomock.Any(), userID).Return(expectedUser, nil)
			mockUserStore.EXPECT().Update(gomock.Any(), &expectedUser).Return(nil)
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Update(context.TODO(), userID, updateUserReq)
			Convey("Then repo should not return an error", func() {
				So(err, ShouldBeNil)
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
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		Convey("When repo's delete method is called", func() {
			mockUserStore.EXPECT().Delete(gomock.Any(), userID).Return(nil)
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Delete(context.TODO(), userID)
			Convey("Then repo should be deleted", func() {
				So(err, ShouldBeNil)
				Convey("Then user updated event should be produced", func() {
					actualEvents := mockProducer.Events[user.KafkaUserDeletedTopic]
					expectedCreatedEvent := []interface{}{user.Event{
						UserID: userID,
					}}
					So(actualEvents, ShouldResemble, expectedCreatedEvent)
				})
			})
		})
	})

	Convey("Given user id", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		Convey("When repo's delete method is called", func() {
			mockUserStore.EXPECT().Delete(gomock.Any(), userID).Return(errors.New("error"))
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			err := service.Delete(context.TODO(), userID)
			Convey("Then repo should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestService_List(t *testing.T) {
	Convey("Given country code request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		req := user.ListUserRequest{
			Country: "UK",
		}
		findUserStore := store.ListCriteria{
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
			mockUserStore.EXPECT().List(gomock.Any(), findUserStore).Return(users, nil)
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			acutalResponse, _ := service.List(context.TODO(), &req)
			Convey("Then repo should return users", func() {
				So(acutalResponse, ShouldResemble, &expectedResponse)
			})
		})

	})
	Convey("Given country code request is valid", t, func() {
		c := gomock.NewController(t)
		defer c.Finish()
		mockUserStore := mocks.NewMockStore(c)
		mockProducer := producerMock.NewEventProducerMock()
		mockUUID := mocks2.NewMockGenerateIDFunc()
		req := user.ListUserRequest{
			Country: "UK",
		}
		findUserStore := store.ListCriteria{
			Country: req.Country,
		}
		Convey("When repo's findBy method is called", func() {
			mockUserStore.EXPECT().List(gomock.Any(), findUserStore).Return(nil, errors.New("error"))
			service := user.NewService(mockUserStore, mockProducer, mockUUID)
			_, err := service.List(context.TODO(), &req)
			Convey("Then repo should return users", func() {
				So(err, ShouldNotBeNil)
			})
		})

	})
}
