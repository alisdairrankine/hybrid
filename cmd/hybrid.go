package main

import (
	"fmt"
	"time"

	"github.com/alisdairrankine/hybrid/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	renderer, err := gfx.NewRenderer("Hybrid", 800, 600)
	defer renderer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	tick := time.Tick(16 * time.Millisecond)

	for {
		select {
		case <-tick:
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					return
				}
			}
		}
		renderer.Render()
	}
}
