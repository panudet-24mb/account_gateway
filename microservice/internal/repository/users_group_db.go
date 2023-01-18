package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserGroupRepositoryDB struct {
	db *mongo.Client
}

func NewUserGroupRepositoryDB(db *mongo.Client) UserGroupRepository {
	return UserGroupRepositoryDB{db: db}
}

func (r UserGroupRepositoryDB) Create(u Group) (*Group, error) {
	collection := r.db.Database("account").Collection("users_group")
	doc := Group{
		GroupName:   u.GroupName,
		GroupCode:   u.GroupCode,
		GroupDesc:   u.GroupDesc,
		GroupType:   u.GroupType,
		GroupStatus: u.GroupStatus,
		GroupActive: u.GroupActive,
		CreateAt:    u.CreateAt,
		UpdateAt:    u.UpdateAt,
		DeleteAt:    u.DeleteAt,
	}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &doc, nil

}
