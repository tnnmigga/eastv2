package example1

import (
	"eastv2/game/modules/play/userops/muser"
	"eastv2/pb"

	"github.com/tnnmigga/corev2/log"
)

func onSayHelloReq(user *muser.Model, req *pb.SayHelloReq) {
	log.Infof("onSayHelloReq %d %s", user.ID, req.Text)
}
