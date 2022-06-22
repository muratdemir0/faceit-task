package user

import (
	"context"
	"github.com/muratdemir0/faceit-task/pkg/store"
)

type Store interface {
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, userID string) error
	Find(ctx context.Context, userID string) (store.User, error)
	FindBy(ctx context.Context, criteria store.FindBy) ([]store.User, error)
}
type service struct {
	store Store
}

func NewService(store Store) Service {
	return &service{store: store}
}

func (s service) Create(ctx context.Context, req *CreateUserRequest) error {
	user := &User{
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
	user := &User{
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

func (s service) FindBy(ctx context.Context, criteria *FindUserRequest) (*Response, error) {
	findBy := store.FindBy{
		Country: criteria.Country,
		Email:   criteria.Email,
	}
	users, err := s.store.FindBy(ctx, findBy)
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
