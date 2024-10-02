package abilities

import (
	"time"
	"tomb_mates/internal/effects"
	"tomb_mates/internal/protos"
)

type FirePizza = Ability

var damageEffect = &effects.Effect{
	Name:         "Fire Pizza Damage",
	Type:         effects.TypeDamage,
	Damage:       10,
	FriendlyFire: false,
}

func NewFirePizza(caster *protos.Unit) *FirePizza {
	return &FirePizza{
		Name:        "Fire Pizza",
		Description: "Cast fire pizza on the gorund to burn your enemies!",
		Caster:      caster,
		Type:        TypeActive,
		Cooldown:    time.Second * 3,
		Cost: AbilityCost{
			Type:  CostTypeNone,
			Value: 0,
		},
		Target: AbilityTarget{
			Type: TargetTypeArea,
			Area: &TargetArea{
				Range:  100,
				Radius: 20,
			},
		},
		Effects: []*effects.Effect{
			damageEffect,
		},
	}
}
