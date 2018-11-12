package main

import (
	"fmt"

	"github.com/alisdairrankine/hybrid"
	"github.com/alisdairrankine/hybrid/data"
)

func main() {

	settings := data.LoadSettings("settings.json")

	game, err := hybrid.NewGame(settings)
	if err != nil {
		fmt.Println("error starting game: %s", err)
		return
	}
	game.Run()
}
