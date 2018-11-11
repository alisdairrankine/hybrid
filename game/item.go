package game

type Item struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`

	Durability      int32 `json:"durability,omitempty"`
	HealthModifier  int32 `json:"health_modifier,omitempty"`
	DefenceModifier int32 `json:"defence_modifier,omitempty"`
	AttackModifier  int32 `json:"attack_modifier,omitempty"`

	DamageRate          int32 `json:"damage_rate,omitempty"`
	HealthModifierRate  int32 `json:"health_modifier_rate,omitempty"`
	DefenceModifierRate int32 `json:"defence_modifier_rate,omitempty"`
	AttackModifierRate  int32 `json:"attack_modifier_rate,omitempty"`
}

func (i *Item) Use(e *Entity, delta int32) {
	i.Durability -= i.DamageRate * delta

	if i.Durability > 0 {
		e.Attack += i.AttackModifier
		e.Defence += i.DefenceModifier

		if (i.HealthModifier + e.Health) < e.MaxHealth {
			e.Health += i.HealthModifier
		} else {
			e.Health = e.MaxHealth
		}
		i.HealthModifier += i.HealthModifierRate * delta
		i.DefenceModifier += i.DefenceModifierRate * delta
		i.AttackModifier += i.AttackModifierRate * delta

	}
}
