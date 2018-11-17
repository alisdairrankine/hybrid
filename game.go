package hybrid

import (
	"time"

	"github.com/alisdairrankine/hybrid/data"
	"github.com/alisdairrankine/hybrid/events"
	"github.com/alisdairrankine/hybrid/game"
	"github.com/alisdairrankine/hybrid/gfx"
)

type Game struct {
	eventHandler *events.Handler
	renderer     *gfx.Renderer
	ui           *gfx.UI
	player       *game.Entity
	monsters     []game.Entity
}

func NewGame(s *data.Settings, gamepack *data.GamePack) (*Game, error) {
	eventHandler, err := events.NewHandler(s)
	if err != nil {
		return nil, err
	}
	renderer, err := gfx.NewRenderer("Hybrid", s.ScreenWidth, s.ScreenHeight, s.Fullscreen, gamepack)
	if err != nil {
		return nil, err
	}

	g := &Game{
		eventHandler: eventHandler,
		renderer:     renderer,
		ui:           gfx.NewUI(s),
		monsters:     []game.Entity{},
		player: &game.Entity{
			Name: "Ali The great",
			X:    20,
			Y:    10,
		},
	}

	return g, nil
}

func (g *Game) Run() {
	defer g.renderer.Close()
	tick := time.Tick(16 * time.Millisecond)
	lastTime := time.Now()
	for {
		delta := time.Since(lastTime)
		lastTime = time.Now()
		moveRate := int32(240 * delta.Seconds())
		XMove := int32(0)
		YMove := int32(0)
		select {
		case <-tick:
			events := g.eventHandler.PollForEvents()
			for _, event := range events {
				switch event.Type {
				case "quit":
					return
				case "inventory":
					g.ui.ToggleInventory()
				case "up":
					YMove--
				case "down":
					YMove++
				case "left":
					XMove--
				case "right":
					XMove++
				}
			}
			g.player.X += XMove * moveRate
			g.player.Y += YMove * moveRate
			g.renderer.Render(append(g.monsters, *g.player), g.ui)
		}
	}
}
