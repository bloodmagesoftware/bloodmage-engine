package engine

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	KeyForward         = sdl.SCANCODE_W
	KeyBack            = sdl.SCANCODE_S
	KeyLeft            = sdl.SCANCODE_A
	KeyRight           = sdl.SCANCODE_D
	KeyEscape          = sdl.SCANCODE_ESCAPE
	speed      float64 = 3
	turnSpeed  float64 = 0.1
)

var (
	MouseX, MouseY           int32
	MouseDeltaX, MouseDeltaY int32
	mouse_state              uint32
)

func getMouseInput() {
	MouseX, MouseY, mouse_state = sdl.GetMouseState()
	MouseDeltaX = MouseX - center_x
	MouseDeltaY = MouseY - center_y
	window.WarpMouseInWindow(center_x, center_y)
}

func KeyDown(key uint8) bool {
	return keystates[key] != 0
}

func MovePlayer() {
	P.Angle += float64(MouseDeltaX) * turnSpeed * DeltaTime

	if KeyDown(KeyForward) {
		P.Speed = speed
	} else if KeyDown(KeyBack) {
		P.Speed = -speed
	} else {
		P.Speed = 0
	}

	if KeyDown(KeyLeft) {
		P.Strafe = -speed
	} else if KeyDown(KeyRight) {
		P.Strafe = speed
	} else {
		P.Strafe = 0
	}

	var x, y int
	x_dir := P.Speed*math.Cos(P.Angle) + P.Strafe*math.Cos(P.Angle+math.Pi/2)
	new_x := P.X + x_dir*DeltaTime
	buffer_x := new_x + x_dir*0.01
	y = int(math.Floor(P.Y))
	x = int(math.Floor(buffer_x))
	if x > 0 && x < len(level[y]) && level[y][x] == 0 {
		P.X = new_x
	}

	y_dir := P.Speed*math.Sin(P.Angle) + P.Strafe*math.Sin(P.Angle+math.Pi/2)
	new_y := P.Y + y_dir*DeltaTime
	buffer_y := new_y + y_dir*0.01
	x = int(math.Floor(P.X))
	y = int(math.Floor(buffer_y))
	if y > 0 && y < len(level) && level[y][x] == 0 {
		P.Y = new_y
	}
}
