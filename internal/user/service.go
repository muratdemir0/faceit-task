package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/muratdemir0/faceit-task/pkg/store"
)

type Store interface {
	Create(ctx context.Context, u *store.User) error
	Update(ctx context.Context, u *store.User) error
	Delete(ctx context.Context, userID string) error
	Get(ctx context.Context, userID string) (store.User, error)
	List(ctx context.Context, criteria store.FindBy) ([]store.User, error)
}
type service struct {
	store Store
	uuid  uuid.UUID
}

func NewService(store Store, u uuid.UUID) Service {
	return &service{store: store, uuid: u}
}

func (s service) Create(ctx context.Context, req *CreateUserRequest) error {
	user := &store.User{
		ID:        s.uuid.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  req.Password, // TODO: hash password
		Email:     req.Email,
		Country:   req.Country,
	}
	return s.store.Create(ctx, user)
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
	return s.store.Update(ctx, user)
}

func (s service) Delete(ctx context.Context, userID string) error {
	return s.store.Delete(ctx, userID)
}

func (s service) List(ctx context.Context, criteria *FindUserRequest) (*Response, error) {
	findBy := store.FindBy{
		Country: criteria.Country,
	}
	users, err := s.store.List(ctx, findBy)
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
