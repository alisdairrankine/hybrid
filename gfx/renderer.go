package gfx

import (
	"fmt"

	"github.com/alisdairrankine/hybrid/data"
	"github.com/alisdairrankine/hybrid/game"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Renderer struct {
	width            int32
	height           int32
	title            string
	window           *sdl.Window
	spriteCollection *SpriteCollection
	tilemap          *Tilemap
	fonts            map[string]*ttf.Font
}

func NewRenderer(title string, width, height int32, isFullScreen bool, gamepack *data.GamePack) (*Renderer, error) {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	err = ttf.Init()
	if err != nil {
		return nil, err
	}

	font, err := ttf.OpenFont("assets/font.ttf", 16)
	if err != nil {
		return nil, err
	}
	fonts := map[string]*ttf.Font{
		"ui": font,
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

	collection, err := NewSpriteCollection(gamepack.Spritesheets...)
	if err != nil {
		return nil, fmt.Errorf("could not load sprite collection: %s", err)
	}

	tilemap, err := LoadTilemap("assets/map1.json")
	if err != nil {
		return nil, fmt.Errorf("could not load tilemap: %s", err)
	}

	rnd := &Renderer{
		window:           window,
		title:            title,
		width:            width,
		height:           height,
		spriteCollection: collection,
		tilemap:          tilemap,
		fonts:            fonts,
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
	sheet := r.spriteCollection.SpriteSheet("terrain")

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

			sheet.RenderToSurface(tile, rect, surface)
		}
	}

	player := entities[len(entities)-1]
	r.spriteCollection.RenderToSurface("player", 0, &sdl.Rect{
		X: player.X,
		Y: player.Y,
		W: 32,
		H: 32,
	}, surface)

	ui.Render(surface, &player, r.fonts["ui"])

	r.window.UpdateSurface()
	return nil
}

func (r *Renderer) Close() {
	for _, font := range r.fonts {
		font.Close()
	}
	r.window.Destroy()
	img.Quit()
	sdl.Quit()
}
