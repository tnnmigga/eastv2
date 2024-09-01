package msg

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/api"
)

type useCase struct {
	*domain.Domain
}

func New(d *domain.Domain) api.IMsg {
	s := &useCase{
		Domain: d,
	}
	return s
}

// func (s *service) Notify(userID uint64, msg any) {
// 	msgbus.Cast(&pb.S2CPackage{
// 		UserID: userID,
// 		Body:   codec.Encode(msg),
// 	})
// }
