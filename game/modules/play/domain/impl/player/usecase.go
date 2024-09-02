package player

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/api"
	"eastv2/game/modules/play/user"
	"eastv2/pb"
)

type useCase struct {
	*domain.Domain
}

func New(d *domain.Domain) api.IUser {
	c := &useCase{
		Domain: d,
	}
	user.HandleMsg[pb.SayHelloReq](c, c.onSayHelloReq)
	return c
}
