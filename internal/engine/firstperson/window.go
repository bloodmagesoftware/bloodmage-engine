package firstperson

import (
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/core"
	"github.com/bloodmagesoftware/bloodmage-engine/internal/engine/mathf"
	"github.com/charmbracelet/log"
)

func Init() {
	f := func() {
		log.Debug("firstperson window resize")
		screenDist = core.HalfWidthF() / mathf.Tan(halfFov)
		numOfRays = core.Width() / core.Options().PixelScale
		deltaAngle = fov / (core.WidthF() / float32(core.Options().PixelScale))
		scale = core.Width() / numOfRays
	}
	core.OnResize(&f)
}
