package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoleRepositoryDB struct {
	db *mongo.Client
}

func NewRoleRepositoryDB(db *mongo.Client) RoleRepository {
	return RoleRepositoryDB{db: db}
}

func (r RoleRepositoryDB) Create(role Roles) (*Roles, error) {
	collection := r.db.Database("account").Collection("role")
	doc := Roles{
		RoleName:      role.RoleName,
		RoleCode:      role.RoleCode,
		RoleDesc:      role.RoleDesc,
		RoleImgUrl:    role.RoleImgUrl,
		RoleIcon:      role.RoleIcon,
		RoleIconColor: role.RoleIconColor,
		RoleType:      role.RoleType,
		RoleActive:    role.RoleActive,
		CreateAt:      role.CreateAt,
		UpdateAt:      role.UpdateAt,
		DeleteAt:      role.DeleteAt,
	}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &doc, nil
}

func (r RoleRepositoryDB) GetAll() ([]Roles, error) {
	collection := r.db.Database("account").Collection("role")
	roles := []Roles{}
	doc, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	for doc.Next(context.TODO()) {
		role := Roles{}
		err := doc.Decode(&role)
		if err != nil {
			log.Fatal(err)
		}
		roles = append(roles, role)
	}

	defer r.db.Disconnect(context.Background())

	return roles, nil
}

func (r RoleRepositoryDB) GetByID(id string) (*Roles, error) {
	collection := r.db.Database("account").Collection("role")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	role := Roles{}
	err = collection.FindOne(context.TODO(), filter).Decode(&role)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &role, nil

}

func (r RoleRepositoryDB) Update(role Roles) (*Roles, error) {
	collection := r.db.Database("account").Collection("role")
	filter := bson.D{{Key: "rolename", Value: role.RoleName}}
	update := bson.M{"$set": bson.M{"roledesc": role.RoleDesc, "roleimgurl": role.RoleImgUrl, "roleicon": role.RoleIcon, "roleiconcolor": role.RoleIconColor, "roletype": role.RoleType, "rolestatus": role.RoleStatus, "roleactive": role.RoleActive, "updateat": role.UpdateAt}}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())

	return &role, nil

}

func (r RoleRepositoryDB) UpdateOne(id string, role Roles) (*Roles, error) {
	collection := r.db.Database("account").Collection("role")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"rolename": role.RoleName, "roledesc": role.RoleDesc, "roleimgurl": role.RoleImgUrl, "roleicon": role.RoleIcon, "roleiconcolor": role.RoleIconColor, "roletype": role.RoleType, "rolestatus": role.RoleStatus, "roleactive": role.RoleActive, "updateat": role.UpdateAt}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())

	return &role, nil
}

func (r RoleRepositoryDB) DeleteOne(roleName string, role Roles) (*Roles, error) {
	collection := r.db.Database("account").Collection("role")
	filter := bson.D{{Key: "rolename", Value: role.RoleName}}
	update := bson.M{"$set": bson.M{"deleteat": role.DeleteAt}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())
	return &role, nil
}

type UserRoleRepositoryDB struct {
	db *mongo.Client
}

func NewUserRoleRepositoryDB(db *mongo.Client) UserRoleRepository {
	return UserRoleRepositoryDB{db: db}
}

func (r UserRoleRepositoryDB) Create(u UsersRole) (*UsersRole, error) {
	collection := r.db.Database("account").Collection("users_role")
	doc := UsersRole{
		UserID:   u.UserID,
		RoleID:   u.RoleID,
		CreateAt: u.CreateAt,
		UpdateAt: u.UpdateAt,
		DeleteAt: u.DeleteAt,
	}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &doc, nil
}

func (r UserRoleRepositoryDB) GetAll() ([]UsersRole, error) {
	collection := r.db.Database("account").Collection("users_role")
	usersRole := []UsersRole{}
	doc, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	for doc.Next(context.TODO()) {
		userRole := UsersRole{}
		err := doc.Decode(&userRole)
		if err != nil {
			return nil, err
		}
		usersRole = append(usersRole, userRole)
	}
	defer r.db.Disconnect(context.Background())

	return usersRole, nil
}

func (r UserRoleRepositoryDB) GetByID(id string) (*UsersRole, error) {
	collection := r.db.Database("account").Collection("users_role")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	userRole := UsersRole{}
	err = collection.FindOne(context.TODO(), filter).Decode(&userRole)
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (r UserRoleRepositoryDB) Update(u UsersRole) (*UsersRole, error) {
	collection := r.db.Database("account").Collection("users_role")
	filter := bson.D{{Key: "userid", Value: u.UserID}}
	update := bson.M{"$set": bson.M{"roleid": u.RoleID, "updateat": u.UpdateAt}}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	defer r.db.Disconnect(context.Background())

	return &u, nil

}
func (r UserRoleRepositoryDB) UpdateOne(id string, u UsersRole) (*UsersRole, error) {
	collection := r.db.Database("account").Collection("users_role")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"userid": u.UserID, "roleid": u.RoleID, "updateat": u.UpdateAt}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.db.Disconnect(context.Background())
	return &u, nil

}

func (r UserRoleRepositoryDB) DeleteOne(id string, u UsersRole) (*UsersRole, error) {
	collection := r.db.Database("account").Collection("users_role")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"deleteat": u.DeleteAt}}
	idFilter, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	if idFilter.MatchedCount == 0 {
		return nil, nil
	}
	defer r.db.Disconnect(context.Background())
	return &u, nil
}
