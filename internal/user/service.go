package user

import (
	"context"
	"github.com/muratdemir0/faceit-task/pkg/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Store interface {
	Create(ctx context.Context, u *store.User) (string, error)
	Update(ctx context.Context, u *store.User) error
	Delete(ctx context.Context, userID string) error
	Get(ctx context.Context, userID string) (store.User, error)
	List(ctx context.Context, criteria store.ListCriteria) ([]store.User, error)
}

type Producer interface {
	Produce(ctx context.Context, topic string, message interface{}) error
}

type service struct {
	store    Store
	producer Producer
	logger   *zap.Logger
}

func NewService(store Store, producer Producer, logger *zap.Logger) Service {
	return &service{store: store, producer: producer, logger: logger}
}

func (s service) Create(ctx context.Context, req *CreateUserRequest) error {
	user := &store.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  req.Password, // TODO: hash password
		Email:     req.Email,
		Country:   req.Country,
	}
	userID, err := s.store.Create(ctx, user)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	kafkaErr := s.producer.Produce(ctx, KafkaUserCreatedTopic, Event{
		UserID:    userID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if kafkaErr != nil {
		s.logger.Error("failed to produce user created event", zap.Any("user", user))
		return errors.Wrap(kafkaErr, "failed to produce user created event")
	}
	return nil
}

func (s service) Update(ctx context.Context, userID string, req *UpdateUserRequest) error {
	u, getErr := s.store.Get(ctx, userID)
	if getErr != nil {
		s.logger.Error("failed to get user", zap.Error(getErr))
		return errors.Wrap(getErr, "failed to get user")
	}

	u.FirstName = req.FirstName
	u.LastName = req.LastName
	u.Nickname = req.Nickname
	u.Email = req.Email
	u.Country = req.Country
	if req.Password != "" {
		u.Password = req.Password
	}

	err := s.store.Update(ctx, &u)
	if err != nil {
		s.logger.Error("failed to update user", zap.Any("userId", userID), zap.Any("request", req))
		return errors.Wrap(err, "failed to update user")
	}

	kafkaErr := s.producer.Produce(ctx, KafkaUserUpdatedTopic, Event{
		UserID:    userID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	})

	if kafkaErr != nil {
		s.logger.Error("failed to produce user updated event", zap.Any("userId", userID))
		return errors.Wrap(kafkaErr, "failed to produce user updated event")
	}
	return nil
}

func (s service) Delete(ctx context.Context, userID string) error {
	err := s.store.Delete(ctx, userID)
	if err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	kafkaErr := s.producer.Produce(ctx, KafkaUserDeletedTopic, Event{
		UserID: userID,
	})
	if kafkaErr != nil {
		s.logger.Error("failed to produce user deleted event", zap.Any("userId", userID))
		return errors.Wrap(kafkaErr, "failed to produce user deleted event")
	}
	return nil
}

func (s service) List(ctx context.Context, criteria *ListUserRequest) (*Response, error) {
	c := store.ListCriteria{
		Country: criteria.Country,
		Page:    criteria.Page,
		Limit:   criteria.Limit,
	}
	users, err := s.store.List(ctx, c)
	if err != nil {
		s.logger.Error("failed to list users", zap.Error(err))
		return nil, err
	}

	var usersResponse Response

	for _, user := range users {
		usersResponse.Users = append(usersResponse.Users, User{
			ID:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Nickname:  user.Nickname,
			Password:  user.Password,
			Email:     user.Email,
			Country:   user.Country,
		})
	}

	return &usersResponse, nil
}
