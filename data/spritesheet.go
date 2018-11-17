package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Spritesheet struct {
	Filename       string           `json:"filename,omitempty"`
	Name           string           `json:"name,omitempty"`
	SpritePosition map[string]int32 `json:"sprite_position,omitempty"`
	SpriteSize     int32            `json:"sprite_size,omitempty"`
}

func LoadSpritesheet(filename string) (*Spritesheet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load gamepack from manifest: %s", err)
	}

	dec := json.NewDecoder(file)
	sheet := &Spritesheet{}
	err = dec.Decode(sheet)
	if err != nil {
		return nil, fmt.Errorf("could not load gamepack from manifest: %s", err)
	}
	return sheet, nil
}
