package main

import (
	"fmt"

	"github.com/alisdairrankine/hybrid"
	"github.com/alisdairrankine/hybrid/data"
)

func main() {

	settings := &data.Settings{}
	game, err := hybrid.NewGame(settings)
	if err != nil {
		fmt.Println("error sarting game: %s", err)
		return
	}
	game.Run()
}
