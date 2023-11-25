package level

type Level struct {
	Collision []uint32
	Textures  [][]byte
}

var (
	currentLevelWidth    int   = 1
	currentLevelWidth32  int32 = 1
	currentLevelHeight   int   = 1
	currentLevelHeight32 int32 = 1
	currentLevel         *Level
)

func Set(level *Level) {
	currentLevel = level
	currentLevelHeight = len(level.Textures)
	for _, row := range level.Textures {
		rowLen := len(row)
		if rowLen > currentLevelWidth {
			currentLevelWidth = rowLen
		}
	}
	currentLevelWidth32 = int32(currentLevelWidth)
	currentLevelHeight32 = int32(currentLevelHeight)
}

func Width() int {
	return currentLevelWidth
}

func Width32() int32 {
	return currentLevelWidth32
}

func Height() int {
	return currentLevelHeight
}

func Height32() int32 {
	return currentLevelHeight32
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

func Texture(x int, y int) byte {
	if !InBounds(x, y) {
		return 0
	}
	return currentLevel.Textures[y][x]
}
