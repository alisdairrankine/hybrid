package events

import (
	"github.com/alisdairrankine/hybrid/data"
	"github.com/veandco/go-sdl2/sdl"
)

type Handler struct {
}

func NewHandler(s *data.Settings) (*Handler, error) {
	handler := &Handler{}
	return handler, nil
}

func (h *Handler) PollForEvents() *Event {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return &Event{Type: "quit"}
		}
	}
	return nil
}

type Event struct {
	Type string
}
