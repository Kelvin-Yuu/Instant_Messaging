package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avartar   string `bson:"avartar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
	// "accout": "账号",
	// "password": "密码",
	// "nickname": "昵称",
	// "sex": "1",
	// "email": "邮箱",
	// "avartar": "头像",
	// "created_at": "1",
	// "updated_at": "1"
}

func (ub UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"identity", identity}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicCountByEmail(email string) (int64, error) {

	return MongoDB.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}

func GetUserBasicCountByAccount(account string) (int64, error) {

	return MongoDB.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"account", account}})
}

func InsertOneUserBasic(ub *UserBasic) error {

	_, err := MongoDB.Collection(UserBasic{}.CollectionName()).
		InsertOne(context.Background(), ub)
	if err != nil {
		return err
	}
	return nil

}

func GetUserBasicByAccount(account string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}}).
		Decode(ub)
	return ub, err
}
