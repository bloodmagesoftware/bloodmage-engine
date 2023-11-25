package firstperson

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
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
