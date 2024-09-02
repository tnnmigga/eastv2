package player

import (
	"eastv2/game/modules/play/user/model"
	"eastv2/pb"
)

func (c *useCase) onSayHelloReq(p *model.Model, req *pb.SayHelloReq) {

}

func (c *useCase) onRPCTest(req *pb.TestRPC, resolve func(any), reject func(error)) {
	resolve(&pb.TestRPCRes{
		V: 22,
	})
}
