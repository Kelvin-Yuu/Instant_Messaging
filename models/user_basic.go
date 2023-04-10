package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity  string `bson:"_id"`
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

func GetUserBasicByIdentity(identity primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(ub)
	return ub, err
}
