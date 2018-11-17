package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alisdairrankine/hybrid/game"
)

type GamePack struct {
	Races        []game.Race   `json:"races,omitempty"`
	Levels       []game.Level  `json:"levels,omitempty"`
	Items        []game.Item   `json:"items,omitempty"`
	Languages    []LanguageSet `json:"languages,omitempty"`
	Spritesheets []Spritesheet `json:"spritesheets,omitempty"`
}

type manifest struct {
	Languages    []string `json:"languages,omitempty"`
	Races        []string `json:"races,omitempty"`
	Items        []string `json:"items,omitempty"`
	Levels       []string `json:"levels,omitempty"`
	Spritesheets []string `json:"spritesheets,omitempty"`
}

func LoadGamePackFromManifest(filename string) (*GamePack, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load gamepack from manifest: %s", err)
	}

	dec := json.NewDecoder(file)
	m := &manifest{}
	err = dec.Decode(m)
	if err != nil {
		return nil, fmt.Errorf("could not load gamepack from manifest: %s", err)
	}

	//load spritesheets
	spritesheets := []Spritesheet{}
	for _, spritesheetFileName := range m.Spritesheets {
		sheet, err := LoadSpritesheet(spritesheetFileName)
		if err != nil {
			return nil, fmt.Errorf("could not load spritesheet: %s", err)
		}
		spritesheets = append(spritesheets, *sheet)
	}

	pack := &GamePack{
		Spritesheets: spritesheets,
	}
	return pack, nil
}
