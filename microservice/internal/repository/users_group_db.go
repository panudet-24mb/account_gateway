package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		GroupName:      u.GroupName,
		GroupCode:      u.GroupCode,
		GroupDesc:      u.GroupDesc,
		GroupImgUrl:    u.GroupImgUrl,
		GroupIcon:      u.GroupIcon,
		GroupIconColor: u.GroupIconColor,
		GroupType:      u.GroupType,
		GroupStatus:    u.GroupStatus,
		GroupActive:    u.GroupActive,
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

func (r UserGroupRepositoryDB) GetAll() ([]Group, error) {
	collection := r.db.Database("account").Collection("users_group")
	groups := []Group{}
	doc, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	for doc.Next(context.TODO()) {
		group := Group{}
		err := doc.Decode(&group)
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, group)
	}

	defer r.db.Disconnect(context.Background())

	return groups, nil

}

func (r UserGroupRepositoryDB) GetByID(id string) (*Group, error) {
	collection := r.db.Database("account").Collection("users_group")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	group := Group{}
	err = collection.FindOne(context.TODO(), filter).Decode(&group)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &group, nil
}

func (r UserGroupRepositoryDB) Update(Group Group) (*Group, error) {
	collection := r.db.Database("account").Collection("users_group")
	filter := bson.D{{Key: "groupname", Value: Group.GroupName}}
	update := bson.M{"$set": bson.M{"groupdesc": Group.GroupDesc, "groupimgurl": Group.GroupImgUrl, "groupicon": Group.GroupIcon, "groupiconcolor": Group.GroupIconColor, "grouptype": Group.GroupType, "groupstatus": Group.GroupType, "groupactive": Group.GroupActive, "updateat": Group.UpdateAt}}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &Group, nil
}

func (r UserGroupRepositoryDB) UpdateOne(id string, Group Group) (*Group, error) {
	collection := r.db.Database("account").Collection("users_group")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"groupimgurl": Group.GroupImgUrl, "groupicon": Group.GroupIcon, "groupiconcolor": Group.GroupIconColor, "grouptype": Group.GroupType, "groupstatus": Group.GroupType, "groupactive": Group.GroupActive, "updateat": Group.UpdateAt}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())

	return &Group, nil

}

func (r UserGroupRepositoryDB) DeleteOne(groupName string, Group Group) (*Group, error) {
	collection := r.db.Database("account").Collection("users_group")
	filter := bson.D{{Key: "groupname", Value: Group.GroupName}}
	update := bson.M{"$set": bson.M{"deleteat": Group.DeleteAt}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())
	return &Group, nil

}

func (r UserGroupRepositoryDB) FindGroup(g Group) ([]Group, error) {
	collection := r.db.Database("account").Collection("users")
	groups := []Group{}
	filter := bson.M{"$or": []interface{}{
		bson.M{"groupname": g.GroupName}, bson.M{"groupcode": g.GroupCode}, bson.M{"groupdesc": g.GroupDesc}, bson.M{"grouptype": g.GroupType}, bson.M{"groupstatus": g.GroupStatus},
	}}
	doc, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for doc.Next(context.TODO()) {
		group := Group{}
		err = doc.Decode(&group)
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, group)

	}
	defer r.db.Disconnect(context.Background())

	return groups, nil
}
