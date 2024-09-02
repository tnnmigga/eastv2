package domain

import (
	"eastv2/game/modules/play/domain/api"

	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/module/domain"
)

type Domain struct {
	iface.IModule
	domain.Root
}

func New(m iface.IModule) *Domain {
	d := &Domain{
		Root:    domain.New(m, MaxCaseIndex),
		IModule: m,
	}
	return d
}

const (
	MsgCaseIndex = iota
	EventCaseIndex
	TimerCaseIndex
	MaxCaseIndex
)

func (d *Domain) MsgCase() api.IMsg {
	return d.GetCase(MsgCaseIndex).(api.IMsg)
}

func (d *Domain) EventCase() api.IEvent {
	return d.GetCase(EventCaseIndex).(api.IEvent)
}

func (d *Domain) TimerCase() api.ITimer {
	return d.GetCase(TimerCaseIndex).(api.ITimer)
}
