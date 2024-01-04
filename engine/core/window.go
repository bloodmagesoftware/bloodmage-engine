// Bloodmage Engine
// Copyright (C) 2024 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

package core

import (
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	targetFrameTime uint32 = 1000 / 60
)

var (
	// DeltaTime is the time in seconds since last frame
	DeltaTime      float32
	bkg            = sdl.Color{R: 0, G: 0, B: 0, A: 255}
	title          string
	width          int32 = 1280
	widthF               = float32(width)
	halfWidthF           = widthF / 2
	height         int32 = 720
	heightF              = float32(height)
	halfHeightF          = heightF / 2
	centerX              = width / 2
	centerY              = height / 2
	renderer       *sdl.Renderer
	window         *sdl.Window
	frameStartTime uint64
	keyStates      = sdl.GetKeyboardState()
	running        bool
	cursorLocked   bool
	onResize       *func() = nil
)

// Width returns the window width as an int32.
func Width() int32 {
	return width
}

// WidthF returns the window width as a float32.
func WidthF() float32 {
	return widthF
}

// HalfWidthF returns half of the window width as a float32.
func HalfWidthF() float32 {
	return halfWidthF
}

// Height returns the window height as an int32.
func Height() int32 {
	return height
}

// HeightF returns the window height as a float32.
func HeightF() float32 {
	return heightF
}

// HalfHeightF returns half of the window height as a float32.
func HalfHeightF() float32 {
	return halfHeightF
}

// CenterX returns the center of the window on the X axis.
// This is equal to Width() / 2.
func CenterX() int32 {
	return centerX
}

// CenterY returns the center of the window on the Y axis.
// This is equal to Height() / 2.
func CenterY() int32 {
	return centerY
}

// Renderer returns the SDL renderer.
func Renderer() *sdl.Renderer {
	return renderer
}

// Window returns the SDL window.
func Window() *sdl.Window {
	return window
}

// KeyStates returns a snapshot of the current state of the keyboard. <https://wiki.libsdl.org/SDL_GetKeyboardState>
func KeyStates() []uint8 {
	return keyStates
}

// IsCursorLocked returns whether the cursor is locked to the window or not.
func IsCursorLocked() bool {
	return cursorLocked
}

// LockCursor sets whether the cursor should be locked to the window or not.
func LockCursor(lock bool) {
	cursorLocked = lock
	if cursorLocked {
		_, err := sdl.ShowCursor(sdl.DISABLE)
		if err != nil {
			log.Error(err)
			return
		}
		sdl.SetRelativeMouseMode(true)
		window.SetGrab(true)
		window.WarpMouseInWindow(centerX, centerY)
	} else {
		_, err := sdl.ShowCursor(sdl.ENABLE)
		if err != nil {
			log.Error(err)
			return
		}
		sdl.SetRelativeMouseMode(false)
		window.SetGrab(false)
	}
}

// OnResize sets the function to be called when the window is resized.
// This is for the game modes like firstperson, topdown, etc.
func OnResize(f *func()) {
	onResize = f
}

// Start initializes the window and renderer.
func Start(t string) {
	if window != nil {
		log.Fatal("window already started")
	}

	var err error

	title = t

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}

	dimension, err := sdl.GetDisplayBounds(0)
	if err != nil {
		log.Fatal(err)
	}
	width = dimension.W
	widthF = float32(width)
	halfWidthF = widthF / 2
	centerX = width / 2
	height = dimension.H
	heightF = float32(height)
	halfHeightF = heightF / 2
	centerY = height / 2

	var windowFlags uint32
	if optionData.Fullscreen {
		if optionData.WindowedFullscreen {
			windowFlags = uint32(sdl.WINDOW_SHOWN)
		} else {
			windowFlags = uint32(sdl.WINDOW_SHOWN | sdl.WINDOW_FULLSCREEN_DESKTOP)
		}
	} else {
		windowFlags = uint32(sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)
	}
	window, err = sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		windowFlags)
	if err != nil {
		log.Fatal(err)
	}
	if optionData.WindowedFullscreen {
		window.SetBordered(false)
		window.SetResizable(false)
		window.SetPosition(dimension.X, dimension.Y)
		window.SetSize(dimension.W, dimension.H)
	}

	rendererFlags := uint32(sdl.RENDERER_ACCELERATED)
	if optionData.Vsync {
		rendererFlags |= sdl.RENDERER_PRESENTVSYNC
	}
	renderer, err = sdl.CreateRenderer(window, -1, rendererFlags)
	if err != nil {
		log.Fatal(err)
	}
	err = renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		log.Fatal(err)
	}
	running = true
	updateWindowSize()
	frameStartTime = sdl.GetTicks64()
}

func updateWindowSize() {
	width, height = window.GetSize()
	widthF = float32(width)
	halfWidthF = widthF / 2
	heightF = float32(height)

	centerX = width / 2
	centerY = height / 2

	if onResize != nil {
		(*onResize)()
	}
}

// Stop destroys the window and renderer.
// Call this after the game loop has exited.
func Stop() {
	running = false
	_ = window.Destroy()
	_ = renderer.Destroy()
	sdl.Quit()
}

func beginRender() {
	if err := renderer.SetDrawColor(bkg.R, bkg.G, bkg.B, bkg.A); err != nil {
		log.Error(err)
	}
	if err := renderer.Clear(); err != nil {
		log.Error(err)
	}
}

// Running returns whether the game loop should continue or not.
// Use this to determine when to exit the game loop.
func Running() bool {
	now := sdl.GetTicks64()
	DeltaTime = float32(now-frameStartTime) / 1000.0
	frameStartTime = now

	eventLoop()
	beginRender()
	return running
}

func eventLoop() {
	keyStates = sdl.GetKeyboardState()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			running = false
		case sdl.WINDOWEVENT:
			updateWindowSize()
		}
	}
}

// Present draws the frame to the screen.
// Call this at the very end of the game loop.
func Present() {
	renderer.Present()
	frameTime := uint32(sdl.GetTicks64() - frameStartTime)
	if frameTime < targetFrameTime {
		sdl.Delay(targetFrameTime - frameTime)
	}
}
