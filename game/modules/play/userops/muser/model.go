package muser

import "fmt"

type Model struct {
	ID  uint64 `bson:"_id"`
	Tag *Tag   `bson:"-"`
}

func New(id uint64) *Model {
	return &Model{
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
