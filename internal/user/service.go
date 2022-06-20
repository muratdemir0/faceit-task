package user

import "context"

type Repository interface {
	Create(ctx context.Context, cu *CreateUser) error
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) Create(ctx context.Context, req *CreateUserRequest) error {
	cu := &CreateUser{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  req.Password, // TODO: hash password
		Email:     req.Email,
		Country:   req.Country,
	}
	return s.repo.Create(ctx, cu)
}
