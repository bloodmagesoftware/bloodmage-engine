package main

import (
	"flag"
	"os"
	"path"

	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/pkg/engine/level"
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	mouseX int32
	mouseY int32
	posX   int32
	posY   int32
)

func main() {
	// get level file from command line arguments
	levelParam := flag.String("level", "", "level file to load")
	flag.Parse()

	// check if level file was provided
	if levelParam == nil || *levelParam == "" {
		log.Fatal("No level file provided")
	}

	levelFile := path.Clean(*levelParam)

	var l *level.Level
	var err error

	// check if file exists on disk
	if _, err = os.Stat(levelFile); os.IsNotExist(err) {
		log.Warn("Level file does not exist")
		l = level.New()
	} else {
		// load level file
		l, err = level.Load(levelFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// set current level
	level.Set(l)

	// start editor
	core.InitOptions()
	core.Options().Fullscreen = false
	core.Start("Bloodmage Engine - Editor")
	defer core.Stop()

	rect := sdl.Rect{X: 0, Y: 0, W: 32, H: 32}

	// game loop
	for core.Running() {
		// control camera
		mx, my, mouseState := sdl.GetMouseState()
		if mouseState&sdl.ButtonMMask() != 0 {
			posX -= int32(mx - mouseX)
			posY -= int32(my - mouseY)

			if posX < 0 {
				posX = 0
			}
			if posY < 0 {
				posY = 0
			}

			if posX > l.Width*32 {
				posX = l.Width * 32
			}
			if posY > l.Height*32 {
				posY = l.Height * 32
			}

			mouseX = mx
			mouseY = my
		} else {
			mouseX = mx
			mouseY = my
		}

		if mouseState&sdl.ButtonLMask() != 0 {
			x := (mouseX + posX) / 32
			y := (mouseY + posY) / 32

			l.SetCollision(int(x), int(y), true)
		} else if mouseState&sdl.ButtonRMask() != 0 {
			x := (mouseX + posX) / 32
			y := (mouseY + posY) / 32

			l.SetCollision(int(x), int(y), false)
		}

		// draw level
		for x := int32(0); x < l.Width; x++ {
			for y := int32(0); y < l.Height; y++ {
				collision := l.Collision[y*l.Width+x] == 1

				rect.X = x*32 - posX

				if rect.X < -32 || rect.X > core.Width() {
					continue
				}

				rect.Y = y*32 - posY

				if rect.Y < -32 || rect.Y > core.Height() {
					continue
				}

				if collision {
					_ = core.Renderer().SetDrawColor(255, 0, 0, 255)
				} else {
					_ = core.Renderer().SetDrawColor(10, 10, 10, 255)
				}
				_ = core.Renderer().FillRect(&rect)
			}
		}

		// draw frame
		core.Present()
	}

	// save level
	err = l.Save(levelFile)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Saved level to " + levelFile)
	}
}
