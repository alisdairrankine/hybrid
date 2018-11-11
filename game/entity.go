package game

type Entity struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	X      int32  `json:"-"`
	Y      int32  `json:"-"`
	Sprite string `json:"sprite,omitempty"`

	MaxHealth  int32 `json:"max_health,omitempty"`
	Health     int32 `json:"health,omitempty"`
	Level      int32 `json:"level,omitempty"`
	Experience int32 `json:"experience,omitempty"`
	Attack     int32 `json:"attack,omitempty"`
	Defence    int32 `json:"defence,omitempty"`

	Inventory []Item `json:"inventory,omitempty"`
}
