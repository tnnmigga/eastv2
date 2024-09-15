package example1

import (
	"eastv2/game/modules/play/userops/userdata"
	"eastv2/pb"

	"github.com/tnnmigga/corev2/log"
)

func onSayHelloReq(user *userdata.Entity, req *pb.SayHelloReq) {
	log.Infof("onSayHelloReq %d %s", user.ID, req.Text)
}
