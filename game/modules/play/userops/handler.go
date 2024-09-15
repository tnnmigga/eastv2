package userops

import (
	"eastv2/game/modules/play/userops/userdata"
	"eastv2/pb"
	"reflect"

	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message/codec"
	"github.com/tnnmigga/corev2/utils"
)

func onC2SPackage(req *pb.C2SPackage) {
	msg, err := codec.Decode(req.Body)
	if err != nil {
		log.Errorf("onC2SPackage decode error %s", err)
		return
	}
	LoadAsync(req.UserID, func(m *userdata.Entity, err error) {
		if err != nil {
			log.Errorf("onC2SPackage LoadAsync error %v", err)
			return
		}
		h, ok := msgHandler[reflect.TypeOf(msg)]
		if !ok {
			log.Error("onC2SPackage msg handler not found %s", utils.TypeName(msg))
			return
		}
		h(m, msg)
	})
}
