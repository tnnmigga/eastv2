package userdata

import "fmt"

type Entity struct {
	ID  uint64 `bson:"_id"`
	Tag *Tag   `bson:"-"`
}

func New(id uint64) *Entity {
	return &Entity{
		ID: id,
	}
}

func groupKey(uid uint64) string {
	return fmt.Sprintf("mongo.userdata.%d", uid)
}

type Tag struct {
	IsOnline bool  // 是否在线
	GwID     int32 // 网关id
}
