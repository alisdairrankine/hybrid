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

func (h *Handler) PollForEvents() []Event {
	events := h.checkMovement()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return []Event{Event{Type: "quit"}}
		case *sdl.KeyboardEvent:
			e := event.(*sdl.KeyboardEvent)
			event := h.translateKeyToEvent(e)
			if event != nil {
				events = append(events, *event)
			}
		}
	}
	return events
}

func (h *Handler) checkMovement() []Event {
	events := []Event{}
	state := sdl.GetKeyboardState()

	for key, on := range state {
		if on > 0 {
			switch sdl.GetKeyFromScancode(sdl.Scancode(key)) {
			case h.settings.KeyUp:
				events = append(events, Event{Type: "up"})
			case h.settings.KeyDown:
				events = append(events, Event{Type: "down"})
			case h.settings.KeyLeft:
				events = append(events, Event{Type: "left"})
			case h.settings.KeyRight:
				events = append(events, Event{Type: "right"})
			case h.settings.KeyAttack:
				events = append(events, Event{Type: "attack"})
			}
		}
	}
	return events
}

func (h *Handler) translateKeyToEvent(e *sdl.KeyboardEvent) *Event {
	if h.settings.KeyInventory == e.Keysym.Sym && e.State == sdl.RELEASED {
		return &Event{Type: "inventory"}
	}
	return nil
}

type Event struct {
	Type string
}
