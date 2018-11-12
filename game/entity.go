package game

type Entity struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	Race Race `json:"race,omitempty"`

	X int32 `json:"-,omitempty"`
	Y int32 `json:"-,omitempty"`

	MaxHealth  int32 `json:"max_health,omitempty"`
	Health     int32 `json:"health,omitempty"`
	Level      int32 `json:"level,omitempty"`
	Experience int32 `json:"experience,omitempty"`
	Attack     int32 `json:"attack,omitempty"`
	Defence    int32 `json:"defence,omitempty"`

	Inventory []Item `json:"inventory,omitempty"`
}

type Race struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Sprite      string   `json:"sprite,omitempty"`
	Attributes  []string `json:"attributes,omitempty"`
}
