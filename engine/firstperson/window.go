// Bloodmage Engine - Retro first person game engine
// Copyright (C) 2024  Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package firstperson

import (
	"github.com/bloodmagesoftware/bloodmage-engine/engine/core"
	"github.com/charmbracelet/log"
	"github.com/chewxy/math32"
)

func Init() {
	f := func() {
		log.Debug("firstperson window resize")
		screenDist = core.HalfWidthF() / math32.Tan(halfFov)
		numOfRays = core.Width() / core.Options().PixelScale
		deltaAngle = fov / (core.WidthF() / float32(core.Options().PixelScale))
		scale = core.Width() / numOfRays
	}
	core.OnResize(&f)
}
