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

func NewGame(s *data.Settings) (*Game, error) {
	eventHandler, err := events.NewHandler(s)
	if err != nil {
		return nil, err
	}
	renderer, err := gfx.NewRenderer("Hybrid", s.ScreenWidth, s.ScreenHeight, s.Fullscreen)
	if err != nil {
		return nil, err
	}

	g := &Game{
		eventHandler: eventHandler,
		renderer:     renderer,
		ui:           gfx.NewUI(s),
		monsters:     []game.Entity{},
		player:       &game.Entity{},
	}

	return g, nil
}

func (g *Game) Run() {
	defer g.renderer.Close()
	tick := time.Tick(16 * time.Millisecond)

	for {
		select {
		case <-tick:
			event := g.eventHandler.PollForEvents()
			if event != nil {
				switch event.Type {
				case "quit":
					return
				case "inventory":
					g.ui.ToggleInventory()

				}
			}
			g.renderer.Render(append(g.monsters, *g.player), g.ui)
		}
	}
}
