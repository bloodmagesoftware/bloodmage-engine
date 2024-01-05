package font

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (f *font) Rune(r rune) (*sdl.Texture, *sdl.Rect, error) {
	// get rune as number
	rn := uint32(r)

	// check if rune is in range
	if rn < f.startChar || rn > f.endChar {
		return nil, nil, fmt.Errorf("rune %v (%v) is not in range %v-%v", rn, r, f.startChar, f.endChar)
	}

	// get rune index
	runeIndex := rn - f.startChar

	// get rune position in texture atlas raster (x, y)
	runeX := runeIndex % uint32(f.collumnCount)
	runeY := runeIndex / uint32(f.collumnCount)

	// get rune position in texture atlas (x, y, w, h)
	runeRect := &sdl.Rect{
		X: int32(runeX) * f.charWidth,
		Y: int32(runeY) * f.charHeight,
		W: f.charWidth,
		H: f.charHeight,
	}

	// get rune texture
	runeTexture, err := f.texture.Texture()
	if err != nil {
		return nil, nil, errors.Join(
			errors.New("Failed to get font texture"),
			err,
		)
	}

	return runeTexture, runeRect, nil
}
