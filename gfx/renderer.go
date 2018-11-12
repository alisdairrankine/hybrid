package gfx

import (
	"fmt"

	"github.com/alisdairrankine/hybrid/game"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	width       int32
	height      int32
	title       string
	window      *sdl.Window
	spritesheet *Spritesheet
	tilemap     *Tilemap
}

func NewRenderer(title string, width, height int32, isFullScreen bool) (*Renderer, error) {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	if i := img.Init(img.INIT_PNG); i != img.INIT_PNG {
		return nil, fmt.Errorf("could not initialise image libraries")
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)

	if isFullScreen {
		window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	}

	if err != nil {
		return nil, fmt.Errorf("could not create window: %s", err)
	}

	spritesheet, err := NewSpriteSheet("assets/terrain.png", int32(32))
	if err != nil {
		return nil, fmt.Errorf("could not load spritesheet: %s", err)
	}

	tilemap, err := LoadTilemap("assets/map1.json")
	if err != nil {
		return nil, fmt.Errorf("could not load tilemap: %s", err)
	}

	rnd := &Renderer{
		window:      window,
		title:       title,
		width:       width,
		height:      height,
		spritesheet: spritesheet,
		tilemap:     tilemap,
	}

	return rnd, nil
}

func (r *Renderer) Render(entities []game.Entity, ui *UI) error {
	surface, err := r.window.GetSurface()
	if err != nil {
		return err
	}

	white := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	surface.FillRect(nil, white.Uint32())

	for _, layer := range r.tilemap.Layers {
		for i, raw := range layer.Data {
			if raw == int32(0) {
				continue
			}
			tile := raw - 1
			rect := &sdl.Rect{
				W: 32,
				H: 32,
				X: (32) * (int32(i) % (layer.Width)),
				Y: (32) * (int32(i) / (layer.Width)),
			}

			r.spritesheet.RenderToSurface(tile, rect, surface)
		}
	}

	r.window.UpdateSurface()
	return nil
}

func (r *Renderer) Close() {
	r.window.Destroy()
	img.Quit()
	sdl.Quit()
}
