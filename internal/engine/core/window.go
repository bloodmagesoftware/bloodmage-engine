// Bloodmage Engine
// Copyright (C) 2023 Frank Mayer
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
	width          int32 = 800
	widthF               = float32(width)
	halfWidthF           = widthF / 2
	height         int32 = 800
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

func Width() int32 {
	return width
}

func WidthF() float32 {
	return widthF
}

func HalfWidthF() float32 {
	return halfWidthF
}

func Height() int32 {
	return height
}

func HeightF() float32 {
	return heightF
}

func HalfHeightF() float32 {
	return halfHeightF
}

func CenterX() int32 {
	return centerX
}

func CenterY() int32 {
	return centerY
}

func Renderer() *sdl.Renderer {
	return renderer
}

func Window() *sdl.Window {
	return window
}

func KeyStates() []uint8 {
	return keyStates
}

func IsCursorLocked() bool {
	return cursorLocked
}

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

func OnResize(f *func()) {
	onResize = f
}

func Start(t string) {
	if window != nil {
		log.Fatal("window already started")
	}
	title = t
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")
	var err error
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	windowFlags := uint32(sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)
	if optionData.Fullscreen {
		windowFlags |= sdl.WINDOW_FULLSCREEN
	}
	window, err = sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		windowFlags)
	if err != nil {
		log.Fatal(err)
	}

	rendererFlags := uint32(sdl.RENDERER_ACCELERATED)
	if optionData.Vsync {
		rendererFlags |= sdl.RENDERER_PRESENTVSYNC
	}
	renderer, err = sdl.CreateRenderer(window, -1, rendererFlags)
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

func Present() {
	renderer.Present()
	frameTime := uint32(sdl.GetTicks64() - frameStartTime)
	if frameTime < targetFrameTime {
		sdl.Delay(targetFrameTime - frameTime)
	}
}
