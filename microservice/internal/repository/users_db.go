package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// filter :=  options.Find()
	users := []User{}
	doc, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	for doc.Next(context.TODO()) {
		user := User{}
		err := doc.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)

	}

	defer r.db.Disconnect(context.Background())

	return users, nil
}

func (r UserRepositoryDB) GetByID(id string) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	user := User{}
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &user, nil
}

func (r UserRepositoryDB) Update(u User) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	filter := bson.D{{Key: "username", Value: u.UserName}}
	update := bson.M{"$set": bson.M{"email": u.Email}}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &u, nil

}

func (r UserRepositoryDB) UpdateOne(id string, u User) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"username": u.UserName, "email": u.Email}}
	// update := bson.D{{Key: "email", Value: u.Email}, {Key: "updateat", Value: u.UpdateAt}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())

	return &u, nil
}

func (r UserRepositoryDB) DeleteOne(username string, u User) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	filter := bson.D{{Key: "username", Value: u.UserName}}
	update := bson.M{"$set": bson.M{"deleteat": u.DeleteAt}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())
	return &u, nil
}

// func (r UserRepositoryDB) FindUser(u User) ([]User, error) {
// 	collection := r.db.Database("account").Collection("users")
// 	users := []User{}
// 	// filter := bson.M{"$or": bson.M{"username": u.UserName, "email": u.Email}}
// 	findQuery := bson.M{"username": u.UserName}
// 	orQuery := []bson.M{}
// 	orQuery = append(orQuery, bson.M{"email": u.Email})
// 	findQuery["$or"] = orQuery

// 	doc, err := collection.Find(context.TODO(), findQuery)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for doc.Next(context.TODO()) {
// 		user := User{}
// 		err = doc.Decode(&user)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, user)

// 	}
// 	defer r.db.Disconnect(context.Background())

// 	return users, nil
// }

func (r UserRepositoryDB) FindUser(u User) (*User, error) {
	collection := r.db.Database("account").Collection("users")
	user := User{}
	// filter := bson.M{"$or": bson.M{"username": u.UserName, "email": u.Email}}
	// filter := bson.M{"$or": []interface{}{
	// 	bson.M{"username": u.UserName}, bson.M{"email": u.Email},
	// }}
	filter := bson.M{"$or": []interface{}{
		bson.M{"username": u.UserName}, bson.M{"email": u.Email},
	}}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &user, nil
}
