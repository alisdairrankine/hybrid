package events

import (
	"github.com/alisdairrankine/hybrid/data"
	"github.com/veandco/go-sdl2/sdl"
)

type Handler struct {
	settings *data.Settings
}

func NewHandler(s *data.Settings) (*Handler, error) {
	handler := &Handler{settings: s}
	return handler, nil
}

func (h *Handler) PollForEvents() *Event {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return &Event{Type: "quit"}
		case *sdl.KeyboardEvent:
			e := event.(*sdl.KeyboardEvent)
			return h.translateKeyToEvent(e)
		}
	}
	return nil
}

func (h *Handler) translateKeyToEvent(e *sdl.KeyboardEvent) *Event {
	switch e.Keysym.Sym {
	case h.settings.KeyInventory:
		if e.State == sdl.RELEASED {
			return &Event{Type: "inventory"}
		}
	case h.settings.KeyUp:
		if e.State == sdl.PRESSED {
			return &Event{Type: "up"}
		}
	case h.settings.KeyDown:
		if e.State == sdl.PRESSED {
			return &Event{Type: "down"}
		}
	case h.settings.KeyLeft:
		if e.State == sdl.PRESSED {
			return &Event{Type: "left"}
		}
	case h.settings.KeyRight:
		if e.State == sdl.PRESSED {
			return &Event{Type: "right"}
		}
	case h.settings.KeyAttack:
		if e.State == sdl.PRESSED {
			return &Event{Type: "attack"}
		}
	}
	return nil
}

type Event struct {
	Type string
}
