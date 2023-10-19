package engine

import (
	"log"

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

type Text struct {
	Content string
	start_x int32
	start_y int32
	end_x   int32
	end_y   int32
}

func (self *Text) MouseDown() bool {
	return mouse_state == sdl.BUTTON_LEFT && self.MouseOver()
}

func (self *Text) MouseOver() bool {
	return MouseX >= self.start_x &&
		MouseX <= self.end_x &&
		MouseY >= self.start_y &&
		MouseY <= self.end_y
}

func (self *Text) Draw() {
	renderer.SetDrawColor(255, 50, 200, 255)
	src_rect := sdl.Rect{0, 0, char_width, char_height}
	dst_rect := sdl.Rect{0, self.start_y, scaled_char_width, scaled_char_height}
	for i, c := range self.Content {
		ascii := int32(c) - 32
		if ascii > 95 || i < 0 {
			log.Println("Character", c, "is not in charmap")
			continue
		}
		charmap_x := ascii % cols
		charmap_y := ascii / cols

		src_rect.X = charmap_x * char_width
		src_rect.Y = charmap_y * char_height

		dst_rect.X = self.start_x + int32(i)*scaled_char_width

		renderer.Copy(
			charmap,
			&src_rect,
			&dst_rect,
		)
	}
}

func CreateAlignedText(content string, rel_x float64, rel_y float64, v_align Alignment, h_align Alignment) UIElement {
	text_width := int32(len(content)) * scaled_char_width
	text_height := scaled_char_height

	var abs_x, abs_y int32

	switch h_align {
	case UI_ALIGN_START:
		abs_x = int32(rel_x * width_f64)
	case UI_ALIGN_CENTER:
		abs_x = int32(rel_x*width_f64) - text_width/2
	case UI_ALIGN_END:
		abs_x = int32(rel_x*width_f64) - text_width
	}

	switch v_align {
	case UI_ALIGN_START:
		abs_y = int32(rel_y * height_f64)
	case UI_ALIGN_CENTER:
		abs_y = int32(rel_y*height_f64) - text_height/2
	case UI_ALIGN_END:
		abs_y = int32(rel_y*height_f64) - text_height
	}

	text := Text{
		Content: content,
		start_x: abs_x,
		start_y: abs_y,
		end_x:   abs_x + text_width,
		end_y:   abs_y + text_height,
	}
	return &text
}
