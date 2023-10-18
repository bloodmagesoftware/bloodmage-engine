package engine

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	// Time in seconds since last frame
	DeltaTime        float64
	bkg              sdl.Color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
	title            string
	width            int32   = 800
	width_f64        float64 = float64(width)
	half_width_f64   float64 = width_f64 / 2
	height           int32   = 800
	height_f64       float64 = float64(height)
	half_height_f64  float64 = height_f64 / 2
	center_x         int32   = width / 2
	center_y         int32   = height / 2
	screen_dist      float64 = 0.5
	renderer         *sdl.Renderer
	window           *sdl.Window
	frame_start_time uint64
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

func Start(t string, fs bool) {
	title = t
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")
	var err error
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	if fs {
		window, err = sdl.CreateWindow(
			title,
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			width, height,
			sdl.WINDOW_SHOWN|sdl.WINDOW_FULLSCREEN)
	} else {
		window, err = sdl.CreateWindow(
			title,
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			width, height,
			sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	}
	if err != nil {
		panic(err)
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	sdl.ShowCursor(sdl.DISABLE)
	sdl.SetRelativeMouseMode(true)
	window.SetGrab(true)
	running = true
	updateWindowSize()
	frame_start_time = sdl.GetTicks64()
}

func updateWindowSize() {
	width, height = window.GetSize()
	width_f64 = float64(width)
	half_width_f64 = width_f64 / 2
	height_f64 = float64(height)
	half_width_f64 = width_f64 / 2
	screen_dist = half_width_f64 / math.Tan(half_fov)

	center_x = width / 2
	center_y = height / 2

	num_of_rays = width / 2
	delta_angle = fov / (width_f64 / 2)
	scale = width / int32(num_of_rays)
}

func Stop() {
	running = false
	window.Destroy()
	renderer.Destroy()
	sdl.Quit()
}

func input() {
	keystates = sdl.GetKeyboardState()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			running = false
		case sdl.WINDOWEVENT:
			updateWindowSize()
		}
	}
	getMouseInput()
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
	DeltaTime = float64(now-frame_start_time) / 1000.0
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
