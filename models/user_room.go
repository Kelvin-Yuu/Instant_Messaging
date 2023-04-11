package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	RoomType     int    `bson:"room_type"` // 房间 类型 【1-独聊房间 2-群聊房间】
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByUserIdRoomId(userId, roomId string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := MongoDB.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"user_identity", userId}, {"room_identity", roomId}}).
		Decode(ur)
	return ur, err
}

func GetUserRoomByRoomId(roomIdentity string) ([]*UserRoom, error) {
	cursor, err := MongoDB.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		return nil, err
	}

	urs := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		if err := cursor.Decode(ur); err != nil {
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}
