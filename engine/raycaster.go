package engine

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	cell_size         = 32
	fov       float64 = math.Pi / 3
	half_fov          = fov / 2
	max_depth         = 20
	epsilon           = 1e-6
)

var (
	num_of_rays int32
	delta_angle float64
	scale       int32

	level = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
)

func RenderViewport() {
	ox := P.X
	oy := P.Y
	x_level := math.Floor(ox)
	y_level := math.Floor(oy)

	ray_angle := P.Angle - half_fov + epsilon
	for ray := int32(0); ray < num_of_rays; ray++ {
		sin_a := math.Sin(ray_angle)
		cos_a := math.Cos(ray_angle)

		var dy float64
		var dx float64
		var delta_depth float64

		// horizontals
		var y_hor float64
		if sin_a > 0 {
			y_hor = y_level + 1
			dy = 1.0
		} else {
			y_hor = y_level - epsilon
			dy = -1.0
		}
		depth_hor := (y_hor - oy) / sin_a
		x_hor := ox + depth_hor*cos_a
		delta_depth = dy / sin_a
		dx = delta_depth * cos_a
		for i := 0; i < max_depth; i++ {
			tile_x := int(math.Floor(x_hor))
			tile_y := int(math.Floor(y_hor))
			if tile_x < 0 || tile_x >= len(level[0]) || tile_y < 0 || tile_y >= len(level) {
				break
			}
			if level[tile_y][tile_x] != 0 {
				break
			}
			x_hor += dx
			y_hor += dy
			depth_hor += delta_depth
		}

		// verticals
		var x_vert float64
		if cos_a > 0 {
			x_vert = x_level + 1
			dx = 1.0
		} else {
			x_vert = x_level - epsilon
			dx = -1.0
		}
		depth_vert := (x_vert - ox) / cos_a
		y_vert := oy + depth_vert*sin_a
		delta_depth = dx / cos_a
		dy = delta_depth * sin_a
		for i := 0; i < max_depth; i++ {
			tile_x := int(math.Floor(x_vert))
			tile_y := int(math.Floor(y_vert))
			if tile_x < 0 || tile_x >= len(level[0]) || tile_y < 0 || tile_y >= len(level) {
				break
			}
			if level[tile_y][tile_x] != 0 {
				break
			}
			x_vert += dx
			y_vert += dy
			depth_vert += delta_depth
		}

		// depth
		var depth float64
		if depth_hor < depth_vert {
			depth = depth_hor
		} else {
			depth = depth_vert
		}

		// remove fish eye
		depth *= math.Cos(P.Angle - ray_angle)

		// projection
		proj_height := screen_dist / (depth + epsilon)

		// draw wall
		rect := sdl.Rect{
			ray * scale, int32(half_height_f64 - proj_height/2),
			scale, int32(proj_height),
		}
		// ditsant walls are darker
		darknes := uint8(255 / (depth + 1))
		renderer.SetDrawColor(darknes, darknes, darknes, 255)
		renderer.FillRect(&rect)

		ray_angle += delta_angle
	}
}
