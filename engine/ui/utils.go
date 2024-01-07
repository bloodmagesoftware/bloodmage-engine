// Bloodmage Engine - Retro first person game engine
// Copyright (C) 2024  Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	// default rect
	defaultRect = &sdl.Rect{X: 0, Y: 0, W: 0, H: 0}
)

// parseColor parses a color string and returns the color channel values.
//
// Possible formats:
//   - #RRGGBB
//   - #RRGGBBAA
//   - rgb(R, G, B)
//   - rgba(R, G, B, A)
func ParseColorChannels(color string) (r uint8, g uint8, b uint8, a uint8, err error) {
	// Remove any whitespace
	color = strings.ReplaceAll(color, " ", "")

	// Check for hex color format
	if strings.HasPrefix(color, "#") {
		color = strings.TrimPrefix(color, "#")

		switch len(color) {
		case 6:
			// #RRGGBB format
			if r, err = parseHex(color[0:2]); err != nil {
				return 0, 0, 0, 0, err
			}
			if g, err = parseHex(color[2:4]); err != nil {
				return 0, 0, 0, 0, err
			}
			if b, err = parseHex(color[4:6]); err != nil {
				return 0, 0, 0, 0, err
			}
			a = 255
		case 8:
			// #RRGGBBAA format
			if r, err = parseHex(color[0:2]); err != nil {
				return 0, 0, 0, 0, err
			}
			if g, err = parseHex(color[2:4]); err != nil {
				return 0, 0, 0, 0, err
			}
			if b, err = parseHex(color[4:6]); err != nil {
				return 0, 0, 0, 0, err
			}
			if a, err = parseHex(color[6:8]); err != nil {
				return 0, 0, 0, 0, err
			}
		default:
			return 0, 0, 0, 0, fmt.Errorf("invalid hex color format: %s", color)
		}
	} else if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
		// cut off the "rgb(" and ")" using indexes
		// split the values by comma
		rgbValues := strings.Split(color[4:len(color)-1], ",")
		if len(rgbValues) != 3 {
			return 0, 0, 0, 0, fmt.Errorf("invalid rgb color format: %s", color)
		}
		if r64, err := strconv.ParseUint(rgbValues[0], 10, 8); err == nil {
			r = uint8(r64)
		} else {
			return 0, 0, 0, 0, err
		}
		if g64, err := strconv.ParseUint(rgbValues[1], 10, 8); err == nil {
			g = uint8(g64)
		} else {
			return 0, 0, 0, 0, err
		}
		if b64, err := strconv.ParseUint(rgbValues[2], 10, 8); err == nil {
			b = uint8(b64)
		} else {
			return 0, 0, 0, 0, err
		}
		a = 255
	} else if strings.HasPrefix(color, "rgba(") && strings.HasSuffix(color, ")") {
		// cut off the "rgba(" and ")" using indexes
		// split the values by comma
		rgbaValues := strings.Split(color[5:len(color)-1], ",")
		if len(rgbaValues) != 4 {
			return 0, 0, 0, 0, fmt.Errorf("invalid rgba color format: %s", color)
		}
		if r64, err := strconv.ParseUint(rgbaValues[0], 10, 8); err == nil {
			r = uint8(r64)
		} else {
			return 0, 0, 0, 0, err
		}
		if g64, err := strconv.ParseUint(rgbaValues[1], 10, 8); err == nil {
			g = uint8(g64)
		} else {
			return 0, 0, 0, 0, err
		}
		if b64, err := strconv.ParseUint(rgbaValues[2], 10, 8); err == nil {
			b = uint8(b64)
		} else {
			return 0, 0, 0, 0, err
		}
		if a64, err := strconv.ParseUint(rgbaValues[3], 10, 8); err == nil {
			a = uint8(a64)
		} else {
			return 0, 0, 0, 0, err
		}
	} else {
		return 0, 0, 0, 0, fmt.Errorf("unsupported color format: %s", color)
	}

	return r, g, b, a, nil
}

// Helper function to parse hexadecimal values
func parseHex(hex string) (uint8, error) {
	value, err := strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return 0, err
	}
	return uint8(value), nil
}

// ParseColor parses a color string into an sdl.Color struct.
func ParseColor(color string) (sdl.Color, error) {
	r, g, b, a, err := ParseColorChannels(color)
	return sdl.Color{R: r, G: g, B: b, A: a}, err
}
