package user

import (
	"context"
	"github.com/muratdemir0/faceit-task/pkg/store"
	"github.com/pkg/errors"
)

type Store interface {
	Create(ctx context.Context, u *store.User) error
	Update(ctx context.Context, u *store.User) error
	Delete(ctx context.Context, userID string) error
	Get(ctx context.Context, userID string) (store.User, error)
	List(ctx context.Context, criteria store.ListCriteria) ([]store.User, error)
}

type Producer interface {
	Produce(ctx context.Context, topic string, message interface{}) error
}

type GenerateUUID func() string

type service struct {
	store      Store
	producer   Producer
	generateID GenerateUUID
}

func NewService(store Store, producer Producer, uuidGenerator GenerateUUID) Service {
	return &service{store: store, producer: producer, generateID: uuidGenerator}
}

func (s service) Create(ctx context.Context, req *CreateUserRequest) error {
	user := &store.User{
		ID:        s.generateID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  req.Password, // TODO: hash password
		Email:     req.Email,
		Country:   req.Country,
	}
	err := s.store.Create(ctx, user)
	if err != nil {
		return err
	}
	kafkaErr := s.producer.Produce(ctx, KafkaUserCreatedTopic, Event{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if kafkaErr != nil {
		return errors.Wrap(kafkaErr, "failed to produce user created event")
	}
	return nil
}

func (s service) Update(ctx context.Context, userID string, req *UpdateUserRequest) error {
	user := &store.User{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  req.Password, // TODO: check password
		Email:     req.Email,
		Country:   req.Country,
	}
	err := s.store.Update(ctx, user)
	if err != nil {
		return err
	}

	kafkaErr := s.producer.Produce(ctx, KafkaUserUpdatedTopic, Event{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})

	if kafkaErr != nil {
		return errors.Wrap(kafkaErr, "failed to produce user updated event")
	}
	return nil
}

func (s service) Delete(ctx context.Context, userID string) error {
	err := s.store.Delete(ctx, userID)
	if err != nil {
		return err
	}
	kafkaErr := s.producer.Produce(ctx, KafkaUserDeletedTopic, Event{
		UserID: userID,
	})
	if kafkaErr != nil {
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
		return nil, err
	}

	var usersResponse Response

	for _, user := range users {
		usersResponse.Users = append(usersResponse.Users, User{
			ID:        user.ID,
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
