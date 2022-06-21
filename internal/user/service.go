package user

import "context"

type Store interface {
	Create(ctx context.Context, cu *User) error
	Update(ctx context.Context, cu *User) error
	Delete(ctx context.Context, userID string) error
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
