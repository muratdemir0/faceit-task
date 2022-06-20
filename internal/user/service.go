package user

import "context"

type Repository interface {
	Create(ctx context.Context, cu *User) error
	Update(ctx context.Context, cu *User) error
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
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
	return s.repo.Create(ctx, user)
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
	return s.repo.Update(ctx, user)
}
