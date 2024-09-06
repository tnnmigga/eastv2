package userops

import (
	"eastv2/game/modules/play/userops/muser"
	"reflect"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/message/codec"
)

var msgHandler = map[reflect.Type]func(p *muser.Model, msg any){}

func HandleMsg[T any](m iface.IModule, handler func(m *muser.Model, msg *T)) {
	mType := reflect.TypeOf(new(T))
	if _, ok := msgHandler[mType]; ok {
		panic("msg handler already exists")
	}
	codec.Register[T]()
	// 注册消息处理函数
	msgHandler[reflect.TypeOf(new(T))] = func(m *muser.Model, msg any) {
		handler(m, msg.(*T))
	}
}
