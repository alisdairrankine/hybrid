package gfx

import (
	"github.com/alisdairrankine/hybrid/data"
	"github.com/alisdairrankine/hybrid/game"

	"github.com/veandco/go-sdl2/sdl"
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

func (ui *UI) Render(surface *sdl.Surface, player *game.Entity) {
	if ui.inventoryScreen {
		ui.renderInventoryScreen(surface, player)
	}
	if ui.characterScreen {
		ui.renderCharacterScreen(surface, player)
	}
}

func (ui *UI) renderInventoryScreen(surface *sdl.Surface, player *game.Entity) {

}

func (ui *UI) renderCharacterScreen(surface *sdl.Surface, player *game.Entity) {

}
