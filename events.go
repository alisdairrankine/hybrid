package hybrid

import "github.com/veandco/go-sdl2/sdl"

type EventHandler struct {
}

func NewEventHandler() (*EventHandler, error) {
	handler := &EventHandler{}
	return handler, nil
}

func (h *EventHandler) PollForEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return
		}
	}
}
