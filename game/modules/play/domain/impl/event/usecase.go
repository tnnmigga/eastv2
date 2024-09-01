package event

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/api"
)

type useCase struct {
	*domain.Domain
}

func New(d *domain.Domain) api.IEvent {
	return &useCase{
		Domain: d,
	}
}
