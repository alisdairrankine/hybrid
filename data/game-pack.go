package data

import "github.com/alisdairrankine/hybrid/game"

type GamePack struct {
	Races []game.Race `json:"races,omitempty"`
}
