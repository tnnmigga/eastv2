package example1

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/api"
	"eastv2/game/modules/play/userops"

	"github.com/tnnmigga/corev2/basic/domainops"
)

var uc *useCase

type useCase struct {
	*domain.Domain
}

func Init(d *domain.Domain) {
	uc = &useCase{
		Domain: d,
	}
	domainops.RegisterCase[api.IExample1](d, domain.Example1Index, uc)
	userops.HandleMsg(d, onSayHelloReq)
}
