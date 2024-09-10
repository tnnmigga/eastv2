package domain

import (
	"eastv2/game/modules/play/domain/api"

	"github.com/tnnmigga/corev2/basic/domainops"
	"github.com/tnnmigga/corev2/iface"
)

type Domain struct {
	iface.IModule
	domainops.Root
}

func New(m iface.IModule) *Domain {
	d := &Domain{
		Root:    domainops.New(m, caseMaxIndex),
		IModule: m,
	}
	return d
}

const (
	caseMinIndex = iota
	Example1Index
	Example2Index
	caseMaxIndex
)

func (d *Domain) Example1Case() api.IExample1 {
	return d.GetCase(Example1Index).(api.IExample1)
}

func (d *Domain) Example2Case() api.IExample2 {
	return d.GetCase(Example2Index).(api.IExample2)
}
