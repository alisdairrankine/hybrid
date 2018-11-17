package gfx

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Spritesheet struct {
	sdlSurface  *sdl.Surface
	spriteSize  int32
	spriteCount int32
	spriteRects map[int32]sdl.Rect
}

func NewSpriteSheet(filename string, spriteSize int32) (*Spritesheet, error) {

	surface, err := img.Load(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load spritesheet: %s", err)
	}

	count := (surface.W / spriteSize) * (surface.H / spriteSize)

	rects := map[int32]sdl.Rect{}

	for i := int32(0); i < count; i++ {
		x := (i % (surface.W / spriteSize))
		y := (i / (surface.W / spriteSize))
		r := sdl.Rect{
			W: spriteSize,
			H: spriteSize,
			X: (spriteSize) * x,
			Y: (spriteSize) * y,
		}

		rects[i] = r
	}

	ss := &Spritesheet{
		sdlSurface:  surface,
		spriteSize:  spriteSize,
		spriteCount: count,
		spriteRects: rects,
	}

	return ss, nil
}

func (s *Spritesheet) spriteRect(spriteID int32) *sdl.Rect {

	if rect, exists := s.spriteRects[spriteID]; exists {
		return &rect
	}

	return &sdl.Rect{}
}

func (s *Spritesheet) RenderToSurface(sprite int32, dst *sdl.Rect, surface *sdl.Surface) {
	s.sdlSurface.Blit(s.spriteRect(sprite), surface, dst)
}
