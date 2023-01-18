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
		UserName:       u.UserName,
		HashPassword:   u.HashPassword,
		Email:          u.Email,
		EmailConfirm:   u.EmailConfirm,
		UserStatus:     u.UserStatus,
		LoginAttempt:   u.LoginAttempt,
		AcceptTerms:    u.AcceptTerms,
		VerifyIdentity: u.VerifyIdentity,
		LastLogin:      u.LastLogin,
		CreateAt:       u.CreateAt,
		UpdateAt:       u.UpdateAt,
		DeleteAt:       u.DeleteAt,
	}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &doc, nil

}

// func (r UserGroupRepositoryDB) GetAll() ([]User, error) {
// 	return nil, nil
// }
