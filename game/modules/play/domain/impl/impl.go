package impl

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/impl/example1"
	"eastv2/game/modules/play/domain/impl/example2"
)

func Init(d *domain.Domain) {
	example1.Init(d)
	example2.Init(d)
}
