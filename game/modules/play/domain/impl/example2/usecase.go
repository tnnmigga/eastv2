package example2

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/api"

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
	domainops.RegisterCase[api.IExample2](d, domain.Example2Index, uc)
}
