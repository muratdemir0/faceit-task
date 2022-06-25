package store

import (
	"context"
	"github.com/muratdemir0/faceit-task/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	client *mongo.Client
	config *config.Mongo
}
type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Nickname  string `json:"nickname" bson:"nickname"`
	Password  string `json:"password" json:"password"`
	Email     string `json:"email" bson:"email"`
	Country   string `json:"country" bson:"country"`
}
type ListCriteria struct {
	Country string `json:"country" bson:"country"`
}

func NewUserStore(client *mongo.Client, config *config.Mongo) UserStore {
	return UserStore{client: client, config: config}
}

func (s UserStore) Create(ctx context.Context, cu *User) error {
	_, err := s.client.
		Database(s.config.Name).
		Collection(s.config.Collections.Users).
		InsertOne(ctx, cu)

	if err != nil {
		return err
	}
	return nil
}
func (s UserStore) Update(ctx context.Context, u *User) error {
	_, err := s.client.
		Database(s.config.Name).
		Collection(s.config.Collections.Users).
		UpdateOne(ctx, u.ID, u)

	if err != nil {
		return err
	}
	return nil
}

func (s UserStore) Delete(ctx context.Context, userID string) error {
	_, err := s.client.
		Database(s.config.Name).
		Collection(s.config.Collections.Users).
		DeleteOne(ctx, bson.D{{"_id", userID}})
	if err != nil {
		return err
	}
	return nil
}
func (s UserStore) Get(ctx context.Context, userID string) (User, error) {
	user := User{}
	err := s.client.
		Database(s.config.Name).
		Collection(s.config.Collections.Users).
		FindOne(ctx, bson.D{{"_id", userID}}).
		Decode(user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (s UserStore) List(ctx context.Context, criteria ListCriteria) ([]User, error) {
	var users []User
	var filter bson.D
	if criteria.Country != "" {
		filter = bson.D{{"country", bson.D{{"$eq", criteria.Country}}}}
	}
	cursor, err := s.client.
		Database(s.config.Name).
		Collection(s.config.Collections.Users).
		Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	cursorErr := cursor.All(ctx, &users)
	if cursorErr != nil {
		return nil, cursorErr
	}

	return users, nil
}
