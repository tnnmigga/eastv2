package impl

import (
	"eastv2/game/modules/play/domain"
	"eastv2/game/modules/play/domain/impl/msg"
	"eastv2/game/modules/play/domain/impl/player"
	"eastv2/game/modules/play/domain/impl/timer"
)

func Init(d *domain.Domain) {
	d.PutCase(domain.MsgCaseIndex, msg.New(d))
	// d.PutCase(domain.EventCaseIndex, event.New(d))
	d.PutCase(domain.TimerCaseIndex, timer.New(d))
	d.PutCase(domain.UserCaseIndex, player.New(d))
}
