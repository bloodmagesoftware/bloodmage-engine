package engine

type player struct {
	X      float64
	Y      float64
	Angle  float64
	Speed  float64
	Strafe float64
}

var (
	P = player{cell_size * 2, cell_size * 1.5, 0, 0, 0}
)
