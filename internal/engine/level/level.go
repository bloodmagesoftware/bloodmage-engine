package level

import (
	"google.golang.org/protobuf/proto"
	"os"
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

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
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
	if !InBounds(x, y) {
		return true
	}
	row := currentLevel.Collision[y]
	return row&(1<<x) != 0
}
