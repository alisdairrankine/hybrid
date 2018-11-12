package data

import (
	"encoding/json"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Settings represents the configuration of the game
type Settings struct {

	//graphics
	ScreenWidth  int32 `json:"screen_width"`
	ScreenHeight int32 `json:"screen_height"`
	Fullscreen   bool  `json:"fullscreen"`

	//audio
	MusicVolume int `json:"music_volume"`
	FXVolume    int `json:"fx_volume"`

	//controls
	KeyInventory sdl.Keycode `json:"key_inventory"`
	KeyUp        sdl.Keycode `json:"key_up"`
	KeyDown      sdl.Keycode `json:"key_down"`
	KeyLeft      sdl.Keycode `json:"key_left"`
	KeyRight     sdl.Keycode `json:"key_right"`
	KeyAttack    sdl.Keycode `json:"key_attack"`

	Language string `json:"language"`
}

// LoadSettings loads a settings file from the given filename, or loads sane
// defaults
func LoadSettings(filename string) *Settings {
	if filename == "" {
		return loadDefaultSettings()
	}
	file, err := os.Open(filename)
	if err != nil {
		return loadDefaultSettings()
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	s := &Settings{}
	dec.Decode(s)
	return s
}

// SaveSettings saves the settings to a file
func SaveSettings(filename string, settings Settings) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	return enc.Encode(settings)
}

func loadDefaultSettings() *Settings {
	return &Settings{
		ScreenWidth:  800,
		ScreenHeight: 600,
		Fullscreen:   false,
		MusicVolume:  50,
		FXVolume:     50,
		KeyInventory: sdl.K_i,
		KeyUp:        sdl.K_w,
		KeyDown:      sdl.K_s,
		KeyLeft:      sdl.K_a,
		KeyRight:     sdl.K_d,
		KeyAttack:    sdl.K_SPACE,
		Language:     "en_GB",
	}
}
