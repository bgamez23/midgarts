package rendercmd

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/project-midgard/midgarts/internal/graphic"
)

type SpriteRenderCommand struct {
	Scale           [2]float32
	Size            mgl32.Vec2
	Position        mgl32.Vec3
	Offset          mgl32.Vec2
	RotationRadians float32
	Texture         *graphic.Texture
	FlipVertically  bool
}
