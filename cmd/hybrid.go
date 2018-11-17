package main

import (
	"fmt"

	"github.com/alisdairrankine/hybrid"
	"github.com/alisdairrankine/hybrid/data"
)

func main() {

	settings := data.LoadSettings("settings.json")
	gamepack, err := data.LoadGamePackFromManifest("data_files/manifest.json")
	if err != nil {
		fmt.Printf("error loading gamepack: %s\n", err)
		return
	}
	game, err := hybrid.NewGame(settings, gamepack)
	if err != nil {
		fmt.Printf("error starting game: %s\n", err)
		return
	}
	game.Run()
}
