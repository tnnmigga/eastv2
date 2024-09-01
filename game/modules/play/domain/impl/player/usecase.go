package player

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/api"
	"eastv2/game/modules/play/pm"
	"eastv2/pb"
)

type useCase struct {
	*domain.Domain
}

func New(d *domain.Domain) api.IMsg {
	c := &useCase{
		Domain: d,
	}
	pm.HandleUserMsg[pb.SayHelloReq](c.onSayHelloReq)
	return c
}
