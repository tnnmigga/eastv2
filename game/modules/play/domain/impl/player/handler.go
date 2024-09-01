package player

import (
	"eastv2/game/modules/play/pm"
	"eastv2/pb"
)

func (c *useCase) onSayHelloReq(p *pm.UserData, req *pb.SayHelloReq) {

}

func (c *useCase) onRPCTest(req *pb.TestRPC, resolve func(any), reject func(error)) {
	resolve(&pb.TestRPCRes{
		V: 22,
	})
}
