package store

import (
	"context"
	"github.com/muratdemir0/faceit-task/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	client *mongo.Client
	config *config.Mongo
}
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func NewUserStore(client *mongo.Client, config *config.Mongo) UserStore {
	return UserStore{client: client, config: config}
}

func (us UserStore) Create(ctx context.Context, cu *User) error {
	_, err := us.client.Database(us.config.Name).Collection(us.config.Collections.Users).InsertOne(ctx, cu)
	if err != nil {
		return err
	}
	return nil
}
