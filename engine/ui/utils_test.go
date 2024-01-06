package ui_test

import (
	"fmt"
	"testing"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/ui"
)

func TestParseColorChannels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		color string
		r     uint8
		g     uint8
		b     uint8
		a     uint8
		err   bool
	}{
		{"#000000", 0, 0, 0, 255, false},
		{"#00000000", 0, 0, 0, 0, false},
		{"#ffffff", 255, 255, 255, 255, false},
		{"#ffffffff", 255, 255, 255, 255, false},
		{"#ff0000", 255, 0, 0, 255, false},
		{"#ff000000", 255, 0, 0, 0, false},
		{"#00ff00", 0, 255, 0, 255, false},
		{"#00ff0000", 0, 255, 0, 0, false},
		{"#0000ff", 0, 0, 255, 255, false},
		{"#0000ff00", 0, 0, 255, 0, false},
		{"rgb(0, 0, 0)", 0, 0, 0, 255, false},
		{"rgb(11,22,23)", 11, 22, 23, 255, false},
		{"rgb(11, 22, 23)", 11, 22, 23, 255, false},
		{"rgb(255, 255, 255)", 255, 255, 255, 255, false},
		{"rgb(255, 0, 0)", 255, 0, 0, 255, false},
		{"rgb(0, 255, 0)", 0, 255, 0, 255, false},
		{"rgb(0, 0, 255)", 0, 0, 255, 255, false},
		{"rgba(0, 0, 0, 0)", 0, 0, 0, 0, false},
		{"rgba(255, 255, 255, 255)", 255, 255, 255, 255, false},
		{"rgba(255, 0, 0, 255)", 255, 0, 0, 255, false},
		{"rgba(0, 255, 0, 255)", 0, 255, 0, 255, false},
		{"rgba(0, 0, 255, 255)", 0, 0, 255, 255, false},
		{"#000", 0, 0, 0, 0, true},
		{"#0000000", 0, 0, 0, 0, true},
		{"#000000000", 0, 0, 0, 0, true},
		{"#0000000000", 0, 0, 0, 0, true},
		{"#00000000000", 0, 0, 0, 0, true},
		{"#üdußnz", 0, 0, 0, 0, true},
		{"#üdußnzßz", 0, 0, 0, 0, true},
		{"rgb(foo, bar, baz)", 0, 0, 0, 0, true},
		{"rgba(foo, bar, baz, qux)", 0, 0, 0, 0, true},
		{"#ff1493", 255, 20, 147, 255, false},
		{"#ff149380", 255, 20, 147, 128, false},
		{"rgb(255, 20, 147)", 255, 20, 147, 255, false},
		{"rgb(255, 20, 147, 128)", 0, 0, 0, 0, true},
		{"rgba(255, 20, 147, 128)", 255, 20, 147, 128, false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d: %s", i, test.color), func(t *testing.T) {
			r, g, b, a, err := ui.ParseColorChannels(test.color)
			if err != nil && !test.err {
				t.Errorf("unexpected error: %s", err)
			} else if err == nil && test.err {
				t.Errorf("expected error, got nil (%d, %d, %d, %d)", r, g, b, a)
			} else if r != test.r || g != test.g || b != test.b || a != test.a {
				t.Errorf("expected (%d, %d, %d, %d), got (%d, %d, %d, %d)", test.r, test.g, test.b, test.a, r, g, b, a)
			}
		})
	}
}
