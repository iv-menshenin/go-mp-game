package game

import (
	"tomb_mates/internal/effects"
	"tomb_mates/internal/protos"
)

type AreaOfEffects struct {
	Area    *protos.Area
	Caster  *protos.Unit
	Effects []*effects.Effect
	Ttd     uint32
}

func (aoe *AreaOfEffects) AddAffectedUnit(unit *protos.Unit) {
	aoe.Area.AffectedUnitIds[unit.Id] = &protos.Empty{}
}

func (aoe *AreaOfEffects) RemoveAffectedUnit(unit *protos.Unit) {
	delete(aoe.Area.AffectedUnitIds, unit.Id)
}
