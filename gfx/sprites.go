package gfx

import (
	"fmt"

	"github.com/alisdairrankine/hybrid/data"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteCollection struct {
	sheets map[string]*Spritesheet
}

func NewSpriteCollection(spritesheets ...data.Spritesheet) (*SpriteCollection, error) {

	collection := &SpriteCollection{
		sheets: make(map[string]*Spritesheet),
	}

	for _, sheet := range spritesheets {
		ss, err := NewSpriteSheet(sheet.Filename, sheet.SpriteSize)
		if err != nil {
			return nil, fmt.Errorf("could not create sprite collection: %s", err)
		}
		collection.sheets[sheet.Name] = ss
	}

	return collection, nil
}

func (c *SpriteCollection) SpriteSheet(sheet string) *Spritesheet {
	if s, exists := c.sheets[sheet]; exists {
		return s
	}
	return nil
}

func (c *SpriteCollection) RenderToSurface(sheet string, sprite int32, dst *sdl.Rect, surface *sdl.Surface) {
	if s, exists := c.sheets[sheet]; exists {
		s.RenderToSurface(sprite, dst, surface)
	}
}
