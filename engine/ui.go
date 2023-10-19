package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Alignment uint8

const (
	UI_ALIGN_START Alignment = iota
	UI_ALIGN_CENTER
	UI_ALIGN_END
)

const (
	char_width  = int32(7)
	char_height = int32(9)
	cols        = 18
	rows        = 6
)

var (
	charmap            *sdl.Texture
	scaled_char_width  int32
	scaled_char_height int32
)

func initUI() {
	charmap_surface, err := sdl.LoadBMP("./assets/textures/charmap.bmp")
	if err != nil {
		panic(err)
	}
	defer charmap_surface.Free()
	charmap, err = renderer.CreateTextureFromSurface(charmap_surface)
	if err != nil {
		panic(err)
	}
	scaled_char_width = char_width * options.PixelScale
	scaled_char_height = char_height * options.PixelScale
}

type UIElement interface {
	Draw()
	MouseDown() bool
	MouseOver() bool
}
