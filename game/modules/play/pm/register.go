package pm

import (
	"reflect"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/message/codec"
)

var msgHandler = map[reflect.Type]func(p *UserData, msg any){}

func RegMsgHandler[T any](m iface.IModule, handler func(p *UserData, msg *T)) {
	mType := reflect.TypeOf(new(T))
	if _, ok := msgHandler[mType]; ok {
		panic("msg handler already exists")
	}
	codec.Register[T]()
	// 注册消息处理函数
	msgHandler[reflect.TypeOf(new(T))] = func(p *UserData, msg any) {
		handler(p, msg.(*T))
	}
}
