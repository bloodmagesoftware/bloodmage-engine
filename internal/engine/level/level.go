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

package level

import (
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

var (
	currentLevelWidth  = 0
	currentLevelHeight = 0
	currentLevel       = &Level{
		Width:           0,
		Height:          0,
		Collision:       []byte{},
		FloorTextures:   []byte{},
		WallTextures:    []byte{},
		CeilingTextures: []byte{},
	}
)

func Set(level *Level) {
	currentLevel = level
	currentLevelWidth = int(level.Width)
	currentLevelHeight = int(level.Height)
}

func Load(path string) (*Level, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	level := &Level{}
	err = proto.Unmarshal(b, level)
	if err != nil {
		return nil, err
	}

	return level, nil
}

func (self *Level) Save(path string) error {
	b, err := proto.Marshal(self)
	if err != nil {
		return err
	}

	// ensure directory exists
	dir := filepath.Dir(path)
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
        return err
    }

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Enlarge the level to the given width and height.
// If the level is already larger than the given width and height, this function does nothing.
// If the level is smaller than the given width and height, the level is enlarged to the given width and height.
// The new cells are filled with the value 0.
// The object is modified in place.
func (self *Level) Enlarge(width int32, height int32) {
	if width <= self.Width && height <= self.Height {
		return
	}

	newWidth := width
	if width < self.Width {
		newWidth = self.Width
	}

	newHeight := height
	if height < self.Height {
		newHeight = self.Height
	}

	newCollision := make([]byte, newWidth*newHeight)
	newFloorTextures := make([]byte, newWidth*newHeight)
	newWallTextures := make([]byte, newWidth*newHeight)
	newCeilingTextures := make([]byte, newWidth*newHeight)

	for x := int32(0); x < newWidth; x++ {
		for y := int32(0); y < newHeight; y++ {
			if x < self.Width && y < self.Height {
				newCollision[y*newWidth+x] = self.Collision[y*self.Width+x]
				newFloorTextures[y*newWidth+x] = self.FloorTextures[y*self.Width+x]
				newWallTextures[y*newWidth+x] = self.WallTextures[y*self.Width+x]
				newCeilingTextures[y*newWidth+x] = self.CeilingTextures[y*self.Width+x]
			} else {
				newCollision[y*newWidth+x] = 0
				newFloorTextures[y*newWidth+x] = 0
				newWallTextures[y*newWidth+x] = 0
				newCeilingTextures[y*newWidth+x] = 0
			}
		}
	}

	self.Width = newWidth
	self.Height = newHeight
	self.Collision = newCollision
	self.FloorTextures = newFloorTextures
	self.WallTextures = newWallTextures
	self.CeilingTextures = newCeilingTextures
}

func (self *Level) SaveCollision(x int, y int) bool {
	if !self.InBounds(x, y) {
		return true
	}
	return self.Collision[y*int(self.Width)+x] == 1
}

func (self *Level) InBounds(x int, y int) bool {
	return x >= 0 && x < int(self.Width) && y >= 0 && y < int(self.Height)
}

func New() *Level {
	return &Level{
		Width:  5,
		Height: 5,
		Collision: []byte{
			1, 1, 1, 1, 1,
			0, 0, 0, 0, 1,
			1, 0, 0, 0, 0,
			0, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
		FloorTextures: []byte{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
		},
		WallTextures: []byte{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
		},
		CeilingTextures: []byte{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
		},
	}
}

func Width() int {
	return currentLevelWidth
}

func Width32() int32 {
	return currentLevel.Width
}

func Height() int {
	return currentLevelHeight
}

func Height32() int32 {
	return currentLevel.Height
}

func InBounds(x int, y int) bool {
	return x >= 0 && x < currentLevelWidth && y >= 0 && y < currentLevelHeight
}

func Collision(x int, y int) bool {
	return currentLevel.SaveCollision(x, y)
}

const round = 0.4921875

func CollisionF(x float32, y float32) bool {
    for y1 := int(y - round); y1 <= int(y+round); y1++ {
        for x1 := int(x - round); x1 <= int(x+round); x1++ {
            if Collision(x1, y1) {
                return true
            }
        }
    }
    return false
}

func (self *Level) SetCollision(x int, y int, collision bool) {
	if !InBounds(x, y) {
		self.Enlarge(int32(x+1), int32(y+1))
	}
	if collision {
		self.Collision[y*int(self.Width)+x] = 1
	} else {
		self.Collision[y*int(self.Width)+x] = 0
	}
}
