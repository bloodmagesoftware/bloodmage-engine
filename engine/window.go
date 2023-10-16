package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	bkg              sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
	title            string
	width            int32   = 800
	width_f64        float64 = float64(width)
	height           int32   = 800
	height_f64       float64 = float64(height)
	renderer         *sdl.Renderer
	window           *sdl.Window
	frame_start_time uint64
	DeltaTime        uint64
	mouse            sdl.Point
	mousestate       uint32
	keystates        = sdl.GetKeyboardState()
	event            sdl.Event
	running          bool
)

const (
	target_frame_time uint64 = 1000 / 60
	pixel_scale       int32  = 4
)

func setColor(r, g, b, a uint8) sdl.Color {
	var c sdl.Color
	c.R = r
	c.G = g
	c.B = b
	c.A = a
	return c
}

func Start(t string) {
	title = t
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")
	var err error
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	window, err = sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		sdl.WINDOW_SHOWN) //|sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		panic(err)
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	running = true
	width, height = window.GetSize()
    width_f64 = float64(width)
    height_f64 = float64(height)
	frame_start_time = sdl.GetTicks64()
}

func Quit() {
	running = false
	window.Destroy()
	renderer.Destroy()
	sdl.Quit()
}

func input() {
	keystates = sdl.GetKeyboardState()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}
	}
	mouse.X, mouse.Y, mousestate = sdl.GetMouseState()
}

func beginRender() {
	if err := renderer.SetDrawColor(bkg.R, bkg.G, bkg.B, bkg.A); err != nil {
		panic(err)
	}
	if err := renderer.Clear(); err != nil {
		panic(err)
	}
}

func Running() bool {
	now := sdl.GetTicks64()
	DeltaTime = now - frame_start_time
	frame_start_time = now

	input()
	beginRender()
	return running
}

func Present() {
	renderer.Present()
	frame_time := sdl.GetTicks64() - frame_start_time
	if frame_time < target_frame_time {
		sdl.Delay(uint32(target_frame_time - frame_time))
	}
}
