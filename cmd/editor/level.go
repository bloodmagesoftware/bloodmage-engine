package main

import (
	"errors"
	"flag"
	"os"
	"path"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/engine/level"
	"github.com/charmbracelet/log"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	// l is the current level
	l *level.Level
	// levelFile is the path to the level file
	levelFile string
	// unsavedChanges is true if the level has unsaved changes
	unsavedChanges bool
)

func InitLevel() error {
	// get level file from command line arguments
	levelParam := flag.String("level", "", "level file to load")
	flag.Parse()

	// check if level file was provided
	if levelParam == nil || *levelParam == "" {
		return errors.New("no level file provided")
	}

	levelFile = path.Clean(*levelParam)

	// check if file exists on disk
	if _, err := os.Stat(levelFile); os.IsNotExist(err) {
		log.Warn("Level file does not exist (creating new level)")
		l = level.New()
		unsavedChanges = true
	} else {
		// load level file
		l, err = level.Load(levelFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	level.Set(l)
	return nil
}

func SaveLevel() {
	// save level file
	err := l.Save(levelFile)
	if err != nil {
		log.Fatal(
			"Failed to save level to "+levelFile,
			"error", err,
		)
	}

	log.Info("Level saved to " + levelFile)
	unsavedChanges = false
}

var (
	posX int32
	posY int32
)

var (
	unit = int32(64)
)

func DrawLevel() error {
	rect := sdl.Rect{X: 0, Y: 0, W: unit, H: unit}

	// control camera
	if core.MouseState&sdl.ButtonMMask() != 0 {
		posX -= core.MouseDeltaX
		posY -= core.MouseDeltaY

		if posX < -unit {
			posX = -unit
		}
		if posY < -unit {
			posY = -unit
		}

		if posX > l.Width*unit-unit {
			posX = l.Width*unit - unit
		}
		if posY > l.Height*unit-unit {
			posY = l.Height*unit - unit
		}
	}

	// draw level
	for x := int32(0); x < l.Width; x++ {
		rect.X = x*unit - posX
		if rect.X < -unit || rect.X > core.Width() {
			continue
		}

		for y := int32(0); y < l.Height; y++ {
			rect.Y = y*unit - posY
			if rect.Y < -unit || rect.Y > core.Height() {
				continue
			}

			collision := l.SaveCollision(int(x), int(y))
			if collision {
				et := level.WallTexture(int(x), int(y))
				t, err := et.Texture()
				if err != nil {
					return err
				}
				if err = core.Renderer().Copy(t, nil, &rect); err != nil {
					return err
				}
			} else {
				if err := core.Renderer().SetDrawColor(20, 20, 20, 255); err != nil {
					return err
				}
				if err := core.Renderer().FillRect(&rect); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func EditLevel() {
	if core.MouseState&sdl.ButtonLMask() != 0 {
		x := (core.MouseX + posX) / unit
		y := (core.MouseY + posY) / unit
		l.SetWall(int(x), int(y), 1)
		unsavedChanges = true
	} else if core.MouseState&sdl.ButtonRMask() != 0 {
		x := (core.MouseX + posX) / unit
		y := (core.MouseY + posY) / unit
		l.SetWall(int(x), int(y), 0)
		unsavedChanges = true
	}
}
