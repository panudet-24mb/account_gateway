package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDB struct {
	db *mongo.Client
}

func NewUserRepositoryDB(db *mongo.Client) UserRepository {
	return UserRepositoryDB{db: db}
}

func (r UserRepositoryDB) Create(u User) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	doc := User{
		UserName: u.UserName,
	}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &doc, nil

}
