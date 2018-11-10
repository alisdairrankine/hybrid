package gfx

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tilemap struct {
	Height int32          `json:"height,omitempty"`
	Width  int32          `json:"width,omitempty"`
	Layers []TilemapLayer `json:"layers,omitempty"`
}

type TilemapLayer struct {
	Data   []int32 `json:"data,omitempty"`
	Height int32   `json:"height,omitempty"`
	Width  int32   `json:"width,omitempty"`
}

func LoadTilemap(filename string) (*Tilemap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open tilemap: %s", err)
	}
	dec := json.NewDecoder(file)

	tilemap := &Tilemap{}

	err = dec.Decode(tilemap)

	if err != nil {
		return nil, fmt.Errorf("could not decode tilemap: %s", err)
	}

	return tilemap, err
}
