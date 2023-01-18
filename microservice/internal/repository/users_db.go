package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r UserRepositoryDB) GetAll() ([]User, error) {
	collection := r.db.Database("account").Collection("users")
	findOption := options.Find()
	users := []User{}
	doc, err := collection.Find(context.TODO(), bson.D{{}}, findOption)
	if err != nil {
		return nil, err
	}
	for doc.Next(context.TODO()) {
		var user User
		err := doc.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)

	}

	defer r.db.Disconnect(context.Background())

	return users, nil
}

func (r UserRepositoryDB) GetById(id string) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	filter := bson.D{{"_id", id}}
	user := User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &user, nil
}
