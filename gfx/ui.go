package gfx

import (
	"github.com/alisdairrankine/hybrid/data"
	"github.com/alisdairrankine/hybrid/game"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type UI struct {
	mainMenu        bool
	inventoryScreen bool
	characterScreen bool
	settings        *data.Settings
}

func NewUI(settings *data.Settings) *UI {
	return &UI{
		settings: settings,
	}
}

func (ui *UI) ToggleInventory() {
	ui.inventoryScreen = !ui.inventoryScreen
}

func (ui *UI) ToggleCharacterScreen() {
	ui.characterScreen = !ui.characterScreen
}

func (ui *UI) Render(surface *sdl.Surface, player *game.Entity, font *ttf.Font) {
	if ui.inventoryScreen {
		ui.renderInventoryScreen(surface, player, font)
	}
	if ui.characterScreen {
		ui.renderCharacterScreen(surface, player, font)
	}

	if player != nil {
		text, _ := font.RenderUTF8Solid(player.Name, sdl.Color{R: 255, G: 255, B: 255, A: 255})
		text.Blit(nil, surface, &sdl.Rect{X: 10, Y: 10})
	}
}

func (ui *UI) renderInventoryScreen(surface *sdl.Surface, player *game.Entity, font *ttf.Font) {

}

func (ui *UI) renderCharacterScreen(surface *sdl.Surface, player *game.Entity, font *ttf.Font) {

}
