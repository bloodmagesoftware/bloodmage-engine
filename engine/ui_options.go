package engine

func CreateOptions(opt ...string) []UIElement {
	opt_elems := make([]UIElement, len(opt))
	opt_elems_height := int32(len(opt)) * scaled_char_height
	opt_elems_abs_y := calcStartPos(opt_elems_height, 0.5, height_f64, UI_ALIGN_CENTER)
	for i, o := range opt {
		text_width := int32(len(o)) * scaled_char_width
		text_height := scaled_char_height
		text_abs_x := calcStartPos(text_width, 0.5, width_f64, UI_ALIGN_CENTER)
		text_abs_y := opt_elems_abs_y + int32(i)*text_height

		text := ui_text{
			o,
			text_abs_x, text_abs_y,
			text_abs_x + text_width, text_abs_y + text_height,
		}
		opt_elems[i] = &text
	}
	return opt_elems
}
